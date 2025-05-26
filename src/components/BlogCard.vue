<template>
  <article
    class="blog-card"
    :class="{ 'featured-card': index === 0 }"
    @mouseenter="handleMouseEnter"
    @mouseleave="handleMouseLeave"
    role="article"
    :aria-labelledby="`blog-title-${post.id || index}`"
  >
    <!-- Card Background Gradient -->
    <div class="card-background" aria-hidden="true"></div>

    <!-- Card Border Glow -->
    <div class="card-border-glow" aria-hidden="true"></div>

    <!-- Card Content -->
    <div class="card-content">
      <!-- Header Section -->
      <header class="card-header">
        <div class="category-badge" aria-hidden="true">
          <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
            <path d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
          <span>Blog Post</span>
        </div>

        <h2
          :id="`blog-title-${post.id || index}`"
          class="card-title"
        >
          {{ post.title || 'Untitled Post' }}
        </h2>
      </header>

      <!-- Content Section -->
      <div class="card-main-content">
        <div
          class="prose-content"
          v-html="post.content || 'No content available.'"
          :aria-describedby="`blog-meta-${post.id || index}`"
        >
        </div>

        <!-- Content Fade Overlay -->
        <div class="content-fade-overlay" aria-hidden="true"></div>
      </div>

      <!-- Footer Section -->
      <footer
        class="card-footer"
        :id="`blog-meta-${post.id || index}`"
      >
        <div class="author-info">
          <div class="author-avatar">
            <span class="avatar-text">{{ getAuthorInitial(post['author-name'], post['author-email']) }}</span>
            <div class="avatar-ring" aria-hidden="true"></div>
          </div>

          <div class="author-details" @mouseenter="handleAuthorHover" @mouseleave="handleAuthorLeave">
            <p class="author-name">{{ post['author-name'] || post['author-email'] || 'Unknown Author' }}</p>
            <time
              class="publish-date"
              :datetime="formatISODate(post['created-at'])"
            >
              {{ formatReadableDate(post['created-at']) }}
            </time>

            <!-- Email Tooltip -->
            <div
              v-if="showEmailTooltip && post['author-email']"
              class="email-tooltip"
              role="tooltip"
            >
              {{ post['author-email'] }}
            </div>
          </div>
        </div>

        <!-- Interactive Elements -->
        <div class="card-actions">
          <button
            class="action-btn read-btn"
            @click="handleReadMore"
            :aria-label="`Read full post: ${post.title}`"
          >
            <span>Read More</span>
            <svg class="w-4 h-4 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3"/>
            </svg>
          </button>
        </div>
      </footer>
    </div>

    <!-- Floating Particles -->
    <div class="floating-particles" aria-hidden="true">
      <div class="particle particle-1"></div>
      <div class="particle particle-2"></div>
      <div class="particle particle-3"></div>
    </div>
  </article>
</template>

<script>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { formatISODate, formatReadableDate, getAuthorInitial } from '@/utils/api.js'

export default {
  name: 'BlogCard',
  props: {
    post: {
      type: Object,
      required: true,
      validator(value) {
        return value && typeof value === 'object'
      }
    },
    index: {
      type: Number,
      default: 0,
      validator(value) {
        return value >= 0
      }
    }
  },
  setup(props) {
    var isHovered = ref(false)
    var showEmailTooltip = ref(false)
    var router = useRouter()

    function handleMouseEnter() {
      isHovered.value = true
    }

    function handleMouseLeave() {
      isHovered.value = false
    }

    function handleAuthorHover() {
      showEmailTooltip.value = true
    }

    function handleAuthorLeave() {
      showEmailTooltip.value = false
    }

    function handleReadMore() {
      // Navigate to the individual blog post page
      if (props.post.id) {
        router.push({ name: 'BlogPost', params: { id: props.post.id } })
      } else {
        // Fallback: use array index if no ID is available
        router.push({ name: 'BlogPost', params: { id: props.index.toString() } })
      }
    }

    return {
      isHovered,
      showEmailTooltip,
      handleMouseEnter,
      handleMouseLeave,
      handleAuthorHover,
      handleAuthorLeave,
      handleReadMore,
      formatISODate,
      formatReadableDate,
      getAuthorInitial
    }
  }
}
</script>

<style scoped>
.blog-card {
  position: relative;
  background: linear-gradient(
    135deg,
    rgba(255, 255, 255, 0.9) 0%,
    rgba(248, 250, 252, 0.95) 100%
  );
  border-radius: 24px;
  padding: 0;
  overflow: hidden;
  transition: all 0.6s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  backdrop-filter: blur(12px);
  border: 1px solid rgba(148, 163, 184, 0.2);
  box-shadow:
    0 4px 6px rgba(0, 0, 0, 0.05),
    0 1px 3px rgba(0, 0, 0, 0.1),
    inset 0 1px 0 rgba(255, 255, 255, 0.6);
  cursor: pointer;
}

.blog-card:hover {
  transform: translateY(-8px) scale(1.02);
  box-shadow:
    0 20px 25px rgba(0, 0, 0, 0.1),
    0 8px 10px rgba(0, 0, 0, 0.04),
    inset 0 1px 0 rgba(255, 255, 255, 0.8);
}

.featured-card {
  background: linear-gradient(
    135deg,
    rgba(16, 185, 129, 0.1) 0%,
    rgba(59, 130, 246, 0.1) 50%,
    rgba(139, 92, 246, 0.1) 100%
  );
  border: 2px solid rgba(16, 185, 129, 0.3);
}

/* Card Background Effects */
.card-background {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: radial-gradient(
    circle at 50% 0%,
    rgba(16, 185, 129, 0.1) 0%,
    transparent 50%
  );
  opacity: 0;
  transition: opacity 0.6s ease;
}

.blog-card:hover .card-background {
  opacity: 1;
}

.card-border-glow {
  position: absolute;
  top: -2px;
  left: -2px;
  right: -2px;
  bottom: -2px;
  background: linear-gradient(
    45deg,
    rgba(16, 185, 129, 0.6),
    rgba(59, 130, 246, 0.6),
    rgba(139, 92, 246, 0.6),
    rgba(16, 185, 129, 0.6)
  );
  border-radius: 26px;
  opacity: 0;
  filter: blur(8px);
  transition: opacity 0.6s ease;
  z-index: -1;
}

.blog-card:hover .card-border-glow {
  opacity: 0.7;
  animation: borderGlow 3s ease-in-out infinite;
}

/* Card Content */
.card-content {
  position: relative;
  padding: 2rem;
  z-index: 2;
  height: 100%;
  display: flex;
  flex-direction: column;
}

/* Header Section */
.card-header {
  margin-bottom: 1.5rem;
}

.category-badge {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 1rem;
  background: linear-gradient(135deg, rgba(16, 185, 129, 0.1), rgba(59, 130, 246, 0.1));
  border: 1px solid rgba(16, 185, 129, 0.2);
  border-radius: 50px;
  font-size: 0.875rem;
  font-weight: 500;
  color: #059669;
  margin-bottom: 1rem;
  width: fit-content;
  transition: all 0.3s ease;
}

.blog-card:hover .category-badge {
  background: linear-gradient(135deg, rgba(16, 185, 129, 0.2), rgba(59, 130, 246, 0.2));
  transform: scale(1.05);
}

.card-title {
  font-size: 1.75rem;
  font-weight: 700;
  line-height: 1.3;
  color: #1e293b;
  margin: 0;
  background: linear-gradient(135deg, #1e293b, #475569);
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  transition: all 0.3s ease;
}

.blog-card:hover .card-title {
  background: linear-gradient(135deg, #059669, #3b82f6);
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

/* Content Section */
.card-main-content {
  position: relative;
  flex: 1;
  margin-bottom: 1.5rem;
  overflow: hidden;
}

.prose-content {
  color: #64748b;
  line-height: 1.7;
  font-size: 1rem;
  max-height: 150px;
  overflow: hidden;
  position: relative;
}

.prose-content :deep(p) {
  margin-bottom: 1rem;
}

.prose-content :deep(h1),
.prose-content :deep(h2),
.prose-content :deep(h3) {
  color: #334155;
  font-weight: 600;
  margin-bottom: 0.75rem;
}

.content-fade-overlay {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 40px;
  background: linear-gradient(
    180deg,
    transparent 0%,
    rgba(255, 255, 255, 0.9) 100%
  );
  pointer-events: none;
}

/* Footer Section */
.card-footer {
  border-top: 1px solid rgba(148, 163, 184, 0.1);
  padding-top: 1.5rem;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
}

.author-info {
  display: flex;
  align-items: center;
  gap: 1rem;
  flex: 1;
}

.author-avatar {
  position: relative;
  width: 48px;
  height: 48px;
  background: linear-gradient(135deg, #10b981, #3b82f6);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;
}

.avatar-text {
  color: white;
  font-weight: 600;
  font-size: 1.125rem;
}

.avatar-ring {
  position: absolute;
  top: -3px;
  left: -3px;
  right: -3px;
  bottom: -3px;
  border: 2px solid rgba(16, 185, 129, 0.3);
  border-radius: 50%;
  opacity: 0;
  transition: all 0.3s ease;
}

.blog-card:hover .avatar-ring {
  opacity: 1;
  transform: scale(1.1);
}

.author-details {
  flex: 1;
  position: relative;
}

.author-name {
  font-weight: 600;
  color: #1e293b;
  margin: 0 0 0.25rem 0;
  font-size: 0.9rem;
}

.publish-date {
  color: #64748b;
  font-size: 0.8rem;
}

/* Email Tooltip */
.email-tooltip {
  position: absolute;
  top: 100%;
  left: 0;
  background: #1e293b;
  color: white;
  padding: 0.5rem 0.75rem;
  border-radius: 6px;
  font-size: 0.75rem;
  white-space: nowrap;
  z-index: 100;
  margin-top: 0.5rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  animation: tooltipFadeIn 0.2s ease-in-out;
}

.email-tooltip::before {
  content: '';
  position: absolute;
  top: -4px;
  left: 1rem;
  width: 0;
  height: 0;
  border-left: 4px solid transparent;
  border-right: 4px solid transparent;
  border-bottom: 4px solid #1e293b;
}

@keyframes tooltipFadeIn {
  from {
    opacity: 0;
    transform: translateY(-4px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Action Button */
.card-actions {
  display: flex;
  gap: 0.75rem;
}

.action-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem 1.25rem;
  border-radius: 12px;
  font-weight: 500;
  font-size: 0.875rem;
  transition: all 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  border: none;
  cursor: pointer;
}

.read-btn {
  background: linear-gradient(135deg, #10b981, #059669);
  color: white;
  box-shadow: 0 2px 4px rgba(16, 185, 129, 0.3);
}

.read-btn:hover {
  background: linear-gradient(135deg, #059669, #047857);
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(16, 185, 129, 0.4);
}

.read-btn:hover svg {
  transform: translateX(4px);
}

/* Floating Particles */
.floating-particles {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  pointer-events: none;
  overflow: hidden;
}

.particle {
  position: absolute;
  width: 4px;
  height: 4px;
  background: rgba(16, 185, 129, 0.6);
  border-radius: 50%;
  opacity: 0;
}

.particle-1 {
  top: 20%;
  left: 15%;
  animation: float 6s ease-in-out infinite;
  animation-delay: 0s;
}

.particle-2 {
  top: 60%;
  right: 20%;
  animation: float 8s ease-in-out infinite;
  animation-delay: -2s;
}

.particle-3 {
  bottom: 30%;
  left: 80%;
  animation: float 7s ease-in-out infinite;
  animation-delay: -4s;
}

.blog-card:hover .particle {
  opacity: 1;
}

/* Animations */
@keyframes borderGlow {
  0%, 100% {
    background-size: 200% 200%;
    background-position: 0% 50%;
  }
  50% {
    background-size: 200% 200%;
    background-position: 100% 50%;
  }
}

@keyframes float {
  0%, 100% {
    transform: translateY(0px) translateX(0px);
    opacity: 0.6;
  }
  25% {
    transform: translateY(-10px) translateX(5px);
    opacity: 1;
  }
  50% {
    transform: translateY(-5px) translateX(-3px);
    opacity: 0.8;
  }
  75% {
    transform: translateY(-15px) translateX(8px);
    opacity: 1;
  }
}

/* Responsive Design */
@media (max-width: 768px) {
  .card-content {
    padding: 1.5rem;
  }

  .card-title {
    font-size: 1.5rem;
  }

  .card-footer {
    flex-direction: column;
    align-items: flex-start;
    gap: 1rem;
  }

  .card-actions {
    width: 100%;
  }

  .action-btn {
    flex: 1;
    justify-content: center;
  }
}

/* Reduced Motion Support */
@media (prefers-reduced-motion: reduce) {
  .blog-card {
    transition: none;
  }

  .blog-card:hover {
    transform: none;
  }

  .particle,
  .card-border-glow {
    animation: none;
  }
}

/* High Contrast Mode */
@media (prefers-contrast: high) {
  .blog-card {
    background: white;
    border: 2px solid black;
  }

  .card-title {
    color: black;
    -webkit-text-fill-color: black;
  }

  .floating-particles,
  .card-background {
    display: none;
  }
}

/* Dark Mode */
@media (prefers-color-scheme: dark) {
  .blog-card {
    background: linear-gradient(
      135deg,
      rgba(30, 41, 59, 0.9) 0%,
      rgba(15, 23, 42, 0.95) 100%
    );
    border-color: rgba(71, 85, 105, 0.3);
  }

  .card-title {
    color: #e2e8f0;
    background: linear-gradient(135deg, #e2e8f0, #cbd5e1);
    background-clip: text;
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
  }

  .prose-content {
    color: #94a3b8;
  }

  .prose-content :deep(h1),
  .prose-content :deep(h2),
  .prose-content :deep(h3) {
    color: #cbd5e1;
  }

  .author-name {
    color: #f1f5f9;
  }

  .publish-date {
    color: #94a3b8;
  }

  .content-fade-overlay {
    background: linear-gradient(
      180deg,
      transparent 0%,
      rgba(30, 41, 59, 0.9) 100%
    );
  }
}
</style>
