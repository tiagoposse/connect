<template>
  <v-toolbar
      dark
      density="compact"
      color="blue-darken-3"
      class="px-2 mb-2"
  >
    <v-text-field
        v-model="search"
        @input="delayedApplyFilter"
        placeholder="Format: field:value"
        del
        clearable
        hide-details
        prepend-inner-icon="mdi-magnify"
        variant="solo"
        density="compact"
    />
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
    <v-btn @click="openCreate">Create</v-btn>
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

<script setup lang="ts">
import { useGlobalsStore } from '@/stores/globals';
import { useDataStore } from '@/stores/data';
import { computed, ref } from 'vue';
import { useDataDialogStore } from '@/stores/dialogs';

const dialogStore = useDataDialogStore();
const globals = useGlobalsStore();
const dataStore = useDataStore();

const search = ref('');

// const open = () => {
//   dataStore.create()
// }

const filters = computed({
  get() {
    return dataStore.filters
  },
  set(fs: { [key: string]: string}) {
    dataStore.setFilters(fs)
  }
})

let timeoutId: number;
const delayedApplyFilter = () => {
  clearTimeout(timeoutId);
  timeoutId = setTimeout(applyFilter, 300);
};

const applyFilter = () => {
  const parts = search.value.split(':');
  if (parts.length === 2) {
    const [field, value] = parts;
    filters.value[field] = value;
  }
};

// const clearFilters = () => {
//   filters.value = {};
// };

function openCreate() {
  useDataDialogStore().open({
    fn: async () => {
      // dialogStore.setLoading(true)
      // await dataStore.create()
      // dialogStore.close()
      // notifications.showSuccess("Removed.")
    },
  })
};
</script>
