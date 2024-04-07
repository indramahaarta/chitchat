"use server";

import { Chat, History, User } from "@/models/model";
import { serverAxios } from "@/utils/axios";
import { cookies } from "next/headers";

export async function getHistory(): Promise<{
  histories: History[];
}> {
  "use server";
  try {
    const accessToken = cookies().get("access_token")?.value;
    const { data } = await serverAxios.get<{
      histories: History[];
    }>("/api/chats/histories", {
      headers: { Authorization: `Bearer ${accessToken}` },
    });

    return data;
  } catch (error) {
    throw error;
  }
}

export async function getChatsDetails(uid: string): Promise<{
  chats: Chat[];
  user: User;
}> {
  "use server";
  try {
    const accessToken = cookies().get("access_token")?.value;
    const { data } = await serverAxios.get<{
      chats: Chat[];
      user: User;
    }>(`/api/chats/${uid}`, {
      headers: { Authorization: `Bearer ${accessToken}` },
    });

    return data;
  } catch (error) {
    throw error;
  }
}
