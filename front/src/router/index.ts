import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import Home from '../views/HomeMain.vue';
import TourDetails from '../views/TourDetails.vue';
import Booking from '../views/BookingPage.vue';
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
  history: createWebHistory(),
  routes
});

export default router;
