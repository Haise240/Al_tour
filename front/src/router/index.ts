import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import Home from '../views/Home.vue';
import TourDetails from '../views/TourDetails.vue';
import Booking from '../views/Booking.vue';
import AdminPanel from '../views/AdminPanel.vue';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/tour/:id',
    name: 'TourDetails',
    component: TourDetails,
    props: true
  },
  {
    path: '/booking',
    name: 'Booking',
    component: Booking
  },
  {
    path: '/admin',
    name: 'AdminPanel',
    component: AdminPanel
  }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});

export default router;
