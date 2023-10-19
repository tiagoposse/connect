<template>
  <DataZone>
    <template v-slot:card="{ item, update, remove }">
      <v-card :title="`${item.firstname} ${item.lastname}`" :update="update" :remove="remove">
        <template v-slot:prepend>
          <v-avatar>
            <v-img :src="item.photoUrl" />
          </v-avatar>
        </template>
        <template v-slot:append>
          <v-avatar size="24">
            <v-img :src="providers[item.provider]" />
          </v-avatar>
        </template>
        <v-card-subtitle>{{ item.id }}</v-card-subtitle>
        <v-card-text>
          <div>
            <v-icon>mdi-at</v-icon> {{ item.email }}
          </div>
          <div>
            <v-icon>mdi-account-group</v-icon> {{ item.groupId }}
          </div>
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn size="medium" icon="mdi-pencil-outline" @click="update" class="mr-2" variant="plain" />
          <v-btn size="medium" icon="mdi-delete-outline" @click="remove" variant="plain" class="text-black" />
        </v-card-actions>
      </v-card>
    </template>
    <template v-slot:dialog-content>
      <DialogUser />
    </template>
  </DataZone>
</template>

<script setup lang="ts">
import DataZone from '@/components/DataZone.vue';
import { GenericUsersAPI } from '@/lib/apis';
import DialogUser from '@/components/DialogUser.vue';
import { useDataStore } from '@/stores/data';
import { onBeforeMount } from 'vue';

const providers = {
  userpass: '/assets/local.png',
  google: '/assets/google.png'
} as { [key:string]: string}

onBeforeMount(() => useDataStore().init('adminUsers', GenericUsersAPI))
</script>
