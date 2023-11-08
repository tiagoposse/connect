import { defineStore } from 'pinia'
import type { CreateGroupRequest, UpdateGroupRequest, GroupList, CreateUserRequest, UpdateUserRequest, UserList } from '@/lib/api/models';
import { GroupsAPI, UsersAPI } from '@/lib/apis';
import type { CardItem } from '@/lib/utils';
import { hash } from '@stablelib/sha256';

export const useIdentitiesStore = defineStore('identities', {
  state: (): {
    groups: GroupList[],
    users: UserList[],
    groupFilter: null | string,
    groupUser: null | string,
    cache: Uint8Array[],
  } => ({
    users: [],
    groups: [],
    groupFilter: null,
    groupUser: null,
    cache: [],
  }),
  getters: {
    filteredUsers(state): UserList[] {
      if (state.groupFilter) {
        return state.users.filter((item) => item.id == state.groupFilter)
      } else {
        return state.users
      }
    },
    gridUsers() {
      const grid = [] as CardItem[];
      this.filteredUsers.forEach(item => {
        grid.push({
          id: item.id,
          title: `${item.firstname} ${item.lastname}`,
          subtitle: item.email,
          fields: [],
        })
      });
    },
    filteredGroups(state): GroupList[] {
      return state.groups.filter((item) => item.id == state.groupFilter)
    },
    gridGroups() {
      const grid = [] as CardItem[];
      this.filteredGroups.forEach(item => {
        grid.push({
          id: item.id,
          title: item.id,
          subtitle: item.cidr,
          fields: [],
        })
      });
    },
  },
  actions: {
    async fetchUsers() {
      const h = hash(new TextEncoder().encode(''))
      if (this.cache.indexOf(h) == -1) {
        this.cache.push(h)
        this.groups = await GroupsAPI.listGroup();
      }
    },
    async addGroup(item: CreateGroupRequest) {
      const user = await GroupsAPI.createGroup({ createGroupRequest: item }) as GroupList
      this.groups.push(user);
    },
    async removeGroup(id: string) {
      await GroupsAPI.deleteGroup({ id })
      this.groups = this.groups.filter((obj: GroupList) => obj.id !== id);
    },
    async updateGroup(id: string, item: UpdateGroupRequest) {
      const user = await GroupsAPI.updateGroup({
        id,
        updateGroupRequest: item
      }) as GroupList

      const index = this.groups.findIndex((obj: GroupList) => obj.id === id)
      this.groups.splice(index, 1, user);
    },
    async fetchGroups() {
      const h = hash(new TextEncoder().encode(''))
      if (this.cache.indexOf(h) == -1) {
        this.cache.push(h)
        this.groups = await GroupsAPI.listGroup();
      }
    },
    async addUser(item: CreateUserRequest) {
      const user = await UsersAPI.createUser({ createUserRequest: item }) as UserList
      this.users.push(user);
    },
    async removeUser(id: string) {
      await UsersAPI.deleteUser({ id })
      this.users = this.users.filter((obj: UserList) => obj.id !== id);
    },
    async updateUser(id: string, item: UpdateUserRequest) {
      const user = await UsersAPI.updateUser({
        id,
        updateUserRequest: item
      }) as UserList

      const index = this.users.findIndex((obj: UserList) => obj.id === id)
      this.users.splice(index, 1, user);
    },
  },
});
