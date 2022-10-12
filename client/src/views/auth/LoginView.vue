<script setup lang="ts"></script>

<template>
  <main class="login-container">
    <div class="login-container2">
      <div class="login-title-container">
        <h1>Login</h1>
      </div>
      <div class="login-form-input-container">
        <label for="email"> Email : </label>
        <input id="email" type="email" placeholder="Email" />
      </div>
      <div class="login-form-input-container">
        <label for="password"> Password : </label>
        <input id="password" type="password" placeholder="Password" />
      </div>
      <div class="login-form-input-container">
        <button>Login</button>
      </div>
      <div class="login-register-container">
        <p>Don't have an account ?</p>
        <router-link to="/register">Register</router-link>
      </div>
    </div>
  </main>
</template>
