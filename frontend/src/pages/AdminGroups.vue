<template>
  <DataZone>
    <template v-slot:dialog-content>
      <DialogGroup />
    </template>
  </DataZone>
</template>

<script setup lang="ts">
import DataZone from '@/components/DataZone.vue';
import { GenericGroupsAPI } from '@/lib/apis';
import DialogGroup from '@/components/DialogGroup.vue';
import { useDataStore } from '@/stores/data';
import { onBeforeMount } from 'vue';

onBeforeMount(() => useDataStore().init('groups', GenericGroupsAPI))
</script>
