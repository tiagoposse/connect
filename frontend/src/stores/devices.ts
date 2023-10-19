import { defineStore } from 'pinia'
// import { type UserList } from '../lib/api/models/UserList';
import { type DeviceList } from '../lib/api/models/index';

export const useDevicesStore = defineStore('devices', {
  state: () => ({
    devices: [] as DeviceList[],
    filter: 'all',
  }),
  getters: {
    deviceUsers(): DeviceList[] {
      if (this.filter === 'finished') {
        return this.devices
      }
      return this.devices
    },
    // userDevices(state): { [key: string]: DeviceList[] } {
    //   const m: { [key: string]: DeviceList[] } = {}

    //   for (const d of state.devices) {
    //     if (!this.users.some((item: UserList) => item.email === d.user)) {
    //       continue
    //     }

    //     if (!m.hasOwnProperty(d.user)) {
    //       m[d.user] = []
    //     }
    //     m[d.user].push(d)
    //   }

    //   return m
    // }
  },
  actions: {
    async setDevices(devs: DeviceList[]) {
      this.devices = devs
    },
    async addDevice(dev: DeviceList) {
      this.devices.push(dev)
    },
    async removeDevice(id: number) {
      this.devices = this.devices.filter((obj: DeviceList) => obj.id !== id);
    }
  }
});
