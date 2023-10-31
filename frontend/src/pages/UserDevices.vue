<template>
  <DataTable
      :headers="headers"
      :fetch="fetch"
      :update="update"
      :remove="remove"
      :convert-to-card="convertToCard"
  >
  <template v-slot:dialog-content="{ onFinish }">
    <DialogDevice :onFinish="onFinish" />
  </template>
  </DataTable>
</template>

<script setup lang="ts">
import DataTable from '@/components/DataTable.vue';
import type { DeviceList, UpdateDeviceRequest } from '@/lib/api';
import { DevicesAPI, type PaginationArgs, type ListResult } from '@/lib/apis';
import { type CardItem, type ReadonlyDataTableHeader } from '@/lib/utils';
import DialogDevice from '@/components/DialogDevice.vue';

async function fetch(pageArgs: PaginationArgs): Promise<ListResult<DeviceList>> {
  const resp = await DevicesAPI.listDeviceRaw(pageArgs)
  return {
    items: await resp.value(),
    total: parseInt(resp.raw.headers.get("x-total")!)
  }
}

function convertToCard(item: DeviceList): CardItem {
  return {
    id: item.id,
    title: item.id,
    subtitle: '',
    fields: [],
  }
}

async function remove(id: string) {
  await DevicesAPI.deleteDevice({ id })
}

async function update(id: string, payload: UpdateDeviceRequest): Promise<DeviceList> {
  return await DevicesAPI.updateDevice({
    id,
    updateDeviceRequest: payload
  })
}

const headers = [
  {
    title: 'ID',
    align: 'start',
    sortable: true,
    key: 'id',
  },
  {
    title: 'Name',
    align: 'end',
    sortable: true,
    key: 'title',
  },
  {
    title: 'Type',
    align: 'end',
    sortable: false,
    key: 'type',
  }
] as ReadonlyDataTableHeader[]
</script>
