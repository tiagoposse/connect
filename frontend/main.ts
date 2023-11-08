import 'vuetify/styles'
import VuetifyPlugin from './src/plugins/vuetify'
import './assets/main.css'
import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import { createPinia } from 'pinia'
import '@mdi/font/css/materialdesignicons.css' // Ensure you are using css-loader

import AdminDevices from './src/pages/AdminDevices.vue'
import AdminUsers from './src/pages/AdminUsers.vue'
import AdminGroups from './src/pages/AdminGroups.vue'
import User from './src/pages/UserDevices.vue'
import Login from './src/pages/Login.vue'
import AdminConnections from './src/pages/AdminConnections.vue'
import App from './src/App.vue'
import axios from 'axios';

const app = createApp(App)

app.use(createPinia())
app.use(VuetifyPlugin)

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

axios.defaults.withCredentials = true

app.use(router)
app.mount('#app')
