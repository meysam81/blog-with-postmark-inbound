<template>
  <footer class="app-footer" role="contentinfo">
    <!-- Background Elements -->
    <div class="footer-background" aria-hidden="true">
      <div class="background-gradient"></div>
      <div class="floating-shapes">
        <div class="shape shape-1"></div>
        <div class="shape shape-2"></div>
        <div class="shape shape-3"></div>
      </div>
    </div>

    <!-- Main Footer Content -->
    <div class="footer-content">
      <div class="footer-grid">
        <!-- Brand Section -->
        <div class="brand-section">
          <div class="brand-logo">
            <div class="logo-icon">
              <svg viewBox="0 0 24 24" fill="currentColor" class="logo-svg">
                <path d="M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5"/>
              </svg>
            </div>
            <div class="brand-info">
              <h3 class="brand-title">Tarzan</h3>
              <p class="brand-tagline">Email to Blog Platform</p>
            </div>
          </div>

          <p class="brand-description">
            Send an email, get a blog post. No complex interfaces, no registration required.
          </p>

          <div class="brand-motto">
            <span class="motto-icon">📧</span>
            <span class="motto-text">Email → Blog Post. It's That Simple.</span>
          </div>
        </div>

        <!-- Quick Links Section -->
        <div class="links-section">
          <h4 class="section-title">Quick Links</h4>
          <nav class="footer-nav" aria-label="Footer navigation">
            <a href="#features" class="footer-link">
              <svg class="link-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
              </svg>
              <span>Features</span>
            </a>
            <a href="#how-it-works" class="footer-link">
              <svg class="link-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
              </svg>
              <span>How it Works</span>
            </a>
            <a href="#getting-started" class="footer-link">
              <svg class="link-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253"/>
              </svg>
              <span>Getting Started</span>
            </a>
            <a href="https://github.com/meysam81/tarzan" class="footer-link" target="_blank" rel="noopener noreferrer">
              <svg class="link-icon" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
              </svg>
              <span>GitHub</span>
            </a>
          </nav>
        </div>

        <!-- Legal Section -->
        <div class="legal-section">
          <h4 class="section-title">Legal</h4>
          <nav class="footer-nav" aria-label="Legal navigation">
            <a
              href="https://github.com/meysam81/tarzan/blob/main/LICENSE"
              class="footer-link"
              target="_blank"
              rel="noopener noreferrer"
            >
              <svg class="link-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
              </svg>
              <span>Apache 2.0 License</span>
            </a>
            <a href="#privacy" class="footer-link">
              <svg class="link-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"/>
              </svg>
              <span>Privacy</span>
            </a>
          </nav>
        </div>
      </div>

      <!-- Footer Bottom -->
      <div class="footer-bottom">
        <div class="footer-divider" aria-hidden="true"></div>

        <div class="bottom-content">
          <div class="copyright-info">
            <p class="copyright-text">
              © {{ currentYear }}
              <a
                href="https://github.com/meysam81"
                class="author-link"
                target="_blank"
                rel="noopener noreferrer"
              >
                Meysam Azad
              </a>
              . All rights reserved.
            </p>
            <p class="build-info">
              Built with ❤️ and lots of ☕
            </p>
          </div>

          <div class="social-links">
            <a
              href="https://github.com/meysam81"
              class="social-link"
              target="_blank"
              rel="noopener noreferrer"
              aria-label="GitHub Profile"
            >
              <svg fill="currentColor" viewBox="0 0 24 24">
                <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
              </svg>
            </a>
          </div>
        </div>
      </div>
    </div>

    <!-- Decorative Elements -->
    <div class="footer-decorations" aria-hidden="true">
      <div class="deco-line deco-line-1"></div>
      <div class="deco-line deco-line-2"></div>
      <div class="deco-dot deco-dot-1"></div>
      <div class="deco-dot deco-dot-2"></div>
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

<style scoped>
.app-footer {
  position: relative;
  background: linear-gradient(
    135deg,
    rgba(15, 23, 42, 0.95) 0%,
    rgba(30, 41, 59, 0.98) 50%,
    rgba(51, 65, 85, 0.95) 100%
  );
  color: #e2e8f0;
  overflow: hidden;
  margin-top: auto;
  border-top: 1px solid rgba(148, 163, 184, 0.1);
}

/* Background Elements */
.footer-background {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  pointer-events: none;
  z-index: 1;
}

.background-gradient {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 100px;
  background: linear-gradient(
    180deg,
    rgba(16, 185, 129, 0.05) 0%,
    transparent 100%
  );
}

.floating-shapes {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
}

.shape {
  position: absolute;
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(16, 185, 129, 0.1), rgba(59, 130, 246, 0.1));
  animation: gentleFloat 20s ease-in-out infinite;
}

.shape-1 {
  width: 150px;
  height: 150px;
  top: 20%;
  left: 5%;
  animation-delay: 0s;
}

.shape-2 {
  width: 100px;
  height: 100px;
  top: 60%;
  right: 10%;
  animation-delay: -7s;
}

.shape-3 {
  width: 80px;
  height: 80px;
  bottom: 20%;
  left: 30%;
  animation-delay: -14s;
}

/* Main Content */
.footer-content {
  position: relative;
  z-index: 2;
  max-width: 1200px;
  margin: 0 auto;
  padding: 4rem 2rem 2rem;
}

.footer-grid {
  display: grid;
  grid-template-columns: 2fr 1fr 1fr;
  gap: 3rem;
  margin-bottom: 3rem;
}

/* Brand Section */
.brand-section {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.brand-logo {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.logo-icon {
  width: 48px;
  height: 48px;
  background: linear-gradient(135deg, #10b981, #059669);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 4px 8px rgba(16, 185, 129, 0.3);
}

.logo-svg {
  width: 24px;
  height: 24px;
  color: white;
}

.brand-title {
  font-size: 1.5rem;
  font-weight: 700;
  color: #f1f5f9;
  margin: 0 0 0.25rem 0;
  background: linear-gradient(135deg, #f1f5f9, #cbd5e1);
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.brand-tagline {
  font-size: 0.875rem;
  color: #94a3b8;
  margin: 0;
}

.brand-description {
  color: #cbd5e1;
  line-height: 1.6;
  font-size: 0.95rem;
  max-width: 400px;
}

.brand-motto {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 1rem;
  background: rgba(16, 185, 129, 0.1);
  border-radius: 12px;
  border: 1px solid rgba(16, 185, 129, 0.2);
  max-width: 400px;
}

.motto-icon {
  font-size: 1.25rem;
}

.motto-text {
  color: #10b981;
  font-weight: 500;
  font-size: 0.875rem;
}

/* Links and Legal Sections */
.links-section,
.legal-section {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.section-title {
  font-size: 1.125rem;
  font-weight: 600;
  color: #f1f5f9;
  margin: 0 0 1rem 0;
  position: relative;
}

.section-title::after {
  content: '';
  position: absolute;
  bottom: -4px;
  left: 0;
  width: 30px;
  height: 2px;
  background: linear-gradient(90deg, #10b981, #059669);
  border-radius: 1px;
}

.footer-nav {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.footer-link {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  color: #cbd5e1;
  text-decoration: none;
  transition: all 0.3s ease;
  padding: 0.5rem;
  border-radius: 8px;
  font-size: 0.9rem;
}

.footer-link:hover {
  color: #10b981;
  background: rgba(16, 185, 129, 0.1);
  transform: translateX(4px);
}

.link-icon {
  width: 16px;
  height: 16px;
  flex-shrink: 0;
  transition: transform 0.3s ease;
}

.footer-link:hover .link-icon {
  transform: scale(1.1);
}

/* Footer Bottom */
.footer-bottom {
  position: relative;
  z-index: 2;
}

.footer-divider {
  height: 1px;
  background: linear-gradient(
    90deg,
    transparent 0%,
    rgba(148, 163, 184, 0.3) 50%,
    transparent 100%
  );
  margin-bottom: 2rem;
}

.bottom-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 1rem;
}

.copyright-info {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.copyright-text {
  color: #94a3b8;
  font-size: 0.875rem;
  margin: 0;
}

.author-link {
  color: #10b981;
  text-decoration: none;
  font-weight: 500;
  transition: all 0.3s ease;
}

.author-link:hover {
  color: #059669;
  text-decoration: underline;
}

.build-info {
  color: #64748b;
  font-size: 0.8rem;
  margin: 0;
  font-style: italic;
}

.social-links {
  display: flex;
  gap: 1rem;
}

.social-link {
  width: 40px;
  height: 40px;
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(148, 163, 184, 0.2);
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #cbd5e1;
  text-decoration: none;
  transition: all 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  backdrop-filter: blur(8px);
}

.social-link:hover {
  background: rgba(16, 185, 129, 0.2);
  border-color: rgba(16, 185, 129, 0.4);
  color: #10b981;
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(16, 185, 129, 0.2);
}

.social-link svg {
  width: 20px;
  height: 20px;
}

/* Decorative Elements */
.footer-decorations {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  pointer-events: none;
  z-index: 1;
}

.deco-line {
  position: absolute;
  height: 1px;
  background: linear-gradient(
    90deg,
    transparent,
    rgba(16, 185, 129, 0.3),
    transparent
  );
  animation: lineGlow 8s ease-in-out infinite;
}

.deco-line-1 {
  width: 200px;
  top: 25%;
  left: 10%;
  animation-delay: 0s;
}

.deco-line-2 {
  width: 150px;
  bottom: 35%;
  right: 15%;
  animation-delay: -4s;
}

.deco-dot {
  position: absolute;
  width: 6px;
  height: 6px;
  background: #10b981;
  border-radius: 50%;
  opacity: 0.6;
  animation: dotPulse 6s ease-in-out infinite;
}

.deco-dot-1 {
  top: 30%;
  right: 25%;
  animation-delay: 0s;
}

.deco-dot-2 {
  bottom: 40%;
  left: 20%;
  animation-delay: -3s;
}

/* Animations */
@keyframes gentleFloat {
  0%, 100% {
    transform: translateY(0px) translateX(0px) rotate(0deg);
  }
  25% {
    transform: translateY(-15px) translateX(8px) rotate(2deg);
  }
  50% {
    transform: translateY(-8px) translateX(-5px) rotate(-1deg);
  }
  75% {
    transform: translateY(-20px) translateX(12px) rotate(3deg);
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
@media (max-width: 1024px) {
  .footer-grid {
    grid-template-columns: 1fr 1fr;
    gap: 2rem;
  }

  .legal-section {
    grid-column: 1 / -1;
  }
}

@media (max-width: 768px) {
  .footer-content {
    padding: 3rem 1.5rem 2rem;
  }

  .footer-grid {
    grid-template-columns: 1fr;
    gap: 2rem;
  }

  .brand-logo {
    flex-direction: column;
    text-align: center;
    gap: 0.75rem;
  }

  .brand-description,
  .brand-motto {
    max-width: none;
  }

  .bottom-content {
    flex-direction: column;
    text-align: center;
    gap: 1.5rem;
  }

  .floating-shapes {
    display: none;
  }
}

@media (max-width: 480px) {
  .footer-content {
    padding: 2rem 1rem 1.5rem;
  }

  .brand-title {
    font-size: 1.25rem;
  }

  .footer-link {
    padding: 0.75rem;
    font-size: 0.875rem;
  }
}

/* Reduced Motion Support */
@media (prefers-reduced-motion: reduce) {
  .shape,
  .deco-line,
  .deco-dot {
    animation: none;
  }

  .footer-link:hover {
    transform: none;
  }

  .social-link:hover {
    transform: none;
  }
}

/* High Contrast Mode */
@media (prefers-contrast: high) {
  .app-footer {
    background: #000000;
    border-top: 2px solid #ffffff;
  }

  .brand-title {
    color: #ffffff;
    -webkit-text-fill-color: #ffffff;
  }

  .footer-link {
    color: #ffffff;
  }

  .footer-link:hover {
    background: #333333;
  }

  .floating-shapes,
  .footer-decorations {
    display: none;
  }
}

/* Light Mode Adjustments */
@media (prefers-color-scheme: light) {
  .app-footer {
    background: linear-gradient(
      135deg,
      rgba(248, 250, 252, 0.95) 0%,
      rgba(241, 245, 249, 0.98) 50%,
      rgba(226, 232, 240, 0.95) 100%
    );
    color: #1e293b;
    border-top-color: rgba(71, 85, 105, 0.2);
  }

  .brand-title {
    color: #1e293b;
    background: linear-gradient(135deg, #1e293b, #475569);
    background-clip: text;
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
  }

  .section-title {
    color: #1e293b;
  }

  .footer-link {
    color: #475569;
  }

  .footer-link:hover {
    color: #059669;
    background: rgba(16, 185, 129, 0.1);
  }

  .copyright-text {
    color: #64748b;
  }

  .build-info {
    color: #94a3b8;
  }

  .social-link {
    background: rgba(0, 0, 0, 0.05);
    border-color: rgba(71, 85, 105, 0.2);
    color: #475569;
  }
}
</style>
