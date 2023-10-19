<!-- src/components/AdminPage.vue -->
<template>
  <v-container class="pa-0 ma-0">
    <v-row no-gutters v-for="devices, user in userDevices" :key="user">
      <v-col>
        <h2>{{ user }}</h2>
        <v-divider />
        <v-container>
          <DeviceList :devices="devices" />
        </v-container>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
import DeviceList from './DeviceList.vue';
import { Device, useDevicesStore } from '../stores/devices'
import { computed } from 'vue';

const store = useDevicesStore()

const userDevices = computed(() => {
  const m: { [key: string]: Device[] } = {}
  for (const d of store.devices) {
    if (!m.hasOwnProperty(d.user)) {
      m[d.user] = []
    }
    m[d.user].push(d)
  }
  return m
})

</script>

<style scoped>
.device-grid {
  display: flex;
  flex-wrap: wrap;
}
</style>
