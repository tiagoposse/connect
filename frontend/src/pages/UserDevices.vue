<template>
  <DataZone>
    <template v-slot:card="{ item, update, remove }">
      <v-card :title="item.name" :append-icon="item.type == 'laptop' ? 'mdi-laptop' : 'mdi-cellphone'" :update="update" :remove="remove">
      <v-card-subtitle>{{ item.id }}</v-card-subtitle>
      <v-card-text><v-icon>mdi-ip-network-outline</v-icon> {{ item.endpoint }}</v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn size="medium" icon="mdi-pencil-outline" @click="update" class="mr-2" variant="plain" />
        <v-btn size="medium" icon="mdi-delete-outline" @click="remove" variant="plain" class="text-black" />
      </v-card-actions>
    </v-card>

    </template>
    <template v-slot:dialog-content>
      <DialogDevice />
    </template>
  </DataZone>
</template>

<script setup lang="ts">
import DataZone from '@/components/DataZone.vue';
import { GenericDevicesAPI } from '@/lib/apis';
import DialogDevice from '@/components/DialogDevice.vue';
import { useDataStore } from '@/stores/data';
import { onBeforeMount } from 'vue';

onBeforeMount(() => useDataStore().init('users', GenericDevicesAPI))
</script>
