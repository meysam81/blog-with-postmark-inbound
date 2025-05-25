<template>
  <div class="github-link-container">
    <a
      href="https://github.com/meysam81/tarzan"
      class="github-link"
      aria-label="View Tarzan project on GitHub"
      target="_blank"
      rel="noopener noreferrer"
      @mouseenter="handleMouseEnter"
      @mouseleave="handleMouseLeave"
    >
      <div class="link-background"></div>
      <div class="link-content">
        <div class="github-icon">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="currentColor"
            aria-hidden="true"
          >
            <path
              d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"
            />
          </svg>
        </div>
        <span class="link-text">View on GitHub</span>
        <div class="link-arrow">
          <svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor" aria-hidden="true">
            <path d="M8.22 2.97a.75.75 0 0 1 1.06 0l4.25 4.25a.75.75 0 0 1 0 1.06l-4.25 4.25a.75.75 0 0 1-1.06-1.06L11.44 8.5H2.75a.75.75 0 0 1 0-1.5h8.69L8.22 4.03a.75.75 0 0 1 0-1.06Z"/>
          </svg>
        </div>
      </div>
      <div class="floating-stars" aria-hidden="true">
        <div class="star star-1">⭐</div>
        <div class="star star-2">⭐</div>
        <div class="star star-3">⭐</div>
      </div>
    </a>
    <div class="tooltip" role="tooltip">
      Open source on GitHub
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'

export default {
  name: 'GitHubLink',
  setup() {
    var isHovered = ref(false)

    function handleMouseEnter() {
      isHovered.value = true
    }

    function handleMouseLeave() {
      isHovered.value = false
    }

    return {
      isHovered,
      handleMouseEnter,
      handleMouseLeave
    }
  }
}
</script>

<style scoped>
.github-link-container {
  position: relative;
  display: inline-block;
}

.github-link {
  position: relative;
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 20px;
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(10px);
  border: 2px solid rgba(102, 126, 234, 0.2);
  border-radius: 50px;
  color: #333;
  text-decoration: none;
  font-weight: 600;
  font-size: 14px;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  overflow: hidden;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
}

.github-link::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(102, 126, 234, 0.1), transparent);
  transition: left 0.6s ease;
}

.github-link:hover::before {
  left: 100%;
}

.github-link:hover {
  transform: translateY(-2px);
  border-color: rgba(102, 126, 234, 0.5);
  box-shadow: 0 8px 32px rgba(102, 126, 234, 0.2);
  color: #667eea;
}

.github-link:focus {
  outline: none;
  box-shadow: 0 0 0 4px rgba(102, 126, 234, 0.3), 0 8px 32px rgba(102, 126, 234, 0.2);
}

.link-background {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.05), rgba(118, 75, 162, 0.05));
  opacity: 0;
  transition: opacity 0.3s ease;
}

.github-link:hover .link-background {
  opacity: 1;
}

.link-content {
  position: relative;
  display: flex;
  align-items: center;
  gap: 12px;
  z-index: 2;
}

.github-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  transition: transform 0.3s ease;
}

.github-link:hover .github-icon {
  transform: rotate(360deg);
}

.link-text {
  white-space: nowrap;
}

.link-arrow {
  display: flex;
  align-items: center;
  transition: transform 0.3s ease;
}

.github-link:hover .link-arrow {
  transform: translateX(4px);
}

.floating-stars {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  pointer-events: none;
  z-index: 1;
}

.star {
  position: absolute;
  font-size: 12px;
  opacity: 0;
  transition: all 0.5s ease;
}

.star-1 {
  top: 20%;
  left: 10%;
}

.star-2 {
  top: 60%;
  right: 15%;
}

.star-3 {
  bottom: 25%;
  left: 70%;
}

.github-link:hover .star {
  opacity: 0.6;
  animation: starFloat 2s ease-in-out infinite;
}

.github-link:hover .star-1 {
  animation-delay: 0s;
}

.github-link:hover .star-2 {
  animation-delay: 0.3s;
}

.github-link:hover .star-3 {
  animation-delay: 0.6s;
}

.tooltip {
  position: absolute;
  bottom: 100%;
  left: 50%;
  transform: translateX(-50%);
  padding: 8px 12px;
  background: rgba(0, 0, 0, 0.8);
  color: white;
  font-size: 12px;
  border-radius: 8px;
  white-space: nowrap;
  opacity: 0;
  pointer-events: none;
  transition: all 0.3s ease;
  margin-bottom: 8px;
  z-index: 9999;
}

.tooltip::after {
  content: '';
  position: absolute;
  top: 100%;
  left: 50%;
  transform: translateX(-50%);
  border: 4px solid transparent;
  border-top-color: rgba(0, 0, 0, 0.8);
}

.github-link:hover + .tooltip {
  opacity: 1;
  transform: translateX(-50%) translateY(-4px);
}

@keyframes starFloat {
  0%, 100% {
    transform: translateY(0) rotate(0deg);
  }
  50% {
    transform: translateY(-8px) rotate(180deg);
  }
}

@media (max-width: 768px) {
  .github-link {
    padding: 10px 16px;
    font-size: 13px;
    gap: 8px;
  }

  .link-text {
    display: none;
  }

  .tooltip {
    display: none;
  }
}

@media (prefers-reduced-motion: reduce) {
  .github-link,
  .github-icon,
  .link-arrow,
  .star,
  .tooltip,
  .link-background {
    transition: none !important;
    animation: none !important;
  }

  .github-link:hover .github-icon {
    transform: none;
  }

  .github-link:hover .link-arrow {
    transform: none;
  }
}

@media (prefers-color-scheme: dark) {
  .github-link {
    background: rgba(26, 32, 44, 0.9);
    border-color: rgba(139, 157, 195, 0.3);
    color: #f7fafc;
  }

  .github-link:hover {
    border-color: rgba(139, 157, 195, 0.6);
    box-shadow: 0 8px 32px rgba(139, 157, 195, 0.2);
    color: #8b9dc3;
  }

  .link-background {
    background: linear-gradient(135deg, rgba(139, 157, 195, 0.1), rgba(221, 179, 240, 0.1));
  }

  .tooltip {
    background: rgba(255, 255, 255, 0.9);
    color: #333;
  }

  .tooltip::after {
    border-top-color: rgba(255, 255, 255, 0.9);
  }
}
</style>
