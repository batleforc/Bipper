<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import { useGlobalStore } from "./stores/Global";
import NavMain from "./components/NavMain.vue";

const Global = useGlobalStore();
if (!Global.inited) Global.init();
// @ts-ignore
window.stores = {
  Global,
};
</script>

<template>
  <div :class="'main-' + String($route.name)">
    <header>
      <NavMain />
      <nav v-if="false">
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/app">app</RouterLink>
        <RouterLink to="/login">login</RouterLink>
        <RouterLink to="/register">register</RouterLink>
      </nav>
    </header>
    <RouterView />
  </div>
</template>

<style>
@import "@/assets/index.scss";
</style>
