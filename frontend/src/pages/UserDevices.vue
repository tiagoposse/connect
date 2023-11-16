<template>
  <DevicesView :same-user="true"/>
</template>

<script setup lang="ts">
import { GenericDevicesAPI } from '@/lib/apis';
import { useDataStore } from '@/stores/data';
import { onBeforeMount } from 'vue';
import DevicesView from '@/components/DevicesView.vue';

onBeforeMount(() => useDataStore().init('users', GenericDevicesAPI))
</script>
