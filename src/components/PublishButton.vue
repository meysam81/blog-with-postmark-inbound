<template>
  <section class="publish-section" role="region" aria-label="Publishing options">
    <!-- Main Publish Button -->
    <div class="publish-button-container">
      <button
        @click="togglePublishInfo"
        class="publish-main-btn"
        :class="{ 'active': isPublishInfoVisible }"
        :aria-expanded="isPublishInfoVisible"
        aria-haspopup="true"
        aria-controls="publishInfoPanel"
        ref="publishButton"
      >
        <!-- Button Background Effects -->
        <div class="btn-background" aria-hidden="true"></div>
        <div class="btn-glow" aria-hidden="true"></div>

        <!-- Button Content -->
        <div class="btn-content">
          <div class="btn-icon-container">
            <svg
              class="btn-icon"
              :class="{ 'rotate': isPublishInfoVisible }"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              aria-hidden="true"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 4v16m8-8H4"
              />
            </svg>
          </div>

          <div class="btn-text">
            <span class="btn-main-text">Start Publishing</span>
            <span class="btn-sub-text">Send an email, get a blog post</span>
          </div>

          <div class="btn-arrow" aria-hidden="true">
            <svg
              class="arrow-icon"
              :class="{ 'expanded': isPublishInfoVisible }"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M19 9l-7 7-7-7"
              />
            </svg>
          </div>
        </div>

        <!-- Floating Particles -->
        <div class="btn-particles" aria-hidden="true">
          <div class="particle particle-1"></div>
          <div class="particle particle-2"></div>
          <div class="particle particle-3"></div>
        </div>
      </button>
    </div>

    <!-- Expanded Publish Panel -->
    <Transition name="panel-slide">
      <div
        v-show="isPublishInfoVisible"
        id="publishInfoPanel"
        ref="publishInfoPanel"
        class="publish-info-panel"
        :aria-hidden="!isPublishInfoVisible"
        role="dialog"
        aria-labelledby="publishInfoTitle"
        aria-describedby="publishInfoDescription"
      >
        <!-- Panel Background -->
        <div class="panel-background" aria-hidden="true"></div>

        <!-- Panel Content -->
        <div class="panel-content">
          <!-- Header -->
          <header class="panel-header">
            <div class="header-icon">
              <svg viewBox="0 0 24 24" fill="currentColor">
                <path d="M2.01 21L23 12 2.01 3 2 10l15 2-15 2z"/>
              </svg>
            </div>

            <div class="header-text">
              <h3 id="publishInfoTitle" class="panel-title">
                Email Publishing Made Simple
              </h3>
              <p id="publishInfoDescription" class="panel-description">
                Transform any email into a beautiful blog post instantly. No registration required.
              </p>
            </div>
          </header>

          <!-- Steps Container -->
          <div class="steps-container">
            <!-- Step 1 -->
            <div class="step-card step-1">
              <div class="step-indicator">
                <div class="step-number">1</div>
                <div class="step-connector" aria-hidden="true"></div>
              </div>

              <div class="step-content">
                <h4 class="step-title">
                  <svg class="step-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"/>
                  </svg>
                  Write Your Blog Post
                </h4>

                <p class="step-description">
                  Compose your content in any email client.
                  The subject becomes your post title, and the body becomes your content.
                </p>

                <div class="step-highlight">
                  <svg class="highlight-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"/>
                  </svg>
                  <span>Tip: Use HTML formatting for rich content and styling!</span>
                </div>
              </div>
            </div>

            <!-- Step 2 -->
            <div class="step-card step-2">
              <div class="step-indicator">
                <div class="step-number">2</div>
              </div>

              <div class="step-content">
                <h4 class="step-title">
                  <svg class="step-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"/>
                  </svg>
                  Send to Blog Address
                </h4>

                <p class="step-description">
                  Send your email to the blog address and watch your content transform into a beautiful blog post instantly.
                </p>

                <div class="email-container">
                  <div class="email-display">
                    <div class="email-icon">
                      <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 12a4 4 0 10-8 0 4 4 0 008 0zm0 0v1.5a2.5 2.5 0 005 0V12a9 9 0 10-9 9m4.5-1.206a8.959 8.959 0 01-4.5 1.207"/>
                      </svg>
                    </div>

                    <div class="email-address">
                      <span class="email-text">{{ emailToCopy }}</span>
                    </div>

                    <button
                      @click="copyEmailToClipboard"
                      class="copy-btn"
                      :class="{
                        'copied': emailCopied,
                        'success-pulse': emailCopied
                      }"
                      :title="copyButtonTitle"
                      :aria-label="copyButtonTitle"
                      ref="copyEmailBtn"
                    >
                      <svg
                        v-if="!emailCopied"
                        class="copy-icon transition-transform duration-200"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                      >
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"/>
                      </svg>

                      <svg
                        v-else
                        class="check-icon transition-all duration-200 scale-110"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                      >
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
                      </svg>

                      <span class="copy-text transition-all duration-200">
                        {{ emailCopied ? 'Copied!' : 'Copy' }}
                      </span>

                      <!-- Success ripple effect -->
                      <div v-if="emailCopied" class="absolute inset-0 bg-white/20 rounded-md animate-ping"></div>
                    </button>
                  </div>

                  <!-- Success notification -->
                  <div
                    v-if="emailCopied"
                    class="success-notification"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
                    </svg>
                    <span>Email address copied successfully!</span>
                  </div>

                  <div class="email-note">
                    <svg class="note-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
                    </svg>
                    <span>Your blog post will appear here instantly!</span>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Action Footer -->
          <div class="panel-footer">
            <div class="footer-content">
              <button @click="closePublishInfo" class="close-btn">
                <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
                </svg>
                <span>Close</span>
              </button>

              <div class="footer-links">
                <a href="#documentation" class="footer-link">
                  <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253"/>
                  </svg>
                  <span>Documentation</span>
                </a>

                <a href="#examples" class="footer-link">
                  <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"/>
                  </svg>
                  <span>Examples</span>
                </a>
              </div>
            </div>
          </div>
        </div>

        <!-- Decorative Elements -->
        <div class="panel-decorations" aria-hidden="true">
          <div class="deco-shape deco-shape-1"></div>
          <div class="deco-shape deco-shape-2"></div>
          <div class="deco-line"></div>
        </div>
      </div>
    </Transition>

    <!-- Overlay -->
    <Transition name="overlay-fade">
      <div
        v-show="isPublishInfoVisible"
        class="panel-overlay"
        @click="closePublishInfo"
        aria-hidden="true"
      ></div>
    </Transition>
  </section>
</template>

<script>
import { ref, onMounted, onUnmounted } from 'vue'
import { VITE_INBOUND_EMAIL_ADDRESS } from '@/utils/env.js'

export default {
  name: 'PublishButton',
  setup() {
    var isPublishInfoVisible = ref(false)
    var emailToCopy = ref(VITE_INBOUND_EMAIL_ADDRESS)
    var emailCopied = ref(false)
    var copyButtonTitle = ref("Copy email address")
    var publishInfoPanel = ref(null)
    var copyEmailBtn = ref(null)
    var publishButton = ref(null)

    function togglePublishInfo() {
      isPublishInfoVisible.value = !isPublishInfoVisible.value

      if (isPublishInfoVisible.value) {
        setTimeout(function focusPanel() {
          if (publishInfoPanel.value) {
            publishInfoPanel.value.focus()
          }
        }, 100)
      }
    }

    function closePublishInfo() {
      isPublishInfoVisible.value = false

      if (publishButton.value) {
        publishButton.value.focus()
      }
    }

    async function copyEmailToClipboard() {
      try {
        await navigator.clipboard.writeText(emailToCopy.value)

        emailCopied.value = true
        copyButtonTitle.value = "Copied!"

        setTimeout(function resetCopyState() {
          emailCopied.value = false
          copyButtonTitle.value = "Copy email address"
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
          textArea.value = emailToCopy.value
          document.body.appendChild(textArea)
          textArea.select()
          document.execCommand('copy')
          document.body.removeChild(textArea)

          emailCopied.value = true
          copyButtonTitle.value = "Copied!"

          setTimeout(function resetFallbackCopyState() {
            emailCopied.value = false
            copyButtonTitle.value = "Copy email address"
          }, 2000)
        } catch (fallbackError) {
          // Silently fail for production
        }
      }
    }

    function handleClickOutside(event) {
      if (!publishInfoPanel.value) {
        return
      }

      var isClickInside = publishInfoPanel.value.contains(event.target) ||
                         event.target.closest('.publish-main-btn')

      if (!isClickInside && isPublishInfoVisible.value) {
        closePublishInfo()
      }
    }

    function handleKeydown(event) {
      if (event.key === "Escape" && isPublishInfoVisible.value) {
        closePublishInfo()
      }
    }

    onMounted(function setupEventListeners() {
      document.addEventListener("click", handleClickOutside)
      document.addEventListener("keydown", handleKeydown)
    })

    onUnmounted(function cleanupEventListeners() {
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
      publishButton,
      togglePublishInfo,
      closePublishInfo,
      copyEmailToClipboard
    }
  }
}
</script>

<style scoped>
/* Main Publish Button */
.publish-section {
  position: relative;
  z-index: 30;
}

.publish-button-container {
  position: relative;
  margin: 2rem 0;
}

.publish-main-btn {
  position: relative;
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 1.25rem 2rem;
  background: linear-gradient(135deg, var(--primary-600) 0%, var(--accent-600) 100%);
  color: white;
  border: none;
  border-radius: var(--radius-2xl);
  font-size: var(--text-lg);
  font-weight: var(--font-semibold);
  cursor: pointer;
  transition: all var(--transition-normal);
  box-shadow: var(--shadow-xl);
  overflow: hidden;
  min-width: 280px;
  z-index: 1;
}

.publish-main-btn:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-2xl);
  background: linear-gradient(135deg, var(--primary-700) 0%, var(--accent-700) 100%);
}

.publish-main-btn:focus {
  outline: none;
  box-shadow: 0 0 0 4px var(--primary-200), var(--shadow-2xl);
}

.publish-main-btn.active {
  background: linear-gradient(135deg, var(--accent-600) 0%, var(--primary-600) 100%);
}

/* Button Background Effects */
.btn-background {
  position: absolute;
  inset: 0;
  background: linear-gradient(45deg, transparent 30%, rgba(255, 255, 255, 0.1) 50%, transparent 70%);
  transform: translateX(-100%);
  transition: transform 0.6s ease;
}

.publish-main-btn:hover .btn-background {
  transform: translateX(100%);
}

.btn-glow {
  position: absolute;
  inset: -2px;
  background: linear-gradient(135deg, var(--primary-400), var(--accent-400));
  border-radius: var(--radius-2xl);
  opacity: 0;
  transition: opacity var(--transition-normal);
  z-index: -1;
  filter: blur(8px);
}

.publish-main-btn:hover .btn-glow {
  opacity: 0.6;
}

/* Button Content */
.btn-content {
  display: flex;
  align-items: center;
  gap: 1rem;
  z-index: 2;
  width: 100%;
}

.btn-icon-container {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 2.5rem;
  height: 2.5rem;
  background: rgba(255, 255, 255, 0.2);
  border-radius: var(--radius-lg);
  backdrop-filter: blur(10px);
}

.btn-icon {
  width: 1.5rem;
  height: 1.5rem;
  transition: transform var(--transition-normal);
}

.btn-icon.rotate {
  transform: rotate(45deg);
}

.btn-text {
  flex: 1;
  text-align: left;
}

.btn-main-text {
  display: block;
  font-size: var(--text-lg);
  font-weight: var(--font-bold);
  line-height: 1.2;
}

.btn-sub-text {
  display: block;
  font-size: var(--text-sm);
  opacity: 0.9;
  margin-top: 0.25rem;
}

.btn-arrow {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 2rem;
  height: 2rem;
}

.arrow-icon {
  width: 1.25rem;
  height: 1.25rem;
  transition: transform var(--transition-normal);
}

.arrow-icon.expanded {
  transform: rotate(180deg);
}

/* Button Particles */
.btn-particles {
  position: absolute;
  inset: 0;
  pointer-events: none;
  overflow: hidden;
  border-radius: var(--radius-2xl);
}

.particle {
  position: absolute;
  width: 4px;
  height: 4px;
  background: rgba(255, 255, 255, 0.6);
  border-radius: 50%;
  animation: float 3s ease-in-out infinite;
}

.particle-1 {
  top: 20%;
  left: 10%;
  animation-delay: 0s;
}

.particle-2 {
  top: 60%;
  right: 15%;
  animation-delay: 1s;
}

.particle-3 {
  bottom: 25%;
  left: 70%;
  animation-delay: 2s;
}

/* Publish Info Panel */
.publish-info-panel {
  position: absolute;
  top: calc(100% + 1rem);
  left: 0;
  right: 0;
  background: var(--surface);
  border-radius: var(--radius-2xl);
  padding: 2rem;
  box-shadow: var(--shadow-2xl);
  border: 1px solid var(--border-light);
  z-index: 40;
  max-width: 600px;
  margin: 0 auto;
  backdrop-filter: blur(20px);
  background: rgba(var(--surface-rgb), 0.95);
}

/* Panel Transitions */
.panel-slide-enter-active,
.panel-slide-leave-active {
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

.panel-slide-enter-from {
  opacity: 0;
  transform: translateY(-20px) scale(0.95);
}

.panel-slide-leave-to {
  opacity: 0;
  transform: translateY(-20px) scale(0.95);
}

/* Panel Background */
.panel-background {
  position: absolute;
  inset: 0;
  background: linear-gradient(135deg, var(--primary-50) 0%, var(--accent-50) 100%);
  border-radius: var(--radius-2xl);
  opacity: 0.5;
}

.panel-content {
  position: relative;
  z-index: 2;
}

/* Panel Header */
.panel-header {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 2rem;
  padding-bottom: 1.5rem;
  border-bottom: 1px solid var(--border-light);
}

.header-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 3rem;
  height: 3rem;
  background: linear-gradient(135deg, var(--primary-600), var(--accent-600));
  color: white;
  border-radius: var(--radius-xl);
  flex-shrink: 0;
}

.header-icon svg {
  width: 1.5rem;
  height: 1.5rem;
}

.header-text {
  flex: 1;
}

.panel-title {
  font-size: var(--text-xl);
  font-weight: var(--font-bold);
  color: var(--text-primary);
  margin: 0 0 0.5rem 0;
  background: linear-gradient(135deg, var(--primary-600), var(--accent-600));
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.panel-description {
  color: var(--text-secondary);
  margin: 0;
  font-size: var(--text-base);
}

/* Steps Container */
.steps-container {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
  margin-bottom: 2rem;
}

.step-card {
  display: flex;
  gap: 1rem;
  padding: 1.5rem;
  background: var(--surface);
  border-radius: var(--radius-xl);
  border: 1px solid var(--border-light);
  transition: all var(--transition-normal);
  position: relative;
  overflow: hidden;
}

.step-card::before {
  content: '';
  position: absolute;
  inset: 0;
  background: linear-gradient(135deg, var(--primary-50), var(--accent-50));
  opacity: 0;
  transition: opacity var(--transition-normal);
}

.step-card:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-lg);
  border-color: var(--primary-200);
}

.step-card:hover::before {
  opacity: 1;
}

/* Step Indicator */
.step-indicator {
  position: relative;
  z-index: 2;
  flex-shrink: 0;
}

.step-number {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 2.5rem;
  height: 2.5rem;
  background: linear-gradient(135deg, var(--primary-600), var(--accent-600));
  color: white;
  border-radius: 50%;
  font-weight: var(--font-bold);
  font-size: var(--text-sm);
}

.step-connector {
  position: absolute;
  top: 100%;
  left: 50%;
  transform: translateX(-50%);
  width: 2px;
  height: 1.5rem;
  background: linear-gradient(to bottom, var(--primary-300), transparent);
}

.step-2 .step-connector {
  display: none;
}

/* Step Content */
.step-content {
  flex: 1;
  position: relative;
  z-index: 2;
}

.step-title {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  font-size: var(--text-lg);
  font-weight: var(--font-semibold);
  color: var(--text-primary);
  margin: 0 0 0.75rem 0;
}

.step-icon {
  width: 1.25rem;
  height: 1.25rem;
  color: var(--primary-600);
}

.step-description {
  color: var(--text-secondary);
  margin: 0 0 1rem 0;
  line-height: 1.6;
}

.step-highlight {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem 1rem;
  background: var(--warning-50);
  border: 1px solid var(--warning-200);
  border-radius: var(--radius-lg);
  font-size: var(--text-sm);
  color: var(--warning-800);
}

.highlight-icon {
  width: 1rem;
  height: 1rem;
  color: var(--warning-600);
  flex-shrink: 0;
}

/* Email Container */
.email-container {
  margin-top: 1rem;
}

.email-display {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 1rem;
  background: var(--gray-50);
  border: 1px solid var(--border-light);
  border-radius: var(--radius-lg);
  margin-bottom: 0.75rem;
}

.email-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 2rem;
  height: 2rem;
  color: var(--primary-600);
  flex-shrink: 0;
}

.email-icon svg {
  width: 1.25rem;
  height: 1.25rem;
}

.email-address {
  flex: 1;
}

.email-text {
  font-family: var(--font-mono);
  font-size: var(--text-sm);
  color: var(--text-primary);
  font-weight: var(--font-medium);
}

.copy-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 0.75rem;
  background: var(--primary-600);
  color: white;
  border: none;
  border-radius: var(--radius-md);
  font-size: var(--text-sm);
  font-weight: var(--font-medium);
  cursor: pointer;
  transition: all var(--transition-normal);
  position: relative;
  overflow: hidden;
}

.copy-btn:hover {
  background: var(--primary-700);
  transform: translateY(-1px);
}

.copy-btn:focus {
  outline: none;
  box-shadow: 0 0 0 3px var(--primary-200);
}

.copy-btn.copied {
  background: var(--success-600);
}

.copy-btn.success-pulse {
  animation: successPulse 0.6s ease-out;
}

@keyframes successPulse {
  0% {
    transform: scale(1) translateY(-1px);
    box-shadow: 0 0 0 0 var(--success-600);
  }
  50% {
    transform: scale(1.05) translateY(-1px);
    box-shadow: 0 0 0 10px transparent;
  }
  100% {
    transform: scale(1) translateY(-1px);
    box-shadow: 0 0 0 0 transparent;
  }
}

.copy-icon,
.check-icon {
  width: 1rem;
  height: 1rem;
}

.copy-text {
  transition: all var(--transition-normal);
}

.email-note {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem;
  background: var(--info-50);
  border: 1px solid var(--info-200);
  border-radius: var(--radius-md);
  font-size: var(--text-sm);
  color: var(--info-800);
}

.note-icon {
  width: 1rem;
  height: 1rem;
  color: var(--info-600);
  flex-shrink: 0;
}

/* Success notification */
.success-notification {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem 1rem;
  background: var(--success-50);
  border: 1px solid var(--success-200);
  border-radius: var(--radius-md);
  font-size: var(--text-sm);
  color: var(--success-800);
  margin-top: 0.75rem;
  animation: slideInFromTop 0.3s ease-out, fadeOut 0.3s ease-in 1.7s forwards;
}

@keyframes slideInFromTop {
  from {
    opacity: 0;
    transform: translateY(-10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes fadeOut {
  from {
    opacity: 1;
    transform: translateY(0);
  }
  to {
    opacity: 0;
    transform: translateY(-5px);
  }
}

/* Panel Footer */
.panel-footer {
  padding-top: 1.5rem;
  border-top: 1px solid var(--border-light);
}

.footer-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
}

.close-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem 1rem;
  background: var(--gray-100);
  color: var(--text-secondary);
  border: none;
  border-radius: var(--radius-md);
  font-size: var(--text-sm);
  cursor: pointer;
  transition: all var(--transition-normal);
}

.close-btn:hover {
  background: var(--gray-200);
  color: var(--text-primary);
}

.close-btn svg {
  width: 1rem;
  height: 1rem;
}

.footer-links {
  display: flex;
  gap: 1rem;
}

.footer-link {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 0.75rem;
  color: var(--text-secondary);
  text-decoration: none;
  font-size: var(--text-sm);
  border-radius: var(--radius-md);
  transition: all var(--transition-normal);
}

.footer-link:hover {
  background: var(--primary-50);
  color: var(--primary-700);
}

.footer-link svg {
  width: 1rem;
  height: 1rem;
}

/* Panel Decorations */
.panel-decorations {
  position: absolute;
  inset: 0;
  pointer-events: none;
  border-radius: var(--radius-2xl);
  overflow: hidden;
}

.deco-shape {
  position: absolute;
  background: linear-gradient(135deg, var(--primary-100), var(--accent-100));
  border-radius: 50%;
  opacity: 0.6;
}

.deco-shape-1 {
  width: 100px;
  height: 100px;
  top: -20px;
  right: -20px;
  animation: float 6s ease-in-out infinite;
}

.deco-shape-2 {
  width: 60px;
  height: 60px;
  bottom: -10px;
  left: -10px;
  animation: float 4s ease-in-out infinite reverse;
}

.deco-line {
  position: absolute;
  top: 50%;
  right: 0;
  width: 100px;
  height: 1px;
  background: linear-gradient(to right, transparent, var(--primary-300), transparent);
  transform: translateY(-50%);
}

/* Panel Overlay */
.panel-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.3);
  backdrop-filter: blur(2px);
  z-index: 35;
}

.overlay-fade-enter-active,
.overlay-fade-leave-active {
  transition: all var(--transition-normal);
}

.overlay-fade-enter-from,
.overlay-fade-leave-to {
  opacity: 0;
}

/* Animations */
@keyframes float {
  0%, 100% {
    transform: translateY(0px) rotate(0deg);
  }
  50% {
    transform: translateY(-10px) rotate(180deg);
  }
}

/* Responsive Design */
@media (max-width: 768px) {
  .publish-main-btn {
    flex-direction: column;
    text-align: center;
    padding: 1rem;
    min-width: auto;
  }

  .btn-content {
    flex-direction: column;
    gap: 0.75rem;
  }

  .btn-text {
    text-align: center;
  }

  .publish-info-panel {
    position: fixed;
    top: 50%;
    left: 1rem;
    right: 1rem;
    transform: translateY(-50%);
    max-height: 80vh;
    overflow-y: auto;
    margin: 0;
  }

  .panel-slide-enter-from,
  .panel-slide-leave-to {
    transform: translateY(-50%) scale(0.95);
  }

  .step-card {
    flex-direction: column;
    text-align: center;
  }

  .step-indicator {
    align-self: center;
  }

  .footer-content {
    flex-direction: column;
    gap: 1rem;
  }

  .footer-links {
    justify-content: center;
  }
}

/* Reduced Motion */
@media (prefers-reduced-motion: reduce) {
  .publish-main-btn,
  .step-card,
  .copy-btn,
  .btn-icon,
  .arrow-icon,
  .btn-background,
  .particle {
    transition: none;
    animation: none;
  }

  .btn-icon.rotate,
  .arrow-icon.expanded {
    transform: none;
  }
}

/* High Contrast Mode */
@media (prefers-contrast: high) {
  .publish-main-btn {
    border: 2px solid currentColor;
  }

  .step-card {
    border-width: 2px;
  }

  .email-display,
  .step-highlight,
  .email-note {
    border-width: 2px;
  }
}

/* Dark Mode */
@media (prefers-color-scheme: dark) {
  .panel-background {
    background: linear-gradient(135deg, var(--gray-800) 0%, var(--gray-900) 100%);
  }

  .email-display {
    background: var(--gray-800);
  }

  .step-highlight {
    background: rgba(251, 191, 36, 0.1);
    border-color: rgba(251, 191, 36, 0.3);
    color: rgb(251, 191, 36);
  }

  .email-note {
    background: rgba(59, 130, 246, 0.1);
    border-color: rgba(59, 130, 246, 0.3);
    color: rgb(147, 197, 253);
  }

  .close-btn {
    background: var(--gray-700);
  }

  .close-btn:hover {
    background: var(--gray-600);
  }

  .footer-link:hover {
    background: var(--gray-700);
  }
}
</style>
