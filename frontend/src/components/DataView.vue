<template>
  <v-data-table 
      fixed-header
      height="calc(100vh - 64px - 50px - 54px)"
      density="compact"
      :headers="headers"
      :items="dataStore.items"
      :items-per-page="dataStore.itemsPerPage"
      :page="dataStore.page"
      v-if="globalsStore.view == 'list'"
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
    <template v-slot:item.disabled="{ item }">
      <v-checkbox density="compact" color="error" v-model="item.raw.disabled" disabled hide-details></v-checkbox>
    </template>

    <template v-slot:item.controls="{ index }">
      <v-btn size="small" density="compact" icon="mdi-pencil" variant="plain" @click="() => openUpdate(index)"/>
      <v-btn size="small" density="compact" icon="mdi-delete" variant="plain" @click="() => openRemove(index)"/>
    </template>
    <template v-slot:bottom>
    </template>
  </v-data-table>
  <v-row style="height: calc(100vh - 64px - 26px - 54px); align-content: flex-start !important; overflow-y: auto;" v-else-if="globalsStore.view == 'grid'">
    <v-col v-for="(item, index) in dataStore.items" :key="item.id" class="v-col-sm-12 v-col-md-4 v-col-l-3 v-col-xl-2 pa-1" style="align-content: 'flex-start';">
      <slot name="card"
        :item="item"
        :update="() => openUpdate(index)"
        :remove="() => openRemove(index)"
      />
    </v-col>
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
import { computed } from 'vue';

const confirmDialog = useConfirmDialogStore();
const dataDialog = useDataDialogStore();
const notifications = useNotificationsStore();
const dataStore = useDataStore();
const globalsStore = useGlobalsStore();

const headers = computed(() => {
  return dataStore.headers.concat(
    {
      title: 'Controls',
      align: 'end',
      sortable: false,
      key: 'controls',
    },)
})

function openRemove(index: number) {
  confirmDialog.open({
    fn: async () => {
      confirmDialog.setLoading(true)
      await dataStore.remove(dataStore.items[index].id)
      confirmDialog.close()
      notifications.showSuccess("Removed.")
    },
  })
};

function openUpdate(index: number) {
  dataDialog.open({
    title: 'Update?',
    item: dataStore.items[index],
  })
};

</script>
