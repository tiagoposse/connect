<template>
  <v-row>
    <v-col cols="6"><v-text-field v-model="payload.name" label="Name" :rules="[(val: string) => validationRules.required(val, 'name')]" /></v-col>
    <v-col cols="6"><v-text-field v-model="payload.cidr" label="CIDR"  :rules="[(val: string) => validationRules.required(val, 'cidr'), validationRules.cidr]" /></v-col>
  </v-row>
  <v-select chips v-model="payload.scopes" label="Scopes" :items="scopeItems" multiple class="pt-4" :rules="[(val: any[]) => validationRules.nonEmptyArray(val, 'scopes')]" />
  <v-sheet>
    <div class="text-h6">Allow Rules</div>
    <v-text-field v-for="(value, index) in payload.allowRules" :key="index"
        label="Rule" hide-details
        :model-value="value"
        @update:modelValue="(val: string) => handleAllowInput(index, val)"
    />
  </v-sheet>
  <v-sheet class="mt-4">
    <div class="text-h6">Deny Rules</div>
    <v-text-field v-for="(value, index) in payload.denyRules" :key="index"
        label="Rule" hide-details
        :model-value="value"
        @update:modelValue="(val: string) => handleDenyInput(index, val)"
    />
  </v-sheet>
</template>

<script setup lang="ts">

import { GroupCreateScopesEnum, type CreateGroupRequestRulesInner, CreateGroupRequestRulesInnerTypeEnum } from '@/lib/api/models';
import { GroupsAPI } from '@/lib/apis';
import { onMounted, ref } from 'vue';
import { validationRules } from '@/lib/utils';
import { ValidationError } from '@/lib/errors';
import { useDataDialogStore } from '@/stores/dialogs';

const dialogStore = useDataDialogStore();

onMounted(() => {
  if (Object.keys(dialogStore.item).length > 0) {
    payload.value = dialogStore.item
  } else {
    dialogStore.item = payload.value
  }
  
  dialogStore.registerCallback(async () => {
    const validation = [] as (string | boolean)[]
    validation.push(validationRules.required(payload.value.name, 'name'))
    validation.push(validationRules.nonEmptyArray(payload.value.scopes, 'scopes'))
    validation.push(validationRules.required(payload.value.cidr, 'cidr'))
    validation.push(validationRules.cidr(payload.value.cidr))

    const errors = [] as string[]
    validation.forEach((v) => {
      if (typeof v === 'string') {
        errors.push(v)
      }
    })

    if (errors.length > 0) {
      throw new ValidationError(errors)
    }

    const rules = [] as CreateGroupRequestRulesInner[]
    for (const index in payload.value.allowRules.slice(0, payload.value.allowRules.length -1)) {
      rules.push({
        target: payload.value.allowRules[index],
        type: CreateGroupRequestRulesInnerTypeEnum.Allow,
      })
    }

    for (const index in payload.value.denyRules.slice(0, payload.value.allowRules.length -1)) {
      rules.push({
        target: payload.value.denyRules[index],
        type: CreateGroupRequestRulesInnerTypeEnum.Deny,
      })
    }
    
    await GroupsAPI.createGroup({
      createGroupRequest: {
        name: payload.value.name,
        rules,
        scopes: payload.value.scopes,
        cidr: payload.value.cidr,
      }
    })

    return true
  })
})

const payload = ref({
  allowRules: [''],
  denyRules: [''],
  scopes: [],
  cidr: '',
  name: '',
})

const scopeItems = [] as { value: any, title: string} []

const vals = Object.values(GroupCreateScopesEnum)
const keys = Object.keys(GroupCreateScopesEnum)

for (var i = 0; i < vals.length; i++) {
  scopeItems.push({ title: keys[i], value: vals[i]} )
}

const handleAllowInput = (index: number, val: string) => {
  if (payload.value.allowRules[index] == '' && val != '') {
    payload.value.allowRules.push('')
  } else if (payload.value.allowRules[index] != '' && val == '') {
    payload.value.allowRules.pop()
  }

  payload.value.allowRules[index] = val;
};

const handleDenyInput = (index: number, val: string) => {
  if (payload.value.denyRules[index] == '' && val != '') {
    payload.value.denyRules.push('')
  } else if (payload.value.denyRules[index] != '' && val == '') {
    payload.value.denyRules.pop()
  }

  payload.value.denyRules[index] = val;
};
</script>
