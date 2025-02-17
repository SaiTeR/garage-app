<script setup>
import { useRouter } from 'vue-router';

// Инициализация роутера
const router = useRouter();

// Метод для навигации
const navigateTo = (route) => {
  router.push(route);
};
</script>

<template>
  <div class="split-container">
    <!-- Блок "Я - КЛИЕНТ" -->
    <div class="block client-block">
      <div class="interactive-area">
        <img @click="navigateTo('/client/login')" src="@/assets/auth/customers.png" alt="Иконка клиента" class="icon" />
        <p>КЛИЕНТ</p>
      </div>
    </div>

    <!-- Блок "Я - СОТРУДНИК" -->
    <div class="block employee-block">
      <div class="interactive-area">
        <img @click="navigateTo('/worker/login')" src="@/assets/auth/worker.png" alt="Иконка сотрудника" class="icon" />
        <p>СОТРУДНИК</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* Контейнер для двух блоков */
.split-container {
  display: flex;
  justify-content: center;
  min-height: 100vh; /* Фон на всю высоту экрана */
  background-color: #f9f9f9; /* Фоновый цвет */
  gap: 200px;
  margin-top: 50px;
}

/* Общий стиль для блоков */
.block {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  padding: 20px;
}

/* Интерактивная область (иконка + текст) */
.interactive-area {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  cursor: pointer; /* Курсор указателя при наведении */
  transition: color 0.3s ease; /* Плавный переход */
}

/* Стили текста */
.interactive-area p {
  font-size: 36px;
  font-family: 'Bahnschrift', sans-serif;
  color: #333;
  margin-top: 10px;
  transition: color 0.3s ease; /* Плавный переход цвета */
}

/* Стили иконок */
.icon {
  width: 200px; /* Размер иконки */
  height: 200px;
  transition: transform 0.3s ease; /* Плавное увеличение размера */
}

/* Изменение цвета текста и увеличение иконки при наведении */
.interactive-area:hover p {
  color: #7675fa; /* Цвет текста при наведении */
}

.interactive-area:hover .icon {
  transform: scale(1.1); /* Увеличение иконки при наведении */
}
</style>
