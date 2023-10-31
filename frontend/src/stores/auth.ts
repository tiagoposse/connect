import { defineStore } from 'pinia'
import { AuthAPI } from '@/lib/apis';

export const useAuthStore = defineStore('auth', {
  state: () => ({
    isAuthenticated: false,
    isAdmin: false,
    email: '',
    userID: '',
    displayName: '',
    photoUrl: '',
    scopes: [] as string[],
  }),
  actions: {
    logout() {
      AuthAPI.logout();
      this.$reset();
    },
    async status() {
      const data = await AuthAPI.status();
      this.scopes = data.scopes;
      this.email = data.email;
      this.userID = data.id
      this.displayName = `${data.firstname} ${data.lastname}`;
      this.photoUrl = data.photoUrl
      this.isAuthenticated = true
    }
  }
});
