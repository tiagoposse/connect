import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import { aliases, mdi } from 'vuetify/iconsets/mdi'
import './assets/main.css'
import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import { createPinia } from 'pinia'
import '@mdi/font/css/materialdesignicons.css' // Ensure you are using css-loader

import AdminDevices from './src/components/AdminDevicesPage.vue'
import AdminUsers from './src/pages/AdminUsers.vue'
import AdminGroups from './src/pages/AdminGroups.vue'
import User from './src/pages/UserDevices.vue'
import Login from './src/pages/Login.vue'
import AdminConnections from './src/pages/AdminConnections.vue'
import App from './src/App.vue'

const vuetify = createVuetify({
  components,
  directives,
  icons: {
    defaultSet: 'mdi',
    aliases,
    sets: {
      mdi,
    }
  }
})

const app = createApp(App)

app.use(createPinia())
app.use(vuetify)

const routes = [
  { path: '/', component: User, name: "My Devices" },
  { path: '/devices', component: User, name: "My Devices" },
  { path: '/admin/devices', component: AdminDevices, name: "Admin Devices" },
  { path: '/admin/connections', component: AdminConnections, name: "Admin Connections" },
  { path: '/admin/users', component: AdminUsers, name: "Admin Users" },
  { path: '/admin/groups', component: AdminGroups, name: "Admin Groups" },
  { path: '/login', component: Login, name: "Login" }
]
const router = createRouter({
  history: createWebHistory(),
  routes
})


app.use(router)
app.mount('#app')
