<template>
  <DevicesView :same-user="true"/>
</template>

<script setup lang="ts">
import { DevicesAPI, GenericDevicesAPI } from '@/lib/apis';
import { useDataStore } from '@/stores/data';
import { onBeforeMount } from 'vue';
import DevicesView from '@/components/DevicesView.vue';
import type { FilterArgs, GenericAPI, PaginationArgs } from '@/lib/utils';
import type { CreateDeviceRequest } from '@/lib/api';
import { useAuthStore } from '@/stores/auth';

const authStore = useAuthStore();
const MyDevicesAPI: GenericAPI = {
  create: async (payload: CreateDeviceRequest) => {
    payload.userId = authStore.userID
    return await GenericDevicesAPI.create(payload)
  },
  remove: GenericDevicesAPI.remove,
  update: async (id: string, payload: CreateDeviceRequest) => {
    payload.userId = authStore.userID
    return await DevicesAPI.updateDevice({ id, updateDeviceRequest: payload })
  },
  fetch: async (params: PaginationArgs, filters: FilterArgs) => {
    filters.userId = authStore.userID
    return await DevicesAPI.listDeviceRaw({ ...params, ...filters }) 
  },
  headers: GenericDevicesAPI.headers,
}

onBeforeMount(() => useDataStore().init('users', MyDevicesAPI))
</script>
