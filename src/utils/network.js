import { ref, onMounted, onUnmounted } from 'vue'
import log from 'loglevel'

const isOnline = ref(navigator.onLine)
const connectionType = ref('unknown')
const networkSpeed = ref('unknown')

function updateNetworkInfo() {
  isOnline.value = navigator.onLine

  if ('connection' in navigator) {
    var connection = navigator.connection || navigator.mozConnection || navigator.webkitConnection
    if (connection) {
      connectionType.value = connection.effectiveType || connection.type || 'unknown'
      networkSpeed.value = connection.downlink ? `${connection.downlink} Mbps` : 'unknown'
    }
  }

  log.debug('Network status updated:', {
    online: isOnline.value,
    type: connectionType.value,
    speed: networkSpeed.value
  })
}

function handleOnline() {
  updateNetworkInfo()
  log.info('Network connection restored')
}

function handleOffline() {
  updateNetworkInfo()
  log.warn('Network connection lost')
}

function useNetworkStatus() {
  onMounted(function onMountedHandler() {
    updateNetworkInfo()

    window.addEventListener('online', handleOnline)
    window.addEventListener('offline', handleOffline)

    if ('connection' in navigator) {
      var connection = navigator.connection || navigator.mozConnection || navigator.webkitConnection
      if (connection) {
        connection.addEventListener('change', updateNetworkInfo)
      }
    }
  })

  onUnmounted(function onUnmountedHandler() {
    window.removeEventListener('online', handleOnline)
    window.removeEventListener('offline', handleOffline)

    if ('connection' in navigator) {
      var connection = navigator.connection || navigator.mozConnection || navigator.webkitConnection
      if (connection) {
        connection.removeEventListener('change', updateNetworkInfo)
      }
    }
  })

  return {
    isOnline,
    connectionType,
    networkSpeed,
    updateNetworkInfo
  }
}

function checkNetworkHealth() {
  return new Promise(function checkNetworkPromise(resolve) {
    if (!navigator.onLine) {
      resolve({ healthy: false, reason: 'offline' })
      return
    }

    var startTime = Date.now()
    var img = new Image()

    img.onload = function onImageLoad() {
      var loadTime = Date.now() - startTime
      var speed = 'slow'
      if (loadTime < 1000) {
        speed = 'fast'
      } else if (loadTime < 3000) {
        speed = 'medium'
      }
      resolve({
        healthy: true,
        latency: loadTime,
        speed: speed
      })
    }

    img.onerror = function onImageError() {
      resolve({ healthy: false, reason: 'connectivity' })
    }

    img.src = '/favicon.ico?' + Date.now()

    setTimeout(function networkTimeout() {
      resolve({ healthy: false, reason: 'timeout' })
    }, 5000)
  })
}

export {
  useNetworkStatus,
  checkNetworkHealth,
  isOnline,
  connectionType,
  networkSpeed
}
