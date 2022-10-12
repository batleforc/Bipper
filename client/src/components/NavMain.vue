<script setup lang="ts">
import Burger from "./icon/burger.vue";
import Bell from "./icon/bell.vue";

// Add in profile-container user picture if authentified
</script>
<template>
  <nav class="nav">
    <div class="burger-container">
      <Burger />
    </div>
    <div>
      <p>Bipper</p>
    </div>
    <div class="profile-container">
      <Bell />
    </div>
  </nav>
</template>
