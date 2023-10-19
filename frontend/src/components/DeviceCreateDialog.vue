<template>
  <v-dialog :model-value="open" width="50vh">
    <v-card>
      <v-toolbar>
        <v-toolbar-title dark>
          <span class="text-h5">Add Device</span>
        </v-toolbar-title>
      </v-toolbar>
      <v-card-text v-if="created">
        <code class="block whitespace-pre overflow-x-scroll">{{ getPeerConfig(peerConfig) }}</code>
      </v-card-text>
      <v-card-text>
        <v-text-field class="my-4" v-model="deviceData.name" label="Name" required hide-details></v-text-field>
        <v-select :items="['laptop', 'mobile']" v-model="deviceData.type" label="Device Type" required />
        <v-text-field v-model="deviceData.notes" label="Notes" hide-details required></v-text-field>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="blue-darken-1" variant="text" @click="close">
          Close
        </v-btn>
        <v-btn :loading="loading" @click="createDevice" color="blue-darken-1" variant="text">
          Save
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
  <v-snackbar v-model="snackbar" timeout="2000">
    {{ errorMessage }}

    <template v-slot:actions>
      <v-btn color="blue" variant="text" @click="snackbar = false">
        Close
      </v-btn>
    </template>
  </v-snackbar>
</template>

<script setup lang="ts">
import { useDevicesStore } from '../stores/devices'
// import { PeerConfig } from '../api/devices';
import { useAuthStore } from '../stores/auth'
import { ref } from 'vue';
import { generateKeyPair } from '@stablelib/x25519'
import { encode as b64encode } from '@stablelib/base64'
import { CreateDeviceRequest } from '../lib/api'
import { DevicesAPI } from '../lib/apis'

const devicesStore = useDevicesStore()
const authStore = useAuthStore()

const props = defineProps({
  open: Boolean,
  close: Function
})

const created = ref(false)
const loading = ref(false);
const snackbar = ref(false);
const errorMessage = ref("");

const deviceData = ref({
  name: '',
  notes: '',
  user: authStore.userID,
  type: 'laptop',
  publicKey: ''
} as CreateDeviceRequest)

const peerConfig = ref({})

const createDevice = async () => {
  loading.value = true;

  const keyPair = generateKeyPair()
  deviceData.value.publicKey = b64encode(keyPair.publicKey)
  deviceData.value.user = authStore.userID;

  try {
    const resp = await DevicesAPI.createDevice({
      createDeviceRequest: deviceData.value
    })
  } catch (e) {
    snackbar.value = true
    errorMessage.value = e
  }

  // if (resp.error !== undefined) {

  // } else {
  //   peerConfig.value = {
  //     ...resp.details,
  //     ...deviceData.value,
  //     privateKey: b64encode(keyPair.secretKey)
  //   } as PeerConfig

  //   created.value = true
  // }

  loading.value = false
}

const getPeerConfig = (peerConfig: any) => {
  return `[Interface]
PrivateKey = ${peerConfig.privateKey}
Address = ${peerConfig.address}
DNS = ${peerConfig.dns.join(",")}

[Peer]
PublicKey = ${peerConfig.publicKey}
PresharedKey = ${peerConfig.presharedKey}
AllowedIPs = ${peerConfig.allowedIps}
Endpoint = ${peerConfig.endpoint}
PersistentKeepalive = ${peerConfig.keepalive}`
}
</script>
