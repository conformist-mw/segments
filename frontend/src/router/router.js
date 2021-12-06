import { createRouter, createWebHistory } from 'vue-router';
import Company from '../components/Company.vue';
import Section from '../components/Section.vue';

const routes = [
  {
    path: '/',
    component: Company,
  },
  {
    path: '/company/:slug/sections',
    component: Section,
  },
];

const router = createRouter({
  routes,
  history: createWebHistory(),
});

export default router;
