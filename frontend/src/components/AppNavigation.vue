<template>
  <v-navigation-drawer permanent v-model="mainSidebarDrawer" :rail="rail" @transitionend="active = false"
    @click="rail = false" app>
    <v-list-item :prepend-avatar="authStore.photoUrl" :title="authStore.displayName" nav>
      <template v-slot:append>
        <v-btn variant="text" icon="mdi-chevron-left" @click.stop="rail = !rail"></v-btn>
      </template>
    </v-list-item>
    <v-divider />
    <v-list v-model="open" density="compact">
      <v-list-item prepend-icon="mdi-account-circle" title="My profile" to="/profile"></v-list-item>
      <v-list-item prepend-icon="mdi-devices" title="My devices" to="/devices"></v-list-item>

      <v-list-group value="Admin">
        <template v-slot:activator="{ props }">
          <v-list-item v-bind="props" prepend-icon="mdi-security" title="Admin"></v-list-item>
        </template>

        <!-- <v-list-item key="overview" value="overview" title="Overview" to="/admin"></v-list-item> -->
        <v-list-item key="devices" prepend-icon="mdi-devices" value="devices" title="Devices" to="/admin/devices"></v-list-item>
        <v-list-item key="users" prepend-icon="mdi-account" value="users" title="Users" to="/admin/users"></v-list-item>
        <v-list-item key="groups" prepend-icon="mdi-account-group" value="groups" title="Groups" to="/admin/groups"></v-list-item>
        <v-list-item key="audit" prepend-icon="mdi-radar" value="audit" title="Audit" to="/admin/audit"></v-list-item>
        <v-list-item key="connections" prepend-icon="mdi-connection" value="connections" title="Connections" to="/admin/connections"></v-list-item>
      </v-list-group>
    </v-list>
    <v-divider />
    <v-spacer />
    <template v-slot:append>
      <v-btn text block variant="plain" @click="logout" v-if="!rail">
        Sign Out
      </v-btn>
      <div class="d-flex mb-2 justify-center align-center" v-else>
        <v-icon icon="mdi-cog" />
      </div>
    </template>
  </v-navigation-drawer>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAuthStore } from '@/stores/auth';
import { useRouter } from 'vue-router'

const router = useRouter()
const authStore = useAuthStore();
const active = ref(false);
const mainSidebarDrawer = ref(true);
const rail = ref(true)
const open = ref(false)

function logout() {
  authStore.logout()
  router.push({ name: "Login" })
}
</script>

<style>
.v-list-group__items .v-list-item {
  padding-inline-start: 42px !important;
}
</style>
