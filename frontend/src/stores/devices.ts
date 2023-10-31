import { defineStore } from 'pinia'
import type { CreateDeviceRequest, DeviceList } from '@/lib/api/models/index';
import { DevicesAPI } from '@/lib/apis';

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
  },
  actions: {
    async setDevices(devs: DeviceList[]) {
      this.devices = devs
    },
    async addDevice(dev: CreateDeviceRequest): Promise<DeviceList> {
      const resp = await DevicesAPI.createDevice({
        createDeviceRequest: dev
      })
      this.devices.push(resp)
      return resp
    },
    async removeDevice(id: string) {
      this.devices = this.devices.filter((obj: DeviceList) => obj.id !== id);
    }
  }
});
