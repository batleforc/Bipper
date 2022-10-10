import type { ModelUser } from "@/api/Api";

export const storeRefreshToken = (token: string) => {
  localStorage.setItem("refresh_token", token);
};

export const storeAccessToken = (token: string) => {
  sessionStorage.setItem("access_token", token);
};

export const storeUserInfo = (userInfo: ModelUser) => {
  sessionStorage.setItem("user_info", JSON.stringify(userInfo));
};

export const getRegreshToken = () => {
  return localStorage.getItem("refresh_token");
};

export const getAccessToken = () => {
  return sessionStorage.getItem("access_token");
};

export const getUserInfo = () => {
  const userInfo = sessionStorage.getItem("user_info");
  if (userInfo === null) return null;
  return JSON.parse(userInfo);
};
