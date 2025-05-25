<template>
  <div class="min-h-screen relative overflow-hidden">
    <!-- Header -->
    <AppHeader />

    <!-- Main Content -->
    <MainContent />

    <!-- Footer -->
    <AppFooter />

    <!-- Scroll Progress Indicator -->
    <ScrollProgress />
  </div>
</template>

<script>
import AppHeader from '@/components/AppHeader.vue'
import MainContent from '@/components/MainContent.vue'
import AppFooter from '@/components/AppFooter.vue'
import ScrollProgress from '@/components/ScrollProgress.vue'

export default {
  name: 'HomePage',
  components: {
    AppHeader,
    MainContent,
    AppFooter,
    ScrollProgress
  },
  mounted: function setHomePageTitle() {
    document.title = 'Tarzan - Email to Blog Platform'
  },
  beforeRouteEnter: function beforeEnterHome(to, from, next) {
    document.title = 'Tarzan - Email to Blog Platform'
    next()
  },
  beforeRouteUpdate: function beforeUpdateHome(to, from, next) {
    document.title = 'Tarzan - Email to Blog Platform'
    next()
  }
}
</script>
