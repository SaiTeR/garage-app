<template>
  <div class="main-page">
    <img src="../assets/main/workers_main_page.png" alt="" style="margin-left: -9px;">
    <div style="margin-left: 8px;">
      <h1>Добро пожаловать в <br><span style="color: #7675fa">Наш Гараж!</span> </h1>
      <p>
        Мы рады видеть вас в нашем автосервисе! Здесь вы найдете самые качественные услуги по ремонту автомобилей,
        а наша команда профессионалов позаботится о вашем автомобиле как о своем.
      </p>
      <p>
        Независимо от того, требуется ли вашему автомобилю небольшая настройка или капитальный ремонт,
        мы всегда готовы помочь. Звоните, записывайтесь онлайн или приезжайте — мы всегда рядом, чтобы сделать
        вашу поездку безопасной и комфортной!
      </p>
      <button @click="navigateTo('/record')" class="cta-button" style="margin-top: 25px">
        Записаться на<br>техническое обслуживание
      </button>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router';

// Использование роутера
const router = useRouter();

function navigateTo(route) {
  router.push(route); // Переход к маршруту
}
</script>

<style scoped>
.main-page {
  text-align: center;
  display: flex;
  flex-direction: row;

}

h1 {
  font-size: 52px;
  color: #333;
  margin-bottom: 20px;
  font-family: 'Bahnschrift', sans-serif;
  font-weight: bold;
}

p {
  font-size: 25px;
  line-height: 1.6;
  color: #555;
  margin-bottom: 20px;
  font-family: Arial, sans-serif;
}

.cta-button {
  padding: 12px 30px;
  background-color: #7675fa;
  color: white;
  border: none;
  border-radius: 15px;
  cursor: pointer;
  font-size: 37px;
  transition: background-color 0.3s;
  font-family: 'Bahnschrift', sans-serif;
}

.cta-button:hover {
  background-color: #1e90ff;
}
</style>
