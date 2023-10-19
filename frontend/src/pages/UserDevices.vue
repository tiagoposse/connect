<template>
  <v-toolbar density="compact" class="pa-0">
    <v-text-field hide-details prepend-icon="mdi-magnify" single-line v-model="filter" />
    <v-spacer />
    <v-btn v-if="authStore.isAuthenticated" color="primary" @click.stop="dialog = true">
      Create Device
    </v-btn>
  </v-toolbar>
  <v-container fluid>
    <DeviceList :devices="devicesStore.devices" :loading="loading" />
  </v-container>
  <DeviceCreateDialog :open="dialog" :close="() => dialog = false" />
</template>

<script setup lang="ts">
import DeviceList from '../components/DeviceList.vue';
import { useDevicesStore } from '../stores/devices'
import { useAuthStore } from '../stores/auth'
import { ref, onMounted } from 'vue';
import DeviceCreateDialog from '../components/DeviceCreateDialog.vue';
import { DevicesAPI } from '../lib/apis';

const filter = ref("")
const dialog = ref(false);
const loading = ref(true);

const devicesStore = useDevicesStore()
const authStore = useAuthStore()


onMounted(async () => {
  const resp = await DevicesAPI.listDevice({
    user: authStore.userID
  })
  devicesStore.setDevices(resp)
  loading.value = false;
})



</script>
