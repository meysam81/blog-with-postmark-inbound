<template>
  <div
    class="title-animation"
    ref="titleElement"
    :aria-label="titleText"
    role="heading"
    aria-level="2"
  >
    <div class="title-text" v-html="wrappedTitle"></div>
    <div class="title-decoration">
      <div class="sparkle sparkle-1" aria-hidden="true">✨</div>
      <div class="sparkle sparkle-2" aria-hidden="true">✨</div>
      <div class="sparkle sparkle-3" aria-hidden="true">✨</div>
    </div>
    <span class="sr-only">{{ titleText }}</span>
  </div>
</template>

<script>
import { ref, onMounted, computed } from 'vue'

export default {
  name: 'TitleAnimation',
  setup() {
    var titleElement = ref(null)
    var titleText = "Transforming Your Emails into Blogs!"
    var isAnimated = ref(false)

    var wrappedTitle = computed(function computeWrappedTitle() {
      var wrapped = ""
      for (var i = 0; i < titleText.length; i++) {
        var char = titleText[i]
        var delay = i * 0.01
        if (char === " ") {
          wrapped += `<span class="char-hover char-space" style="animation-delay: ${delay}s"> </span>`
        } else {
          wrapped += `<span class="char-hover" style="animation-delay: ${delay}s">${char}</span>`
        }
      }
      return wrapped
    })

    function startAnimation() {
      if (!titleElement.value) {return}
      isAnimated.value = true
      titleElement.value.classList.add('animated')
    }

    onMounted(function() {
      setTimeout(startAnimation, 500)
    })

    return {
      titleElement,
      titleText,
      wrappedTitle,
      isAnimated
    }
  }
}
</script>

<style scoped>
.title-animation {
  position: relative;
  text-align: center;
  margin: 32px 0;
  padding: 24px;
  font-size: clamp(24px, 4vw, 48px);
  font-weight: 700;
  line-height: 1.2;
  perspective: 1000px;
}

.title-text {
  position: relative;
  z-index: 2;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 50%, #f093fb 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  background-size: 200% 200%;
  animation: gradientShift 4s ease-in-out infinite;
}

.char-hover {
  display: inline-block;
  opacity: 0;
  transform: translateY(30px) rotateX(90deg);
  transition: all 0.6s cubic-bezier(0.4, 0, 0.2, 1);
}

.title-animation.animated .char-hover {
  opacity: 1;
  transform: translateY(0) rotateX(0deg);
  animation: charFloat 3s ease-in-out infinite;
}

.char-hover:hover {
  transform: translateY(-5px) scale(1.1);
  text-shadow: 0 5px 15px rgba(102, 126, 234, 0.4);
  transition: all 0.3s ease;
}

.char-space {
  width: 0.3em;
}

.title-decoration {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  pointer-events: none;
  z-index: 1;
}

.sparkle {
  position: absolute;
  font-size: 24px;
  opacity: 0;
  animation: sparkleFloat 2s ease-in-out infinite;
}

.sparkle-1 {
  top: 10%;
  left: 15%;
  animation-delay: 0s;
}

.sparkle-2 {
  top: 20%;
  right: 20%;
  animation-delay: 0.7s;
}

.sparkle-3 {
  bottom: 15%;
  left: 25%;
  animation-delay: 1.4s;
}

.title-animation.animated .sparkle {
  opacity: 1;
}

@keyframes gradientShift {
  0%, 100% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
}

@keyframes charFloat {
  0%, 100% {
    transform: translateY(0) rotateX(0deg);
  }
  50% {
    transform: translateY(-3px) rotateX(5deg);
  }
}

@keyframes sparkleFloat {
  0%, 100% {
    opacity: 0;
    transform: translateY(0) scale(0.8);
  }
  50% {
    opacity: 0.8;
    transform: translateY(-10px) scale(1.2);
  }
}

@media (max-width: 768px) {
  .title-animation {
    padding: 16px;
    margin: 24px 0;
  }

  .sparkle {
    font-size: 16px;
  }
}

@media (prefers-reduced-motion: reduce) {
  .char-hover,
  .sparkle,
  .title-text {
    animation: none !important;
    transition: none !important;
  }

  .title-animation.animated .char-hover {
    opacity: 1;
    transform: none;
  }

  .char-hover:hover {
    transform: none;
  }
}

@media (prefers-color-scheme: dark) {
  .title-text {
    background: linear-gradient(135deg, #8b9dc3 0%, #ddb3f0 50%, #f093fb 100%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    background-clip: text;
  }

  .char-hover:hover {
    text-shadow: 0 5px 15px rgba(139, 157, 195, 0.4);
  }
}

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
