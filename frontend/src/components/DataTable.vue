<template>
  <v-data-table-server v-model:items-per-page="itemsPerPage" :headers="props.headers" :items="items"
    :items-length="totalItems" :loading="loading" item-value="name" class="elevation-1" @update:options="loadItems">
    <template v-slot:item.controls="tableHeaders">
      <v-btn size="medium" icon="mdi-edit" @click="edit(index)"></v-btn>
      <v-btn size="medium" icon="mdi-delete" @click="remove(item.id)"></v-btn>
    </template>
  </v-data-table-server>
</template>

<script setup lang="ts" generic="T">
import { defineProps, ref, computed, type PropType } from 'vue';
import {
  VDataTableServer
} from "vuetify/labs/VDataTable";
import type { ApiResponse } from '@/lib/api';
import type { Store, StoreDefinition } from 'pinia';
const loading = ref(true);
const itemsPerPage = ref(20);

interface Api {
  fetch: (req: any) => Promise<ApiResponse<T>>,
  add: (req: any) => Promise<T>
  remove: (id: number) => Promise<void>
}

const props = defineProps(
  {
    headers: {
      type: Object,
      required: true,
    },
    api: {
      type: Object as Api,
      required: true,
    },
    data: {
      type: Object as PropType<StoreDefinition>,
      required: true
    }
  }
)

const tableHeaders = computed(() => {
  const newHeaders = props.headers;
  newHeaders.push(
    {
      title: 'Controls',
      align: 'end',
      value: 'controls',
      key: 'controls'
    })


  return newHeaders;
})

const items = ref([] as T[])
const totalItems = ref(50)

const loadItems = async ({ page, itemsPerPage, sortBy }) => {
  const resp = await props.api.fetch({
    page,
    itemsPerPage,
    sort: sortBy,
  })

  items.value = await resp.value()
  totalItems.value = resp.raw.headers.get("X-Total")
  loading.value = false
}

const remove = async (id: number) => {
  await props.api.remove(id)
  props.store.remove(id)
}

</script>
