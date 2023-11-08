
import { DefaultApiConfig, type FilterArgs, type GenericAPI, type PaginationArgs } from '@/lib/utils';
import { AuditApi } from '@/lib/api';

export const AuditAPI = new AuditApi(DefaultApiConfig);

export const GenericGroupsAPI: GenericAPI = {
  create: async () => {},
  update: async () => {},
  remove: async () => {},
  fetch: async (params: PaginationArgs, filters: FilterArgs) => { return await AuditAPI.listAuditRaw({ ...params, ...filters }) },
  headers: [
    {
      title: 'Name',
      align: 'start',
      sortable: true,
      key: 'id',
    },
    {
      title: 'Cidr',
      align: 'end',
      sortable: true,
      key: 'cidr',
      width: "150px",
    },
    {
      title: 'Scopes',
      align: 'end',
      sortable: false,
      key: 'scopes',
    },
    {
      title: 'Rules',
      align: 'end',
      sortable: true,
      key: 'rules',
    },
  ],
}
