<template>
  <main id="main-content" class="container mx-auto px-4 py-10 flex-grow">
    <div class="max-w-4xl mx-auto">
      <BlogPosts v-if="posts.length > 0" :posts="posts" />
      <EmptyState v-else />
    </div>
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

    async function loadPosts() {
      try {
        var fetchedPosts = await fetchPosts()
        posts.value = fetchedPosts
      } catch (error) {
        log.error("Error loading posts:", error)
        posts.value = []
      }
    }

    onMounted(function() {
      loadPosts()
    })

    return {
      posts
    }
  }
}
</script>
