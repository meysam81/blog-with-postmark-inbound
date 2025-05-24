import { createApp } from 'vue'
import App from './App.vue'
import './styles.css'
import log from 'loglevel'

if (import.meta.env.DEV) {
  log.setDefaultLevel("debug");
} else {
  log.setDefaultLevel("warn");
}

log.info("🚀 Tarzan: Where Words Swing Free - Initializing revolutionary blogging experience...");

var app = createApp(App)

app.config.errorHandler = function handlerError(err, instance, info) {
  log.error('Vue error:', err, info)
}

app.config.warnHandler = function handlerWarn(msg, instance, trace) {
  if (import.meta.env.DEV) {
    log.warn('Vue warning:', msg, trace)
  }
}

app.mount('#app')
