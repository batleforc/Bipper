import { Api } from "@/api/Api";
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
    },
  },
});
