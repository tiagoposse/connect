
import { DefaultApiConfig, type FilterArgs, type GenericAPI, type PaginationArgs } from '@/lib/utils';
import { AuditApi } from '@/lib/api';

export const AuditAPI = new AuditApi(DefaultApiConfig);

export const GenericAuditAPI: GenericAPI = {
  create: async () => {},
  update: async () => {},
  remove: async () => {},
  fetch: async (params: PaginationArgs, filters: FilterArgs) => { return await AuditAPI.listAuditRaw({ ...params, ...filters }) },
  headers: [
    {
      title: 'Author',
      align: 'start',
      sortable: true,
      key: 'user.email',
    },
    {
      title: 'Timestamp',
      align: 'start',
      sortable: true,
      key: 'timestamp',
    },
    {
      title: 'Action',
      align: 'start',
      sortable: true,
      key: 'action',
    },
  ],
}
