
import { ApiKeyApi, AuthApi, Configuration, DeviceApi, GroupApi, UserApi, type DeviceList } from './api';
import { BaseAPI } from '@/lib/api';
import type { TemplatedStore } from '@/lib/template';
import { useDevicesStore } from '@/stores/devices';
// @ts-ignore
const API_URL = `${window.WG_API}/api/v1`;

type ApiFunctions<T> = {
  fetch: () => Promise<T[]>;
  create: (item: T) => Promise<T>;
  remove: (id: number) => Promise<void>;
  update: (item: T) => Promise<T>;
}

class API<T> {
  private store: TemplatedStore<T>
  private fns: ApiFunctions<T>
  constructor(fns: ApiFunctions<T>, store: TemplatedStore<T>) {
    this.store = store
    this.fns = fns
  }

  async create(item: T) {
    const added = await this.fns.create(item)
    this.store.prototype.add(added)
  }

  async list() {
    const items = await this.fns.fetch()
    this.store.prototype.set(items)
  }

  async remove(id: number) {
    await this.fns.remove(id)
    this.store.prototype.remove(id)
  }
  // async update(item: T) {
  //   const updated = await this.fns.update(item)
  // }
}

export const DefaultApiConfig = new Configuration({
  basePath: API_URL,
  credentials: 'include',
},);

const dev = new DeviceApi(DefaultApiConfig)

// export const DevicesAPI = new API<DeviceList>({
//   create: 
// }, useDevicesStore());
export const DevicesAPI = dev;
export const AuthAPI = new AuthApi(DefaultApiConfig);
export const GroupsAPI = new GroupApi(DefaultApiConfig);
export const UsersAPI = new UserApi(DefaultApiConfig);
export const ApiKeyAPI = new ApiKeyApi(DefaultApiConfig);

// export class RawApi extends BaseAPI {
//   rawRequest(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<DeviceCreate>> {
//     if (requestParameters.createDeviceRequest === null || requestParameters.createDeviceRequest === undefined) {
//       throw new runtime.RequiredError('createDeviceRequest', 'Required parameter requestParameters.createDeviceRequest was null or undefined when calling createDevice.');
//     }

//     const queryParameters: any = {};

//     const headerParameters: runtime.HTTPHeaders = {};

//     headerParameters['Content-Type'] = 'application/json';

//     if (this.configuration && this.configuration.apiKey) {
//       headerParameters["x-api-key"] = this.configuration.apiKey("x-api-key"); // ApiKeyAuth authentication
//     }
//     if (requestParameters.after === null || requestParameters.after === undefined) {
//       throw new runtime.RequiredError('after', 'Required parameter requestParameters.after was null or undefined when calling googleAuthStart.');
//     }

//     const queryParameters: any = {};

//     if (requestParameters.after !== undefined) {
//       queryParameters['after'] = requestParameters.after;
//     }

//     const headerParameters: runtime.HTTPHeaders = {};

//     const response = await this.request({
//       path: `/auth/google/start`,
//       method: 'GET',
//       headers: headerParameters,
//       query: queryParameters,
//     }, initOverrides);

//     return new runtime.VoidApiResponse(response);
//   }
// }
// export const RawAPI = new RawApi();
