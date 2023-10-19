<template>
  <v-toolbar density="compact" plain>
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
          {{ dataStore.itemsPerPage }}
        </v-btn>
      </template>
      <v-list>
        <v-list-item
          v-for="(number, index) in itemsPerPageArray"
          :key="index"
          :title="number"
          @click="dataStore.itemsPerPage = number"
        ></v-list-item>
      </v-list>
    </v-menu>
    <span class="mr-4 grey--text">Page {{ dataStore.page }} of {{ numberOfPages }}</span>
    <v-btn icon size="small" flat @click="() => dataStore.setPage(1)">
      <v-icon>mdi-page-first</v-icon>
    </v-btn>
    <v-btn icon size="small" flat @click="prevPage">
      <v-icon>mdi-chevron-left</v-icon>
    </v-btn>
    <v-btn icon size="small" flat @click="nextPage">
      <v-icon>mdi-chevron-right</v-icon>
    </v-btn>
    <v-btn icon size="small" flat @click="() => dataStore.setPage(numberOfPages)">
      <v-icon>mdi-page-last</v-icon>
    </v-btn>
  </v-toolbar>
</template>

<script setup lang="ts">
import { useDataStore } from '@/stores/data';
import { computed } from 'vue';

const dataStore = useDataStore();

const itemsPerPageArray = [10, 20, 50, 100];
const numberOfPages = computed(() => Math.ceil(dataStore.total / dataStore.itemsPerPage))

function nextPage () {
  if (dataStore.page + 1 <= numberOfPages.value) dataStore.setPage(dataStore.page + 1)
}
function prevPage () {
  if (dataStore.page - 1 >= 1) dataStore.setPage(dataStore.page - 1)
}
</script>
