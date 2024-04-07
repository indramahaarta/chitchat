package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type getSocialUserResponse struct {
	Message string         `json:"message"`
	Users   []UserResponse `json:"users"`
}

// @Summary Get social users
// @Description Retrieves a list of social users excluding the current user
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} getSocialUserResponse "Successful retrieval of users data"
// @Failure 401 {object} ErrorResponse "Unauthorized access"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /users [get]
func (server *Server) getSocialUser(ctx *gin.Context) {
	payload, err := getUserPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	users, err := server.store.GetAllUser(ctx, payload.Uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var userResponse []UserResponse

	for _, user := range users {
		userResponse = append(userResponse, *ReturnUserResponse(&user))
	}

	res := getSocialUserResponse{
		Message: "success to retrive users data",
		Users:   userResponse,
	}

	ctx.JSON(http.StatusOK, res)
}

// type getUserDetailsParam struct {
// 	Uid string `uri:"uid"`
// }

// type getUserDetailsResponse struct {
// 	Message string `json:"message"`
// 	User UserResponse `json:"user"`
// }

// func (server *Server) getUserDetails(ctx *gin.Context) {
// 	var param getUserDetailsParam
// 	if err := ctx.ShouldBindUri(&param); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	user, err := server.store.GetUserById(ctx, uuid.MustParse(param.Uid))
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	res := getUserDetailsResponse {
// 		Message: "success to get user information",

// 	}
// }
