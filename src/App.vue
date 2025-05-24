<template>
  <div class="min-h-screen relative overflow-hidden">
    <!-- Skip Navigation -->
    <SkipLink />

    <!-- Header -->
    <AppHeader />

    <!-- Main Content -->
    <MainContent />

    <!-- Footer -->
    <AppFooter />

    <!-- Floating Publish Button -->
    <PublishButton />

    <!-- Scroll Progress Indicator -->
    <ScrollProgress />
  </div>
</template>

<script>
import SkipLink from '@/components/SkipLink.vue'
import AppHeader from '@/components/AppHeader.vue'
import MainContent from '@/components/MainContent.vue'
import AppFooter from '@/components/AppFooter.vue'
import PublishButton from '@/components/PublishButton.vue'
import ScrollProgress from '@/components/ScrollProgress.vue'

export default {
  name: 'App',
  components: {
    SkipLink,
    AppHeader,
    MainContent,
    AppFooter,
    PublishButton,
    ScrollProgress
  }
}
</script>
