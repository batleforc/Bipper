import { Api, type RouteLogoutBody, type RouteRegisterBody } from "@/api/Api";
import { defineStore } from "pinia";

export const useGlobalStore = defineStore({
  id: "global",
  state: () => ({
    Api: new Api(),
    inited: false,
  }),
  actions: {
    init() {
      this.inited = true;
      if (import.meta.env.VITE_API !== undefined)
        this.Api.baseUrl = import.meta.env.VITE_API;
      // TODO : Check if renew token is valid
      // TODO : If not valid go to /login
      // TODO : If valid go to /home and fetch access token + userinfo
    },
    login({ email, password }: { email: string; password: string }) {
      this.Api.auth
        .loginCreate({ email, password })
        .then((res) => {
          // TODO : Save Refresh token
          // TODO : Save Access token and fetch userinfo
        })
        .catch((err) => {
          console.log(err);
          // TODO : Show error to login screen
        });
    },
    logout() {
      this.Api.auth
        .logoutCreate({
          renew_token: "Renew Token",
        })
        .then((res) => {
          // TODO : Remove Refresh token, call logout endpoint
          // TODO : Remove Access token and userinfo
        })
        .catch((err) => {
          // TODO : Remove refresh token and access token
          // TODO : Notify user that logout failed but device is logged out
        });
    },
    register(registrerBody: RouteRegisterBody) {
      // TODO : Call register endpoint

      this.Api.auth
        .registerCreate(registrerBody)
        .then((res) => {
          // TODO : If success go to /login
        })
        .catch((err) => {
          // TODO : If error show error to register screen
        });
    },
    renewToken() {
      // TODO : Call renew token endpoint
      this.Api.auth
        .renewCreate({
          renew_token: "Renew Token",
        })
        .then((res) => {
          // TODO : Save new access token
        })
        .catch((err) => {
          // TODO : Remove refresh token and access token
          // TODO : Notify user that login failed but device is logged out
        });
    },
  },
});
