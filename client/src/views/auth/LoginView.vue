<script setup lang="ts">
import { ref } from "vue";
import { useGlobalStore } from "@/stores/Global";
import { validateMail, validatePassword } from "@/helper/validateString";

const Global = useGlobalStore();

const email = ref("");
const password = ref("");
const mailError = ref("");
const passwordError = ref("");

const login = () => {
  var error = false;
  if (validateMail(email.value)) {
    mailError.value = "";
  } else {
    mailError.value = "Mail invalid";
    error = true;
  }

  if (validatePassword(password.value)) {
    passwordError.value = "";
  } else {
    passwordError.value = "Password invalid";
    error = true;
  }
  if (error) return;
  Global.login({ email: email.value, password: password.value });
};
</script>

<template>
  <main class="login-container">
    <div class="login-container2">
      <div class="login-title-container">
        <h1>Login</h1>
      </div>
      <div class="form-auth-input-container">
        <label for="email"> Email : </label>
        <input id="email" type="email" placeholder="Email" v-model="email" />
        <span class="error-text" v-if="mailError !== ''">{{ mailError }}</span>
      </div>
      <div class="form-auth-input-container">
        <label for="password"> Password : </label>
        <input
          id="password"
          type="password"
          placeholder="Password"
          v-model="password"
        />
        <span class="error-text" v-if="passwordError !== ''">{{
          passwordError
        }}</span>
      </div>
      <div class="form-auth-input-container">
        <button @click="login()">Login</button>
      </div>
      <div class="form-auth-link-container">
        <p>Don't have an account ?</p>
        <router-link to="/register">Register</router-link>
      </div>
    </div>
  </main>
</template>
