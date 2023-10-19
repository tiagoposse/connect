import { defineStore } from 'pinia'
import type { CreateUserRequest, UpdateUserRequest, UserList } from '@/lib/api/models';
import { UsersAPI } from '@/lib/apis';

export const useUsersStore = defineStore('users', {
  state: (): { users: UserList[] } => ({
    users: []
  }),

  actions: {
    async fetch() {
      this.users = await UsersAPI.listUser();
    },
    async add(item: CreateUserRequest) {
      const user = await UsersAPI.createUser({ createUserRequest: item }) as UserList
      this.users.push(user);
    },
    async remove(id: string) {
      await UsersAPI.deleteUser({ id })
      this.users = this.users.filter((obj: UserList) => obj.id !== id);
    },
    async update(id: string, item: UpdateUserRequest) {
      const user = await UsersAPI.updateUser({
        id,
        updateUserRequest: item
      }) as UserList

      const index = this.users.findIndex((obj: UserList) => obj.id === id)
      this.users.splice(index, 1, user);
    },
  },
});
