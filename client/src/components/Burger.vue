<script setup lang="ts">
import { useGlobalStore } from "@/stores/Global";

const Global = useGlobalStore();
// Make a two way burger workflow
// If on small screen, show burger menu in the middle
// If on large screen, show burger menu on the left

// LARGE on active:
// container :top-0 right-0 flex

// SMALL on active:
// container : rounded-b-lg flex justify-center mx-2 (smaller thatn nav)

// test
</script>

<template>
  <div class="NavBurger active">
    <div>
      <p>Welcome Maxime Leriche</p>
    </div>
    <div>
      <div><p>Home</p></div>
      <div><p>Mon compte</p></div>
      <div><p>Mes channel</p></div>
      <div><p>Mes channel</p></div>
    </div>
  </div>
</template>
