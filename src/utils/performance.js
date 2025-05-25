import { ref, onMounted } from 'vue'
import log from 'loglevel'

const performanceMetrics = ref({
  loadTime: 0,
  firstContentfulPaint: 0,
  largestContentfulPaint: 0,
  firstInputDelay: 0,
  cumulativeLayoutShift: 0,
  memoryUsage: 0,
  bundleSize: 0
})

const isPerformanceSupported = 'performance' in window
const isObserverSupported = 'PerformanceObserver' in window

function measureLoadTime() {
  if (!isPerformanceSupported) {
    return 0
  }

  var navigationEntry = performance.getEntriesByType('navigation')[0]
  if (navigationEntry) {
    var loadTime = navigationEntry.loadEventEnd - navigationEntry.loadEventStart
    performanceMetrics.value.loadTime = Math.round(loadTime)
    return loadTime
  }
  return 0
}

function measureWebVitals() {
  if (!isObserverSupported) {
    log.warn('PerformanceObserver not supported')
    return
  }

  // First Contentful Paint
  try {
    var paintObserver = new PerformanceObserver(function handlePaintObserver(list) {
      list.getEntries().forEach(function processPaintEntry(entry) {
        if (entry.name === 'first-contentful-paint') {
          performanceMetrics.value.firstContentfulPaint = Math.round(entry.startTime)
          log.debug('First Contentful Paint:', entry.startTime)
        }
      })
    })
    paintObserver.observe({ entryTypes: ['paint'] })
  } catch (error) {
    log.warn('Failed to observe paint metrics:', error)
  }

  // Largest Contentful Paint
  try {
    var lcpObserver = new PerformanceObserver(function handleLcpObserver(list) {
      var entries = list.getEntries()
      var lastEntry = entries[entries.length - 1]
      if (lastEntry) {
        performanceMetrics.value.largestContentfulPaint = Math.round(lastEntry.startTime)
        log.debug('Largest Contentful Paint:', lastEntry.startTime)
      }
    })
    lcpObserver.observe({ entryTypes: ['largest-contentful-paint'] })
  } catch (error) {
    log.warn('Failed to observe LCP metrics:', error)
  }

  // First Input Delay
  try {
    var fidObserver = new PerformanceObserver(function handleFidObserver(list) {
      list.getEntries().forEach(function processFidEntry(entry) {
        var fid = entry.processingStart - entry.startTime
        performanceMetrics.value.firstInputDelay = Math.round(fid)
        log.debug('First Input Delay:', fid)
      })
    })
    fidObserver.observe({ entryTypes: ['first-input'] })
  } catch (error) {
    log.warn('Failed to observe FID metrics:', error)
  }

  // Cumulative Layout Shift
  try {
    var clsObserver = new PerformanceObserver(function handleClsObserver(list) {
      var clsValue = 0
      list.getEntries().forEach(function processClsEntry(entry) {
        if (!entry.hadRecentInput) {
          clsValue += entry.value
        }
      })
      performanceMetrics.value.cumulativeLayoutShift = Math.round(clsValue * 1000) / 1000
      log.debug('Cumulative Layout Shift:', clsValue)
    })
    clsObserver.observe({ entryTypes: ['layout-shift'] })
  } catch (error) {
    log.warn('Failed to observe CLS metrics:', error)
  }
}

function measureMemoryUsage() {
  if ('memory' in performance) {
    var memory = performance.memory
    performanceMetrics.value.memoryUsage = {
      used: Math.round(memory.usedJSHeapSize / 1048576), // MB
      total: Math.round(memory.totalJSHeapSize / 1048576), // MB
      limit: Math.round(memory.jsHeapSizeLimit / 1048576) // MB
    }
    log.debug('Memory usage:', performanceMetrics.value.memoryUsage)
  }
}

function measureBundleSize() {
  if (!isPerformanceSupported) {
    return
  }

  var resourceEntries = performance.getEntriesByType('resource')
  var totalSize = 0

  resourceEntries.forEach(function processResourceEntry(entry) {
    if (entry.transferSize) {
      totalSize += entry.transferSize
    }
  })

  performanceMetrics.value.bundleSize = Math.round(totalSize / 1024) // KB
  log.debug('Total bundle size:', performanceMetrics.value.bundleSize, 'KB')
}

function markPerformance(name, startTime) {
  if (!isPerformanceSupported) {
    return
  }

  try {
    if (startTime) {
      performance.measure(name, startTime)
    } else {
      performance.mark(name)
    }
    log.debug('Performance mark:', name)
  } catch (error) {
    log.warn('Failed to mark performance:', error)
  }
}

function getPerformanceReport() {
  return {
    metrics: performanceMetrics.value,
    support: {
      performance: isPerformanceSupported,
      observer: isObserverSupported,
      memory: 'memory' in (performance || {})
    },
    timestamp: new Date().toISOString(),
    userAgent: navigator.userAgent,
    connection: getConnectionInfo()
  }
}

function getConnectionInfo() {
  if ('connection' in navigator) {
    var connection = navigator.connection || navigator.mozConnection || navigator.webkitConnection
    if (connection) {
      return {
        effectiveType: connection.effectiveType,
        downlink: connection.downlink,
        rtt: connection.rtt,
        saveData: connection.saveData
      }
    }
  }
  return null
}

function logPerformanceReport() {
  var report = getPerformanceReport()
  log.info('Performance Report:', report)

  // Log warnings for poor performance
  if (report.metrics.firstContentfulPaint > 3000) {
    log.warn('First Contentful Paint is slow:', report.metrics.firstContentfulPaint, 'ms')
  }

  if (report.metrics.largestContentfulPaint > 4000) {
    log.warn('Largest Contentful Paint is slow:', report.metrics.largestContentfulPaint, 'ms')
  }

  if (report.metrics.firstInputDelay > 100) {
    log.warn('First Input Delay is high:', report.metrics.firstInputDelay, 'ms')
  }

  if (report.metrics.cumulativeLayoutShift > 0.25) {
    log.warn('Cumulative Layout Shift is high:', report.metrics.cumulativeLayoutShift)
  }

  return report
}

function usePerformanceMonitoring() {
  onMounted(function onMountedHandler() {
    // Wait for page load to start measurements
    if (document.readyState === 'complete') {
      startMeasurements()
    } else {
      window.addEventListener('load', startMeasurements)
    }
  })

  function startMeasurements() {
    measureLoadTime()
    measureWebVitals()
    measureMemoryUsage()
    measureBundleSize()

    // Log initial report after a delay to capture more metrics
    setTimeout(function delayedReport() {
      logPerformanceReport()
    }, 5000)
  }

  return {
    performanceMetrics,
    markPerformance,
    getPerformanceReport,
    logPerformanceReport
  }
}

export {
  usePerformanceMonitoring,
  markPerformance,
  getPerformanceReport,
  logPerformanceReport,
  performanceMetrics
}
