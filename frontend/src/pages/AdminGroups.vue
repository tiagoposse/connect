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
      <DialogGroup :onFinish="onFinish" />
    </template>
  </DataTable>
</template>

<script setup lang="ts">
import DataTable from '@/components/DataTable.vue';
import type { CreateGroupRequest, GroupList, UpdateGroupRequest } from '@/lib/api';
import { GroupsAPI, type PaginationArgs, type ListResult } from '@/lib/apis';
import type { CardItem, ReadonlyDataTableHeader } from '@/lib/utils';
import DialogGroup from '@/components/DialogGroup.vue';

async function fetch(pageArgs: PaginationArgs): Promise<ListResult<GroupList>> {
  const resp = await GroupsAPI.listGroupRaw(pageArgs)
  return {
    items: await resp.value(),
    total: parseInt(resp.raw.headers.get("x-total")!)
  }
}

function convertToCard(item: GroupList): CardItem {
  return {
    id: item.id,
    title: item.id,
    subtitle: item.cidr,
    fields: [],
  }
}

async function create(payload: CreateGroupRequest): Promise<GroupList> {
  return (await GroupsAPI.createGroup({ createGroupRequest: payload }) as GroupList)
}

async function remove(id: string) {
  await GroupsAPI.deleteGroup({ id })
}

async function update(id: string, payload: UpdateGroupRequest): Promise<GroupList> {
  return await GroupsAPI.updateGroup({
    id,
    updateGroupRequest: payload
  })
}

const headers = [
  {
    title: 'Name',
    align: 'start',
    sortable: true,
    key: 'id',
  },
  {
    title: 'Cidr',
    align: 'end',
    sortable: true,
    key: 'cidr',
  },
  {
    title: 'Scopes',
    align: 'end',
    sortable: false,
    key: 'scopes',
  },
  {
    title: 'Rules',
    align: 'end',
    sortable: true,
    key: 'rules',
  },
] as ReadonlyDataTableHeader[]
</script>
