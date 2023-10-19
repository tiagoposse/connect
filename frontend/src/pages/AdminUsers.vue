<template>
  <v-data-table :headers="headers" :items="users" show-select>
    <!-- <template v-slot:item.actions="{ item }">
        <v-icon small @click="deleteItem(item)">mdi-delete</v-icon>
      </template> -->
  </v-data-table>
</template>

<script setup lang="ts">
import { useDevicesStore } from '../stores/devices'
import { computed, ref } from 'vue';

const store = useDevicesStore()
const storeDevices = ref(store.devices)

const users = computed(() => {
  const us: string[] = []
  const res: any[] = []
  for (const d of storeDevices.value) {
    if (!us.includes(d.user)) {
      us.push(d.user)
      res.push({ name: d.user })
    }
  }

  return res
})

const headers = [
  { title: 'Name', align: 'start', key: 'name' }
]
</script>
