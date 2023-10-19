<!-- src/components/UserPage.vue -->
<template>
  <v-row no-gutters>
    <v-col :key="device.id" v-for="device in props.devices">
      <DeviceCard :device="device" />
    </v-col>
  </v-row>
    <v-overlay
        :model-value="loading"
        class="align-center justify-center"
      >
      <v-progress-circular
        color="primary"
        indeterminate
        size="64"
      />
    </v-overlay>
</template>

<script setup lang="ts">
// Import required functions and components
import DeviceCard from './DeviceCard.vue';
import { type Device } from '../api/devices'

const props = defineProps<{
  devices: Device[],
  loading: true,
}>()

</script>
