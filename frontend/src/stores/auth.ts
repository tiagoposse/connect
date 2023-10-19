import { defineStore } from 'pinia'
import { AuthAPI, UsersAPI } from '@/lib/apis';
import type { UpdateUserRequest } from '@/lib/api';

export const useAuthStore = defineStore('auth', {
  state: () => ({
    isAuthenticated: false,
    isAdmin: false,
    email: '',
    userID: '',
    lastname: '',
    firstname: '',
    photoUrl: '',
    provider: '',
    group: '',
    scopes: [] as string[],
  }),
  getters: {
    displayName: (state) => {
      return `${state.firstname} ${state.lastname}`
    }
  },
  actions: {
    async logout() {
      await AuthAPI.logout();
      this.$reset();
    },
    async status() {
      const data = await AuthAPI.status();
      this.scopes = data.scopes;
      this.email = data.email;
      this.userID = data.id
      this.lastname = data.lastname;
      this.firstname = data.firstname;
      this.photoUrl = data.photoUrl
      this.provider = data.provider
      this.group = data.group
      this.isAuthenticated = true
    },
    async update(fields: UpdateUserRequest) {
      await UsersAPI.updateUser({
        id: this.userID,
        updateUserRequest: fields
      })
      Object.assign(this, fields)
    }
  }
});
