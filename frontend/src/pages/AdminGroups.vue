<template>
  <DataTable :headers="headers" :api="api" :store="useGroupsStore()">
  </DataTable>
</template>

<script setup lang="ts">
import type { CreateGroupOperationRequest, GroupCreate, GroupList, ListGroupRequest, ApiResponse } from '@/lib/api';
import DataTable from '@/components/DataTable.vue';
import { GroupsAPI } from '@/lib/apis';
import { useGroupsStore } from '@/stores/users';

const api = {
  fetch: async (req: ListGroupRequest): Promise<ApiResponse<GroupList[]>> => {
    return await GroupsAPI.listGroupRaw(req)
  },
  add: async (req: CreateGroupOperationRequest): Promise<GroupCreate> => {
    return await GroupsAPI.createGroup(req)
  },
  remove: async (id: string) => {
    return await GroupsAPI.deleteGroup({ id })
  }
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
]
</script>
