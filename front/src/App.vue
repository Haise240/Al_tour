<template>
    <div id="app">
      <HeaderMain/>
      <IntroductionMain />
      <ToursMain />
      <ReviewsMain />
      <FaqMain />
      <ContactMain />
      <FooterMain />
    </div>
  </template>
  
  <script>
  import HeaderMain from './components/HeaderMain.vue'
  import IntroductionMain from './components/IntroductionMain.vue'
  import ToursMain from './components/ToursMain.vue'
  import ReviewsMain from './components/RevuewsMain.vue'
  import FaqMain from './components/FaqMain.vue'
  import ContactMain from './components/ContactMain.vue'
  import FooterMain from './components/FooterMain.vue'
  
  export default {
    name: 'App',
    components: {
      HeaderMain,
      IntroductionMain,
      ToursMain,
      ReviewsMain,
      FaqMain,
      ContactMain,
      FooterMain
    }
  }
  </script>
  
  <style>
  /* Глобальные стили */
  body {
    font-family: Arial, sans-serif;
    margin: 0;
    padding: 0;
    color: #333;
  }

header {
    background-image: url('header-background.jpg');
    background-size: cover;
    color: white;
    padding: 20px 0;
    text-align: center;
}

.header-content {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0 20px;
}

.logo {
    height: 50px;
}

nav ul {
    list-style: none;
    margin: 0;
    padding: 0;
    display: flex;
}

nav ul li {
    margin-left: 20px;
}

nav ul li a {
    color: white;
    text-decoration: none;
}

.introduction {
    text-align: center;
    padding: 20px;
}

.introduction img {
    max-width: 100%;
    height: auto;
}

.tours {
    padding: 20px;
    text-align: center;
}

.tour-cards {
    display: flex;
    justify-content: space-around;
    flex-wrap: wrap;
}

.tour-card {
    width: 45%;
    margin: 10px 0;
    border: 1px solid #ddd;
    padding: 10px;
}

.tour-card img {
    max-width: 100%;
    height: auto;
}

.reviews, .faq, .contact {
    padding: 20px;
}

footer {
    background-color: #333;
    color: white;
    text-align: center;
    padding: 10px 0;
}

  </style>
  