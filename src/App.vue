<template>
  <div class="min-h-screen relative overflow-hidden">
    <!-- Accessibility Announcements -->
    <div
      v-for="announcement in announcements"
      :key="announcement.id"
      :aria-live="announcement.priority"
      :aria-atomic="true"
      class="sr-only"
    >
      {{ announcement.message }}
    </div>

    <!-- Skip Navigation -->
    <SkipLink />

    <!-- Global Error Boundary -->
    <ErrorBoundary @error="handleGlobalError" @retry="handleGlobalRetry">
      <!-- Header -->
      <AppHeader />

      <!-- Main Content -->
      <MainContent />

      <!-- Footer -->
      <AppFooter />

      <!-- Floating Publish Button -->
      <PublishButton />

      <!-- Scroll Progress Indicator -->
      <ScrollProgress />
    </ErrorBoundary>
  </div>
</template>

<script>
import SkipLink from '@/components/SkipLink.vue'
import AppHeader from '@/components/AppHeader.vue'
import MainContent from '@/components/MainContent.vue'
import AppFooter from '@/components/AppFooter.vue'
import PublishButton from '@/components/PublishButton.vue'
import ScrollProgress from '@/components/ScrollProgress.vue'
import ErrorBoundary from '@/components/ErrorBoundary.vue'
import { useAccessibility } from '@/utils/accessibility.js'
import { usePerformanceMonitoring } from '@/utils/performance.js'
import { useNetworkStatus } from '@/utils/network.js'
import log from 'loglevel'

export default {
  name: 'App',
  components: {
    SkipLink,
    AppHeader,
    MainContent,
    AppFooter,
    PublishButton,
    ScrollProgress,
    ErrorBoundary
  },
  setup() {
    var { announcements } = useAccessibility()
    var { markPerformance } = usePerformanceMonitoring()
    var { isOnline } = useNetworkStatus()

    function handleGlobalError(error) {
      log.error('Global error caught:', error)
    }

    function handleGlobalRetry() {
      log.info('Global retry triggered')
      window.location.reload()
    }

    // Mark app initialization
    markPerformance('app-initialized')

    return {
      announcements,
      isOnline,
      handleGlobalError,
      handleGlobalRetry
    }
  }
}
</script>

<style>
/* Screen reader only content */
.sr-only {
  position: absolute;
  width: 1px;
  height: 1px;
  padding: 0;
  margin: -1px;
  overflow: hidden;
  clip: rect(0, 0, 0, 0);
  white-space: nowrap;
  border: 0;
}

/* Focus styles for accessibility */
*:focus {
  outline: 2px solid #10b981;
  outline-offset: 2px;
}

/* High contrast mode support */
@media (prefers-contrast: high) {
  *:focus {
    outline: 3px solid #000;
    outline-offset: 2px;
  }
}

/* Reduced motion support */
@media (prefers-reduced-motion: reduce) {
  * {
    animation-duration: 0.01ms !important;
    animation-iteration-count: 1 !important;
    transition-duration: 0.01ms !important;
  }
}
</style>
