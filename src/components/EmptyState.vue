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
          <span class="title-secondary">Be the First to Send an Email!</span>
        </h2>

        <p class="empty-description">
          This is where blog posts created from emails will appear. Send an email to our blog address and watch your post appear here instantly.
        </p>

        <!-- Action Buttons -->
        <div class="action-buttons">
          <button class="primary-action-btn" @click="handleLearnMore">
            <svg class="btn-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
            </svg>
            <span>Send Email</span>
          </button>

          <button class="secondary-action-btn" @click="handleScrollToInstructions">
            <svg class="btn-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14l-7 7m0 0l-7-7m7 7V3"/>
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
          Ready to get started? Send your first email and watch your blog post appear ✨
        </p>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: 'EmptyState',
  setup() {
    function handleLearnMore() {
      // Scroll to or navigate to email instructions
      var instructionsElement = document.querySelector('#publishing-instructions')
      if (instructionsElement) {
        instructionsElement.scrollIntoView({
          behavior: 'smooth',
          block: 'start'
        })
      }
    }

    function handleScrollToInstructions() {
      // Alternative scroll to instructions
      var mainContent = document.querySelector('#main-content')
      if (mainContent) {
        var instructionsSection = mainContent.querySelector('.instructions-section')
        if (instructionsSection) {
          instructionsSection.scrollIntoView({
            behavior: 'smooth',
            block: 'start'
          })
        }
      }
    }

    return {
      handleLearnMore,
      handleScrollToInstructions
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
}
</style>
