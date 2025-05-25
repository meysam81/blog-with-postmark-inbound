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
              Where Words Swing Free
            </p>
            <div class="h-px bg-gradient-to-r from-transparent via-slate-400 to-transparent flex-1 max-w-32"></div>
          </div>
        </div>

        <!-- Main Tagline -->
        <div class="mb-12">
          <h2 class="text-2xl md:text-4xl lg:text-5xl font-bold text-slate-100 leading-tight mb-6">
            <span class="block">Transform Your</span>
            <span class="text-transparent bg-clip-text bg-gradient-to-r from-emerald-300 to-cyan-300 font-black drop-shadow-sm">
              Inbox
            </span>
            <span class="block">Into a Publishing Powerhouse</span>
          </h2>

          <p class="text-lg md:text-xl text-slate-300 leading-relaxed max-w-3xl mx-auto mb-8">
            Experience the revolution in content creation. No interfaces, no complexity—just pure literary freedom through the elegance of email.
          </p>

          <!-- Publishing Instructions -->
          <div
            class="bg-gradient-to-r from-emerald-900/30 to-blue-900/30 backdrop-blur-sm border border-emerald-500/30 rounded-2xl p-6 mb-10 max-w-2xl mx-auto"
            aria-label="Publishing instructions"
            tabindex="0"
          >
            <div class="flex items-center justify-center gap-3 mb-4">
              <div class="w-2 h-2 bg-emerald-400 rounded-full animate-pulse"></div>
              <h3 class="text-lg font-semibold text-emerald-300">Dead Simple Publishing</h3>
              <div class="w-2 h-2 bg-emerald-400 rounded-full animate-pulse"></div>
            </div>

            <p class="text-slate-200 mb-4 leading-relaxed">
              Send an email to <strong class="text-emerald-300 font-mono">blog@tarzan.meysam.io</strong> and see your content live in seconds
            </p>

            <div class="flex items-center justify-center gap-2">
              <button
                @click="copyEmailAddress"
                :class="{ 'copied': emailCopied }"
                class="group inline-flex items-center gap-2 bg-emerald-600/20 hover:bg-emerald-600/30 border border-emerald-500/40 text-emerald-300 px-4 py-2 rounded-lg text-sm font-medium transition-all duration-300 hover:scale-105"
                :aria-label="emailCopied ? 'Email address copied!' : 'Copy email address'"
              >
                <svg v-if="!emailCopied" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"/>
                </svg>
                <svg v-else class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
                </svg>
                <span>{{ emailCopied ? 'Copied!' : 'Copy Address' }}</span>
              </button>
            </div>
          </div>

          <!-- Feature Highlights -->
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6 max-w-4xl mx-auto text-sm">
            <div class="bg-slate-800/30 backdrop-blur-sm border border-slate-700/50 rounded-xl p-4">
              <div class="text-emerald-400 mb-2">
                <svg class="w-6 h-6 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
                </svg>
              </div>
              <h3 class="font-semibold text-slate-200 mb-1">Email-Powered Publishing</h3>
              <p class="text-slate-400 text-xs">Send an email, publish a post. It's that revolutionary.</p>
            </div>

            <div class="bg-slate-800/30 backdrop-blur-sm border border-slate-700/50 rounded-xl p-4">
              <div class="text-blue-400 mb-2">
                <svg class="w-6 h-6 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"></path>
                </svg>
              </div>
              <h3 class="font-semibold text-slate-200 mb-1">Developer-First Experience</h3>
              <p class="text-slate-400 text-xs">Built for those who think in code and write in markdown.</p>
            </div>

            <div class="bg-slate-800/30 backdrop-blur-sm border border-slate-700/50 rounded-xl p-4">
              <div class="text-purple-400 mb-2">
                <svg class="w-6 h-6 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
                </svg>
              </div>
              <h3 class="font-semibold text-slate-200 mb-1">Lightning Fast</h3>
              <p class="text-slate-400 text-xs">From compose to publish in seconds, not minutes.</p>
            </div>
          </div>
        </div>

        <!-- Navigation Actions -->
        <div class="flex flex-col sm:flex-row items-center justify-center gap-4">
          <GitHubLink />
          <button
            @click="scrollToContent"
            class="group bg-gradient-to-r from-emerald-500 to-blue-500 hover:from-emerald-400 hover:to-blue-400 text-white font-semibold px-8 py-3 rounded-full transition-all duration-300 transform hover:scale-105 hover:shadow-xl hover:shadow-emerald-500/25"
            aria-label="Explore published stories"
          >
            <span class="flex items-center gap-2">
              Explore Stories
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

export default {
  name: 'AppHeader',
  components: {
    GitHubLink,
    TitleAnimation
  },
  setup() {
    var emailCopied = ref(false)
    var emailAddress = 'blog@tarzan.meysam.io'

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
        announcement.textContent = "Email address copied to clipboard"
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
</style>
