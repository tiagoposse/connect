<template>
  <v-navigation-drawer permanent v-model="mainSidebarDrawer" :rail="rail" @transitionend="active = false"
    @click="rail = false" app>
    <v-list-item :prepend-avatar="authStore.photoUrl" :title="authStore.displayName" nav>
      <template v-slot:append>
        <v-btn variant="text" icon="mdi-chevron-left" @click.stop="rail = !rail"></v-btn>
      </template>
    </v-list-item>
    <v-divider />
    <v-list v-model="open">
      <v-list-item prepend-icon="mdi-devices" title="My devices" to="/devices"></v-list-item>

      <v-list-group value="Admin">
        <template v-slot:activator="{ props }">
          <v-list-item v-bind="props" prepend-icon="mdi-account-circle" title="Admin"></v-list-item>
        </template>

        <v-list-item key="overview" value="overview" title="Overview" to="/admin"></v-list-item>
        <v-list-item key="devices" value="devices" title="Devices" to="/admin/devices"></v-list-item>
        <v-list-item key="users" value="users" title="Users" to="/admin/users"></v-list-item>
        <v-list-item key="connections" value="connections" title="Connections" to="/admin/connections"></v-list-item>
        <v-list-item key="groups" value="groups" title="Groups" to="/admin/groups"></v-list-item>
      </v-list-group>
    </v-list>
  </v-navigation-drawer>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAuthStore } from '../stores/auth';

const authStore = useAuthStore();
const active = ref(false);
const mainSidebarDrawer = ref(true);
const rail = ref(true)
const open = ref(false)
</script>
