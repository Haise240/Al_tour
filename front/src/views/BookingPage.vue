<template>
<main>
    <section class="intro">
      <h2>Наши номера</h2>
      <form class="booking-form">
        <label for="checkin">Заезд</label>
        <input type="date" id="checkin">
        <label for="checkout">Выезд</label>
        <input type="date" id="checkout">
        <label for="adults">1 взрослых</label>
        <select id="adults">
          <option value="1">1 взрослый</option>
          <option value="2">2 взрослых</option>
        </select>
        <label for="children">0 детей</label>
        <select id="children">
          <option value="0">0 детей</option>
          <option value="1">1 ребенок</option>
        </select>
        <button type="submit">Подобрать</button>
      </form>
    </section>

    <section class="rooms">
      <div class="room">
        <div class="room-details">
          <h3>Полулюкс</h3>
          <p>Из наших полулюксов вы увидите захватывающую панораму города.</p>
          <ul>
            <li>Площадь: 260 кв. ft</li>
            <li>Кровати: полутороспальных 2</li>
          </ul>
          <p class="price">от $250</p>
          <button>О номере</button>
        </div>
      </div>

      <div class="room">
        <div class="room-details">
          <h3>Стандартный</h3>
          <p>В наших стандартных номерах комфорт отлично сочетается с функциональностью.</p>
          <ul>
            <li>Площадь: 230 кв. ft</li>
            <li>Кровати: полутороспальных 1</li>
          </ul>
          <p class="price">от $150</p>
          <button>О номере</button>
        </div>
      </div>

      <div class="room">
        <div class="room-details">
          <h3>Люкс</h3>
          <p>Номера этого типа — удобные, просторные и элегантные.</p>
          <ul>
            <li>Площадь: 280 кв. ft</li>
            <li>Кровати: больших двуспальных 1</li>
          </ul>
          <p class="price">от $350</p>
          <button>О номере</button>
        </div>
      </div>
    </section>
  </main>
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
body {
    font-family: Arial, sans-serif;
    margin: 0;
    padding: 0;
    background-color: #f5f5f5;
}

header {
    background-color: #fff;
    padding: 20px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

h1 {
    margin: 0;
    display: inline-block;
}

nav {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

nav ul {
    list-style: none;
    padding: 0;
    margin: 0;
    display: flex;
    gap: 20px;
}

nav ul li {
    display: inline;
}

nav a {
    text-decoration: none;
    color: #333;
    font-weight: bold;
}

.book-button {
    padding: 10px 20px;
    background-color: #000;
    color: #fff;
    border: none;
    cursor: pointer;
}

.intro {
    text-align: center;
    padding: 50px 20px;
    background-color: #fff;
    margin-bottom: 20px;
}

.booking-form {
    display: flex;
    gap: 10px;
    justify-content: center;
    align-items: center;
}

.booking-form label {
    margin-right: 5px;
}

.rooms {
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 20px;
}

.room {
    display: flex;
    background-color: #fff;
    padding: 20px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.room img {
    max-width: 150px;
    margin-right: 20px;
}

.room-details {
    flex: 1;
}

.room-details h3 {
    margin-top: 0;
}

.room-details ul {
    list-style: none;
    padding: 0;
}

.price {
    font-weight: bold;
    color: #000;
}

button {
    padding: 10px 20px;
    background-color: #000;
    color: #fff;
    border: none;
    cursor: pointer;
}
  </style>
  