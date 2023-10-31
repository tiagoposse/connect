<template>
  <DataTable
      :headers="headers"
      :fetch="fetch"
      :update="update"
      :remove="remove"
      :create="create"
      :convert-to-card="convertToCard"
  >
    <template v-slot:dialog-content="{ onFinish }">
      <DialogUser :onFinish="onFinish" />
    </template>
  </DataTable>
</template>

<script setup lang="ts">
import DataTable from '@/components/DataTable.vue';
import type { UserList, UpdateUserRequest, CreateUserRequest } from '@/lib/api';
import { UsersAPI, type PaginationArgs, type ListResult } from '@/lib/apis';
import type { CardItem, ReadonlyDataTableHeader } from '@/lib/utils';
import DialogUser from '@/components/DialogUser.vue';

async function fetch(pageArgs: PaginationArgs): Promise<ListResult<UserList>> {
  const resp = await UsersAPI.listUserRaw(pageArgs)
  return {
    items: await resp.value(),
    total: parseInt(resp.raw.headers.get("x-total")!)
  }
}

async function create(item: CreateUserRequest): Promise<UserList> {
  const user = await UsersAPI.createUser({ createUserRequest: item }) as UserList
  return user
}

function convertToCard(item: UserList): CardItem {
  return {
    id: item.id,
    title: `${item.firstname} ${item.lastname}`,
    subtitle: item.email,
    fields: [],
  }
}

async function remove(id: string) {
  await UsersAPI.deleteUser({ id })
}

async function update(id: string, payload: UpdateUserRequest): Promise<UserList> {
  return await UsersAPI.updateUser({
    id,
    updateUserRequest: payload
  })
}

const headers = [
  {
    title: 'ID',
    align: 'start',
    sortable: true,
    key: 'id',
  },
  {
    title: 'Email',
    align: 'start',
    sortable: true,
    key: 'email',
  },
  {
    title: 'First Name',
    align: 'start',
    sortable: true,
    key: 'firstname',
  },
  {
    title: 'Last Name',
    align: 'start',
    sortable: true,
    key: 'lastname',
  },
  {
    title: 'Group',
    align: 'start',
    sortable: true,
    key: 'group',
  },
  {
    title: 'Provider',
    align: 'start',
    sortable: true,
    key: 'provider',
  },
  {
    title: 'Disabled',
    align: 'start',
    sortable: true,
    key: 'disabled',
  },
  {
    title: 'Reason',
    align: 'start',
    sortable: false,
    key: 'disabledReason',
  },
] as ReadonlyDataTableHeader[]
</script>
