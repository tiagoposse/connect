
import { UserApi, type CreateUserRequest, type UserList } from '@/lib/api';
import { DefaultApiConfig, type FilterArgs, type GenericAPI, type PaginationArgs } from '@/lib/utils';

export const UsersAPI = new UserApi(DefaultApiConfig);

export const GenericUsersAPI: GenericAPI = {
  create: async (payload: CreateUserRequest) => { return await UsersAPI.createUser({ createUserRequest: payload }) },
  update: async (id: string, payload: CreateUserRequest) => { return await UsersAPI.updateUser({ id, updateUserRequest: payload }) },
  remove: async (id: string) => { return await UsersAPI.deleteUser({ id }) },
  fetch: async (params: PaginationArgs, filters: FilterArgs) => { return await UsersAPI.listUserRaw({ ...params, ...filters }) },
  toCard: (item: UserList) => {
    return {
      id: item.id,
      title: `${item.firstname} ${item.lastname}`,
      subtitle: '',
      fields: [],
    }
  },
  headers: [
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
  ]
}
