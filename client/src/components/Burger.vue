<script setup lang="ts">
import { getAssetsUrl } from "@/helper/assetsUrl";
import { useGlobalStore } from "@/stores/Global";
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
        <p v-if="Global.loggedIn">
          {{ Global.user.name }} {{ Global.user.surname }}
        </p>
      </div>
      <div class="burger-cat">
        <div><p>Home</p></div>
        <div><p>Mon compte</p></div>
        <div><p>Mes channel</p></div>
        <div><p>Mes channel</p></div>
      </div>
      <div class="burger-cat">
        <div><p>Logout</p></div>
        <div class="burger-credit"><p>Made by batleforc</p></div>
      </div>
    </div>
  </div>
</template>
