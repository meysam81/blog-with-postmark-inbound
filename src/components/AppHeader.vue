<template>
  <header
    class="relative z-20 backdrop-blur-xl bg-slate-950/30 border-b border-slate-800/50"
    role="banner"
  >
    <!-- Hero Section -->
    <div class="container mx-auto px-6 py-16 md:py-24">
      <div class="text-center max-w-5xl mx-auto">
        <!-- Logo & Brand -->
        <div class="mb-8">
          <h1 class="text-6xl md:text-8xl lg:text-9xl font-black text-transparent bg-clip-text bg-gradient-to-r from-emerald-400 via-blue-400 to-purple-400 leading-none tracking-tight font-display mb-4">
            <span class="block relative">
              TARZAN
              <div class="absolute inset-0 text-emerald-400/20 blur-xl">TARZAN</div>
            </span>
          </h1>
          <div class="flex items-center justify-center gap-4 mb-6">
            <div class="h-px bg-gradient-to-r from-transparent via-slate-400 to-transparent flex-1 max-w-32"></div>
            <p class="text-lg md:text-xl text-slate-300 font-light tracking-wider uppercase">
              Email-to-Blog Publishing
            </p>
            <div class="h-px bg-gradient-to-r from-transparent via-slate-400 to-transparent flex-1 max-w-32"></div>
          </div>
        </div>

        <!-- Main Tagline -->
        <div class="mb-12">
          <h2 class="text-2xl md:text-4xl lg:text-5xl font-bold text-slate-100 leading-tight mb-6">
            <span class="block">Email</span>
            <span class="text-transparent bg-clip-text bg-gradient-to-r from-emerald-300 to-cyan-300 font-black drop-shadow-sm">
              Becomes Blog
            </span>
            <span class="block text-xl md:text-2xl lg:text-3xl text-slate-300 font-normal mt-2">Hit send. Post goes live.</span>
          </h2>

          <p class="text-lg md:text-xl text-slate-300 leading-relaxed max-w-3xl mx-auto mb-8">
            Send email. Get blog post. That's it. No accounts, no setup, no complexity.
          </p>

          <!-- Real-time Feature Highlight -->
          <div class="bg-gradient-to-r from-blue-900/40 to-purple-900/40 backdrop-blur-sm border border-blue-500/30 rounded-xl p-4 mb-8 max-w-xl mx-auto">
            <div class="flex items-center justify-center gap-2 mb-2">
              <div class="w-2 h-2 bg-blue-400 rounded-full animate-pulse"></div>
              <span class="text-blue-300 font-semibold text-sm">INSTANT PUBLISHING</span>
              <div class="w-2 h-2 bg-blue-400 rounded-full animate-pulse"></div>
            </div>
            <p class="text-slate-200 text-sm">
              Email goes out, blog post goes live. Watch it happen in real time.
            </p>
          </div>

          <!-- Email Instructions -->
          <div
            class="bg-gradient-to-r from-emerald-900/30 to-blue-900/30 backdrop-blur-sm border border-emerald-500/30 rounded-2xl p-6 mb-10 max-w-2xl mx-auto"
            aria-label="Email publishing instructions"
            tabindex="0"
          >
            <div class="flex items-center justify-center gap-3 mb-4">
              <div class="w-2 h-2 bg-emerald-400 rounded-full animate-pulse"></div>
              <h3 class="text-lg font-semibold text-emerald-300">Start Publishing</h3>
              <div class="w-2 h-2 bg-emerald-400 rounded-full animate-pulse"></div>
            </div>

            <p class="text-slate-200 mb-4 leading-relaxed">
              Email <strong class="text-emerald-300 font-mono text-lg">{{ emailAddress }}</strong> and your post goes live instantly.
            </p>

            <div class="bg-slate-800/40 border border-slate-600/30 rounded-lg p-3 mb-4">
              <div class="flex items-center gap-2 mb-1">
                <svg class="w-4 h-4 text-blue-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/>
                </svg>
                <span class="text-blue-300 text-sm font-medium">Markdown Ready</span>
              </div>
              <p class="text-slate-300 text-sm">Use **bold**, *italic*, # headers for rich formatting</p>
            </div>

            <div class="flex items-center justify-center gap-2 relative">
              <button
                @click="copyEmailAddress"
                :class="{
                  'copied': emailCopied,
                  'success-pulse': emailCopied
                }"
                class="group inline-flex items-center gap-2 bg-emerald-600/20 hover:bg-emerald-600/30 border border-emerald-500/40 text-emerald-300 px-6 py-3 rounded-lg font-medium transition-all duration-300 hover:scale-105 relative overflow-hidden"
                :aria-label="emailCopied ? 'Email copied!' : 'Copy email'"
              >
                <svg v-if="!emailCopied" class="w-4 h-4 transition-transform duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"/>
                </svg>
                <svg v-else class="w-4 h-4 transition-all duration-200 scale-110" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
                </svg>
                <span class="transition-all duration-200">{{ emailCopied ? 'Copied!' : 'Copy Email' }}</span>

                <!-- Success ripple effect -->
                <div v-if="emailCopied" class="absolute inset-0 bg-emerald-400/30 rounded-lg animate-ping"></div>
              </button>

              <!-- Success tooltip -->
              <div
                v-if="emailCopied"
                class="email-copy-tooltip"
              >
                ✓ Email copied!
                <div class="absolute -top-1 left-1/2 transform -translate-x-1/2 w-2 h-2 bg-emerald-900/90 border-l border-t border-emerald-500/50 rotate-45"></div>
              </div>
            </div>
          </div>

          <!-- Feature Highlights -->
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6 max-w-4xl mx-auto text-sm">
            <div class="bg-slate-800/30 backdrop-blur-sm border border-slate-700/50 rounded-xl p-4">
              <div class="text-emerald-400 mb-2">
                <svg class="w-6 h-6 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7"></path>
                </svg>
              </div>
              <h3 class="font-semibold text-slate-200 mb-1">Email to Web</h3>
              <p class="text-slate-400 text-xs">Send email, get live blog post in seconds</p>
            </div>

            <div class="bg-slate-800/30 backdrop-blur-sm border border-slate-700/50 rounded-xl p-4">
              <div class="text-blue-400 mb-2">
                <svg class="w-6 h-6 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 5.636l-3.536 3.536m0 5.656l3.536 3.536M9.172 9.172L5.636 5.636m3.536 9.192L5.636 18.364M12 2.25a9.75 9.75 0 109.75 9.75A9.75 9.75 0 0012 2.25z"></path>
                </svg>
              </div>
              <h3 class="font-semibold text-slate-200 mb-1">Zero Setup</h3>
              <p class="text-slate-400 text-xs">No accounts. No dashboards. Just email.</p>
            </div>

            <div class="bg-slate-800/30 backdrop-blur-sm border border-slate-700/50 rounded-xl p-4">
              <div class="text-purple-400 mb-2">
                <svg class="w-6 h-6 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"></path>
                </svg>
              </div>
              <h3 class="font-semibold text-slate-200 mb-1">Your Control</h3>
              <p class="text-slate-400 text-xs">Open source. Self-host anywhere.</p>
            </div>
          </div>
        </div>

        <!-- Navigation Actions -->
        <div class="flex flex-col sm:flex-row items-center justify-center gap-4">
          <GitHubLink />            <button
            @click="scrollToContent"
            class="group bg-gradient-to-r from-emerald-500 to-blue-500 hover:from-emerald-400 hover:to-blue-400 text-white font-semibold px-8 py-3 rounded-full transition-all duration-300 transform hover:scale-105 hover:shadow-xl hover:shadow-emerald-500/25"
            aria-label="View published blog posts"
          >
            <span class="flex items-center gap-2">
              See Live Blog Posts
              <svg class="w-4 h-4 group-hover:translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14l-7 7m0 0l-7-7m7 7V3"></path>
              </svg>
            </span>
          </button>
        </div>
      </div>
    </div>

    <!-- Animated Title Effect -->
    <TitleAnimation />

    <!-- Scroll Indicator -->
    <div class="absolute bottom-8 left-1/2 transform -translate-x-1/2 animate-bounce">
      <div class="w-6 h-10 border-2 border-slate-400 rounded-full flex justify-center">
        <div class="w-1 h-3 bg-slate-400 rounded-full mt-2 animate-pulse"></div>
      </div>
    </div>
  </header>
</template>

<script>
import { ref } from 'vue'
import GitHubLink from './GitHubLink.vue'
import TitleAnimation from './TitleAnimation.vue'
import { VITE_INBOUND_EMAIL_ADDRESS } from '@/utils/env.js'

export default {
  name: 'AppHeader',
  components: {
    GitHubLink,
    TitleAnimation
  },
  setup() {
    var emailCopied = ref(false)
    var emailAddress = ref(VITE_INBOUND_EMAIL_ADDRESS)

    function scrollToContent() {
      var mainContent = document.getElementById('main-content')
      if (mainContent) {
        mainContent.scrollIntoView({
          behavior: 'smooth',
          block: 'start'
        })
      }
    }

    async function copyEmailAddress() {
      try {
        await navigator.clipboard.writeText(emailAddress)
        emailCopied.value = true

        setTimeout(function resetCopyState() {
          emailCopied.value = false
        }, 2000)

        var announcement = document.createElement("div")
        announcement.setAttribute("aria-live", "polite")
        announcement.className = "sr-only"
        announcement.textContent = "Email copied to clipboard"
        document.body.appendChild(announcement)

        setTimeout(function removeAnnouncement() {
          document.body.removeChild(announcement)
        }, 3000)
      } catch (error) {
        try {
          var textArea = document.createElement("textarea")
          textArea.value = emailAddress
          document.body.appendChild(textArea)
          textArea.select()
          document.execCommand('copy')
          document.body.removeChild(textArea)

          emailCopied.value = true
          setTimeout(function resetFallbackCopyState() {
            emailCopied.value = false
          }, 2000)
        } catch (fallbackError) {
          // Silently fail for production
        }
      }
    }

    return {
      emailCopied,
      emailAddress,
      scrollToContent,
      copyEmailAddress
    }
  }
}
</script>

<style scoped>
/* Screen reader only content */
.sr-only {
  position: absolute;
  width: 1px;
  height: 1px;
  padding: 0;
  margin: -1px;
  overflow: hidden;
  clip: rect(0, 0, 0, 0);
  white-space: nowrap;
  border: 0;
}

/* Success button styles */
.copied {
  background: rgba(16, 185, 129, 0.3) !important;
  border-color: rgba(16, 185, 129, 0.6) !important;
  color: rgb(52, 211, 153) !important;
}

.success-pulse {
  animation: successPulse 0.6s ease-out;
}

@keyframes successPulse {
  0% {
    transform: scale(1);
    box-shadow: 0 0 0 0 rgba(16, 185, 129, 0.7);
  }
  50% {
    transform: scale(1.05);
    box-shadow: 0 0 0 10px rgba(16, 185, 129, 0);
  }
  100% {
    transform: scale(1);
    box-shadow: 0 0 0 0 rgba(16, 185, 129, 0);
  }
}

@keyframes fade-in {
  0% {
    opacity: 0;
    transform: translateX(-50%) translateY(-5px);
  }
  100% {
    opacity: 1;
    transform: translateX(-50%) translateY(0);
  }
}

.animate-fade-in {
  animation: fade-in 0.2s ease-out;
}

.email-copy-tooltip {
  position: absolute;
  top: 100%;
  left: 50%;
  transform: translateX(-50%);
  margin-top: 8px;
  background: rgba(6, 78, 59, 0.9);
  backdrop-filter: blur(4px);
  border: 1px solid rgba(16, 185, 129, 0.5);
  color: rgb(167, 243, 208);
  padding: 8px 12px;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 500;
  white-space: nowrap;
  z-index: 10;
  opacity: 0;
  animation: tooltipFadeIn 0.2s ease-out forwards;
}

@keyframes tooltipFadeIn {
  0% {
    opacity: 0;
    transform: translateX(-50%) translateY(-5px);
  }
  100% {
    opacity: 1;
    transform: translateX(-50%) translateY(0);
  }
}
</style>
