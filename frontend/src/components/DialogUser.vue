<template>
  <v-row>
    <v-col><v-text-field v-model="payload.firstname" label="First Name" :rules="[(val: string) => validationRules.required(val, 'firstname'), (val: string) => validationRules.onlyLetters(val, 'firstname')]" /></v-col>
    <v-col><v-text-field v-model="payload.lastname" label="Last Name" :rules="[(val: string) => validationRules.required(val, 'lastname'), (val: string) => validationRules.onlyLetters(val, 'lastname')]" /></v-col>
  </v-row>
  <v-row>
    <v-col><v-text-field v-model="payload.email" label="Email" :rules="[(val: string) => validationRules.required(val, 'email'), validationRules.email]" /></v-col>
    <v-col>
      <v-select v-model="payload.group" label="Group" :items="existingGroups" :rules="[(val: string) => validationRules.required(val, 'group')]" />
    </v-col>
  </v-row>
  <v-row>
    <v-col cols="3"><v-checkbox v-model="payload.disabled" label="Disabled" hide-details /></v-col>
    <v-col><v-text-field v-model="payload.disabledReason" label="Reason" hide-details v-if="payload.disabled" /></v-col>
  </v-row>
  <v-row>
    <v-col cols="4">
      <v-checkbox v-model="setPassword" hide-details>
        <template v-slot:label>
          <div>
            Set Password
            <v-tooltip location="bottom">
              <template v-slot:activator="{ props }">
                <v-icon
                  v-bind="props"
                  @click.stop
                  size="x-small"
                  icon="mdi mdi-information-outline"
                />
              </template>
              If not set, one will be automatically generated
            </v-tooltip>
          </div>
        </template>
      </v-checkbox>
    </v-col>
    <v-col><v-text-field v-model="payload.password" label="Password" hide-details v-if="setPassword" /></v-col>
  </v-row>
  <v-row>
    <v-col><v-text-field v-model="payload.photoUrl" label="Avatar url" hide-details /></v-col>
  </v-row>
</template>

<script setup lang="ts">
import { UsersAPI } from '@/lib/apis';
import { ref, onMounted } from 'vue';
import crypto from 'crypto';
import type { CreateUserRequest } from '@/lib/api';
import argon2 from 'argon2-wasm-esm';
import { useIdentitiesStore } from '@/stores/identities';
import { useDataDialogStore } from '@/stores/dialogs';
import { validationRules } from '@/lib/utils';
import { ValidationError } from '@/lib/errors';

const dialogStore = useDataDialogStore();
const identities = useIdentitiesStore()

onMounted(async () => {
  await identities.fetchGroups()
  identities.groups.forEach((item) => {
    existingGroups.value.push({ value: item.id, title: item.name })
  })
})

const payload = ref({
  email: '',
  firstname: '',
  lastname: '',
  photoUrl: '',
  group: '',
  disabled: false,
  disabledReason: '',
  password: '',
})

const setPassword = ref(false)
const existingGroups = ref<{ value: string, title:string} []>([])

onMounted(() => {
  dialogStore.registerCallback(async () => {
  const validation = [] as (string | boolean)[]
  validation.push(validationRules.required(payload.value.email, 'email'))
  validation.push(validationRules.email(payload.value.email))
  validation.push(validationRules.required(payload.value.firstname, 'firstname'))
  validation.push(validationRules.required(payload.value.lastname, 'lastname'))
  validation.push(validationRules.required(payload.value.group, 'group'))

  const errors = [] as string[]
  validation.forEach((v) => {
    if (typeof v === 'string') {
      errors.push(v)
    }
  })
  if (errors.length > 0) {
    throw new ValidationError(errors)
  }
  
  let createUserRequest: CreateUserRequest = {
    email: payload.value.email,
    firstname: payload.value.firstname,
    lastname: payload.value.lastname,
    photoUrl: payload.value.photoUrl,
    group: payload.value.group,
    disabled: payload.value.disabled,
    disabledReason: payload.value.disabledReason,
    provider: 'userpass',
  }

  if (payload.value.password !== '') {
    createUserRequest.salt = crypto.randomBytes(16).toString();
    createUserRequest.password = await argon2.hash(`${payload.value.password}.${createUserRequest.salt}`)
  }

  await UsersAPI.createUser({ createUserRequest })
  return true
})})

</script>
