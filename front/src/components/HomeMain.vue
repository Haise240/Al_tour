<template>
    <div>
      <h1>Туры</h1>
      <div v-if="tours.length">
        <div v-for="tour in tours" :key="tour.id">
          <router-link :to="{ name: 'TourDetails', params: { id: tour.id } }">
            <h2>{{ tour.title }}</h2>
            <p>{{ tour.description }}</p>
          </router-link>
        </div>
      </div>
      <div v-else>
        <p>Нет доступных туров</p>
      </div>
    </div>
  </template>
  
  <script>
  import axios from 'axios'
  
  export default {
    data() {
      return {
        tours: []
      }
    },
    created() {
      axios.get('http://localhost:8080/tours')
        .then(response => {
          this.tours = response.data
        })
        .catch(error => {
          console.error("There was an error fetching the tours!", error)
        })
    }
  }
  </script>
  
  <style scoped>
  /* Ваши стили */
  </style>
  