<template>
  <v-data-table 
      fixed-header
      height="calc(100vh - 64px - 50px - 54px)"
      style="--v-table-header-height: 42px; --v-table-row-height: 42px;"
      v-if="globalsStore.view == 'list'"
      :headers="dataStore.headers"
      :items="dataStore.items"
      :items-per-page="dataStore.itemsPerPage"
      :page="dataStore.page"
  >
    <template v-slot:item.scopes="{ item }">
      <v-chip-group class="align-content-end">
        <v-chip v-for="(val, index) in item.raw.scopes" :key="index">{{ val }}</v-chip>
      </v-chip-group>
    </template>
    <template v-slot:item.rules="{ item }">
      <v-chip-group class="align-content-end">
        <v-chip v-for="(val, index) in item.raw.rules" :key="index" :color="val.type === 'allow' ? 'success' : 'error'">{{ val }}</v-chip>
      </v-chip-group>
    </template>
    <template v-slot:bottom>
    </template>
  </v-data-table>
  <v-row style="height: calc(100vh - 64px - 50px - 54px);" v-else-if="globalsStore.view == 'grid'">
    <v-card class="v-col-sm-12 v-col-md-4 v-col-l-3 v-col-xl-2" v-for="item in dataStore.gridItems" :key="item.id">
      <v-card-title>{{ item.title }}</v-card-title>
      <v-card-subtitle v-if="item.subtitle != ''">{{ item.subtitle }}</v-card-subtitle>
      <v-card-actions>
        <v-spacer />
        <v-btn size="medium" icon="mdi-pencil-outline" @click="openUpdate(item.id)" class="mr-2" variant="plain" />
        <v-btn size="medium" icon="mdi-delete-outline" @click="openRemove(item.id)" variant="plain" class="text-black" />
      </v-card-actions>
    </v-card>
  </v-row>
</template>

<script setup lang="ts">
import { useDataStore } from '@/stores/data';
import { useGlobalsStore } from '@/stores/globals';
import { useConfirmDialogStore, useDataDialogStore } from '@/stores/dialogs';
import { useNotificationsStore } from '@/stores/notifications';
import {
  VDataTable,
} from "vuetify/labs/VDataTable";

const confirmDialog = useConfirmDialogStore();
const dataDialog = useDataDialogStore();
const notifications = useNotificationsStore();
const dataStore = useDataStore();
const globalsStore = useGlobalsStore();

function openRemove(id: string) {
  confirmDialog.open({
    fn: async () => {
      confirmDialog.setLoading(true)
      await dataStore.remove(id)
      confirmDialog.close()
      notifications.showSuccess("Removed.")
    },
  })
};

function openUpdate(id: string) {
  let item: any;
  for (const index of dataStore.items) {
    if (dataStore.items[index].id === id) {
      item = dataStore.items[index]
      break
    }
  }

  dataDialog.open({
    title: 'Update?',
    item: item,
    fn: async () => {
      dataDialog.setLoading(true)
      await dataStore.update(id, item)
      dataDialog.close()
    },
  })
};

</script>
