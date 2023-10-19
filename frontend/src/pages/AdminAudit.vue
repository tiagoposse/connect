<template>
  <DataZone>
    <template v-slot:dialog-content>
      <DialogGroup />
    </template>
    <template v-slot:card="{ item, update, remove }">
      <v-card :title="item.name" :append-icon="item.type == 'laptop' ? 'mdi-laptop' : 'mdi-cellphone'" :update="update" :remove="remove">
        <v-card-subtitle>{{ item.id }}</v-card-subtitle>
        <v-card-text><v-icon>mdi-ip-network-outline</v-icon> {{ item.cidr }}</v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn size="medium" icon="mdi-pencil-outline" @click="update" class="mr-2" variant="plain" />
          <v-btn size="medium" icon="mdi-delete-outline" @click="remove" variant="plain" class="text-black" />
        </v-card-actions>
      </v-card>
    </template>
  </DataZone>
</template>

<script setup lang="ts">
import DataZone from '@/components/DataZone.vue';
import { GenericAuditAPI } from '@/lib/apis';
import DialogGroup from '@/components/DialogGroup.vue';
import { useDataStore } from '@/stores/data';
import { onBeforeMount } from 'vue';

onBeforeMount(() => useDataStore().init('audit', GenericAuditAPI))
</script>
