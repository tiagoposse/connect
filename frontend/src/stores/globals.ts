import { defineStore } from 'pinia'

export const useGlobalsStore = defineStore('globals', {
  state: () => ({
    progress: false,
    view: 'list'
  }),
  actions: {
    setProgress(val: boolean) {
      this.progress = val
    }
  }
});
