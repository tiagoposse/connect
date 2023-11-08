<template>
  <v-layout ref="app" style="height: 100vh; width: 100vw;">
    <v-app-bar v-if="auth.isAuthenticated">
      <v-toolbar density="compact">
        <v-spacer></v-spacer>
        <v-toolbar-title>{{ useRoute().name }}</v-toolbar-title>
      </v-toolbar>
    </v-app-bar>

    <Navigation v-if="auth.isAuthenticated" />
    <v-main app class="mt-0 pr-0" >
      <router-view></router-view>
    </v-main>

  </v-layout>
  <v-overlay
      :model-value="globals.progress"
      class="align-center justify-center"
    >
    <v-progress-circular
      color="primary"
      indeterminate
      size="64"
    />
  </v-overlay>

  <v-snackbar
      v-for="(notification) in notificationsStore.notifications"
      :key="notification.id"
      v-model="notification.isShow"
      location="bottom end"
      variant="flat"
      :color="notification.color"
      :style="{ bottom: `${notification.position}px` }"
      timeout="2000"
      vertical
    >
    
    <div v-for="(val, index) in notification.messages" :key="index">{{ val }}</div>

    <template v-slot:actions>
      <v-btn color="blue" variant="text" @click="notificationsStore.remove(notification.id)">
        Close
      </v-btn>
    </template>
  </v-snackbar>

  <v-dialog :model-value="confirmStore.isOpen" width="40vw">
    <v-card>
      <v-card-text><span class="text-h5">{{ confirmStore.title }}</span></v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="blue-darken-1" variant="text" @click="confirmStore.close">Cancel</v-btn>
        <v-btn @click="confirmStore.callback" color="blue-darken-1" variant="text" >Confirm</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>

</template>

<script setup lang="ts">
import Navigation from '@/components/AppNavigation.vue'
import { useAuthStore } from '@/stores/auth';
import { onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router'
import { useGlobalsStore } from '@/stores/globals';
import { useNotificationsStore } from '@/stores/notifications';
import { useConfirmDialogStore } from '@/stores/dialogs';

const auth = useAuthStore();
const globals = useGlobalsStore();
const confirmStore = useConfirmDialogStore();
const notificationsStore = useNotificationsStore();
const router = useRouter();

onMounted(async () => {
  try {
    await auth.status()
  } catch {
    auth.logout()
    router.push('/login')
  }
})

</script>

<style>

#app {
  max-width: 100vw !important;
  min-width: 100vw !important;
  padding: 0px !important;
}

.v-application {
  max-width: 100vw !important;
  min-width: 100vw !important;
}
</style>
