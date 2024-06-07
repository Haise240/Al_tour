<template>
    <div>
      <h1>Бронирование тура</h1>
      <form @submit.prevent="submitBooking">
        <div>
          <label for="tour">Тур:</label>
          <select v-model="booking.tourId">
            <option v-for="tour in tours" :key="tour.id" :value="tour.id">{{ tour.title }}</option>
          </select>
        </div>
        <div>
          <label for="date">Дата:</label>
          <input type="date" v-model="booking.date" />
        </div>
        <button type="submit">Забронировать</button>
      </form>
    </div>
  </template>
  
  <script>
  import axios from 'axios'
  
  export default {
    data() {
      return {
        tours: [],
        booking: {
          tourId: null,
          date: ''
        }
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
    },
    methods: {
      submitBooking() {
        axios.post('http://localhost:8080/bookings', this.booking)
          .then(response => {
            console.log('Booking successful', response.data)
          })
          .catch(error => {
            console.error("There was an error making the booking!", error)
          })
      }
    }
  }
  </script>
  
  <style scoped>
  /* Ваши стили */
  </style>
  