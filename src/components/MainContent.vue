<template>
  <main
    id="main-content"
    class="relative z-10 py-16 md:py-24"
    role="main"
    aria-label="Blog posts and content"
  >
    <div class="container mx-auto px-6">
      <!-- Section Header -->
      <div class="text-center mb-16">
        <div class="inline-flex items-center gap-2 bg-slate-800/50 backdrop-blur-sm border border-slate-700/50 rounded-full px-6 py-2 mb-8">
          <div class="w-2 h-2 bg-emerald-400 rounded-full animate-pulse"></div>
          <span class="text-sm text-slate-300 font-medium">Latest Publications</span>
        </div>

        <h2 class="text-3xl md:text-5xl font-bold text-slate-100 mb-6">
          <span class="block">Stories That</span>
          <span class="text-transparent bg-clip-text bg-gradient-to-r from-emerald-400 to-blue-400">
            Transcend Boundaries
          </span>
        </h2>

        <p class="text-lg text-slate-300 max-w-2xl mx-auto leading-relaxed">
          Discover groundbreaking thoughts and revolutionary ideas from our community of digital pioneers who chose freedom over interfaces.
        </p>
      </div>

      <!-- Content Area -->
      <div class="max-w-6xl mx-auto">
        <BlogPosts
          v-if="posts.length > 0"
          :posts="posts"
          class="fade-in-up"
        />
        <EmptyState
          v-else
          class="fade-in-up"
        />
      </div>

      <!-- Call to Action Section -->
      <div v-if="posts.length > 0" class="text-center mt-20">
        <div class="max-w-3xl mx-auto">
          <h3 class="text-2xl md:text-3xl font-bold text-slate-100 mb-6">
            Ready to Join the Revolution?
          </h3>
          <p class="text-slate-300 mb-8 leading-relaxed">
            Your voice matters. Your stories deserve to be heard. Break free from traditional publishing constraints and let your words swing free with Tarzan.
          </p>
          <button
            @click="scrollToPublishButton"
            class="group bg-gradient-to-r from-purple-500 to-blue-500 hover:from-purple-400 hover:to-blue-400 text-white font-semibold px-8 py-4 rounded-full transition-all duration-300 transform hover:scale-105 hover:shadow-xl hover:shadow-purple-500/25"
            aria-label="Start publishing your content"
          >
            <span class="flex items-center gap-2">
              Start Your Journey
              <svg class="w-5 h-5 group-hover:translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3"></path>
              </svg>
            </span>
          </button>
        </div>
      </div>
    </div>

    <!-- Floating Elements for Visual Interest -->
    <div class="absolute top-1/4 left-8 w-32 h-32 bg-emerald-500/5 rounded-full blur-xl pointer-events-none"></div>
    <div class="absolute bottom-1/3 right-12 w-24 h-24 bg-blue-500/5 rounded-full blur-xl pointer-events-none"></div>
  </main>
</template>

<script>
import { ref, onMounted } from 'vue'
import BlogPosts from './BlogPosts.vue'
import EmptyState from './EmptyState.vue'
import { fetchPosts } from '../utils/api.js'
import log from 'loglevel'

export default {
  name: 'MainContent',
  components: {
    BlogPosts,
    EmptyState
  },
  setup() {
    var posts = ref([])
    var isLoading = ref(true)

    async function loadPosts() {
      try {
        isLoading.value = true
        var fetchedPosts = await fetchPosts()
        posts.value = fetchedPosts
      } catch (error) {
        log.error("Error loading posts:", error)
        posts.value = []
      } finally {
        isLoading.value = false
      }
    }

    function scrollToPublishButton() {
      var publishButton = document.querySelector('.publish-btn')
      if (publishButton) {
        publishButton.scrollIntoView({
          behavior: 'smooth',
          block: 'center'
        })
        // Add focus for accessibility
        publishButton.focus()
      }
    }

    onMounted(function() {
      loadPosts()
    })

    return {
      posts,
      isLoading,
      scrollToPublishButton
    }
  }
}
</script>
