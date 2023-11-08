<template>
  <DataZone>
    <template v-slot:dialog-content>
      <DialogDevice />
    </template>
  </DataZone>
</template>

<script setup lang="ts">
import DataZone from '@/components/DataZone.vue';
import { GenericDevicesAPI } from '@/lib/apis';
import DialogDevice from '@/components/DialogDevice.vue';
import { useDataStore } from '@/stores/data';
import { onBeforeMount } from 'vue';

onBeforeMount(() => useDataStore().init('users', GenericDevicesAPI))
</script>
