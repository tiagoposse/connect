
interface Notification {
  id: number
  messages: string[]
  color: string
  isShow: boolean
  position: number
}

import { defineStore } from 'pinia'

const theSizeOfSnackbar = 48;

export const useNotificationsStore = defineStore('notifications', {
  state: () => ({
    notifications: [] as Notification[],
    ids: 0,
  }),
  actions: {
    show(messages: string[], color: string) {
      this.notifications.push({
        id: ++this.ids,
        messages,
        color,
        isShow: true,
        position: theSizeOfSnackbar * this.notifications.length,
      })
    },
    showError(message: string) {
      this.show([message], 'error')
    },
    showMultiError(messages: string[]) {
      this.show(messages, 'error')
    },
    showSuccess(message: string) {
      this.show([message], 'success')
    },
  
    remove(id: number) {
      const removedIdx = this.notifications.findIndex(x => x.id === id)
  
      this.notifications.splice(removedIdx, 1)
      this.notifications.forEach((x, idx) => (x.position = theSizeOfSnackbar * idx))
    }
  }
});
