import { createRouter, createWebHistory } from 'vue-router';
import Company from '@/components/Company.vue';

const routes = [
  {
    path: '/',
    component: Company,
  },
];

const router = createRouter({
  routes,
  history: createWebHistory('http://localhost:8080'),
});

export default router;
