<template>
    <section class="tours">
      <h2>Однодневные туры</h2>
      <div class="tour-cards">
        <div class="tour-card" v-for="tour in oneDayTours" :key="tour.id">
          <p>{{ tour.title }}</p>
        </div>
      </div>
  
      <h2>Автоэкскурсии</h2>
      <div class="tour-cards">
        <div class="tour-card" v-for="tour in autoTours" :key="tour.id">
          <p>{{ tour.title }}</p>
        </div>
      </div>
    </section>
  </template>
  
  <script>
  export default {
    name: 'ToursMain',
    data() {
      return {
        oneDayTours: [
          { id: 1, title: 'Тур 1', image: '@/assets/tour1.jpg' },
          { id: 2, title: 'Тур 2', image: '@/assets/tour2.jpg' },
          // Add more tours as needed
        ],
        autoTours: [
          { id: 1, title: 'Тур 3', image: '@/assets/tour3.jpg' },
          { id: 2, title: 'Тур 4', image: '@/assets/tour4.jpg' },
          // Add more tours as needed
        ]
      };
    }
  }
  </script>
  
  <style scoped>
  /* Ваши стили здесь */
  </style>
  