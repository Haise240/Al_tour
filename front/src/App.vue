<template>
  <div id="app">
    <nav>
      <router-link to="/">Home</router-link>
      <router-link to="/booking">Booking</router-link>
    </nav>
    <router-view/>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'App'
});
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

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
