<template>
    <section class="faq">
      <h2>Частые вопросы</h2>
      <div class="faq-item" v-for="faq in faqs" :key="faq.id">
        <h3>{{ faq.question }}</h3>
        <p>{{ faq.answer }}</p>
      </div>
    </section>
  </template>
  
  <script>
  export default {
    name: 'FaqMain',
    data() {
      return {
        faqs: [
          { id: 1, question: 'Вопрос 1', answer: 'Ответ 1' },
          // Add more FAQs as needed
        ]
      };
    }
  }
  </script>
  
  <style scoped>
  /* Ваши стили здесь */
  </style>
  