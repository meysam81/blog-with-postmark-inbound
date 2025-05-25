<template>
  <section
    class="empty-state-container"
    role="region"
    aria-label="No blog posts available"
  >
    <!-- Background Elements -->
    <div class="background-elements" aria-hidden="true">
      <div class="floating-shape shape-1"></div>
      <div class="floating-shape shape-2"></div>
      <div class="floating-shape shape-3"></div>
      <div class="gradient-orb orb-1"></div>
      <div class="gradient-orb orb-2"></div>
    </div>

    <!-- Main Content -->
    <div class="empty-state-content">
      <!-- Animated Icon -->
      <div class="icon-container">
        <div class="icon-backdrop" aria-hidden="true"></div>
        <svg
          class="main-icon"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
          aria-hidden="true"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253"
          />
        </svg>

        <!-- Floating Particles -->
        <div class="particles" aria-hidden="true">
          <div class="particle particle-1"></div>
          <div class="particle particle-2"></div>
          <div class="particle particle-3"></div>
          <div class="particle particle-4"></div>
        </div>
      </div>

      <!-- Text Content -->
      <div class="text-content">
        <h2 class="empty-title">
          <span class="title-primary">No Blog Posts Yet</span>
          <span class="title-secondary">Start Publishing with a Single Email</span>
        </h2>

        <p class="empty-description">
          Transform any email into a beautiful blog post instantly. Send an email to our blog address and watch your content go live in seconds—no complex setup required.
        </p>

        <!-- Action Buttons -->
        <div class="action-buttons">
          <button class="primary-action-btn" @click="handleSendEmail">
            <svg class="btn-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 7.89a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"/>
            </svg>
            <span>Send Email</span>
          </button>

          <button class="secondary-action-btn" @click="handleShowInstructions">
            <svg class="btn-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
            </svg>
            <span>How It Works</span>
          </button>
        </div>
      </div>

      <!-- Decorative Elements -->
      <div class="decorative-elements" aria-hidden="true">
        <div class="deco-line deco-line-1"></div>
        <div class="deco-line deco-line-2"></div>
        <div class="deco-dot deco-dot-1"></div>
        <div class="deco-dot deco-dot-2"></div>
        <div class="deco-dot deco-dot-3"></div>
      </div>
    </div>

    <!-- Call to Action Footer -->
    <div class="cta-footer">
      <div class="cta-content">
        <p class="cta-text">
          Ready to publish your first post? Send an email and see the magic happen ✨
        </p>
      </div>
    </div>

    <!-- Instructions Modal -->
    <div
      v-if="isModalOpen"
      class="instructions-modal-backdrop"
      role="dialog"
      aria-labelledby="instructions-title"
      aria-describedby="instructions-description"
      @click="handleBackdropClick"
      @keydown.esc="closeInstructionsModal"
    >
      <div class="instructions-modal" @click.stop>
        <div class="instructions-header">
          <h3 id="instructions-title" class="instructions-title">
            <svg class="title-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 7.89a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"/>
            </svg>
            How Email Blogging Works
          </h3>
          <button class="close-btn" aria-label="Close instructions" @click="closeInstructionsModal">
            <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
            </svg>
          </button>
        </div>

        <div class="instructions-content">
          <p id="instructions-description" class="instructions-intro">
            Publishing a blog post is as simple as sending an email. Here's how:
          </p>

          <div class="instructions-steps">
            <div class="instruction-step">
              <div class="step-number">1</div>
              <div class="step-content">
                <h4>Send an Email</h4>
                <p>Compose your blog post in an email and send it to: <strong>{{ blogEmail }}</strong></p>
                <button
                  class="copy-email-btn"
                  :class="{ 'copied': emailCopied, 'success-pulse': emailCopied }"
                  @click="copyBlogEmail"
                  :aria-label="emailCopied ? 'Email address copied!' : 'Copy email address'"
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
                    class="copy-icon transition-all duration-200 scale-110"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
                  </svg>
                  <span class="transition-all duration-200">{{ emailCopied ? 'Copied!' : 'Copy Email' }}</span>

                  <div v-if="emailCopied" class="absolute inset-0 bg-emerald-400/20 rounded-lg animate-ping"></div>

                  <div
                    v-if="emailCopied"
                    class="email-copy-tooltip"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
                    </svg>
                    <span>Email copied to clipboard!</span>
                  </div>
                </button>
              </div>
            </div>

            <div class="instruction-step">
              <div class="step-number">2</div>
              <div class="step-content">
                <h4>Automatic Publishing</h4>
                <p>Your email will be automatically converted into a beautiful blog post and published instantly.</p>
              </div>
            </div>

            <div class="instruction-step">
              <div class="step-number">3</div>
              <div class="step-content">
                <h4>Share & Enjoy</h4>
                <p>Your blog post is now live and ready to be shared with the world!</p>
              </div>
            </div>
          </div>

          <div class="instructions-footer">
            <button class="try-now-btn" @click="handleTryNow">
              <svg class="btn-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 7.89a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"/>
              </svg>
              Try It Now
            </button>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import { ref } from 'vue'

export default {
  name: 'EmptyState',
  setup() {
    var blogEmail = 'blog@tarzan.meysam.io'
    var isModalOpen = ref(false)
    var emailCopied = ref(false)

    function handleSendEmail() {
      var subject = encodeURIComponent('My First Blog Post')
      var body = encodeURIComponent('Write your blog post content here...\n\nThis email will become a beautiful blog post! Markdown is supported.')
      var mailtoUrl = 'mailto:' + blogEmail + '?subject=' + subject + '&body=' + body

      window.location.href = mailtoUrl
    }

    function handleShowInstructions() {
      isModalOpen.value = true
    }

    function closeInstructionsModal() {
      isModalOpen.value = false
    }

    function handleBackdropClick(event) {
      if (event.target.classList.contains('instructions-modal-backdrop')) {
        closeInstructionsModal()
      }
    }

    function copyBlogEmail() {
      navigator.clipboard.writeText(blogEmail).then(function onSuccess() {
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
      }).catch(function onError() {
        // Fallback for older browsers
        var textArea = document.createElement('textarea')
        textArea.value = blogEmail
        document.body.appendChild(textArea)
        textArea.select()
        document.execCommand('copy')
        document.body.removeChild(textArea)

        emailCopied.value = true
        setTimeout(function resetFallbackCopyState() {
          emailCopied.value = false
        }, 2000)
      })
    }

    function handleTryNow() {
      var subject = encodeURIComponent('My First Blog Post')
      var body = encodeURIComponent('Write your blog post content here...\n\nThis email will become a beautiful blog post!')
      var mailtoUrl = 'mailto:' + blogEmail + '?subject=' + subject + '&body=' + body

      window.location.href = mailtoUrl
      closeInstructionsModal()
    }

    return {
      blogEmail,
      isModalOpen,
      emailCopied,
      handleSendEmail,
      handleShowInstructions,
      closeInstructionsModal,
      handleBackdropClick,
      copyBlogEmail,
      handleTryNow
    }
  }
}
</script>

<style scoped>
.empty-state-container {
  position: relative;
  min-height: 600px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 4rem 2rem;
  overflow: hidden;
  background: linear-gradient(
    135deg,
    rgba(255, 255, 255, 0.9) 0%,
    rgba(248, 250, 252, 0.95) 100%
  );
  border-radius: 24px;
  margin: 2rem 0;
}

/* Background Elements */
.background-elements {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  pointer-events: none;
  z-index: 1;
}

.floating-shape {
  position: absolute;
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(16, 185, 129, 0.1), rgba(59, 130, 246, 0.1));
  animation: gentleFloat 15s ease-in-out infinite;
}

.shape-1 {
  width: 200px;
  height: 200px;
  top: 10%;
  left: 5%;
  animation-delay: 0s;
}

.shape-2 {
  width: 120px;
  height: 120px;
  top: 70%;
  right: 10%;
  animation-delay: -5s;
}

.shape-3 {
  width: 80px;
  height: 80px;
  bottom: 20%;
  left: 15%;
  animation-delay: -10s;
}

.gradient-orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(60px);
  opacity: 0.4;
  animation: pulse 12s ease-in-out infinite;
}

.orb-1 {
  width: 300px;
  height: 300px;
  background: radial-gradient(circle, rgba(16, 185, 129, 0.3), transparent);
  top: 20%;
  right: 20%;
  animation-delay: -2s;
}

.orb-2 {
  width: 250px;
  height: 250px;
  background: radial-gradient(circle, rgba(139, 92, 246, 0.3), transparent);
  bottom: 30%;
  left: 20%;
  animation-delay: -7s;
}

/* Main Content */
.empty-state-content {
  position: relative;
  z-index: 3;
  text-align: center;
  max-width: 600px;
}

/* Icon Container */
.icon-container {
  position: relative;
  margin-bottom: 3rem;
  display: inline-block;
}

.icon-backdrop {
  position: absolute;
  top: -20px;
  left: -20px;
  right: -20px;
  bottom: -20px;
  background: linear-gradient(135deg, rgba(16, 185, 129, 0.1), rgba(59, 130, 246, 0.1));
  border-radius: 50%;
  animation: iconPulse 4s ease-in-out infinite;
}

.main-icon {
  width: 80px;
  height: 80px;
  color: #10b981;
  position: relative;
  z-index: 2;
  filter: drop-shadow(0 4px 8px rgba(16, 185, 129, 0.3));
  animation: iconFloat 6s ease-in-out infinite;
}

.particles {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  pointer-events: none;
}

.particle {
  position: absolute;
  width: 6px;
  height: 6px;
  background: #10b981;
  border-radius: 50%;
  animation: particleFloat 8s ease-in-out infinite;
}

.particle-1 {
  top: 20%;
  left: 10%;
  animation-delay: 0s;
}

.particle-2 {
  top: 30%;
  right: 15%;
  animation-delay: -2s;
}

.particle-3 {
  bottom: 25%;
  left: 20%;
  animation-delay: -4s;
}

.particle-4 {
  bottom: 35%;
  right: 10%;
  animation-delay: -6s;
}

/* Text Content */
.text-content {
  margin-bottom: 3rem;
}

.empty-title {
  margin-bottom: 1.5rem;
  line-height: 1.2;
}

.title-primary {
  display: block;
  font-size: 2.5rem;
  font-weight: 700;
  background: linear-gradient(135deg, #1e293b, #475569);
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  margin-bottom: 0.5rem;
}

.title-secondary {
  display: block;
  font-size: 1.25rem;
  font-weight: 500;
  color: #64748b;
  opacity: 0.8;
}

.empty-description {
  font-size: 1.125rem;
  line-height: 1.7;
  color: #64748b;
  max-width: 500px;
  margin: 0 auto 2rem auto;
}

/* Action Buttons */
.action-buttons {
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
}

.primary-action-btn,
.secondary-action-btn {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  padding: 1rem 2rem;
  border-radius: 16px;
  font-weight: 600;
  font-size: 1rem;
  transition: all 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  border: none;
  cursor: pointer;
  text-decoration: none;
}

.primary-action-btn {
  background: linear-gradient(135deg, #10b981, #059669);
  color: white;
  box-shadow:
    0 4px 6px rgba(16, 185, 129, 0.25),
    0 1px 3px rgba(16, 185, 129, 0.12);
}

.primary-action-btn:hover {
  background: linear-gradient(135deg, #059669, #047857);
  transform: translateY(-2px);
  box-shadow:
    0 8px 12px rgba(16, 185, 129, 0.3),
    0 2px 6px rgba(16, 185, 129, 0.15);
}

.secondary-action-btn {
  background: rgba(255, 255, 255, 0.8);
  color: #374151;
  border: 2px solid rgba(148, 163, 184, 0.2);
  backdrop-filter: blur(8px);
}

.secondary-action-btn:hover {
  background: rgba(255, 255, 255, 0.95);
  transform: translateY(-2px);
  box-shadow:
    0 4px 8px rgba(0, 0, 0, 0.1),
    0 1px 3px rgba(0, 0, 0, 0.05);
}

.btn-icon {
  width: 20px;
  height: 20px;
  transition: transform 0.3s ease;
}

.primary-action-btn:hover .btn-icon,
.secondary-action-btn:hover .btn-icon {
  transform: scale(1.1);
}

/* Decorative Elements */
.decorative-elements {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  pointer-events: none;
  z-index: 2;
}

.deco-line {
  position: absolute;
  height: 2px;
  background: linear-gradient(90deg, transparent, rgba(16, 185, 129, 0.3), transparent);
  animation: lineGlow 8s ease-in-out infinite;
}

.deco-line-1 {
  width: 150px;
  top: 30%;
  left: 10%;
  animation-delay: 0s;
}

.deco-line-2 {
  width: 100px;
  bottom: 40%;
  right: 15%;
  animation-delay: -4s;
}

.deco-dot {
  position: absolute;
  width: 8px;
  height: 8px;
  background: #10b981;
  border-radius: 50%;
  opacity: 0.6;
  animation: dotPulse 6s ease-in-out infinite;
}

.deco-dot-1 {
  top: 25%;
  right: 20%;
  animation-delay: 0s;
}

.deco-dot-2 {
  top: 60%;
  left: 25%;
  animation-delay: -2s;
}

.deco-dot-3 {
  bottom: 30%;
  right: 30%;
  animation-delay: -4s;
}

/* CTA Footer */
.cta-footer {
  position: relative;
  z-index: 3;
  margin-top: 2rem;
  padding: 1.5rem;
  background: rgba(255, 255, 255, 0.6);
  border-radius: 16px;
  border: 1px solid rgba(148, 163, 184, 0.2);
  backdrop-filter: blur(8px);
}

.cta-text {
  font-size: 1rem;
  color: #64748b;
  margin: 0;
  font-weight: 500;
}

/* Instructions Modal Styles */
.instructions-modal-backdrop {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.75);
  backdrop-filter: blur(8px);
  z-index: 1000;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 2rem;
  animation: modalBackdropFadeIn 0.3s ease-out;
}

.instructions-modal {
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.95), rgba(248, 250, 252, 0.98));
  border-radius: 24px;
  border: 1px solid rgba(148, 163, 184, 0.2);
  backdrop-filter: blur(16px);
  max-width: 600px;
  width: 100%;
  max-height: 90vh;
  overflow-y: auto;
  animation: modalSlideIn 0.3s ease-out;
  box-shadow:
    0 25px 50px -12px rgba(0, 0, 0, 0.25),
    0 4px 6px -2px rgba(0, 0, 0, 0.05);
}

.instructions-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 2rem 2rem 1rem 2rem;
  border-bottom: 1px solid rgba(148, 163, 184, 0.1);
}

.instructions-title {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  font-size: 1.5rem;
  font-weight: 700;
  color: #1e293b;
  margin: 0;
}

.title-icon {
  width: 24px;
  height: 24px;
  color: #10b981;
}

.close-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
  border: none;
  background: rgba(148, 163, 184, 0.1);
  border-radius: 12px;
  color: #64748b;
  cursor: pointer;
  transition: all 0.2s ease;
}

.close-btn:hover {
  background: rgba(239, 68, 68, 0.1);
  color: #ef4444;
}

.close-btn svg {
  width: 20px;
  height: 20px;
}

.instructions-content {
  padding: 1.5rem 2rem 2rem 2rem;
}

.instructions-intro {
  font-size: 1.125rem;
  color: #64748b;
  margin-bottom: 2rem;
  line-height: 1.6;
  text-align: center;
}

.instructions-steps {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
  margin-bottom: 2rem;
}

.instruction-step {
  display: flex;
  gap: 1rem;
  align-items: flex-start;
}

.step-number {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background: linear-gradient(135deg, #10b981, #059669);
  color: white;
  font-weight: 700;
  font-size: 0.875rem;
  flex-shrink: 0;
}

.step-content {
  position: relative;
}

.step-content h4 {
  font-size: 1.125rem;
  font-weight: 600;
  color: #1e293b;
  margin: 0 0 0.5rem 0;
}

.step-content p {
  color: #64748b;
  margin: 0;
  line-height: 1.5;
}

.copy-email-btn {
  display: inline-flex;
  align-items: center;
  gap: 0.25rem;
  margin-top: 0.5rem;
  padding: 0.5rem 0.75rem;
  background: rgba(16, 185, 129, 0.1);
  border: 1px solid rgba(16, 185, 129, 0.2);
  border-radius: 8px;
  color: #059669;
  font-size: 0.875rem;
  cursor: pointer;
  transition: all 0.3s ease;
  position: relative;
  overflow: visible;
}

.copy-email-btn:hover {
  background: rgba(16, 185, 129, 0.2);
  transform: scale(1.05);
}

.copy-email-btn:hover:not(.copied) {
  transform: scale(1.05);
}

.copy-email-btn.copied {
  background: rgba(16, 185, 129, 0.2);
  border-color: rgba(16, 185, 129, 0.4);
  color: #047857;
}

.copy-email-btn.copied:hover {
  transform: none;
}

.copy-email-btn.success-pulse {
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

.copy-email-btn svg {
  width: 16px;
  height: 16px;
}

/* Email copy success tooltip */
.email-copy-tooltip {
  position: absolute;
  top: 100%;
  left: 50%;
  transform: translateX(-50%);
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-top: 0.5rem;
  padding: 0.5rem 0.75rem;
  background: rgba(16, 185, 129, 0.95);
  border: 1px solid rgba(16, 185, 129, 0.3);
  border-radius: 8px;
  color: white;
  font-size: 0.875rem;
  white-space: nowrap;
  z-index: 50;
  opacity: 0;
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3);
  animation: tooltipFadeIn 0.2s ease-out forwards;
}

.email-copy-tooltip::before {
  content: '';
  position: absolute;
  bottom: 100%;
  left: 50%;
  transform: translateX(-50%);
  border: 4px solid transparent;
  border-bottom-color: rgba(16, 185, 129, 0.95);
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

.instructions-footer {
  text-align: center;
  padding-top: 1.5rem;
  border-top: 1px solid rgba(148, 163, 184, 0.1);
}

.try-now-btn {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.875rem 1.75rem;
  background: linear-gradient(135deg, #10b981, #059669);
  color: white;
  border: none;
  border-radius: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  box-shadow:
    0 4px 6px rgba(16, 185, 129, 0.25),
    0 1px 3px rgba(16, 185, 129, 0.12);
}

.try-now-btn:hover {
  background: linear-gradient(135deg, #059669, #047857);
  transform: translateY(-2px);
  box-shadow:
    0 8px 12px rgba(16, 185, 129, 0.3),
    0 2px 6px rgba(16, 185, 129, 0.15);
}

.try-now-btn svg {
  width: 18px;
  height: 18px;
}

/* Animations */
@keyframes gentleFloat {
  0%, 100% {
    transform: translateY(0px) translateX(0px) rotate(0deg);
  }
  25% {
    transform: translateY(-10px) translateX(5px) rotate(2deg);
  }
  50% {
    transform: translateY(-5px) translateX(-3px) rotate(-1deg);
  }
  75% {
    transform: translateY(-15px) translateX(8px) rotate(3deg);
  }
}

@keyframes pulse {
  0%, 100% {
    transform: scale(1);
    opacity: 0.4;
  }
  50% {
    transform: scale(1.1);
    opacity: 0.6;
  }
}

@keyframes iconPulse {
  0%, 100% {
    transform: scale(1);
    opacity: 0.3;
  }
  50% {
    transform: scale(1.2);
    opacity: 0.6;
  }
}

@keyframes iconFloat {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-10px);
  }
}

@keyframes particleFloat {
  0%, 100% {
    transform: translateY(0px) scale(1);
    opacity: 0.6;
  }
  25% {
    transform: translateY(-15px) scale(1.2);
    opacity: 1;
  }
  50% {
    transform: translateY(-8px) scale(0.8);
    opacity: 0.8;
  }
  75% {
    transform: translateY(-20px) scale(1.1);
    opacity: 1;
  }
}

@keyframes lineGlow {
  0%, 100% {
    opacity: 0.3;
    transform: scaleX(1);
  }
  50% {
    opacity: 0.8;
    transform: scaleX(1.2);
  }
}

@keyframes dotPulse {
  0%, 100% {
    opacity: 0.6;
    transform: scale(1);
  }
  50% {
    opacity: 1;
    transform: scale(1.5);
  }
}

@keyframes modalBackdropFadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

@keyframes modalSlideIn {
  from {
    opacity: 0;
    transform: scale(0.95) translateY(-20px);
  }
  to {
    opacity: 1;
    transform: scale(1) translateY(0);
  }
}

/* Responsive Design */
@media (max-width: 768px) {
  .empty-state-container {
    padding: 3rem 1rem;
    min-height: 500px;
  }

  .title-primary {
    font-size: 2rem;
  }

  .title-secondary {
    font-size: 1.125rem;
  }

  .empty-description {
    font-size: 1rem;
  }

  .action-buttons {
    flex-direction: column;
    align-items: center;
  }

  .primary-action-btn,
  .secondary-action-btn {
    width: 100%;
    max-width: 280px;
    justify-content: center;
  }

  .floating-shape,
  .gradient-orb {
    display: none;
  }

  .instructions-modal-backdrop {
    padding: 1rem;
  }

  .instructions-modal {
    border-radius: 16px;
  }

  .instructions-header {
    padding: 1.5rem 1.5rem 1rem 1.5rem;
  }

  .instructions-title {
    font-size: 1.25rem;
  }

  .instructions-content {
    padding: 1rem 1.5rem 1.5rem 1.5rem;
  }

  .instruction-step {
    flex-direction: column;
    gap: 0.75rem;
  }

  .step-number {
    align-self: flex-start;
  }
}

/* Dark Mode */
@media (prefers-color-scheme: dark) {
  .empty-state-container {
    background: linear-gradient(
      135deg,
      rgba(30, 41, 59, 0.9) 0%,
      rgba(15, 23, 42, 0.95) 100%
    );
  }

  .title-primary {
    background: linear-gradient(135deg, #e2e8f0, #cbd5e1);
    background-clip: text;
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
  }

  .title-secondary {
    color: #94a3b8;
  }

  .empty-description {
    color: #94a3b8;
  }

  .secondary-action-btn {
    background: rgba(30, 41, 59, 0.8);
    color: #cbd5e1;
    border-color: rgba(71, 85, 105, 0.3);
  }

  .secondary-action-btn:hover {
    background: rgba(30, 41, 59, 0.95);
  }

  .cta-footer {
    background: rgba(30, 41, 59, 0.6);
    border-color: rgba(71, 85, 105, 0.3);
  }

  .cta-text {
    color: #94a3b8;
  }

  .instructions-modal {
    background: linear-gradient(135deg, rgba(30, 41, 59, 0.95), rgba(15, 23, 42, 0.98));
    border-color: rgba(71, 85, 105, 0.3);
  }

  .instructions-header {
    border-bottom-color: rgba(71, 85, 105, 0.2);
  }

  .instructions-title {
    color: #e2e8f0;
  }

  .close-btn {
    background: rgba(71, 85, 105, 0.2);
    color: #94a3b8;
  }

  .close-btn:hover {
    background: rgba(239, 68, 68, 0.2);
    color: #ef4444;
  }

  .instructions-intro,
  .step-content p {
    color: #94a3b8;
  }

  .step-content h4 {
    color: #e2e8f0;
  }

  .copy-email-btn {
    background: rgba(16, 185, 129, 0.2);
    border-color: rgba(16, 185, 129, 0.3);
    color: #10b981;
  }

  .copy-email-btn:hover {
    background: rgba(16, 185, 129, 0.3);
  }

  .instructions-footer {
    border-top-color: rgba(71, 85, 105, 0.2);
  }
}

/* Reduced Motion Support */
@media (prefers-reduced-motion: reduce) {
  .floating-shape,
  .gradient-orb,
  .main-icon,
  .particle,
  .deco-line,
  .deco-dot,
  .icon-backdrop {
    animation: none;
  }

  .primary-action-btn:hover,
  .secondary-action-btn:hover {
    transform: none;
  }
}

/* High Contrast Mode */
@media (prefers-contrast: high) {
  .empty-state-container {
    background: white;
    border: 2px solid black;
  }

  .title-primary {
    color: black;
    -webkit-text-fill-color: black;
  }

  .floating-shape,
  .gradient-orb,
  .decorative-elements {
    display: none;
  }
}

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
