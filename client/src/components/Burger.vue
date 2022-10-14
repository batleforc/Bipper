<script setup lang="ts">
import { getAssetsUrl } from "@/helper/assetsUrl";
import { BurgerLink } from "@/model/BurgerLink";
import { useGlobalStore } from "@/stores/Global";
import Bell from "./icon/bell.vue";
const Global = useGlobalStore();

const props = defineProps<{
  doClose: () => void;
  visible: boolean;
}>();

const exit = () => {
  props.doClose();
};

window.addEventListener("keydown", (e) => {
  if (props.visible && e.key === "Escape") {
    exit();
  }
});
window.addEventListener("click", (e) => {
  if (
    props.visible &&
    e.target !== null &&
    !document.getElementById("nav")?.contains(e.target as Node) &&
    !document.getElementById("burger")?.contains(e.target as Node)
  ) {
    exit();
  }
});
</script>

<template>
  <div id="burger" :class="visible ? 'active' : ''" class="NavBurger">
    <div class="burger-content">
      <div class="burger-cat burger-title">
        <img
          v-if="Global.user.picture !== ''"
          class="icon"
          :src="getAssetsUrl(Global.user.picture)"
        />
        <Bell v-else />
        <p v-if="Global.loggedIn">
          {{ Global.user.name }} {{ Global.user.surname }}
        </p>
      </div>
      <div class="burger-cat">
        <RouterLink
          v-if="Global.loggedIn"
          v-for="link in BurgerLink.LoggedIn"
          :to="link.link"
          @click="exit"
        >
          {{ link.name }}
        </RouterLink>
        <RouterLink
          v-else
          v-for="link in BurgerLink.LoggedOut"
          :to="link.link"
          @click="exit"
        >
          {{ link.name }}
        </RouterLink>
      </div>
      <div class="burger-cat">
        <div v-if="Global.loggedIn" @click="Global.logout()"><p>Logout</p></div>
        <div v-if="!Global.loggedIn">
          <RouterLink @click="exit" to="/login">Login</RouterLink>
        </div>
        <div v-if="!Global.loggedIn">
          <RouterLink @click="exit" to="/register">Register</RouterLink>
        </div>
        <div class="burger-credit"><p>Made by batleforc</p></div>
      </div>
    </div>
  </div>
</template>
