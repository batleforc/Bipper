<script setup lang="ts">
import { getAssetsUrl } from "@/helper/assetsUrl";
import { useGlobalStore } from "@/stores/Global";

const Global = useGlobalStore();
</script>

<template>
  <div>
    <div>
      <img
        v-if="Global.user.picture !== ''"
        class="icon"
        :src="getAssetsUrl(Global.user.picture)"
      />
    </div>
    <div>Role</div>
    <div>
      <div>Email</div>
      <div>Username</div>
      <div>Name</div>
      <div>Surname</div>
    </div>
  </div>
</template>
