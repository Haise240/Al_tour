<template>
    <div>
      <h1>Админ панель</h1>
      <form @submit.prevent="createTour">
        <div>
          <label for="title">Название тура:</label>
          <input type="text" v-model="newTour.title" />
        </div>
        <div>
          <label for="description">Описание:</label>
          <textarea v-model="newTour.description"></textarea>
        </div>
        <div>
          <label for="price">Цена:</label>
          <input type="number" v-model="newTour.price" />
        </div>
        <div>
          <label for="duration">Длительность (дни):</label>
          <input type="number" v-model="newTour.duration" />
        </div>
        <button type="submit">Создать тур</button>
      </form>
      <div v-if="tours.length">
        <h2>Существующие туры:</h2>
        <div v-for="tour in tours" :key="tour.id">
          <h3>{{ tour.title }}</h3>
          <p>{{ tour.description }}</p>
          <button @click="deleteTour(tour.id)">Удалить</button>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import axios from 'axios'
  
  export default {
    data() {
      return {
        newTour: {
          title: '',
          description: '',
          price: 0,
          duration: 0
        },
        tours: []
      }
    },
    created() {
      this.fetchTours()
    },
    methods: {
      fetchTours() {
        axios.get('http://localhost:8080/tours')
          .then(response => {
            this.tours = response.data
          })
          .catch(error => {
            console.error("There was an error fetching the tours!", error)
          })
      },
      createTour() {
        axios.post('http://localhost:8080/tours', this.newTour)
          .then(response => {
            this.tours.push(response.data)
          })
          .catch(error => {
            console.error("There was an error creating the tour!", error)
          })
      },
      deleteTour(id) {
        axios.delete(`http://localhost:8080/tours/${id}`)
          .then(response => {
            this.tours = this.tours.filter(tour => tour.id !== id)
          })
          .catch(error => {
            console.error("There was an error deleting the tour!", error)
          })
      }
    }
  }
  </script>
  
  <style scoped>
  /* Ваши стили */
  </style>
  