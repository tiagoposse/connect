<template>
  <v-data-iterator
    :items-per-page="dataStore.itemsPerPage"
    :page="dataStore.page"
    :items="dataStore.items as T[]"
    :sort-by="[]"
    class="fill-height"
  >
    <template v-slot:no-data>
      <v-alert class="ma-2" type="warning">No results</v-alert>
    </template>
    <template v-slot:header>
      <DataToolbar />
    </template>
    <template v-slot:default="{ }">
      <DataView />
    </template>
    <template v-slot:footer>
      <DataPagination />
    </template>
  </v-data-iterator>
  <v-dialog :model-value="dialogStore.isOpen" width="70vw">
    <v-card>
      <v-toolbar>
        <v-toolbar-title dark>
          <span class="text-h5">{{ dialogStore.title }}</span>
        </v-toolbar-title>
      </v-toolbar>
      <v-card-text>
        <slot name="dialog-content" />
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="blue-darken-1" variant="text" @click="dialogStore.close">Close</v-btn>
        <v-btn color="blue-darken-1" variant="text"
            :loading="dialogStore.loading"
            @click="dialogStore.submit"
            v-if="!dialogStore.hideOkButton"
        >
          Save
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup lang="ts" generic="T extends WithID">
import type { WithID } from '@/lib/template';
import { useDataStore } from '@/stores/data';
import { ref, watch, onMounted, computed } from 'vue';
import { VDataIterator } from 'vuetify/labs/VDataIterator'
import DataPagination from './DataPagination.vue';
import DataToolbar from './DataToolbar.vue';
import DataView from './DataView.vue';
import { useDataDialogStore } from '@/stores/dialogs';

const dataStore = useDataStore();
const dialogStore = useDataDialogStore();

const loading = ref(true);
const sortKey = ref('id');

const page = computed(() => dataStore.page)
const itemsPerPage = computed(() => dataStore.itemsPerPage)

watch(page, loadItems)
watch(itemsPerPage, loadItems)
watch(dataStore.filters, loadItems)

async function loadItems() {
  await dataStore.fetch({
    page: dataStore.page,
    itemsPerPage: dataStore.itemsPerPage,
    sortBy: sortKey.value,
  }, dataStore.filters)
  loading.value = false
}

onMounted(loadItems)
</script>
