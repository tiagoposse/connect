<template>
  <DataZone>
    <template v-slot:dialog-content>
      <DialogUser />
    </template>
  </DataZone>
</template>

<script setup lang="ts">
import DataZone from '@/components/DataZone.vue';
import { GenericUsersAPI } from '@/lib/apis';
import DialogUser from '@/components/DialogUser.vue';
import { useDataStore } from '@/stores/data';
import { onBeforeMount } from 'vue';

onBeforeMount(() => useDataStore().init('adminUsers', GenericUsersAPI))
</script>
