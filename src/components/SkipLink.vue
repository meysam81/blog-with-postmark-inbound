<template>
  <a
    href="#main-content"
    class="skip-link"
    @click="handleSkipLinkClick"
    @keydown="handleKeydown"
  >
    Skip to main content
  </a>
</template>

<script>
export default {
  name: 'SkipLink',
  methods: {
    handleSkipLinkClick: function handleSkipLinkClick(event) {
      var target = document.getElementById('main-content')
      if (target) {
        target.focus()
        target.scrollIntoView({ behavior: 'smooth', block: 'start' })
      }
    },

    handleKeydown: function handleKeydown(event) {
      if (event.key === 'Enter' || event.key === ' ') {
        event.preventDefault()
        this.handleSkipLinkClick(event)
      }
    }
  }
}
</script>

<style scoped>
.skip-link {
  position: absolute;
  top: -40px;
  left: 6px;
  background: #000;
  color: #fff;
  padding: 8px 16px;
  text-decoration: none;
  font-weight: 600;
  font-size: 14px;
  border-radius: 4px;
  border: 2px solid #fff;
  z-index: 10000;
  transition: all 0.3s ease;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
}

.skip-link:focus {
  top: 6px;
  outline: none;
  background: #667eea;
  border-color: #667eea;
  box-shadow: 0 0 0 4px rgba(102, 126, 234, 0.3), 0 8px 24px rgba(0, 0, 0, 0.4);
}

.skip-link:hover:focus {
  background: #5a67d8;
  transform: translateY(-2px);
}

.skip-link:active {
  transform: translateY(0);
}

@media (prefers-reduced-motion: reduce) {
  .skip-link {
    transition: none;
  }

  .skip-link:hover:focus {
    transform: none;
  }
}

@media (prefers-color-scheme: dark) {
  .skip-link {
    background: #fff;
    color: #000;
    border-color: #000;
  }

  .skip-link:focus {
    background: #8b9dc3;
    color: #fff;
    border-color: #8b9dc3;
    box-shadow: 0 0 0 4px rgba(139, 157, 195, 0.3), 0 8px 24px rgba(255, 255, 255, 0.2);
  }

  .skip-link:hover:focus {
    background: #7c8eb8;
  }
}

@media (prefers-contrast: high) {
  .skip-link {
    background: #000;
    color: #fff;
    border: 3px solid #fff;
    font-weight: 700;
  }

  .skip-link:focus {
    background: #fff;
    color: #000;
    border-color: #000;
    box-shadow: 0 0 0 6px #000;
  }
}
</style>
