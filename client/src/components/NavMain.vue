<script setup lang="ts">
import Burger from "./icon/burger.vue";
import Bell from "./icon/bell.vue";
import { useGlobalStore } from "@/stores/Global";
import { getAssetsUrl } from "@/helper/assetsUrl";
import MenuBurger from "./Burger.vue";
import { ref } from "vue";
const Global = useGlobalStore();

const seeBurger = ref(false);
</script>
<template>
  <nav id="nav" class="nav">
    <div class="burger-container" @click="seeBurger = !seeBurger">
      <Burger />
    </div>
    <div>
      <RouterLink v-if="Global.loggedIn" to="/app">Bipper</RouterLink>
      <RouterLink v-else to="/">Bipper</RouterLink>
    </div>
    <div class="profile-container">
      <Bell v-if="Global.user.picture === '' || seeBurger" />
      <img class="icon" v-else :src="getAssetsUrl(Global.user.picture)" />
    </div>
  </nav>
  <nav v-if="Global.noInternetMessage" class="nav2">
    <h1>No Internet connection</h1>
  </nav>
  <MenuBurger
    :visible="seeBurger"
    :do-close="
      () => {
        seeBurger = !seeBurger;
      }
    "
  />
</template>
