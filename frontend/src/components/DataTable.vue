<template>
  <v-data-iterator
    v-model:items-per-page="itemsPerPage"
    v-model:page="page"
    :items="items"
    :search="filter"
    :sort-by="[]"
    class="fill-height"
  >
    <template v-slot:no-data>
      <v-alert class="ma-2" type="warning">No results</v-alert>
    </template>
    <template v-slot:header>
      <v-toolbar
        dark
        density="compact"
        color="blue-darken-3"
        class="px-2 mb-2"
      >
        <v-text-field
          v-model="filter"
          clearable
          hide-details
          prepend-inner-icon="mdi-magnify"
          placeholder="Search"
          variant="solo"
          density="compact"
        ></v-text-field>
        <v-spacer></v-spacer>
        <!-- <v-select
          v-model="sortKey"
          hide-details
          :items="keys"
          :item-value="item => item.toLowerCase()"
          prepend-inner-icon="mdi-sort"
          label="Sort by"
          density="comfortable"
        ></v-select> -->
        <v-spacer></v-spacer>
        <v-btn @click="dialog = true">Create</v-btn>
        <v-btn icon flat @click="globals.view = 'list'" :disabled="globals.view == 'list'">
          <v-icon>mdi-view-list</v-icon>
        </v-btn>
        <v-btn icon flat @click="globals.view = 'grid'" :disabled="globals.view == 'grid'">
          <v-icon>mdi-view-grid</v-icon>
        </v-btn>
        <v-btn icon flat @click="globals.view = 'grid-compact'" :disabled="globals.view == 'grid-compact'">
          <v-icon>mdi-view-grid-compact</v-icon>
        </v-btn>
      </v-toolbar>
    </template>
    <template v-slot:default="{ }">
      <v-table
          fixed-header
          height="calc(100vh - 64px - 50px - 54px)"
          style="--v-table-header-height: 42px; --v-table-row-height: 42px;"
        >
        <thead>
          <tr>
            <th class="text-left" v-for="h in headers" :key="h.key">
              {{ h.title }}
            </th>
            <th key="controls" class="text-right">
              Controls
            </th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in items" :key="item.id">
            <td v-for="h in headers" :key="h.key">{{ item[h.key] }}</td>
            <td key="controls">
              <v-btn size="medium" icon="mdi-edit" @click="update(item.id, item)"></v-btn>
              <v-btn size="medium" icon="mdi-delete" @click="openRemove(item.id)"></v-btn>
            </td>
          </tr>
        </tbody>
      </v-table>
    </template>
    <template v-slot:footer>
      <v-toolbar density="compact" plain>
        <!-- class="d-flex align-center justify-space-around pa-4"-->
        <v-spacer></v-spacer>
        <span class="grey--text">Items per page</span>
        <v-menu>
          <template v-slot:activator="{ props }">
            <v-btn
              variant="text"
              color="primary"
              class="ml-2"
              append-icon="mdi-chevron-down"
              v-bind="props"
            >
              {{ itemsPerPage }}
            </v-btn>
          </template>
          <v-list>
            <v-list-item
              v-for="(number, index) in itemsPerPageArray"
              :key="index"
              :title="number"
              @click="itemsPerPage = number"
            ></v-list-item>
          </v-list>
        </v-menu>
        <span class="mr-4 grey--text">Page {{ page }} of {{ numberOfPages }}</span>
        <v-btn icon size="small" flat @click="page = 1">
          <v-icon>mdi-page-first</v-icon>
        </v-btn>
        <v-btn icon size="small" flat @click="prevPage">
          <v-icon>mdi-chevron-left</v-icon>
        </v-btn>
        <v-btn icon size="small" flat @click="nextPage">
          <v-icon>mdi-chevron-right</v-icon>
        </v-btn>
        <v-btn icon size="small" flat @click="page = numberOfPages">
          <v-icon>mdi-page-last</v-icon>
        </v-btn>
      </v-toolbar>
    </template>
  </v-data-iterator>
  <v-dialog :model-value="dialog" width="70vw">
    <v-card>
      <v-toolbar>
        <v-toolbar-title dark>
          <span class="text-h5">Add Device</span>
        </v-toolbar-title>
      </v-toolbar>
      <v-card-text>
        <slot name="dialog-content" :onFinish="() => dialog = false" />
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="blue-darken-1" variant="text" @click="dialog = false">
          Close
        </v-btn>
        <v-btn :loading="loading" @click="onSave" color="blue-darken-1" variant="text" >
          Save
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup lang="ts" generic="T extends WithID">
import type { ListResult, PaginationArgs } from '@/lib/apis';
import type { WithID } from '@/lib/template';
import type { CardItem, ReadonlyDataTableHeader } from '@/lib/utils';
import { useGlobalsStore } from '@/stores/globals';
import { useNotificationsStore } from '@/stores/notifications';
import { defineProps, ref, computed, type PropType, onMounted, watch, provide } from 'vue';

import { VDataIterator } from 'vuetify/labs/VDataIterator'

const globals = useGlobalsStore();
const loading = ref(true);
const itemsPerPage = ref(50);
const page = ref(1);
const itemsPerPageArray = [10, 20, 50, 100];
const sortKey = ref('id');
const filter = ref('');
const dialog = ref(false);
const createMethod = ref<() => {}>();

const notificationsStore = useNotificationsStore();

const onSave = async () => {
  if (createMethod.value) {
    createMethod.value();
    notificationsStore.showSuccess("Saved.")
  }
};

provide('registerCreateMethod', (method: () => Promise<void> ) => {
  createMethod.value = method;
});

watch(page, loadItems)
watch(itemsPerPage, loadItems)
watch(filter, loadItems)

function nextPage () {
  if (page.value + 1 <= numberOfPages.value) page.value += 1
}
function prevPage () {
  if (page.value - 1 >= 1) page.value -= 1
}
const numberOfPages = computed(() => {
  return Math.ceil(items.value.length / itemsPerPage.value)
})

const props = defineProps(
  {
    headers: {
      type: Array as PropType<ReadonlyDataTableHeader[]>,
      required: true,
    },
    remove: {
      type: Function as PropType<(id: string) => Promise<void>>,
      required: true,
    },
    update: {
      type: Function as PropType<(id: string, payload: { [key:string]: any }) => Promise<T>>,
      required: true,
    },
    fetch: {
      type: Function as PropType<({ page, itemsPerPage, sortBy }: PaginationArgs) => Promise<ListResult<T>>>,
      required: true,
    },
    convertToCard: {
      type: Function as PropType<(item: T) => CardItem>,
      required: true,
    }
  }
)

function openRemove(id: string) {
  globals.openConfirm({
    title: 'Confirm?',
    callback: () => {
      props.remove(id)
      globals.closeConfirm()
      notificationsStore.showSuccess("Removed.")
    },
  })
};

const gridItems = computed(() => {
  const newItems = [] as CardItem[]
  items.value.forEach((item) => newItems.push(props.convertToCard(item as T)))
  return newItems
})

const items = ref([] as T[])
const totalItems = ref(0)

async function loadItems() {
  const resp = await props.fetch({
    page: page.value,
    itemsPerPage: itemsPerPage.value,
    sortBy: sortKey.value,
  })

  items.value = ref<T[]>(resp.items).value
  totalItems.value = resp.total
  loading.value = false
}

onMounted(loadItems)
</script>
