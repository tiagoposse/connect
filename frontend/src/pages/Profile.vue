<template>
  <v-row class="ma-0 mr-2 mt-2 pa-2">
    <v-col style="width: 160px !important; max-width: 160px !important;">
      <v-container class="fill-height justify-center align-start">
        <v-avatar size="104"><v-img :src="user.photoUrl" /></v-avatar>
        <v-btn @click="() => open = true" variant="solo" v-if="canChange">Change avatar</v-btn>
      </v-container>
    </v-col>
    <v-col>
      <v-row>
        <v-col><v-text-field variant="outlined" label="First Name" v-model="user.firstname" hide-details :disabled="!canChange" class="no-opaque" /></v-col>
        <v-col><v-text-field variant="outlined" label="Last Name" v-model="user.lastname" hide-details :disabled="!canChange" class="no-opaque" /></v-col>
      </v-row>
      <v-row>
        <v-col><v-text-field variant="outlined" label="Group" disabled v-model="auth.group" class="no-opaque" hide-details /></v-col>
        <v-col><v-text-field variant="outlined" label="Provider" disabled v-model="auth.provider" class="no-opaque" hide-details /></v-col>
      </v-row>
      <v-row class="pl-4">
        <div v-if="!canChange">
          <v-container style="background-color: cyan;">
            <v-icon icon="mdi-information-outline" /> Your profile is automatically managed by your provider, no edits are possible.
          </v-container>
        </div>
      </v-row>
    </v-col>
  </v-row>
  <v-btn v-if="hasChanges">Save</v-btn>
  <v-btn v-if="hasChanges" @click="reset">Cancel</v-btn>
  <v-dialog >
    <v-card>
      <v-toolbar>
        <v-toolbar-title dark>
          <span class="text-h5">Change Photo URL</span>
        </v-toolbar-title>
      </v-toolbar>
      <v-card-text>
        <v-text-field variant="outlined" label="URL" v-model="user.photoUrl" hide-details />
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="blue-darken-1" variant="text" @click="() => open = false">Close</v-btn>
        <v-btn color="blue-darken-1" variant="text" @click="changeAvatar">Save</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import { ref, computed } from 'vue'
const auth = useAuthStore();
const open = ref(false);

const user = ref({
  email: auth.email,
  firstname: auth.firstname,
  lastname: auth.lastname,
  photoUrl: auth.photoUrl,
})

const reset =() => {
  user.value ={
    email: auth.email,
    firstname: auth.firstname,
    lastname: auth.lastname,
    photoUrl: auth.photoUrl,
  }
}

const hasChanges = computed(() => {
  return (
    user.value.email != auth.email ||
    user.value.firstname != auth.firstname ||
    user.value.lastname != auth.lastname
  )
})

const changeAvatar = () => {
  auth.update({
    photoUrl: user.value.photoUrl,
  })
}

const canChange = computed(() => auth.provider == 'userpass')
</script>
