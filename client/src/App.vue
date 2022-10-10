<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import { useGlobalStore } from "./stores/Global";

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
      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/about">About</RouterLink>
      </nav>
    </div>
  </header>
  <h1>Hello truc</h1>
  <RouterView />
</template>

<style>
@import "@/assets/index.scss";
</style>
