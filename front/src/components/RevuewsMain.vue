<template>
    <section class="reviews">
      <h2>Отзывы</h2>
      <div class="review-card" v-for="review in reviews" :key="review.id">
        <p>{{ review.text }}</p>
      </div>
    </section>
  </template>
  
  <script>
  export default {
    name: 'ReviewsMain',
    data() {
      return {
        reviews: [
          { id: 1, text: 'Отзыв 1' },
          // Add more reviews as needed
        ]
      };
    }
  }
  </script>
  
  <style scoped>
  /* Ваши стили здесь */
  </style>
  