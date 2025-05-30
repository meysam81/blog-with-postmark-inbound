import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './styles.css'
import log from 'loglevel'

if (import.meta.env.DEV) {
  log.setDefaultLevel("debug");
} else {
  log.setDefaultLevel("warn");
}

log.info("🚀 Tarzan: Email to Blog Platform - Initializing application...");

var app = createApp(App)

app.use(router)

app.config.errorHandler = function handlerError(err, instance, info) {
  log.error('Vue error:', err, info)
}

app.config.warnHandler = function handlerWarn(msg, instance, trace) {
  if (import.meta.env.DEV) {
    log.warn('Vue warning:', msg, trace)
  }
}

app.mount('#app')
