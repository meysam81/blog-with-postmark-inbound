<template>
  <div>
    <!-- Publish Button -->
    <button
      @click="togglePublishInfo"
      class="publish-btn bg-gradient-to-br from-primary-500 to-secondary-600 hover:from-primary-600 hover:to-secondary-700 text-white font-bold py-3 px-6 rounded-full flex items-center gap-2"
      :aria-expanded="isPublishInfoVisible"
      aria-haspopup="true"
      aria-controls="publishInfoPanel"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-5 w-5"
        viewBox="0 0 20 20"
        fill="currentColor"
        aria-hidden="true"
      >
        <path
          fill-rule="evenodd"
          d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-11a1 1 0 10-2 0v2H7a1 1 0 100 2h2v2a1 1 0 102 0v-2h2a1 1 0 100-2h-2V7z"
          clip-rule="evenodd"
        />
      </svg>
      Share Your Story
    </button>

    <!-- Publish Info Panel -->
    <div
      id="publishInfoPanel"
      ref="publishInfoPanel"
      :class="['publish-info-panel', 'bg-white', 'dark:bg-gray-800', 'rounded-xl', 'shadow-xl', 'p-6', 'border', 'border-gray-200', 'dark:border-gray-700', { 'active': isPublishInfoVisible }]"
      :aria-hidden="!isPublishInfoVisible"
      role="dialog"
      aria-labelledby="publishInfoTitle"
    >
      <h3
        id="publishInfoTitle"
        class="text-xl font-bold text-gray-900 dark:text-white mb-5 font-display"
      >
        Publish Your Story in Two Simple Steps
      </h3>

      <div class="space-y-6">
        <!-- Step 1 -->
        <div class="flex gap-4">
          <div
            class="step-number bg-primary-100 dark:bg-primary-900 text-primary-700 dark:text-primary-300"
            aria-hidden="true"
          >
            1
          </div>
          <div>
            <h4 class="font-semibold text-gray-800 dark:text-gray-200 mb-2 text-lg">
              Get authorized as a contributor
            </h4>
            <p class="text-gray-600 dark:text-gray-400 text-sm mb-3">
              Open a quick pull request to join our writing community:
            </p>
            <a
              href="https://github.com/meysam81/tarzan/pulls/new?quick_pull=1"
              class="inline-flex items-center gap-2 text-sm bg-gray-100 dark:bg-gray-700 py-2 px-4 rounded-md text-primary-700 dark:text-primary-300 hover:bg-gray-200 dark:hover:bg-gray-600 transition-colors focus-outline"
              target="_blank"
              rel="noopener noreferrer"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-4 w-4"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                aria-hidden="true"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"
                />
              </svg>
              Submit Authorization Request
            </a>
          </div>
        </div>

        <!-- Step 2 -->
        <div class="flex gap-4">
          <div
            class="step-number bg-secondary-100 dark:bg-secondary-900 text-secondary-700 dark:text-secondary-300"
            aria-hidden="true"
          >
            2
          </div>
          <div>
            <h4 class="font-semibold text-gray-800 dark:text-gray-200 mb-2 text-lg">
              Write and send your post by email
            </h4>
            <p class="text-gray-600 dark:text-gray-400 text-sm mb-3">
              Once approved, compose your post and send it to:
            </p>
            <div class="bg-gray-100 dark:bg-gray-700 py-3 px-4 rounded-md flex justify-between items-center">
              <code class="text-sm text-secondary-700 dark:text-secondary-300 font-medium">
                {{ emailToCopy }}
              </code>
              <button
                @click="copyEmailToClipboard"
                ref="copyEmailBtn"
                class="text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200 focus-outline rounded p-1"
                :title="copyButtonTitle"
                aria-label="Copy email address to clipboard"
              >
                <svg
                  v-if="!emailCopied"
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-5 w-5"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                  aria-hidden="true"
                >
                  <path d="M8 3a1 1 0 011-1h2a1 1 0 110 2H9a1 1 0 01-1-1z" />
                  <path
                    d="M6 3a2 2 0 00-2 2v11a2 2 0 002 2h8a2 2 0 002-2V5a2 2 0 00-2-2 3 3 0 01-3 3H9a3 3 0 01-3-3z"
                  />
                </svg>
                <svg
                  v-else
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-5 w-5 text-primary-500"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                  aria-hidden="true"
                >
                  <path
                    fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd"
                  />
                </svg>
              </button>
            </div>
            <div class="mt-3 p-3 bg-primary-50 dark:bg-gray-900 rounded-md border border-primary-100 dark:border-gray-700">
              <p class="text-gray-700 dark:text-gray-300 text-sm mb-2 font-medium">
                Email Format:
              </p>
              <ul class="text-gray-600 dark:text-gray-400 text-sm space-y-2 ml-5 list-disc">
                <li>
                  <span class="font-medium">Subject:</span> Your post title
                </li>
                <li>
                  <span class="font-medium">Body:</span> Your post content (markdown supported)
                </li>
              </ul>
            </div>
          </div>
        </div>
      </div>

      <!-- Close Button -->
      <button
        @click="closePublishInfo"
        class="absolute top-3 right-3 text-gray-400 hover:text-gray-600 dark:hover:text-gray-200 p-1 focus-outline rounded"
        aria-label="Close publishing information"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-5 w-5"
          viewBox="0 0 20 20"
          fill="currentColor"
          aria-hidden="true"
        >
          <path
            fill-rule="evenodd"
            d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
            clip-rule="evenodd"
          />
        </svg>
      </button>
    </div>
  </div>
</template>

<script>
import { ref, onMounted, onUnmounted } from 'vue'

export default {
  name: 'PublishButton',
  setup() {
    var isPublishInfoVisible = ref(false)
    var emailToCopy = ref("blog@tarzan.meysam.io")
    var emailCopied = ref(false)
    var copyButtonTitle = ref("Copy email address")
    var publishInfoPanel = ref(null)
    var copyEmailBtn = ref(null)

    function togglePublishInfo() {
      isPublishInfoVisible.value = !isPublishInfoVisible.value
    }

    function closePublishInfo() {
      isPublishInfoVisible.value = false
    }

    async function copyEmailToClipboard() {
      try {
        await navigator.clipboard.writeText(emailToCopy.value)

        emailCopied.value = true
        copyButtonTitle.value = "Copied!"

        setTimeout(function() {
          emailCopied.value = false
          copyButtonTitle.value = "Copy email address"
        }, 2000)

        var announcement = document.createElement("div")
        announcement.setAttribute("aria-live", "polite")
        announcement.className = "sr-only"
        announcement.textContent = "Email address copied to clipboard"
        document.body.appendChild(announcement)

        setTimeout(function() {
          document.body.removeChild(announcement)
        }, 3000)
      } catch (error) {
        console.error("Failed to copy email:", error)
      }
    }

    function handleClickOutside(event) {
      if (!publishInfoPanel.value) return

      var isClickInside = publishInfoPanel.value.contains(event.target) ||
                         event.target.closest('.publish-btn')

      if (!isClickInside && isPublishInfoVisible.value) {
        closePublishInfo()
      }
    }

    function handleKeydown(event) {
      if (event.key === "Escape" && isPublishInfoVisible.value) {
        closePublishInfo()
      }
    }

    onMounted(function() {
      document.addEventListener("click", handleClickOutside)
      document.addEventListener("keydown", handleKeydown)
    })

    onUnmounted(function() {
      document.removeEventListener("click", handleClickOutside)
      document.removeEventListener("keydown", handleKeydown)
    })

    return {
      isPublishInfoVisible,
      emailToCopy,
      emailCopied,
      copyButtonTitle,
      publishInfoPanel,
      copyEmailBtn,
      togglePublishInfo,
      closePublishInfo,
      copyEmailToClipboard
    }
  }
}
</script>
