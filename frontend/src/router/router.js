import { createRouter, createWebHistory } from 'vue-router';
import Company from '../components/Company.vue';
import Section from '../components/Section.vue';
import SegmentsPage from '../pages/SegmentsPage.vue';
import SegmentEdit from '../components/SegmentEdit.vue';

const routes = [
  {
    path: '/',
    component: Company,
  },
  {
    path: '/companies/:companySlug/sections',
    component: Section,
  },
  {
    path: '/companies/:companySlug/sections/:sectionSlug/segments',
    component: SegmentsPage,
  },
  {
    path: '/companies/:companySlug/sections/:sectionSlug/segments/:segmentId',
    component: SegmentEdit,
  },
];

const router = createRouter({
  routes,
  history: createWebHistory(),
});

export default router;
