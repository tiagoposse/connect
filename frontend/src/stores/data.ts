import type { FilterArgs, PaginationArgs } from '@/lib/utils';
import type { GenericAPI } from '@/lib/utils';
import { defineStore } from 'pinia'
import { useNotificationsStore } from './notifications';
import type { ReadonlyDataTableHeader } from '@/lib/utils';
import { RequestError } from '@/lib/errors';

export const useDataStore = defineStore('data', {
  state: () => ({
    items: [] as any[],
    total: 0,
    title: '',
    api: null as GenericAPI | null,
    filters: {} as { [key: string]: string },
    itemsPerPage: 50,
    page: 1,
    editAllowed: true,
  }),
  getters: {
    headers(): ReadonlyDataTableHeader[] {
      return this.api!.headers
    },
  },
  actions: {
    init(title: string, api: GenericAPI) {
      if (this.title != title) {
        this.title = title
        this.api = api
        this.page = 1
        this.total = 0
        this.items = []
      }
    },
    async fetch(params: PaginationArgs, filters: FilterArgs) {
      try {
        const resp = await this.api!.fetch(params, filters)
        this.items = await resp.value()
        this.total = parseInt(resp.raw.headers.get('X-Total')!);
      } catch (e: any) {
        useNotificationsStore().showError(e)
      }
    },
    async create(payload: { [key:string]: any}): Promise<any> {
      try {
        const item = await this.api!.create(payload)
        console.log(item)
        this.items.push(item)
        this.total++
        return item
      } catch (e: any) {
        const { message }: { message: string } = e
        throw new RequestError(message)
      }
    },
    async remove(id: string) {
      try {
        await this.api!.remove(id)
        const newItems = this.items.filter((item) => item.id !== id)
        this.items = newItems
        this.total--
      } catch (e: any) {
        useNotificationsStore().showError(e)
      }
    },
    async update(id: string, payload: { [key:string]: any}): Promise<any> {
      let newItem;
      try {
        await this.api!.update(id, payload)
        for (const item of this.items) {
          if (item.id === id) {
            for (const k of Object.keys(payload)) {
              item[k] = payload[k]
            }

            newItem = item
          }
        }
      } catch (e: any) {
        useNotificationsStore().showError(e)
      }
      return newItem
    },
    setPage(p: number) {
      this.page = p
    },
    setItemsPerPage(p: number) {
      this.itemsPerPage = p
    },
    setFilters(fs: { [key:string]: string}) {
      this.filters = fs
    }
  }
});
