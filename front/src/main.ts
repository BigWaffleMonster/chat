import { createApp } from 'vue'
import App from './App.vue'

import router from './router'

import './assets/tailwind.css'
import components from './components'
import store from './store'

const app = createApp(App)

components.forEach(component => {
  app.component(component.name, component)
})

app.use(router)
app.use(store)

app.mount('#app')
