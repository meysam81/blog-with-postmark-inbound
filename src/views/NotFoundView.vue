<template>
  <div class="min-h-screen relative overflow-hidden">
    <!-- Background Elements -->
    <div class="background-elements" aria-hidden="true">
      <div class="floating-shape shape-1"></div>
      <div class="floating-shape shape-2"></div>
      <div class="floating-shape shape-3"></div>
      <div class="gradient-orb orb-1"></div>
      <div class="gradient-orb orb-2"></div>
    </div>

    <!-- Main Content -->
    <main class="relative z-10 flex items-center justify-center min-h-screen px-6 py-16">
      <div class="max-w-4xl mx-auto text-center">
        <!-- 404 Animation Container -->
        <div class="mb-12">
          <!-- Large 404 Text -->
          <div class="relative mb-8">
            <h1 class="text-8xl md:text-9xl lg:text-[12rem] font-black text-transparent bg-clip-text bg-gradient-to-r from-emerald-400 via-blue-400 to-purple-400 leading-none tracking-tight font-display">
              <span class="block relative animate-swing">
                404
                <div class="absolute inset-0 text-emerald-400/20 blur-xl">404</div>
              </span>
            </h1>
          </div>

          <!-- Vine/Rope Animation -->
          <div class="flex justify-center mb-8">
            <div class="relative">
              <!-- Vine rope -->
              <div class="w-1 h-32 bg-gradient-to-b from-green-600 to-green-800 mx-auto animate-float"></div>

              <!-- Tarzan figure -->
              <div class="text-6xl animate-swing" style="animation-delay: 0.5s;">
                🦍
              </div>

              <!-- Floating leaves -->
              <div class="absolute -top-4 -left-8 text-2xl animate-float" style="animation-delay: 1s;">🍃</div>
              <div class="absolute -top-2 -right-6 text-xl animate-float" style="animation-delay: 2s;">🍃</div>
              <div class="absolute top-8 -left-6 text-lg animate-float" style="animation-delay: 1.5s;">🍃</div>
            </div>
          </div>
        </div>

        <!-- Error Message -->
        <div class="mb-12">
          <h2 class="text-3xl md:text-5xl lg:text-6xl font-bold text-slate-100 leading-tight mb-6">
            <span class="block">Oops! This page has</span>
            <span class="text-transparent bg-clip-text bg-gradient-to-r from-emerald-300 to-cyan-300 font-black drop-shadow-sm">
              swung away
            </span>
          </h2>

          <p class="text-lg md:text-xl text-slate-300 leading-relaxed max-w-2xl mx-auto mb-8">
            The page you're looking for seems to have vanished into the digital jungle. Let's get you back to safety.
          </p>

          <!-- Status Message -->
          <div class="bg-gradient-to-r from-red-900/40 to-orange-900/40 backdrop-blur-sm border border-red-500/30 rounded-xl p-4 mb-8 max-w-xl mx-auto">
            <div class="flex items-center justify-center gap-2 mb-2">
              <div class="w-2 h-2 bg-red-400 rounded-full animate-pulse"></div>
              <span class="text-red-300 font-semibold text-sm">PAGE NOT FOUND</span>
              <div class="w-2 h-2 bg-red-400 rounded-full animate-pulse"></div>
            </div>
            <p class="text-slate-200 text-sm">
              The requested page was not found!
            </p>
          </div>
        </div>

        <!-- Action Buttons -->
        <div class="flex flex-col sm:flex-row gap-4 justify-center items-center mb-12">
          <router-link
            to="/"
            class="group relative px-8 py-4 bg-gradient-to-r from-emerald-600 to-cyan-600 hover:from-emerald-500 hover:to-cyan-500 text-white font-semibold rounded-xl transition-all duration-300 transform hover:scale-105 hover:shadow-2xl hover:shadow-emerald-500/25 focus:outline-none focus:ring-4 focus:ring-emerald-500/50"
          >
            <span class="relative z-10 flex items-center gap-2">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"/>
              </svg>
              Return Home
            </span>
            <div class="absolute inset-0 bg-gradient-to-r from-emerald-600/50 to-cyan-600/50 rounded-xl blur opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
          </router-link>
        </div>

        <!-- Helpful Links -->
        <div class="border-t border-slate-800/50 pt-8">
          <p class="text-slate-400 mb-4">Looking for something specific?</p>
          <div class="flex flex-wrap justify-center gap-4 text-sm">
            <router-link to="/" class="text-emerald-400 hover:text-emerald-300 transition-colors duration-200">
              Home
            </router-link>
            <span class="text-slate-600">•</span>
            <a href="/rss.xml" class="text-emerald-400 hover:text-emerald-300 transition-colors duration-200">
              RSS Feed
            </a>
            <span class="text-slate-600">•</span>
            <a href="https://github.com/meysam81/tarzan" class="text-emerald-400 hover:text-emerald-300 transition-colors duration-200">
              GitHub
            </a>
          </div>
        </div>
      </div>
    </main>

    <!-- Footer -->
    <AppFooter />

    <!-- Scroll Progress Indicator -->
    <ScrollProgress />
  </div>
</template>

<script>
import AppFooter from '@/components/AppFooter.vue'
import ScrollProgress from '@/components/ScrollProgress.vue'
import { useAccessibility } from '@/utils/accessibility.js'

export default {
  name: 'NotFoundView',
  components: {
    AppFooter,
    ScrollProgress
  },
  setup() {
    var { announceToScreenReader } = useAccessibility()

    function goBack() {
      if (window.history.length > 1) {
        window.history.back()
      } else {
        this.$router.push('/')
      }
    }

    return {
      goBack,
      announceToScreenReader
    }
  },
  mounted: function setNotFoundTitle() {
    document.title = '404 - Page Not Found | Tarzan'

    // Announce to screen readers
    this.announceToScreenReader('Page not found. You are now on the 404 error page.', 'polite')
  },
  beforeRouteEnter: function beforeEnterNotFound(to, from, next) {
    document.title = '404 - Page Not Found | Tarzan'
    next()
  }
}
</script>

<style scoped>
/* Background Elements */
.background-elements {
  position: fixed;
  inset: 0;
  pointer-events: none;
  z-index: 0;
}

.floating-shape {
  position: absolute;
  border-radius: 50%;
  background: linear-gradient(45deg, rgba(59, 130, 246, 0.1), rgba(139, 92, 246, 0.1));
  animation: float 8s ease-in-out infinite;
}

.shape-1 {
  width: 300px;
  height: 300px;
  top: 10%;
  left: 10%;
  animation-delay: 0s;
}

.shape-2 {
  width: 200px;
  height: 200px;
  top: 60%;
  right: 15%;
  animation-delay: 2s;
}

.shape-3 {
  width: 150px;
  height: 150px;
  bottom: 20%;
  left: 70%;
  animation-delay: 4s;
}

.gradient-orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(40px);
  animation: pulse 6s ease-in-out infinite;
}

.orb-1 {
  width: 400px;
  height: 400px;
  background: radial-gradient(circle, rgba(16, 185, 129, 0.15) 0%, transparent 70%);
  top: 20%;
  left: 20%;
  animation-delay: 1s;
}

.orb-2 {
  width: 300px;
  height: 300px;
  background: radial-gradient(circle, rgba(139, 92, 246, 0.15) 0%, transparent 70%);
  bottom: 30%;
  right: 20%;
  animation-delay: 3s;
}

/* Custom animations */
@keyframes swing {
  0%, 100% { transform: rotate(-5deg) scale(1); }
  50% { transform: rotate(5deg) scale(1.05); }
}

@keyframes float {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-20px); }
}

.animate-swing {
  animation: swing 3s ease-in-out infinite;
}

.animate-float {
  animation: float 6s ease-in-out infinite;
}

/* Reduced motion support */
@media (prefers-reduced-motion: reduce) {
  .floating-shape,
  .gradient-orb,
  .animate-swing,
  .animate-float {
    animation: none;
  }
}
</style>
