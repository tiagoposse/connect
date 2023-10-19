
import type { StoreDefinition } from 'pinia';
import type { Ref } from 'vue';

export interface WithID {
  id:string
}

export type TemplatedStore<T> = StoreDefinition<string, { items: Ref<T[]>; }, {}, { set(newItems: T[]): void; add(item: T): void; remove(id: number): void; }>
