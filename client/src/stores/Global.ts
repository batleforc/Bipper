import { Api, type RouteRegisterBody } from "@/api/Api";
import { getQuerryParam, hasQuerryParam } from "@/helper/query";
import {
  getAccessToken,
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
      this.router.beforeEach((to, from, next) => {
        if (to.matched.some((record) => record.meta.requiresAuth)) {
          // this route requires auth, check if logged in
          // if not, redirect to login page.
          if (!this.loggedIn) {
            //this.router.push({ name: "login" });
            next({ name: "login", query: { redirect: to.fullPath } });
          } else {
            next(); // go to wherever I'm going
          }
        } else if (to.matched.some((record) => record.meta.hiddenIfLoggedIn)) {
          if (this.loggedIn) {
            next({ name: "home" });
          } else {
            next();
          }
        } else {
          next(); // does not require auth, make sure to always call next()!
        }
      });
      if (import.meta.env.VITE_API !== undefined)
        this.Api.baseUrl = import.meta.env.VITE_API;
      if (getRefreshToken() != null) {
        this.renewToken().then((res) => {
          if (res) {
            console.log(res);
          }
          this.fetchUserInfo().then(() => {
            this.router.push({ name: "home" });
          });
        });
      }
    },
    login({ email, password }: { email: string; password: string }) {
      return this.Api.auth
        .loginCreate({ email, password })
        .then((res) => res.json())
        .then((data) => {
          if (data.renew_token === null && data.access_token === null) {
            this.router.push({ name: "login" });
          }
          storeRefreshToken(data.renew_token || "");
          storeAccessToken(data.access_token || "");
          return this.fetchUserInfo().then(() => {
            console.log("redirecting to home or previous path");
            this.loggedIn = true;
            if (hasQuerryParam("redirect")) {
              this.router.push({ path: getQuerryParam("redirect") || "" });
            } else {
              this.router.push({ name: "home" });
            }
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
          this.router.push({ name: "login", query: { registered: "true" } });
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
          .then((res) => res.json())
          .then((data) => {
            if (data.access_token !== undefined) {
              this.loggedIn = true;
              storeAccessToken(data.access_token);
            }
            return { status: "success" };
          })
          .catch((err) => {
            removeRefreshToken();
            removeAccessToken();
            return { status: "failure", err };
            // TODO : Notify user that login failed but device is logged out
          });
      else {
        return new Promise((resolve, reject) => {
          removeRefreshToken();
          removeAccessToken();
          this.router.push({ name: "login" });
          resolve("Redirect to /login");
        });
      }
    },
    fetchUserInfo() {
      return this.Api.user
        .userList({
          headers: {
            Authorization: `Bearer ${getAccessToken()}`,
          },
        })
        .then((res) => res.json())
        .then((data) => {
          if (data.id !== undefined) this.user.id = data.id;
          if (data.email !== undefined) this.user.email = data.email;
          if (data.pseudo !== undefined) this.user.pseudo = data.pseudo;
          if (data.name !== undefined) this.user.name = data.name;
          if (data.surname !== undefined) this.user.surname = data.surname;
          if (data.picture !== undefined) this.user.picture = data.picture;
          if (data.role !== undefined) this.user.role = data.role;
          if (data.Channels !== undefined) this.chan.chan = data.Channels;
          if (data.MyChannels !== undefined) this.chan.myChan = data.MyChannels;
          this.user.loaded = true;
        });
    },
  },
});
