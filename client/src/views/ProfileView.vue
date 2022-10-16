<script setup lang="ts">
import { getAssetsUrl } from "@/helper/assetsUrl";
import { useGlobalStore } from "@/stores/Global";

const Global = useGlobalStore();
</script>

<template>
  <div class="profile-container">
    <div>
      <img
        v-if="Global.user.picture !== ''"
        class="icon"
        :src="getAssetsUrl(Global.user.picture)"
      />
    </div>
    <div class="role-pile">Role : {{ Global.user.role }}</div>
    <div>
      <div>Email : {{ Global.user.email }}</div>
      <div>Pseudo : {{ Global.user.pseudo }}</div>
      <div>Name : {{ Global.user.name }}</div>
      <div>Surname : {{ Global.user.surname }}</div>
    </div>
  </div>
</template>
