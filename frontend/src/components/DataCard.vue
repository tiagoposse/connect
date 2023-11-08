
<template>
  <v-card class="v-col-sm-12 v-col-md-4 v-col-l-3 v-col-xl-2">
    <v-card-title>{{ device.name }}</v-card-title>
    <v-card-subtitle>{{ device.type }}</v-card-subtitle>
    <p>{{ device.description }}</p>
    <v-card-actions>
      <v-spacer />
      <v-btn size="medium" icon="mdi-file-find" @click="open(device)"></v-btn>
      <v-btn size="medium" icon="mdi-delete" @click="removeDevice(device.id)"></v-btn>
    </v-card-actions>
  </v-card>

  <v-dialog v-model="dialog" width="50vh">
    <v-card>
      <v-toolbar>
        <v-toolbar-title dark>
          <span class="text-h5">VPN Config for {{ device.name }}</span>
        </v-toolbar-title>
      </v-toolbar>
      <v-card-text>
        <pre><code v-text="configValue"></code></pre>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="blue-darken-1" variant="text" @click="dialog = false">
          Close
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup lang="ts">
import { type PropType, ref } from 'vue';
import { useDevicesStore } from '../stores/devices'

const deviceStore = useDevicesStore()

const dialog = ref(false)
const configValue = ref('')
defineProps(
  {
    device: {
      type: Object as PropType<DeviceList>,
      required: true
    },
  }
)

const removeDevice =  async (id: number) => {
  await DevicesAPI.deleteDevice({id})
  deviceStore.removeDevice(id)
}

const open = async (device: DeviceList) => {
  // configValue.value = device
  dialog.value = true
}
</script>
