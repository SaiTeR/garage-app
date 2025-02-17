<script setup>
import { useRouter } from 'vue-router';
const router = useRouter();
function navigateTo(route) {
  router.push(route); // Переход к маршруту
}
</script>

<template>
  <div class="contacts-container">
    <!-- Блок с номерами телефонов -->
    <div class="contact-block">
      <h2>Телефоны</h2>
      <p>+7 (123) 456-78-90</p>
      <p>+7 (987) 654-32-10</p>
    </div>

    <!-- Вертикальная разделительная полоса -->
    <div class="vertical-divider"></div>

    <!-- Блок с мессенджерами -->
    <div class="contact-block">
      <h2>Мессенджеры</h2>
      <div class="messengers">
        <a href="https://t.me/your_telegram" target="_blank">
          <img src="../assets/contacts/telegram.png" alt="Telegram" />
        </a>
        <a href="https://wa.me/your_whatsapp" target="_blank">
          <img src="../assets/contacts/whatsapp.png" alt="WhatsApp" />
        </a>
        <a href="viber://chat?number=your_number" target="_blank">
          <img src="../assets/contacts/viber.png" alt="Viber" />
        </a>
      </div>
    </div>

    <!-- Вертикальная разделительная полоса -->
    <div class="vertical-divider"></div>

    <!-- Блок с кнопкой записи -->
    <div class="contact-block">
      <h2>Онлайн запись</h2>
      <button  @click="navigateTo('/record')" class="record-button">
        Записаться на обслуживание
      </button>
    </div>
  </div>
</template>

<style scoped>
/* Контейнер для всех блоков */
.contacts-container {
  display: flex;
  align-items: flex-start; /* Выравнивание блоков по верхнему краю */
  justify-content: center;
  gap: 20px; /* Расстояние между блоками и разделителями */
  background-color: #f5f5f5; /* Общий фон сайта */
  padding: 20px;
  min-height: 100vh;
}

/* Блок одной группы контактов */
.contact-block {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start; /* Выравнивание содержимого по верхнему краю */
  background-color: #f5f5f5; /* Фон совпадает с фоном сайта */
  text-align: center;
  padding: 20px;
  width: 300px;
  min-height: 200px; /* Фиксированная минимальная высота для равномерного отображения */

}

/* Заголовки блоков */
.contact-block h2 {
  font-size: 36px;
  color: #333;
  margin-bottom: 15px;
  font-family: 'Bahnschrift', sans-serif;
  height: 40px; /* Фиксированная высота заголовка для выравнивания */
  display: flex;
  align-items: center;
  justify-content: center;
}

/* Стили для текста */
.contact-block p {
  font-size: 24px;
  color: #555;
  margin: 5px 0;
  font-family: 'Bahnschrift', sans-serif;
}

/* Вертикальные разделительные полосы */
.vertical-divider {
  width: 2px;
  background-color: #7675fa; /* Цвет полосы */
  height: 150px; /* Высота полосы */
  margin-top: 40px;
}

/* Стили для мессенджеров */
.messengers {
  display: flex;
  justify-content: center;
  gap: 15px;
}

.messengers img {
  width: 70px;
  height: 70px;
  cursor: pointer;
  transition: transform 0.3s ease;
}

.messengers img:hover {
  transform: scale(1.2); /* Увеличение при наведении */
}

/* Кнопка записи */
.record-button {
  padding: 10px 20px;
  background-color: #7675fa;
  color: white;
  border: none;
  border-radius: 10px;
  font-size: 24px;
  cursor: pointer;
  font-family: 'Bahnschrift', sans-serif;
  transition: background-color 0.3s;
}

.record-button:hover {
  background-color: #1e90ff; /* Синий цвет при наведении */
}
</style>

