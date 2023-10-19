import { defineStore } from 'pinia'
import { useNotificationsStore } from './notifications';
import { RequestError, ValidationError } from '@/lib/errors';

export const useDataDialogStore = defineStore('dataDialog', {
  state: () => ({
    callback: async () => true,
    isOpen: false,
    hideOkButton: false,
    item: {} as any,
    title: '',
    loading: false,
  }),
  actions: {
    setLoading(val: boolean) {
      this.loading = val
    },
    open({ item, title }: { item?: any, title?: string }) {
      this.isOpen = true
      this.loading = false
      this.item = item
      this.title = title || 'Create'
      this.hideOkButton = false
    },
    close() {
      this.$reset()
    },
    hideOk() {
      this.hideOkButton = true
    },
    async submit(): Promise<boolean> {
      const notificationsStore = useNotificationsStore()

      try {
        await this.callback();
        notificationsStore.showSuccess("Saved.")
        return true
      } catch(e: any) {
        if (e instanceof ValidationError) {
          notificationsStore.showMultiError(e.errors)
        } else if (e instanceof RequestError) {
          notificationsStore.showError("error saving: " + e.message)
        } else {
          notificationsStore.showError(e)
        }

        return false
      }
    },
    registerCallback(fn: () => Promise<boolean>) {
      this.callback = fn
    }
  }
});

export const useConfirmDialogStore = defineStore('confirmDialog', {
  state: () => ({
    callback: async () => {},
    isOpen: false,
    title: '',
    loading: false,
  }),
  actions: {
    setLoading(val: boolean) {
      this.loading = val
    },
    open({ fn, title }: { item?: any, fn: () => Promise<void>, title?: string}) {
      this.isOpen = true
      this.loading = false
      this.callback = fn
      this.title = title || 'Confirm?'
    },
    close() {
      this.$reset()
    },
    async submit(): Promise<boolean> {
      const notificationsStore = useNotificationsStore()

      try {
        await this.callback();
        return true
      } catch(e) {
        notificationsStore.showError("error confirming: " + e)
        return false
      }
    }
  }
});
