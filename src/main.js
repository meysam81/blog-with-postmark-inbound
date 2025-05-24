import { createApp } from 'vue'
import App from './App.vue'
import './styles.css'
import log from 'loglevel'

if (import.meta.env.DEV) {
  log.setDefaultLevel("debug");
} else {
  log.setDefaultLevel("warn");
}

log.info("Tarzan Vue.js app is starting!");

var app = createApp(App)
app.mount('#app')
