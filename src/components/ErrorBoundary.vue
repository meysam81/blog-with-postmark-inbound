<template>
  <div v-if="hasError" class="error-boundary" role="alert" aria-live="assertive">
    <div class="error-content">
      <div class="error-icon" aria-hidden="true">
        <svg class="w-16 h-16" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.732-.833-2.464 0L4.35 16.5c-.77.833.192 2.5 1.732 2.5z"/>
        </svg>
      </div>

      <div class="error-text">
        <h2 class="error-title">Something went wrong</h2>
        <p class="error-message">{{ errorMessage }}</p>
        <p class="error-details" v-if="showDetails">{{ errorDetails }}</p>
      </div>

      <div class="error-actions">
        <button
          class="retry-btn"
          @click="handleRetry"
          :disabled="isRetrying"
          :aria-label="isRetrying ? 'Retrying...' : 'Retry'"
        >
          <svg v-if="isRetrying" class="w-4 h-4 animate-spin" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          <span>{{ isRetrying ? 'Retrying...' : 'Try Again' }}</span>
        </button>

        <button
          class="details-btn"
          @click="toggleDetails"
          :aria-expanded="showDetails"
          aria-controls="error-details"
        >
          {{ showDetails ? 'Hide Details' : 'Show Details' }}
        </button>

        <button
          class="reload-btn"
          @click="handleReload"
          aria-label="Reload page"
        >
          Reload Page
        </button>
      </div>
    </div>
  </div>

  <slot v-else></slot>
</template>

<script>
import { ref, onErrorCaptured } from 'vue'
import log from 'loglevel'

export default {
  name: 'ErrorBoundary',
  emits: ['error', 'retry'],
  setup(props, { emit }) {
    var hasError = ref(false)
    var errorMessage = ref('')
    var errorDetails = ref('')
    var showDetails = ref(false)
    var isRetrying = ref(false)

    function handleError(error, instance, errorInfo) {
      hasError.value = true
      errorMessage.value = error.message || 'An unexpected error occurred'
      errorDetails.value = `${error.stack || error}\n\nComponent: ${instance?.$options.name || 'Unknown'}\nError Info: ${errorInfo}`

      log.error('Error captured by ErrorBoundary:', error)
      log.debug('Error details:', { error, instance, errorInfo })

      emit('error', { error, instance, errorInfo })
    }

    function handleRetry() {
      if (isRetrying.value) {
        return
      }

      isRetrying.value = true
      log.debug('ErrorBoundary: Attempting retry')

      setTimeout(function retryTimeout() {
        hasError.value = false
        errorMessage.value = ''
        errorDetails.value = ''
        showDetails.value = false
        isRetrying.value = false

        emit('retry')
      }, 500)
    }

    function handleReload() {
      window.location.reload()
    }

    function toggleDetails() {
      showDetails.value = !showDetails.value
    }

    onErrorCaptured(function onErrorCapturedHandler(error, instance, errorInfo) {
      handleError(error, instance, errorInfo)
      return false
    })

    return {
      hasError,
      errorMessage,
      errorDetails,
      showDetails,
      isRetrying,
      handleRetry,
      handleReload,
      toggleDetails
    }
  }
}
</script>

<style scoped>
.error-boundary {
  min-height: 400px;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 2rem;
  background: linear-gradient(135deg, rgba(239, 68, 68, 0.05), rgba(239, 68, 68, 0.1));
  border: 2px solid rgba(239, 68, 68, 0.2);
  border-radius: 16px;
  margin: 1rem 0;
}

.error-content {
  text-align: center;
  max-width: 500px;
}

.error-icon {
  color: #ef4444;
  margin-bottom: 1.5rem;
  display: flex;
  justify-content: center;
}

.error-text {
  margin-bottom: 2rem;
}

.error-title {
  font-size: 1.5rem;
  font-weight: 600;
  color: #1f2937;
  margin-bottom: 0.5rem;
}

.error-message {
  color: #6b7280;
  margin-bottom: 1rem;
  font-size: 1rem;
}

.error-details {
  background: #f9fafb;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  padding: 1rem;
  font-family: 'Courier New', monospace;
  font-size: 0.875rem;
  color: #374151;
  white-space: pre-wrap;
  text-align: left;
  max-height: 200px;
  overflow-y: auto;
}

.error-actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
}

.retry-btn,
.details-btn,
.reload-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 8px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
  font-size: 0.875rem;
}

.retry-btn {
  background: #ef4444;
  color: white;
}

.retry-btn:hover:not(:disabled) {
  background: #dc2626;
  transform: translateY(-2px);
}

.retry-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none;
}

.details-btn {
  background: #f3f4f6;
  color: #374151;
  border: 1px solid #d1d5db;
}

.details-btn:hover {
  background: #e5e7eb;
}

.reload-btn {
  background: #6b7280;
  color: white;
}

.reload-btn:hover {
  background: #4b5563;
  transform: translateY(-2px);
}

.animate-spin {
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

/* Responsive Design */
@media (max-width: 640px) {
  .error-boundary {
    padding: 1rem;
    margin: 0.5rem 0;
  }

  .error-actions {
    flex-direction: column;
    width: 100%;
  }

  .retry-btn,
  .details-btn,
  .reload-btn {
    width: 100%;
    justify-content: center;
  }
}

/* Dark Mode */
@media (prefers-color-scheme: dark) {
  .error-boundary {
    background: linear-gradient(135deg, rgba(239, 68, 68, 0.1), rgba(239, 68, 68, 0.15));
    border-color: rgba(239, 68, 68, 0.3);
  }

  .error-title {
    color: #f9fafb;
  }

  .error-message {
    color: #d1d5db;
  }

  .error-details {
    background: #1f2937;
    border-color: #374151;
    color: #e5e7eb;
  }

  .details-btn {
    background: #374151;
    color: #d1d5db;
    border-color: #4b5563;
  }

  .details-btn:hover {
    background: #4b5563;
  }
}

/* High Contrast Mode */
@media (prefers-contrast: high) {
  .error-boundary {
    background: white;
    border: 3px solid #ef4444;
  }

  .error-title {
    color: black;
  }

  .error-message {
    color: #ef4444;
  }
}

/* Reduced Motion */
@media (prefers-reduced-motion: reduce) {
  .retry-btn:hover,
  .reload-btn:hover {
    transform: none;
  }

  .animate-spin {
    animation: none;
  }
}
</style>
