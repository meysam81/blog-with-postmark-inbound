<template>
  <article class="blog-card bg-white dark:bg-gray-800 rounded-xl shadow-md p-8 fade-in border-l-4 border-primary-500">
    <h2 class="text-2xl md:text-3xl font-semibold text-gray-900 dark:text-gray-100 mb-3 font-display">
      {{ post.title }}
    </h2>
    <div class="prose prose-green dark:prose-invert max-w-none my-4" v-html="post.content">
    </div>
    <div class="flex items-center mt-6 pt-4 border-t border-gray-100 dark:border-gray-700">
      <div class="bg-primary-100 dark:bg-primary-900 text-primary-700 dark:text-primary-300 rounded-full h-10 w-10 flex items-center justify-center font-medium">
        {{ getAuthorInitial(post['author-email']) }}
      </div>
      <div class="ml-3">
        <p class="text-sm text-gray-700 dark:text-gray-300 font-medium">
          {{ post['author-email'] }}
        </p>
        <p class="text-xs text-gray-500 dark:text-gray-400">
          <time :datetime="formatISODate(post['created-at'])">
            {{ formatReadableDate(post['created-at']) }}
          </time>
        </p>
      </div>
    </div>
  </article>
</template>

<script>
import { formatISODate, formatReadableDate, getAuthorInitial } from '../utils/api.js'

export default {
  name: 'BlogCard',
  props: {
    post: {
      type: Object,
      required: true
    }
  },
  setup() {
    return {
      formatISODate,
      formatReadableDate,
      getAuthorInitial
    }
  }
}
</script>
