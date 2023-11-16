
import { DefaultApiConfig } from '@/lib/utils';
import { ApiKeyApi, AuthApi } from '@/lib/api';

export * from './devices';
export * from './groups';
export * from './users';
export * from './audit';

export const AuthAPI = new AuthApi(DefaultApiConfig);
export const ApiKeyAPI = new ApiKeyApi(DefaultApiConfig);
