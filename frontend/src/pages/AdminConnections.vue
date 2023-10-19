<template>
  <v-expansion-panels>
    <v-expansion-panel key="wireguard" title="Wireguard">
      <v-expansion-panel-text>
        <v-row>
          <v-col cols="12" sm="6">
            <v-text-field label="Interface" v-model="wg.interface" variant="outlined"></v-text-field>
          </v-col>

          <v-col cols="12" sm="6">
            <v-text-field label="Cidr" v-model="wg.cidr" variant="outlined"></v-text-field>
          </v-col>

          <v-col cols="12" sm="6">
            <v-text-field label="DNS Servers" v-model="wg.dnsServers" variant="outlined"></v-text-field>
          </v-col>

          <v-col cols="12" sm="6">
            <v-text-field label="Self Provisioning" v-model="wg.selfProvisioning" variant="outlined"></v-text-field>
          </v-col>
        </v-row>
      </v-expansion-panel-text>
    </v-expansion-panel>
    <v-expansion-panel key="saml" title="SAML">
      <v-expansion-panel-text>
        <AdminSettingsSaml />
      </v-expansion-panel-text>
    </v-expansion-panel>
    <v-expansion-panel key="db" title="Database">
      <v-expansion-panel-text>
        <v-row>
          <v-col cols="12" sm="6">
            <v-text-field label="Host" v-model="database.hostname" variant="outlined"></v-text-field>
          </v-col>

          <v-col cols="12" sm="6">
            <v-text-field label="Port" v-model="database.port" variant="outlined"></v-text-field>
          </v-col>

          <v-col cols="12" sm="6">
            <v-text-field label="Database" v-model="database.database" variant="outlined"></v-text-field>
          </v-col>

          <v-col cols="12" sm="6">
            <v-text-field label="Username" v-model="database.username" variant="outlined"></v-text-field>
          </v-col>

          <v-col cols="12" sm="6">
            <v-text-field label="Password" v-model="database.password" variant="outlined"></v-text-field>
          </v-col>
        </v-row>
      </v-expansion-panel-text>
    </v-expansion-panel>
  </v-expansion-panels>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const wg = ref({
  interface: '',
  cidr: '10.254.0.0/16',
  dnsServers: ['1.1.1.1'],
  selfProvisioning: false
})

const database = ref({
  database: '',
  hostname: 'localhost',
  port: 5432,
  username: 'wireguard',
  password: ''
})
</script>

<style scoped>
.device-grid {
  display: flex;
  flex-wrap: wrap;
}
</style>
