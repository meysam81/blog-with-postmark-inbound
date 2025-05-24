<template>
  <header class="bg-white dark:bg-gray-800 shadow-md py-6">
    <div class="container mx-auto px-4 flex justify-between items-center">
      <h1 class="text-3xl font-bold font-display text-primary-600 dark:text-primary-400 block">
        <a class="no-underline" href="/">Tarzan</a>
      </h1>
      <div class="flex items-center space-x-4">
        <GitHubLink />
      </div>
    </div>
    <TitleAnimation />
    <p class="text-center text-gray-600 dark:text-gray-400 text-lg mt-2 italic font-light px-4">
      For Those Who Still Remember SMTP Commands
    </p>
  </header>
</template>

<script>
import GitHubLink from './GitHubLink.vue'
import TitleAnimation from './TitleAnimation.vue'

export default {
  name: 'AppHeader',
  components: {
    GitHubLink,
    TitleAnimation
  }
}
</script>
