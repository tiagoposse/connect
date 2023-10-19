<!-- src/components/AdminPage.vue -->
<template>
  <v-form @submit="login">
    <v-card>
      <v-card-text>
        <v-text-field label="Username" variant="outlined" v-model="loginData.username" />
        <v-text-field label="Password" variant="outlined" v-model="loginData.password" />
      </v-card-text>
      <v-card-actions>
        <v-btn type="submit">Submit</v-btn>
      </v-card-actions>
    </v-card>
  </v-form>
  <v-btn color="blue" @click="googleRedirect">Login with Google</v-btn>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { DefaultApiConfig, AuthAPI } from '../lib/apis';

const loginData = ref({
  username: '',
  password: ''
})

const login = async () => {
  AuthAPI.userpassLogin({
    userpassLoginRequest: {
      username: loginData.value.username,
      password: loginData.value.password,
    }

  })
}

const googleRedirect = () => {
  const current = `${window.location.protocol}//${window.location.host}`
  console.log(`navigate to: ${DefaultApiConfig.basePath}/auth/google/start?after=${encodeURIComponent(current)}`)
  window.location.href = `${DefaultApiConfig.basePath}/auth/google/start?after=${encodeURIComponent(current)}`
}
</script>
