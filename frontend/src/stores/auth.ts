import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    isAuthenticated: false,
    isAdmin: false,
    email: '',
    userID: '',
    displayName: '',
    token: '',
    photoUrl: ''
  }),
  actions: {
    setUser(data: any) {
      this.isAdmin = data.isAdmin;
      this.email = data.email
      this.userID = data.id
      this.displayName = `${data.firstname} ${data.lastname}`;
      this.photoUrl = data.photoUrl
      this.isAuthenticated = true;
    },
    logout() {
      this.$reset()
    },
    login(token: string) {
      this.isAuthenticated = true
      this.token = token
    }
  }
});
