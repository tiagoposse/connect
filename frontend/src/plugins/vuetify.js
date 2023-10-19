import { createVuetify } from 'vuetify'
import { aliases, mdi } from 'vuetify/iconsets/mdi'

import {
  VDataTable
} from "vuetify/labs/VDataTable";

export default createVuetify({
  components: {
    VDataTable,
  },
  icons: {
    defaultSet: 'mdi',
    aliases,
    sets: {
      mdi,
    }
  },
  display: {
    mobileBreakpoint: 'sm',
  }
})
