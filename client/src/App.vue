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
  <header>
    <div class="wrapper">
      <NavMain />
      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/app">app</RouterLink>
        <RouterLink to="/login">login</RouterLink>
        <RouterLink to="/register">register</RouterLink>
      </nav>
    </div>
  </header>
  <h1>Hello truc</h1>
  <RouterView />
</template>

<style>
@import "@/assets/index.scss";
</style>
