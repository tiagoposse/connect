import { defineStore } from 'pinia'
import type { WithID } from '@/lib/template';
import type { UserList, GroupList } from '../lib/api/models';
import type { Ref } from 'vue';
import { ref } from 'vue';


export const storeFromTemplate = <T extends WithID>(name: string) => {
  return defineStore({
    id: name,

    state: (): { items: Ref<T[]> } => ({
      items: ref<T[]>([]) as Ref<T[]>,
    }),

    actions: {
      set(newItems: T[]) {
        this.items = newItems;
      },
      add(item: T) {
        this.items.push(item);
      },
      remove(id: number) {
        this.items = this.items.filter((obj: T) => obj.id !== id);
      },
    },
  });
};

export const useGroupsStore = storeFromTemplate<GroupList>("groups");
export const useUsersStore = storeFromTemplate<UserList>("users");
