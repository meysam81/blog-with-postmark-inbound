<template>
  <div class="min-h-screen relative overflow-hidden">
    <!-- Header Navigation -->
    <nav class="relative z-20 backdrop-blur-xl bg-slate-950/30 border-b border-slate-800/50" role="navigation">
      <div class="container mx-auto px-6 py-4">
        <div class="flex items-center justify-between">
          <!-- Back to Home -->
          <router-link
            to="/"
            class="group flex items-center gap-3 text-slate-300 hover:text-emerald-300 transition-colors duration-300"
            aria-label="Back to homepage"
          >
            <div class="w-10 h-10 bg-gradient-to-r from-emerald-500 to-blue-500 rounded-lg flex items-center justify-center">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5"/>
              </svg>
            </div>
            <div>
              <div class="font-bold text-lg">TARZAN</div>
              <div class="text-sm opacity-75">Where Words Swing Free</div>
            </div>
          </router-link>

          <!-- Navigation Actions -->
          <div class="flex items-center gap-4">
            <router-link
              to="/"
              class="flex items-center gap-2 text-slate-300 hover:text-emerald-300 transition-colors"
              aria-label="Back to all posts"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
              </svg>
              <span>All Posts</span>
            </router-link>
          </div>
        </div>
      </div>
    </nav>

    <!-- Main Content -->
    <main id="main-content" class="relative z-10 py-8 md:py-16" role="main">
      <!-- Loading State -->
      <div
        v-if="isLoading"
        class="container mx-auto px-6 flex flex-col items-center justify-center py-20"
        role="status"
        aria-label="Loading blog post"
      >
        <div class="relative">
          <div class="w-16 h-16 border-4 border-slate-700 border-t-emerald-400 rounded-full animate-spin"></div>
          <div class="absolute inset-0 w-16 h-16 border-4 border-transparent border-t-blue-400 rounded-full animate-spin" style="animation-delay: 0.15s;"></div>
        </div>
        <p class="text-slate-300 mt-6 text-lg">Loading your story...</p>
      </div>

      <!-- Error State -->
      <div
        v-else-if="hasError"
        class="container mx-auto px-6 flex flex-col items-center justify-center py-20"
        role="alert"
        aria-live="polite"
      >
        <div class="text-center max-w-md mx-auto">
          <div class="w-16 h-16 mx-auto mb-6 text-red-400">
            <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"/>
            </svg>
          </div>
          <h2 class="text-xl font-semibold text-slate-100 mb-4">Post Not Found</h2>
          <p class="text-slate-300 mb-6">{{ errorMessage }}</p>
          <div class="space-y-3">
            <router-link
              to="/"
              class="w-full bg-gradient-to-r from-emerald-500 to-blue-500 hover:from-emerald-400 hover:to-blue-400 text-white font-semibold py-3 px-6 rounded-lg transition-all duration-300 transform hover:scale-105 inline-block text-center"
            >
              <span class="flex items-center justify-center gap-2">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
                </svg>
                Back to All Posts
              </span>
            </router-link>
          </div>
        </div>
      </div>

      <!-- Post Content -->
      <article v-else-if="post" class="container mx-auto px-6 max-w-4xl">
        <!-- Post Header -->
        <header class="text-center mb-12">
          <!-- Category Badge -->
          <div class="inline-flex items-center gap-2 bg-slate-800/50 backdrop-blur-sm border border-slate-700/50 rounded-full px-4 py-2 mb-6">
            <div class="w-2 h-2 bg-emerald-400 rounded-full animate-pulse"></div>
            <span class="text-sm text-slate-300 font-medium">Blog Post</span>
          </div>

          <!-- Title -->
          <h1 class="text-3xl md:text-5xl lg:text-6xl font-bold text-slate-100 mb-8 leading-tight">
            {{ post.title || 'Untitled Post' }}
          </h1>

          <!-- Author and Meta Info -->
          <div class="flex flex-col sm:flex-row items-center justify-center gap-6 text-slate-300">
            <div class="flex items-center gap-3">
              <div class="w-12 h-12 bg-gradient-to-r from-emerald-500 to-blue-500 rounded-full flex items-center justify-center text-white font-semibold">
                {{ getAuthorInitial(post['author-email']) }}
              </div>
              <div class="text-left">
                <p class="font-medium">{{ post['author-email'] || 'Unknown Author' }}</p>
                <p class="text-sm opacity-75">Author</p>
              </div>
            </div>

            <div class="w-px h-8 bg-slate-600 hidden sm:block"></div>

            <div class="text-center sm:text-left">
              <time
                class="font-medium"
                :datetime="formatISODate(post['created-at'])"
              >
                {{ formatReadableDate(post['created-at']) }}
              </time>
              <p class="text-sm opacity-75">Published</p>
            </div>
          </div>
        </header>

        <!-- Post Content -->
        <div class="prose prose-slate prose-lg mx-auto max-w-none">
          <div
            class="blog-post-content"
            v-html="post.content || 'No content available.'"
          ></div>
        </div>

        <!-- Post Footer -->
        <footer class="mt-16 pt-8 border-t border-slate-700/50">
          <div class="flex flex-col sm:flex-row items-center justify-between gap-6">
            <!-- Share Section -->
            <div class="flex items-center gap-4">
              <span class="text-slate-300 font-medium">Share this post:</span>
              <div class="flex items-center gap-2">
                <button
                  @click="copyPostUrl"
                  :class="{ 'text-emerald-400': urlCopied }"
                  class="p-2 text-slate-400 hover:text-emerald-400 transition-colors border border-slate-700 rounded-lg hover:border-emerald-500/50"
                  :aria-label="urlCopied ? 'URL copied!' : 'Copy post URL'"
                >
                  <svg v-if="!urlCopied" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"/>
                  </svg>
                  <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
                  </svg>
                </button>
              </div>
            </div>

            <!-- Navigation -->
            <router-link
              to="/"
              class="group bg-gradient-to-r from-emerald-500 to-blue-500 hover:from-emerald-400 hover:to-blue-400 text-white font-semibold px-6 py-3 rounded-full transition-all duration-300 transform hover:scale-105"
            >
              <span class="flex items-center gap-2">
                <svg class="w-4 h-4 group-hover:-translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
                </svg>
                Explore More Stories
              </span>
            </router-link>
          </div>
        </footer>
      </article>
    </main>

    <!-- Floating Elements for Visual Interest -->
    <div class="absolute top-1/4 left-8 w-32 h-32 bg-emerald-500/5 rounded-full blur-xl pointer-events-none"></div>
    <div class="absolute bottom-1/3 right-12 w-24 h-24 bg-blue-500/5 rounded-full blur-xl pointer-events-none"></div>
  </div>
</template>

<script>
import { ref, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import { fetchPostById, formatISODate, formatReadableDate, getAuthorInitial } from '@/utils/api.js'
import log from 'loglevel'

export default {
  name: 'BlogPostView',
  props: {
    id: {
      type: String,
      required: true
    }
  },
  beforeRouteEnter: function beforeEnterBlogPost(to, from, next) {
    document.title = 'Loading Post... - Tarzan'
    next()
  },
  beforeRouteUpdate: function beforeUpdateBlogPost(to, from, next) {
    document.title = 'Loading Post... - Tarzan'
    next()
  },
  setup(props) {
    var post = ref(null)
    var isLoading = ref(true)
    var hasError = ref(false)
    var errorMessage = ref('')
    var urlCopied = ref(false)
    var route = useRoute()

    async function loadPost(postId) {
      try {
        isLoading.value = true
        hasError.value = false
        errorMessage.value = ''

        log.debug(`Loading post with ID: ${postId}`)
        var fetchedPost = await fetchPostById(postId)

        if (!fetchedPost) {
          throw new Error('Post not found')
        }

        post.value = fetchedPost
        log.info(`Successfully loaded post: ${fetchedPost.title}`)

        // Update document title
        document.title = `${fetchedPost.title || 'Blog Post'} - Tarzan`

      } catch (error) {
        log.error(`Error loading post ${postId}:`, error)
        hasError.value = true
        errorMessage.value = error.message || 'Failed to load the blog post. It may have been removed or is temporarily unavailable.'
        post.value = null

        // Update document title for error state
        document.title = 'Post Not Found - Tarzan'
      } finally {
        isLoading.value = false
      }
    }

    async function copyPostUrl() {
      try {
        var currentUrl = window.location.href
        await navigator.clipboard.writeText(currentUrl)
        urlCopied.value = true

        setTimeout(function resetCopyState() {
          urlCopied.value = false
        }, 2000)

        log.debug('Post URL copied to clipboard')
      } catch (error) {
        try {
          var textArea = document.createElement('textarea')
          textArea.value = window.location.href
          document.body.appendChild(textArea)
          textArea.select()
          document.execCommand('copy')
          document.body.removeChild(textArea)

          urlCopied.value = true
          setTimeout(function resetFallbackCopyState() {
            urlCopied.value = false
          }, 2000)
        } catch (fallbackError) {
          log.error('Failed to copy URL:', fallbackError)
        }
      }
    }

    // Watch for route changes
    watch(() => route.params.id, function handleRouteChange(newId) {
      if (newId) {
        loadPost(newId)
      }
    }, { immediate: true })

    onMounted(function onMountedHandler() {
      if (props.id) {
        loadPost(props.id)
      }
    })

    return {
      post,
      isLoading,
      hasError,
      errorMessage,
      urlCopied,
      copyPostUrl,
      formatISODate,
      formatReadableDate,
      getAuthorInitial
    }
  }
}
</script>

<style scoped>
/* Blog Post Content Styling */
.blog-post-content {
  line-height: 1.8;
  font-size: 1.125rem;
  color: #e2e8f0;
}

.blog-post-content :deep(h1),
.blog-post-content :deep(h2),
.blog-post-content :deep(h3),
.blog-post-content :deep(h4),
.blog-post-content :deep(h5),
.blog-post-content :deep(h6) {
  color: #f1f5f9;
  font-weight: 600;
  margin-top: 2rem;
  margin-bottom: 1rem;
}

.blog-post-content :deep(h1) { font-size: 2.25rem; }
.blog-post-content :deep(h2) { font-size: 1.875rem; }
.blog-post-content :deep(h3) { font-size: 1.5rem; }

.blog-post-content :deep(p) {
  margin-bottom: 1.5rem;
  color: #cbd5e1;
}

.blog-post-content :deep(a) {
  color: #10b981;
  text-decoration: underline;
  transition: color 0.3s ease;
}

.blog-post-content :deep(a:hover) {
  color: #34d399;
}

.blog-post-content :deep(blockquote) {
  border-left: 4px solid #10b981;
  padding-left: 1.5rem;
  margin: 2rem 0;
  font-style: italic;
  background: rgba(16, 185, 129, 0.1);
  padding: 1rem 1.5rem;
  border-radius: 0.5rem;
}

.blog-post-content :deep(ul),
.blog-post-content :deep(ol) {
  margin: 1.5rem 0;
  padding-left: 2rem;
}

.blog-post-content :deep(li) {
  margin-bottom: 0.5rem;
  color: #cbd5e1;
}

.blog-post-content :deep(code) {
  background: #1e293b;
  color: #10b981;
  padding: 0.25rem 0.5rem;
  border-radius: 0.25rem;
  font-family: 'Fira Code', 'JetBrains Mono', 'SF Mono', monospace;
  font-size: 0.875rem;
}

.blog-post-content :deep(pre) {
  background: #0f172a;
  border: 1px solid #334155;
  border-radius: 0.5rem;
  padding: 1.5rem;
  overflow-x: auto;
  margin: 2rem 0;
}

.blog-post-content :deep(pre code) {
  background: none;
  padding: 0;
  color: #e2e8f0;
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .blog-post-content {
    font-size: 1rem;
  }

  .blog-post-content :deep(h1) { font-size: 1.875rem; }
  .blog-post-content :deep(h2) { font-size: 1.5rem; }
  .blog-post-content :deep(h3) { font-size: 1.25rem; }
}
</style>
