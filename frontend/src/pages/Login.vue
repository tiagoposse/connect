<template>
  <v-sheet class="d-flex flex-wrap fill-height align-center justify-center">
    <v-sheet class="ma-2 pa-2" :width="getWidth()" rounded>
      <v-form @submit="login">
        <v-card>
          <v-card-text>
            <v-text-field label="Username" variant="outlined" v-model="loginData.username" />
            <v-text-field label="Password" variant="outlined" v-model="loginData.password" />
          </v-card-text>
          <v-card-actions>
            <v-btn type="submit">Submit</v-btn>
            <v-btn color="blue" @click="googleRedirect">Login with Google</v-btn>
          </v-card-actions>
        </v-card>
      </v-form>
    </v-sheet>
  </v-sheet>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { AuthAPI } from '@/lib/apis';
import { DefaultApiConfig } from '@/lib/utils';
import { useDisplay } from 'vuetify'

const { mobile } = useDisplay()

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

function getWidth() {
  if (mobile.value) {
    return "100vw"
  } else {
    return "50vw"
  }
}
</script>
