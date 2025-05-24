<template>
  <div class="text-center py-16 px-4">
    <svg
      class="w-16 h-16 mx-auto text-gray-400"
      fill="none"
      stroke="currentColor"
      viewBox="0 0 24 24"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10"
      />
    </svg>
    <h2 class="mt-4 text-2xl font-semibold text-gray-700 dark:text-gray-300">
      No posts yet
    </h2>
    <p class="mt-2 text-gray-500 dark:text-gray-400 max-w-md mx-auto">
      Be the first to publish something amazing! Check out the publishing
      instructions below.
    </p>
  </div>
</template>

<script>
export default {
  name: 'EmptyState'
}
</script>
