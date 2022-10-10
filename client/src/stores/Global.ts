import { Api, type RouteLogoutBody, type RouteRegisterBody } from "@/api/Api";
import {
  getRefreshToken,
  removeAccessToken,
  removeRefreshToken,
  storeAccessToken,
  storeRefreshToken,
} from "@/helper/tokenHandler";
import { defineStore } from "pinia";

export const useGlobalStore = defineStore({
  id: "global",
  state: () => ({
    Api: new Api(),
    inited: false,
    loggedIn: false,
    user: {
      id: 0,
      pseudo: "",
      email: "",
      name: "",
      surname: "",
      picture: "",
      role: "",
      loaded: false,
    },
    chan: {
      myChan: [],
      chan: [],
    },
  }),
  actions: {
    init() {
      this.inited = true;
      if (import.meta.env.VITE_API !== undefined)
        this.Api.baseUrl = import.meta.env.VITE_API;
      if (getRefreshToken() == null) {
        window.location.href = "/login";
      } else {
        this.renewToken().then((res) => {
          this.fetchUserInfo().then(() => {
            window.location.href = "/home";
          });
        });
      }
    },
    login({ email, password }: { email: string; password: string }) {
      return this.Api.auth
        .loginCreate({ email, password })
        .then(({ data }) => {
          if (data.renew_token === null && data.access_token === null) {
            window.location.href = "/login";
          }
          storeRefreshToken(data.renew_token || "");
          storeAccessToken(data.access_token || "");
          return this.fetchUserInfo().then(() => {
            window.location.href = "/home";
          });
        })
        .catch((err) => {
          console.log(err);
          // TODO : Show error to login screen
        });
    },
    logout() {
      var token = getRefreshToken();
      if (token != null)
        return this.Api.auth
          .logoutCreate({
            renew_token: token,
          })
          .then((res) => {
            removeRefreshToken();
            removeAccessToken();
            window.location.href = "/";
            window.location.reload();
          })
          .catch((err) => {
            removeRefreshToken();
            removeAccessToken();
            this.user = {
              id: 0,
              pseudo: "",
              email: "",
              name: "",
              surname: "",
              picture: "",
              role: "",
              loaded: false,
            };
            console.log(err);
            return { status: "failure", err };
          });
      else {
        return new Promise((resolve, reject) => {
          window.location.href = "/login";
          window.location.reload();
          resolve("Redirect to /login");
        });
      }
    },
    register(registrerBody: RouteRegisterBody) {
      return this.Api.auth
        .registerCreate(registrerBody)
        .then((res) => {
          window.location.href = "/login";
        })
        .catch((err) => {
          // TODO : If error show error to register screen
        });
    },
    renewToken() {
      var token = getRefreshToken();
      if (token != null)
        return this.Api.auth
          .renewCreate({
            renew_token: token,
          })
          .then((res) => {
            if (res.data.access_token !== undefined)
              storeAccessToken(res.data.access_token);
          })
          .catch((err) => {
            removeRefreshToken();
            removeAccessToken();
            // TODO : Notify user that login failed but device is logged out
          });
      else {
        return new Promise((resolve, reject) => {
          removeRefreshToken();
          removeAccessToken();
          resolve("Redirect to /login");
        });
      }
    },
    fetchUserInfo() {
      return this.Api.user
        .userList({
          headers: {
            Authorization: `Bearer ${getRefreshToken()}`,
          },
        })
        .then(({ data }) => {
          if (data.id !== undefined) this.user.id = data.id;
          if (data.email !== undefined) this.user.email = data.email;
          if (data.pseudo !== undefined) this.user.pseudo = data.pseudo;
          if (data.name !== undefined) this.user.name = data.name;
          if (data.surname !== undefined) this.user.surname = data.surname;
          if (data.picture !== undefined) this.user.picture = data.picture;
          if (data.role !== undefined) this.user.role = data.role;
          this.user.loaded = true;
        });
    },
  },
});
