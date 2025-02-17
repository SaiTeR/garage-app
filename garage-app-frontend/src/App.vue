<template>
  <div class="container">
    <header>
      <div class="logo" @click="navigateTo('/')">
        <img src="@/assets/logo.png" alt="Наш Гараж" />
      </div>
      <div class="nav-buttons">
        <button @click="navigateTo('/record')">ЗАПИСЬ</button>
        <button @click="navigateTo('/services')">УСЛУГИ</button>
        <button @click="navigateTo('/reviews')">ОТЗЫВЫ</button>
        <button @click="navigateTo('/contacts')">КОНТАКТЫ</button>
        <button @click="navigateTo('/team')">КОМАНДА</button>
        <button @click="navigateTo('/auth')" style="margin-left: 350px;">ВХОД</button>
      </div>

    </header>

    <main>
      <router-view />
    </main>
  </div>
</template>

<script>
export default {
  name: "App",
  methods: {
    navigateTo(route) {
      this.$router.push(route); // Переключение маршрута
    },
  },
};
</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
}

body {
  font-family: Arial, sans-serif;
  background-color: #f8f8f8;
  color: #000;
  min-height: 100vh;
}

template {
  min-height: 100vh;
  background-color: #f8f8f8;
}

.container {
  margin: 0;
  background-color: #f8f8f8;

}



header {
  display: flex;
  align-items: center;
  padding: 20px 0;
  border-bottom: 2px solid rgba(118, 117, 250, 0.25); /* Снижена насыщенность с помощью альфа-канала (0.5) */
  height: 100px;
}



.nav-buttons {
  display: flex;
  gap: 30px;
  margin-left: 40px;

}



.nav-buttons button {
  padding: 10px 20px;
  color: #333;
  background-color: transparent; /* Прозрачный фон */
  border: none; /* Убираем всю рамку */
  border-bottom: 4px solid transparent; /* Сделаем нижнюю рамку прозрачной по умолчанию */
  cursor: pointer;
  font-size: 34px;
  transition: color 0.3s, border-bottom-color 0.3s; /* Плавный переход только для цвета текста и нижней рамки */

  font-family: 'Bahnschrift', sans-serif;
  font-weight: bold;
}

.nav-buttons button:hover {
  color: #7675fa; /* Синий цвет текста */
  border-bottom-color: #7675fa; /* Синяя полоска снизу кнопки */
}

.logo img {
  height: 120px; /* Высота изображения 120px */
  width: auto; /* Ширина будет подстраиваться пропорционально */
}

#app {
  min-height: 100vh; /* Делаем так, чтобы содержимое заполняло всю высоту окна */
  background-color: #f8f8f8; /* Убедитесь, что главный контейнер также имеет этот фон */
}
</style>

