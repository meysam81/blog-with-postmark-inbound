<template>
  <footer class="bg-white dark:bg-gray-800 py-8 mt-auto border-t border-gray-200 dark:border-gray-700">
    <div class="container mx-auto px-4">
      <div class="flex flex-col md:flex-row justify-between items-center">
        <div class="mb-4 md:mb-0">
          <p class="text-primary-600 dark:text-primary-400 font-semibold text-lg">
            Tarzan
          </p>
          <p class="text-gray-600 dark:text-gray-400 text-sm mt-1">
            Email-powered blogging for People Who Think GUIs Are Overrated
          </p>
          <p class="text-gray-500 dark:text-gray-400 text-xs mt-2">
            Because Real Geeks Send Blog Posts via Postfix
          </p>
        </div>
        <div class="flex flex-col space-y-2 text-center md:text-right">
          <a
            href="https://github.com/meysam81/tarzan/blob/main/LICENSE"
            class="text-gray-600 dark:text-gray-400 hover:text-primary-600 dark:hover:text-primary-400 focus-outline rounded"
          >
            Apache 2.0 License
          </a>
        </div>
      </div>
      <div class="mt-6 border-t border-gray-200 dark:border-gray-700 pt-6 text-center">
        <p class="text-gray-500 dark:text-gray-400 text-sm">
          © <span>{{ currentYear }}</span>
          <a
            href="https://github.com/meysam81"
            class="text-primary-600 dark:text-primary-400 hover:underline focus-outline rounded"
          >
            Meysam Azad
          </a>
          . All rights reserved.
        </p>
      </div>
    </div>
  </footer>
</template>

<script>
import { ref, onMounted } from 'vue'

export default {
  name: 'AppFooter',
  setup() {
    var currentYear = ref(new Date().getFullYear())

    onMounted(function() {
      currentYear.value = new Date().getFullYear()
    })

    return {
      currentYear
    }
  }
}
</script>
