import {
  VDataTableServer,
} from "vuetify/labs/VDataTable";

import { Configuration, type ApiResponse, type ResponseContext } from './api';
import { RequestError } from "./errors";

type ApiError = {
  error_message: string;
}

// @ts-ignore
export const API_URL = `${window.WG_API}/api/v1`;

export const DefaultApiConfig = new Configuration({
  basePath: API_URL,
  credentials: 'include',
  middleware: [{
    post: async (context: ResponseContext) => {
      if (!context.response.ok) {
        throw new RequestError((await context.response.json() as ApiError).error_message)
      }

      return context.response
    }
  }],
},);

export interface GenericAPI {
  create: (payload: any) => Promise<any>
  update: (id: string, payload: any) => Promise<any>
  remove: (id: string) => Promise<void>
  fetch: (params: PaginationArgs, filters: FilterArgs) => Promise<ApiResponse<any[]>>
  headers: ReadonlyDataTableHeader[]
}

export type PaginationArgs = { page: number, itemsPerPage: number, sortBy: string }
export type FilterArgs = { [key: string]: string }
export type ListResult<T> = {
  items: T[];
  total: number;
}

type UnwrapReadonlyArrayType<A> = A extends Readonly<Array<infer I>> ? UnwrapReadonlyArrayType<I> : A
type DT = InstanceType<typeof VDataTableServer>;
export type ReadonlyDataTableHeader = UnwrapReadonlyArrayType<DT['headers']>;

export type CardItem = {
  id: string;
  title: string;
  subtitle: string;
  fields: any[];
}

export const validationRules = {
  required: (value: string, fieldName: string) => { return !!value || `${fieldName} Required.` },
  nonEmptyArray: (value: any[], fieldName: string) => { return value.length > 0 || `${fieldName} cannot be empty.` },
  email: (value: string) => {
    const pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
    return pattern.test(value) || 'Invalid e-mail.'
  },
  onlyLetters: (value: string, fieldName: string) => {
    const pattern = /^[a-zA-Z]+$/
    return pattern.test(value) || `${fieldName} can only contain letters`
  },
  cidr: (value: string) => {
    const pattern = /^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})\/(\d{1,2})$/
    const matches = pattern.exec(value)
    if (matches == null) {
      return 'Invalid cidr'
    }

    for (let i = 1; i < 5; i++) {
      const octet = parseInt(matches[i])
      if (octet < 0 || octet > 254) {
        return `${matches[i]} is not a valid octet`
      }  
    }

    const mask = parseInt(matches[5])
    if (mask < 0 || mask > 32) {
      return 'Invalid mask'
    }
    
    return true
  },
}
