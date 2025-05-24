<template>
  <section
    class="blog-posts-container"
    aria-label="Blog posts"
  >
    <!-- Posts Grid -->
    <div class="posts-grid">
      <BlogCard
        v-for="(post, index) in posts"
        :key="post.id || post.title"
        :post="post"
        :index="index"
        class="blog-card-item"
        :style="{
          animationDelay: `${index * 0.1}s`,
          '--card-index': index
        }"
      />
    </div>

    <!-- Floating Background Elements -->
    <div class="floating-elements" aria-hidden="true">
      <div class="floating-shape floating-shape-1"></div>
      <div class="floating-shape floating-shape-2"></div>
      <div class="floating-shape floating-shape-3"></div>
      <div class="floating-gradient floating-gradient-1"></div>
      <div class="floating-gradient floating-gradient-2"></div>
    </div>

    <!-- Visual Enhancement Overlays -->
    <div class="visual-overlay visual-overlay-top" aria-hidden="true"></div>
    <div class="visual-overlay visual-overlay-bottom" aria-hidden="true"></div>
  </section>
</template>

<script>
import BlogCard from './BlogCard.vue'

export default {
  name: 'BlogPosts',
  components: {
    BlogCard
  },
  props: {
    posts: {
      type: Array,
      required: true
    }
  }
}
</script>

<style scoped>
.blog-posts-container {
  position: relative;
  padding: 3rem 0;
  overflow: hidden;
}

.posts-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(380px, 1fr));
  gap: 2.5rem;
  max-width: 1400px;
  margin: 0 auto;
  padding: 0 1.5rem;
  position: relative;
  z-index: 2;
}

.blog-card-item {
  opacity: 0;
  transform: translateY(30px);
  animation: slideInUp 0.8s cubic-bezier(0.175, 0.885, 0.32, 1.275) forwards;
}

/* Floating Background Elements */
.floating-elements {
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
  animation: float 20s ease-in-out infinite;
}

.floating-shape-1 {
  width: 300px;
  height: 300px;
  top: 10%;
  left: -5%;
  animation-delay: 0s;
}

.floating-shape-2 {
  width: 200px;
  height: 200px;
  top: 60%;
  right: -3%;
  animation-delay: -7s;
}

.floating-shape-3 {
  width: 150px;
  height: 150px;
  top: 80%;
  left: 20%;
  animation-delay: -14s;
}

.floating-gradient {
  position: absolute;
  width: 400px;
  height: 400px;
  border-radius: 50%;
  filter: blur(100px);
  opacity: 0.3;
  animation: pulse 15s ease-in-out infinite;
}

.floating-gradient-1 {
  background: radial-gradient(circle, rgba(16, 185, 129, 0.2), transparent);
  top: 20%;
  right: 10%;
  animation-delay: -3s;
}

.floating-gradient-2 {
  background: radial-gradient(circle, rgba(139, 92, 246, 0.2), transparent);
  bottom: 20%;
  left: 10%;
  animation-delay: -8s;
}

/* Visual Enhancement Overlays */
.visual-overlay {
  position: absolute;
  left: 0;
  right: 0;
  height: 150px;
  pointer-events: none;
  z-index: 3;
}

.visual-overlay-top {
  top: 0;
  background: linear-gradient(
    180deg,
    rgba(15, 23, 42, 0.1) 0%,
    transparent 100%
  );
}

.visual-overlay-bottom {
  bottom: 0;
  background: linear-gradient(
    0deg,
    rgba(15, 23, 42, 0.1) 0%,
    transparent 100%
  );
}

/* Animations */
@keyframes slideInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes float {
  0%, 100% {
    transform: translateY(0px) translateX(0px) rotate(0deg);
  }
  25% {
    transform: translateY(-20px) translateX(10px) rotate(5deg);
  }
  50% {
    transform: translateY(-10px) translateX(-5px) rotate(-3deg);
  }
  75% {
    transform: translateY(-30px) translateX(15px) rotate(7deg);
  }
}

@keyframes pulse {
  0%, 100% {
    transform: scale(1);
    opacity: 0.3;
  }
  50% {
    transform: scale(1.1);
    opacity: 0.5;
  }
}

/* Responsive Design */
@media (max-width: 768px) {
  .posts-grid {
    grid-template-columns: 1fr;
    gap: 1.5rem;
    padding: 0 1rem;
  }

  .blog-posts-container {
    padding: 2rem 0;
  }

  .floating-shape {
    display: none;
  }

  .floating-gradient {
    width: 250px;
    height: 250px;
  }
}

@media (min-width: 1200px) {
  .posts-grid {
    grid-template-columns: repeat(auto-fit, minmax(420px, 1fr));
    gap: 3rem;
  }
}

/* Reduced Motion Support */
@media (prefers-reduced-motion: reduce) {
  .blog-card-item {
    opacity: 1;
    transform: none;
    animation: none;
  }

  .floating-shape,
  .floating-gradient {
    animation: none;
  }
}

/* High Contrast Mode */
@media (prefers-contrast: high) {
  .floating-shape,
  .floating-gradient,
  .visual-overlay {
    display: none;
  }
}

/* Dark Mode Adjustments */
@media (prefers-color-scheme: dark) {
  .visual-overlay-top {
    background: linear-gradient(
      180deg,
      rgba(2, 6, 23, 0.2) 0%,
      transparent 100%
    );
  }

  .visual-overlay-bottom {
    background: linear-gradient(
      0deg,
      rgba(2, 6, 23, 0.2) 0%,
      transparent 100%
    );
  }
}
</style>
