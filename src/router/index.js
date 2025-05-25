import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '@/views/HomePage.vue'
import BlogPostView from '@/views/BlogPostView.vue'

var routes = [
  {
    path: '/',
    name: 'Home',
    component: HomePage
  },
  {
    path: '/post/:id',
    name: 'BlogPost',
    component: BlogPostView,
    props: true
  }
]

var router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    }

    return { top: 0, behavior: 'smooth' }
  }
})

export default router
