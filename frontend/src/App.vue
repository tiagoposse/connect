<template>
  <v-app app>
    <v-app-bar v-if="authStore.isAuthenticated">
      <v-toolbar>
        <v-toolbar>
          <v-spacer></v-spacer>
          <v-toolbar-title>{{ useRoute().name }}</v-toolbar-title>
        </v-toolbar>
      </v-toolbar>
    </v-app-bar>

    <Navigation v-if="authStore.isAuthenticated" />
    <v-main app class="app mt-0 pr-0">
      <v-container fluid class="pt-0 pl-0 ma-0">
        <v-row align="stretch">
          <v-col cols="12">
            <router-view></router-view>
          </v-col>
        </v-row>
      </v-container>
    </v-main>

    <v-footer>
      <div>Copyright 2023</div>
    </v-footer>
  </v-app>
</template>

<script setup lang="ts">
import Navigation from './components/AppNavigation.vue'
import { AuthAPI } from './lib/apis';
import { useAuthStore } from './stores/auth';
import { onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router'

const authStore = useAuthStore();
const router = useRouter();

onMounted(async () => {
  try {
    const user = await AuthAPI.status();
    authStore.setUser(user)
  } catch {
    authStore.logout()
    router.push('/login')
  }
})

</script>

<style>
html,
body {
  height: 100vh;
  width: 100vw;
  margin: 0;
}


#app {
  max-width: 100vw !important;
  margin-left: 6px !important;
  padding: 0px !important;
}

.app {
  max-width: 100vw !important;
  min-width: 100vw !important;
}
</style>
