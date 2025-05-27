import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '@/views/HomePage.vue'
import BlogPostView from '@/views/BlogPostView.vue'
import NotFoundView from '@/views/NotFoundView.vue'

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
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: NotFoundView
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
