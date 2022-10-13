<script setup lang="ts">
import { reactive } from "vue";
import { useGlobalStore } from "@/stores/Global";
import {
  validateMail,
  validatePassword,
  validatePseudoNameSurname,
} from "@/helper/validateString";

const Global = useGlobalStore();

const email = reactive({
  value: "",
  error: "",
});
const pseudo = reactive({
  value: "",
  error: "",
});
const name = reactive({
  value: "",
  error: "",
});
const surname = reactive({
  value: "",
  error: "",
});
const password = reactive({
  value: "",
  error: "",
});
const password2 = reactive({
  value: "",
  error: "",
});

const register = () => {
  var error = false;
  if (validateMail(email.value)) {
    email.error = "";
  } else {
    email.error = "Mail invalid";
    error = true;
  }
  if (validatePassword(password.value)) {
    password.error = "";
  } else {
    password.error = "Password invalid";
    error = true;
  }
  if (validatePassword(password2.value)) {
    password2.error = "";
  } else {
    password2.error = "Password invalid";
    error = true;
  }
  if (validatePseudoNameSurname(pseudo.value)) {
    pseudo.error = "";
  } else {
    pseudo.error = "Pseudo invalid";
    error = true;
  }
  if (validatePseudoNameSurname(name.value)) {
    name.error = "";
  } else {
    name.error = "Name invalid";
    error = true;
  }
  if (validatePseudoNameSurname(surname.value)) {
    surname.error = "";
  } else {
    surname.error = "Surname invalid";
    error = true;
  }
  if (error) return;
  if (password.value === password2.value) {
    password2.error = "";
  } else {
    password2.error = "Password not match";
    return;
  }

  Global.register({
    email: email.value,
    pseudo: pseudo.value,
    name: name.value,
    surname: surname.value,
    password: password.value,
  });
};
</script>

<template>
  <main class="register-container">
    <div class="register-container2">
      <div class="login-title-container">
        <h1>Register</h1>
      </div>
      <div class="form-auth-input-container">
        <label for="email"> Email : </label>
        <input
          id="email"
          type="email"
          placeholder="Email"
          v-model="email.value"
        />
        <span class="error-text" v-if="email.error !== ''">{{
          email.error
        }}</span>
      </div>
      <div class="form-auth-input-container">
        <label for="pseudo"> Pseudo : </label>
        <input
          id="pseudo"
          type="text"
          placeholder="Pseudo"
          v-model="pseudo.value"
        />
        <span class="error-text" v-if="pseudo.error !== ''">{{
          pseudo.error
        }}</span>
      </div>
      <div class="form-auth-input-container">
        <label for="name"> Name : </label>
        <input id="name" type="text" placeholder="Name" v-model="name.value" />
        <span class="error-text" v-if="name.error !== ''">{{
          name.error
        }}</span>
      </div>
      <div class="form-auth-input-container">
        <label for="Surname"> Surname : </label>
        <input
          id="Surname"
          type="text"
          placeholder="Surname"
          v-model="surname.value"
        />
        <span class="error-text" v-if="surname.error !== ''">{{
          surname.error
        }}</span>
      </div>
      <div class="form-auth-input-container">
        <label for="password"> Password : </label>
        <input
          id="password"
          type="password"
          placeholder="Password"
          v-model="password.value"
        />
        <span class="error-text" v-if="password.error !== ''">{{
          password.error
        }}</span>
      </div>
      <div class="form-auth-input-container">
        <label for="password2"> Retype Password : </label>
        <input
          id="password2"
          type="password"
          placeholder="Retype Password"
          v-model="password2.value"
        />
        <span class="error-text" v-if="password2.error !== ''">{{
          password2.error
        }}</span>
      </div>
      <div class="form-auth-input-container">
        <button @click="register()">Register</button>
      </div>
      <div class="form-auth-link-container">
        <p>You have an account ?</p>
        <router-link to="/login">Login</router-link>
      </div>
    </div>
  </main>
</template>
