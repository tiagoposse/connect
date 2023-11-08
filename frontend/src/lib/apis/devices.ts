
import { DefaultApiConfig, type FilterArgs, type GenericAPI, type PaginationArgs, type DataTableHeader } from '@/lib/utils';
import { DeviceApi, type CreateDeviceRequest, type DeviceList } from '@/lib/api';

export const DevicesAPI = new DeviceApi(DefaultApiConfig)

const headers = [
  {
    title: 'ID',
    align: 'start',
    sortable: true,
    key: 'id',
    width: '356px'
  },
  {
    title: 'Name',
    align: 'start',
    sortable: true,
    key: 'name',
  },
  {
    title: 'Endpoint',
    align: 'end',
    sortable: true,
    key: 'endpoint',
    width: '140px'
  },
  {
    title: 'Type',
    align: 'end',
    sortable: false,
    key: 'type',
    width: '82px'
  }
] as DataTableHeader[]


export const GenericDevicesAPI: GenericAPI = {
  create: async (payload: CreateDeviceRequest) => await DevicesAPI.createDevice(payload),
  update: async (id: string, payload: CreateDeviceRequest) => { return await DevicesAPI.updateDevice(id, payload) },
  remove: async (id: string) => { return await DevicesAPI.deleteDevice(id) },
  fetch: async (params: PaginationArgs, filters: FilterArgs) => {
    return await DevicesAPI.listDevice(
      params.page, params.itemsPerPage, params.sortBy,
      filters.id, filters.user, undefined, filters.type, filters.endpoint, filters.allowedIps, filters.publicKey
    ) 
  },
  toCard: (item: DeviceList) => {
    return {
      id: item.id,
      title: item.name,
      subtitle: '',
      fields: [],
    }
  },
  headers,
}
