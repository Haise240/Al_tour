<template>
    <div>
      <h1>{{ tour.title }}</h1>
      <p>{{ tour.description }}</p>
      <p>Цена: {{ tour.price }}</p>
      <p>Длительность: {{ tour.duration }} дней</p>
      <button @click="bookTour">Забронировать тур</button>
    </div>
  </template>
  
  <script>
  import axios from 'axios'
  
  export default {
    data() {
      return {
        tour: {}
      }
    },
    created() {
      const tourId = this.$route.params.id
      axios.get(`http://localhost:8080/tours/${tourId}`)
        .then(response => {
          this.tour = response.data
        })
        .catch(error => {
          console.error("There was an error fetching the tour details!", error)
        })
    },
    methods: {
      bookTour() {
        // Логика бронирования тура
      }
    }
  }
  </script>
  
  <style scoped>
  /* Ваши стили */
  </style>
  