<template>
  <div class="grid gap-8">
    <BlogCard
      v-for="post in posts"
      :key="post.id || post.title"
      :post="post"
    />
  </div>
</template>

<script>
import BlogCard from './BlogCard.vue'

export default {
  name: 'BlogPosts',
  components: {
    BlogCard
  },
  props: {
    posts: {
      type: Array,
      required: true
    }
  }
}
</script>
