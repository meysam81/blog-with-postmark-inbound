<template>
  <div
    class="scroll-progress-container"
    role="progressbar"
    :aria-valuenow="scrollPercentage"
    aria-valuemin="0"
    aria-valuemax="100"
    aria-label="Page scroll progress"
  >
    <!-- Progress Bar -->
    <div
      class="scroll-progress-bar"
      :style="{ transform: `scaleX(${scrollPercentage / 100})` }"
    ></div>

    <!-- Scroll to Top Button -->
    <Transition name="scroll-top-fade">
      <button
        v-show="showScrollTop"
        @click="scrollToTop"
        class="scroll-to-top-btn"
        aria-label="Scroll to top of page"
        title="Back to top"
      >
        <svg
          class="w-5 h-5"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          aria-hidden="true"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M5 15l7-7 7 7"
          />
        </svg>
      </button>
    </Transition>
  </div>
</template>

<script>
import { ref, onMounted, onUnmounted } from 'vue'

export default {
  name: 'ScrollProgress',
  setup() {
    var scrollPercentage = ref(0)
    var showScrollTop = ref(false)

    function updateScrollProgress() {
      var scrollTop = window.pageYOffset || document.documentElement.scrollTop
      var scrollHeight = document.documentElement.scrollHeight - document.documentElement.clientHeight

      if (scrollHeight > 0) {
        scrollPercentage.value = Math.min((scrollTop / scrollHeight) * 100, 100)
      } else {
        scrollPercentage.value = 0
      }

      // Show scroll to top button when user scrolls down more than 20% of the page
      showScrollTop.value = scrollPercentage.value > 20
    }

    function scrollToTop() {
      window.scrollTo({
        top: 0,
        behavior: 'smooth'
      })

      // Focus management for accessibility
      var skipLink = document.querySelector('a[href="#main-content"]')
      if (skipLink) {
        skipLink.focus()
      }
    }

    function handleScroll() {
      // Use requestAnimationFrame for smooth performance
      requestAnimationFrame(updateScrollProgress)
    }

    onMounted(function() {
      window.addEventListener('scroll', handleScroll, { passive: true })
      window.addEventListener('resize', updateScrollProgress, { passive: true })
      updateScrollProgress()
    })

    onUnmounted(function() {
      window.removeEventListener('scroll', handleScroll)
      window.removeEventListener('resize', updateScrollProgress)
    })

    return {
      scrollPercentage,
      showScrollTop,
      scrollToTop
    }
  }
}
</script>

<style scoped>
.scroll-progress-container {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  z-index: 50;
  background: rgba(30, 41, 59, 0.3);
  backdrop-filter: blur(8px);
}

.scroll-progress-bar {
  height: 100%;
  width: 100%;
  background: linear-gradient(
    90deg,
    #10b981 0%,
    #3b82f6 50%,
    #8b5cf6 100%
  );
  transform-origin: left;
  transition: transform 0.1s ease-out;
  box-shadow: 0 0 10px rgba(16, 185, 129, 0.3);
}

.scroll-to-top-btn {
  position: fixed;
  bottom: 6rem;
  right: 2rem;
  width: 48px;
  height: 48px;
  background: linear-gradient(135deg, #1e293b, #334155);
  border: 1px solid rgba(148, 163, 184, 0.2);
  border-radius: 50%;
  color: #e2e8f0;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  backdrop-filter: blur(12px);
  box-shadow:
    0 4px 6px rgba(0, 0, 0, 0.1),
    0 1px 3px rgba(0, 0, 0, 0.08);
  z-index: 40;
}

.scroll-to-top-btn:hover {
  background: linear-gradient(135deg, #334155, #475569);
  color: #10b981;
  transform: translateY(-2px) scale(1.05);
  box-shadow:
    0 8px 12px rgba(0, 0, 0, 0.15),
    0 3px 6px rgba(0, 0, 0, 0.1),
    0 0 0 1px rgba(16, 185, 129, 0.2);
}

.scroll-to-top-btn:focus-visible {
  outline: 2px solid #10b981;
  outline-offset: 2px;
}

.scroll-to-top-btn:active {
  transform: translateY(-1px) scale(1.02);
}

/* Transition animations */
.scroll-top-fade-enter-active,
.scroll-top-fade-leave-active {
  transition: all 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275);
}

.scroll-top-fade-enter-from {
  opacity: 0;
  transform: translateY(20px) scale(0.8);
}

.scroll-top-fade-leave-to {
  opacity: 0;
  transform: translateY(20px) scale(0.8);
}

/* Reduced motion support */
@media (prefers-reduced-motion: reduce) {
  .scroll-progress-bar {
    transition: none;
  }

  .scroll-to-top-btn {
    transition: none;
  }

  .scroll-top-fade-enter-active,
  .scroll-top-fade-leave-active {
    transition: opacity 0.2s ease;
  }

  .scroll-top-fade-enter-from,
  .scroll-top-fade-leave-to {
    transform: none;
  }
}

/* High contrast mode support */
@media (prefers-contrast: high) {
  .scroll-progress-bar {
    background: #00ff00;
  }

  .scroll-to-top-btn {
    background: #000000;
    border: 2px solid #ffffff;
    color: #ffffff;
  }

  .scroll-to-top-btn:hover {
    background: #ffffff;
    color: #000000;
  }
}

/* Dark mode adjustments */
@media (prefers-color-scheme: dark) {
  .scroll-progress-container {
    background: rgba(15, 23, 42, 0.5);
  }
}
</style>
