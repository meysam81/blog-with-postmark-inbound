<template>
  <div class="min-h-screen flex flex-col">
    <SkipLink />
    <AppHeader />
    <MainContent />
    <AppFooter />
    <PublishButton />
  </div>
</template>

<script>
import SkipLink from '@/components/SkipLink.vue'
import AppHeader from '@/components/AppHeader.vue'
import MainContent from '@/components/MainContent.vue'
import AppFooter from '@/components/AppFooter.vue'
import PublishButton from '@/components/PublishButton.vue'

export default {
  name: 'App',
  components: {
    SkipLink,
    AppHeader,
    MainContent,
    AppFooter,
    PublishButton
  }
}
</script>
