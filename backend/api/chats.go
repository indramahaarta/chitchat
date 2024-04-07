package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	db "github.com/indramahaarta/chitchat/db/sqlc"
)

type getChatsDetailsParams struct {
	Uid string `uri:"uid"`
}

type getChatsDetailsResponse struct {
	Message string       `json:"message"`
	User    UserResponse `json:"user"`
	Chats   []Chats      `json:"chats"`
}

// @Summary Get chat details
// @Description Retrieves the details of the chat history between the logged-in user and the user specified by UID
// @Tags chat
// @Accept json
// @Produce json
// @Param uid path string true "User ID to get chat details with"
// @Security ApiKeyAuth
// @Success 200 {object} getChatsDetailsResponse "Successful retrieval of chat details"
// @Failure 400 {object} ErrorResponse "Bad request - Invalid UID"
// @Failure 401 {object} ErrorResponse "Unauthorized - Invalid or missing JWT token"
// @Failure 404 {object} ErrorResponse "Not Found - User not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /chats/{uid} [get]
func (server *Server) getChatsDetails(ctx *gin.Context) {
	payload, err := getUserPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	var params getChatsDetailsParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if payload.Uid.String() == params.Uid {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("payload uid and params uid should be different")))
	}

	arg := db.GetChatsDetailsParams{
		SenderUid:   payload.Uid,
		ReceiverUid: uuid.MustParse(params.Uid),
	}

	chats, err := server.store.GetChatsDetails(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var parsedChats []Chats
	for _, chat := range chats {
		var sender UserResponse
		var receiver UserResponse
		json.Unmarshal([]byte(chat.Sender), &sender)
		json.Unmarshal([]byte(chat.Receiver), &receiver)

		parsedChats = append(parsedChats, Chats{
			ID:          chat.ID,
			Content:     chat.Content,
			CreatedAt:   chat.CreatedAt,
			SenderUid:   chat.SenderUid,
			ReceiverUid: chat.ReceiverUid,
			Sender:      sender,
			Receiver:    receiver,
		})
	}

	user, err := server.store.GetUserById(ctx, uuid.MustParse(params.Uid))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := getChatsDetailsResponse{
		Message: "success to get chats details",
		User:    *ReturnUserResponse(&user),
		Chats:   parsedChats,
	}

	ctx.JSON(http.StatusOK, res)
}

type History struct {
	ID              uuid.UUID `json:"id"`
	SenderUid       uuid.UUID `json:"sender_uid"`
	SenderName      string    `json:"sender_name"`
	ReceiverUid     uuid.UUID `json:"receiver_uid"`
	ReceiverName    string    `json:"receiver_name"`
	LatestCreatedAt time.Time `json:"latest_created_at"`
	LatestContent   string    `json:"latest_content"`
}

type getChatsHistoriesResponse struct {
	Message   string    `json:"message"`
	Histories []History `json:"histories"`
}

// @Summary Get chat histories
// @Description Retrieves the chat histories for the logged-in user
// @Tags chat
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} getChatsHistoriesResponse "Successful retrieval of chat histories"
// @Failure 401 {object} ErrorResponse "Unauthorized - Invalid or missing JWT token"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /chats/histories [get]
func (server *Server) getChatsHistories(ctx *gin.Context) {
	payload, err := getUserPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	histories, err := server.store.GetChatsHistories(ctx, payload.Uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var parsedHistories []History
	for _, history := range histories {
		parsedHistories = append(parsedHistories, History{
			ID:              history.ID,
			SenderUid:       history.SenderUid,
			SenderName:      history.SenderName.String,
			ReceiverUid:     history.ReceiverUid,
			ReceiverName:    history.ReceiverName.String,
			LatestCreatedAt: history.LatestCreatedAt,
			LatestContent:   history.LatestContent,
		})
	}

	res := getChatsHistoriesResponse{
		Message:   "success to get chat history",
		Histories: parsedHistories,
	}

	ctx.JSON(http.StatusOK, res)
}

var clients = make(map[uuid.UUID]*websocket.Conn)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ChatParam struct {
	Uid string `uri:"uid"`
}

type Chats struct {
	ID          uuid.UUID    `json:"id"`
	Content     string       `json:"content"`
	CreatedAt   time.Time    `json:"created_at"`
	SenderUid   uuid.UUID    `json:"sender_uid"`
	ReceiverUid uuid.UUID    `json:"receiver_uid"`
	Sender      UserResponse `json:"sender"`
	Receiver    UserResponse `json:"receiver"`
}

type ChatBody struct {
	Content     string `json:"content"`
	SenderUid   string `json:"sender_uid"`
	ReceiverUid string `json:"receiver_uid"`
}

type WSData struct {
	History History `json:"history"`
	Chat    Chats   `json:"chat"`
}

func (server *Server) chat(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println("masuk", err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	defer conn.Close()

	accessToken, err := ctx.Request.Cookie("access_token")
	if err != nil {
		sendWebSocketError(conn, "access_token is not found in cookies")
		return
	}

	payload, err := server.tokenMaker.VerifyToken(accessToken.Value)
	if err != nil {
		sendWebSocketError(conn, "can't parsing payload")
		return
	}

	clients[payload.Uid] = conn

	for {
		var chat *ChatBody
		err = conn.ReadJSON(&chat)
		if err != nil {
			fmt.Printf("error reading chat body: %v\n", err)
			if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				fmt.Println("webSocket connection closed normally")
				break
			} else {
				sendWebSocketError(conn, "error reading chat body")
				continue
			}
		}

		if chat == nil {
			sendWebSocketError(conn, "invalid chat body or sender uid")
			break
		}

		senderUID, err := uuid.Parse(chat.SenderUid)
		if err != nil {
			sendWebSocketError(conn, "invalid sender uid")
			continue
		}

		receiverUID, err := uuid.Parse(chat.ReceiverUid)
		if err != nil {
			sendWebSocketError(conn, "invalid receiver uid")
			continue
		}

		arg := db.CreateChatParams{
			Content:     chat.Content,
			SenderUid:   senderUID,
			ReceiverUid: receiverUID,
		}

		sender, err := server.store.GetUserById(ctx, senderUID)
		if err != nil {
			sendWebSocketError(conn, err.Error())
			continue
		}

		receiver, err := server.store.GetUserById(ctx, receiverUID)
		if err != nil {
			sendWebSocketError(conn, err.Error())
			continue
		}

		dbChat, err := server.store.CreateChat(ctx, arg)
		if err != nil {
			sendWebSocketError(conn, err.Error())
			continue
		}
		if receiverConn, ok := clients[receiverUID]; ok && receiverConn != conn {
			sendChat(receiverConn, dbChat, sender, receiver)
		}
		sendChat(conn, dbChat, sender, receiver)
	}
}

func sendWebSocketError(conn *websocket.Conn, message string) {
	if err := conn.WriteJSON(map[string]string{"error": message}); err != nil {
		fmt.Printf("error sending WebSocket error: %v\n", err)
		conn.Close()
	}
}

func sendChat(conn *websocket.Conn, chat db.Chats, sender db.Users, receiver db.Users) {
	history := History{
		ID:              chat.ID,
		SenderUid:       chat.SenderUid,
		SenderName:      sender.Name.String,
		ReceiverUid:     chat.ReceiverUid,
		ReceiverName:    receiver.Name.String,
		LatestCreatedAt: chat.CreatedAt,
		LatestContent:   chat.Content,
	}

	senderData := ReturnUserResponse(&sender)
	receiverData := ReturnUserResponse(&receiver)

	newChat := Chats{
		ID:          chat.ID,
		SenderUid:   chat.SenderUid,
		Sender:      *senderData,
		ReceiverUid: chat.ReceiverUid,
		Receiver:    *receiverData,
		Content:     chat.Content,
		CreatedAt:   chat.CreatedAt,
	}

	if err := conn.WriteJSON(WSData{History: history, Chat: newChat}); err != nil {
		fmt.Printf("error sending chat: %v\n", err)
		conn.Close()
	}
}
