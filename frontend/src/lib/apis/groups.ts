
import { DefaultApiConfig, type FilterArgs, type GenericAPI, type PaginationArgs } from '@/lib/utils';
import { GroupApi, type CreateGroupRequest, type GroupList } from '@/lib/api';

export const GroupsAPI = new GroupApi(DefaultApiConfig);

export const GenericGroupsAPI: GenericAPI = {
  create: async (payload: CreateGroupRequest) => { return await GroupsAPI.createGroup({ createGroupRequest: payload }) },
  update: async (id: string, payload: CreateGroupRequest) => { return await GroupsAPI.updateGroup({ id, updateGroupRequest: payload }) },
  remove: async (id: string) => { return await GroupsAPI.deleteGroup({ id }) },
  fetch: async (params: PaginationArgs, filters: FilterArgs) => { return await GroupsAPI.listGroupRaw({ ...params, ...filters }) },
  toCard: (item: GroupList) => {
    return {
      id: item.id,
      title: item.name,
      subtitle: '',
      fields: [],
    }
  },
  headers: [
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
      width: "150px",
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
  ],
}
