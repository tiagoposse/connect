<!-- src/components/UserPage.vue -->
<template>
  <v-row no-gutters>
    <v-col :key="device.id" v-for="device in props.devices">
      <DeviceCard :device="device" />
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
import type { DeviceList } from '@/lib/api';
import DeviceCard from '@/components/DeviceCard.vue';

const props = defineProps<{
  devices: DeviceList[],
  loading: true,
}>()

</script>
