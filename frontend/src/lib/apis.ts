
import { ApiKeyApi, AuthApi, Configuration, DeviceApi, GroupApi, UserApi, type CreateDeviceRequest, type ApiResponse } from './api';

// @ts-ignore
const API_URL = `${window.WG_API}/api/v1`;

export const DefaultApiConfig = new Configuration({
  basePath: API_URL,
  credentials: 'include',
},);

// export interface GenericAPI {
//   create: (payload: { [key:string]: any }) => Promise<any>
//   update: (id: string, payload: { [key:string]: any }) => Promise<any>
//   remove: (id: string) => Promise<void>
//   fetch: (params: PaginationArgs, filters: FilterArgs) => Promise<any[]>
// }
export interface GenericAPI {
  create: (payload: any) => Promise<any>
  update: (id: string, payload: any) => Promise<any>
  remove: (id: string) => Promise<void>
  fetch: (params: PaginationArgs, filters: FilterArgs) => Promise<ApiResponse<any[]>>
}

export const DevicesAPI = new DeviceApi(DefaultApiConfig)
export const AuthAPI = new AuthApi(DefaultApiConfig);
export const GroupsAPI = new GroupApi(DefaultApiConfig);
export const UsersAPI = new UserApi(DefaultApiConfig);
export const ApiKeyAPI = new ApiKeyApi(DefaultApiConfig);

export const GenericDevicesAPI: GenericAPI = {
  create: async (payload: CreateDeviceRequest) => { return await DevicesAPI.createDevice({ createDeviceRequest: payload }) },
  update: async (id: string, payload: CreateDeviceRequest) => { return await DevicesAPI.updateDevice({ id, updateDeviceRequest: payload }) },
  remove: async (id: string) => { return await DevicesAPI.deleteDevice({ id }) },
  fetch: async (params: PaginationArgs, _filters: FilterArgs) => { return await DevicesAPI.listDeviceRaw(params) }
}

export type PaginationArgs = { page: number, itemsPerPage: number, sortBy: string }
export type FilterArgs = { [key: string]: string }
export type ListResult<T> = {
  items: T[];
  total: number;
}
