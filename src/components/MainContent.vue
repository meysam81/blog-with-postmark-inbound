<template>
  <main id="main-content" class="relative z-10 py-16 md:py-24" role="main" aria-label="Blog posts and content">
    <div class="container mx-auto px-6">
      <!-- Network Status Indicator -->
      <div v-if="!isOnline" class="network-status offline" role="alert" aria-live="assertive">
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
            d="M18.364 5.636l-12.728 12.728m0 0L12 12m-6.364 6.364L12 12m6.364-6.364L12 12" />
        </svg>
        <span>You are currently offline. Some features may be limited.</span>
      </div>

      <!-- Error Boundary Wrapper -->
      <ErrorBoundary @retry="handleErrorRetry">
        <!-- Section Header -->
        <div class="text-center mb-16">
          <div
            class="inline-flex items-center gap-2 bg-slate-800/50 backdrop-blur-sm border border-slate-700/50 rounded-full px-6 py-2 mb-8">
            <div class="w-2 h-2 bg-emerald-400 rounded-full animate-pulse"></div>
            <span class="text-sm text-slate-300 font-medium">Latest Blog Posts</span>
          </div>

          <h2 class="text-3xl md:text-5xl font-bold text-slate-100 mb-6">
            <span class="block">Blog Posts From</span>
            <span class="text-transparent bg-clip-text bg-gradient-to-r from-emerald-300 to-blue-300">
              Email
            </span>
          </h2>

          <p class="text-lg text-slate-300 max-w-2xl mx-auto leading-relaxed">
            Experience the magic of email-powered blogging. Each post below was created by simply sending an email—no
            complex interfaces or setup required.
          </p>
        </div>

        <!-- Content Area -->
        <div class="max-w-6xl mx-auto">
          <!-- Loading State -->
          <div v-if="isLoading" class="flex flex-col items-center justify-center py-20" role="status"
            aria-label="Loading blog posts">
            <div class="relative">
              <div class="w-16 h-16 border-4 border-slate-700 border-t-emerald-400 rounded-full animate-spin"></div>
              <div
                class="absolute inset-0 w-16 h-16 border-4 border-transparent border-t-blue-400 rounded-full animate-spin"
                style="animation-delay: 0.15s;"></div>
            </div>
            <p class="text-slate-300 mt-6 text-lg">Loading blog posts...</p>
            <div class="flex space-x-1 mt-4">
              <div class="w-2 h-2 bg-emerald-400 rounded-full animate-bounce"></div>
              <div class="w-2 h-2 bg-blue-400 rounded-full animate-bounce" style="animation-delay: 0.1s;"></div>
              <div class="w-2 h-2 bg-purple-400 rounded-full animate-bounce" style="animation-delay: 0.2s;"></div>
            </div>
          </div>

          <!-- Error State -->
          <div v-else-if="hasError" class="flex flex-col items-center justify-center py-20" role="alert"
            aria-live="polite">
            <div class="text-center max-w-md mx-auto">
              <div class="w-16 h-16 mx-auto mb-6 text-red-400">
                <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
                </svg>
              </div>
              <h3 class="text-xl font-semibold text-slate-100 mb-4">Oops! Something went wrong</h3>
              <p class="text-slate-300 mb-6">{{ errorMessage }}</p>
              <div class="space-y-3">
                <button @click="retryLoadPosts"
                  class="w-full bg-gradient-to-r from-emerald-500 to-blue-500 hover:from-emerald-400 hover:to-blue-400 text-white font-semibold py-3 px-6 rounded-lg transition-all duration-300 transform hover:scale-105">
                  <span class="flex items-center justify-center gap-2">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                    </svg>
                    Try Again
                  </span>
                </button>
                <p class="text-sm text-slate-400">
                  {{ retryCount > 0 ? `Attempted ${retryCount}/${maxRetries} times` : 'First attempt failed' }}
                </p>
              </div>
            </div>
          </div>

          <!-- Content -->
          <template v-else>
            <BlogPosts v-if="posts.length > 0" :posts="posts" class="fade-in-up" />
            <EmptyState v-else class="fade-in-up" />
          </template>
        </div>

        <!-- Call to Action Section -->
        <div v-if="posts.length > 0" class="text-center mt-20">
          <div class="max-w-3xl mx-auto">
            <h3 class="text-2xl md:text-3xl font-bold text-slate-100 mb-6">
              Ready to Try It?
            </h3>
            <p class="text-slate-300 mb-8 leading-relaxed">
              Send an email to our blog address and watch your post appear here instantly. No registration, no complex
              interface—just email.
            </p>
            <button @click="scrollToTop"
              class="group bg-gradient-to-r from-purple-500 to-blue-500 hover:from-purple-400 hover:to-blue-400 text-white font-semibold px-8 py-4 rounded-full transition-all duration-300 transform hover:scale-105 hover:shadow-xl hover:shadow-purple-500/25"
              aria-label="Scroll to email instructions">
              <span class="flex items-center gap-2">
                Send Your First Email
                <svg class="w-5 h-5 group-hover:-translate-y-1 transition-transform" fill="none" stroke="currentColor"
                  viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16l4-4m0 0l4 4m-4-4v18">
                  </path>
                </svg>
              </span>
            </button>
          </div>
        </div>
      </ErrorBoundary>
    </div>

    <!-- Floating Elements for Visual Interest -->
    <div class="absolute top-1/4 left-8 w-32 h-32 bg-emerald-500/5 rounded-full blur-xl pointer-events-none"></div>
    <div class="absolute bottom-1/3 right-12 w-24 h-24 bg-blue-500/5 rounded-full blur-xl pointer-events-none"></div>
  </main>
</template>

<script>
import { ref, onMounted, onUnmounted } from 'vue'
import BlogPosts from '@/components/BlogPosts.vue'
import EmptyState from '@/components/EmptyState.vue'
import ErrorBoundary from '@/components/ErrorBoundary.vue'
import { fetchPosts } from '@/utils/api.js'
import { useNetworkStatus } from '@/utils/network.js'
import { VITE_WEBSOCKET_URL } from '@/utils/env.js'
import log from 'loglevel'

export default {
  name: 'MainContent',
  components: {
    BlogPosts,
    EmptyState,
    ErrorBoundary
  },
  setup() {
    var posts = ref([])
    var isLoading = ref(true)
    var hasError = ref(false)
    var errorMessage = ref('')
    var retryCount = ref(0)
    var maxRetries = 3
    var websocket = ref(null)
    var reconnectTimeout = ref(null)
    var reconnectAttempts = ref(0)
    var maxReconnectAttempts = 5

    var { isOnline } = useNetworkStatus()

    async function loadPosts() {
      try {
        isLoading.value = true
        hasError.value = false
        errorMessage.value = ''

        log.debug("Starting to load posts...")
        var fetchedPosts = await fetchPosts()

        posts.value = fetchedPosts
        retryCount.value = 0
        log.info(`Successfully loaded ${fetchedPosts.length} posts`)

      } catch (error) {
        log.error("Error loading posts:", error)
        hasError.value = true
        errorMessage.value = error.message || 'Failed to load posts. Please try again.'
        posts.value = []

        if (retryCount.value < maxRetries) {
          retryCount.value++
          var delay = 2 ** retryCount.value * 1000
          log.info(`Retrying in ${delay}ms (attempt ${retryCount.value}/${maxRetries})`)

          setTimeout(function retryLoadPosts() {
            loadPosts()
          }, delay)
        } else {
          log.error("Max retries exceeded, giving up")
        }
      } finally {
        isLoading.value = false
      }
    }

    function connectWebSocket() {
      if (websocket.value && websocket.value.readyState === WebSocket.OPEN) {
        log.debug('WebSocket already connected')
        return
      }

      try {
        log.debug('Attempting to connect WebSocket...')
        websocket.value = new WebSocket(VITE_WEBSOCKET_URL)

        websocket.value.onopen = function wsOnOpen() {
          log.info('WebSocket connection established')
          reconnectAttempts.value = 0
        }

        websocket.value.onmessage = async function wsOnMessage(event) {
          if (event.data === 'ping') {
            websocket.value.send('pong')
            return
          }
          log.info("WebSocket message received:", event.data)
          await loadPosts()
        }

        websocket.value.onclose = function wsOnClose(event) {
          log.warn('WebSocket connection closed:', event.code, event.reason)
          websocket.value = null
          if (event.code !== 1000) {
            scheduleReconnect()
          }
        }

        websocket.value.onerror = function wsOnError(error) {
          log.error('WebSocket error:', error)
          if (websocket.value) {
            websocket.value.close()
          }
        }

      } catch (error) {
        log.error('Failed to create WebSocket connection:', error)
        scheduleReconnect()
      }
    }

    function scheduleReconnect() {
      if (reconnectAttempts.value >= maxReconnectAttempts) {
        log.error('Max WebSocket reconnection attempts exceeded')
        return
      }

      if (reconnectTimeout.value) {
        clearTimeout(reconnectTimeout.value)
      }

      reconnectAttempts.value++
      var delay = Math.min(1000 * Math.pow(2, reconnectAttempts.value), 30000)

      log.info(`Scheduling WebSocket reconnection in ${delay}ms (attempt ${reconnectAttempts.value}/${maxReconnectAttempts})`)

      reconnectTimeout.value = setTimeout(function reconnectWebSocket() {
        connectWebSocket()
      }, delay)
    }

    function closeWebSocket() {
      if (reconnectTimeout.value) {
        clearTimeout(reconnectTimeout.value)
        reconnectTimeout.value = null
      }

      if (websocket.value) {
        websocket.value.close()
        websocket.value = null
      }
    }

    function retryLoadPosts() {
      retryCount.value = 0
      loadPosts()
    }

    function scrollToTop() {
      try {
        // First try to scroll to the email instructions in the header
        var publishingInstructions = document.querySelector('[aria-label="Email instructions"]')
        if (publishingInstructions) {
          publishingInstructions.scrollIntoView({
            behavior: 'smooth',
            block: 'start'
          })
          publishingInstructions.focus()
          log.debug('Scrolled to email instructions')
          return
        }

        // Fallback: scroll to the top of the page where header with instructions is
        window.scrollTo({
          top: 0,
          behavior: 'smooth'
        })
        log.debug('Scrolled to top of page')

        // Add accessibility announcement
        var announcement = document.createElement('div')
        announcement.setAttribute('aria-live', 'polite')
        announcement.className = 'sr-only'
        announcement.textContent = 'Scrolled to email instructions at the top of the page'
        document.body.appendChild(announcement)

        setTimeout(function removeAnnouncement() {
          document.body.removeChild(announcement)
        }, 3000)

      } catch (error) {
        log.error('Error scrolling to email instructions:', error)
        // Fallback: just scroll to top
        window.scrollTo(0, 0)
      }
    }

    function handleErrorRetry() {
      retryLoadPosts()
    }

    onMounted(function onMountedHandler() {
      log.debug("Mounting the main content...")
      connectWebSocket()
      loadPosts()
    })

    onUnmounted(function onUnmountedHandler() {
      log.debug("Unmounting main content, cleaning up WebSocket...")
      closeWebSocket()
    })

    return {
      posts,
      isLoading,
      hasError,
      errorMessage,
      retryCount,
      maxRetries,
      scrollToTop,
      retryLoadPosts,
      isOnline,
      handleErrorRetry
    }
  }
}
</script>

<style scoped>
/* Network Status Indicator */
.network-status {
  position: fixed;
  top: 1rem;
  right: 1rem;
  z-index: 50;
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem 1rem;
  border-radius: 8px;
  font-size: 0.875rem;
  font-weight: 500;
  backdrop-filter: blur(8px);
  transition: all 0.3s ease;
}

.network-status.offline {
  background: rgba(239, 68, 68, 0.9);
  color: white;
  border: 1px solid rgba(239, 68, 68, 0.3);
  animation: pulseRed 2s ease-in-out infinite;
}

@keyframes pulseRed {

  0%,
  100% {
    opacity: 0.9;
  }

  50% {
    opacity: 1;
  }
}

/* Animation utilities */
.fade-in-up {
  animation: fadeInUp 0.6s ease-out;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }

  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
