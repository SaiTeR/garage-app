<script setup>
import { ref } from 'vue';

const phone = ref('');
const firstName = ref('');
const lastName = ref('');
const email = ref('');
const password = ref('');

const register = () => {
  console.log({
    phone: phone.value,
    firstName: firstName.value,
    lastName: lastName.value,
    email: email.value,
    password: password.value,
  });
  alert('Регистрация успешна!');
};
</script>

<template>
  <div class="registration-container">
    <div class="form">
      <h1>РЕГИСТРАЦИЯ</h1>
      <form @submit.prevent="register">
        <label>
          Телефон:
          <input type="tel" v-model="phone" placeholder="+7 (___) ___-__-__" required />
        </label>
        <div class="name-fields">
          <label>
            Имя:
            <input type="text" v-model="firstName" placeholder="Имя" required />
          </label>
          <label>
            Фамилия:
            <input type="text" v-model="lastName" placeholder="Фамилия" required />
          </label>
        </div>
        <label>
          Email:
          <input type="email" v-model="email" placeholder="example@mail.com" required />
        </label>
        <label>
          Пароль:
          <input type="password" v-model="password" placeholder="Введите пароль" required />
        </label>
        <button class="register-button" type="submit">РЕГИСТРАЦИЯ</button>
      </form>
    </div>
  </div>
</template>

<style scoped>
/* Контейнер для страницы */
.registration-container {
  display: flex;
  justify-content: center;

  min-height: 100vh;
  background-color: #f9f9f9; /* Цвет фона, используемый ранее */
  margin-top: 50px;

}

/* Форма регистрации */
.form {
  display: flex;
  flex-direction: column;
  background-color: white;
  border-radius: 15px;
  border: 2px solid #7675fa;
  padding: 30px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  width: 400px;
  height: 500px;
}

h1 {
  font-size: 36px;
  text-align: center;
  margin-bottom: 20px;
  font-family: 'Bahnschrift', sans-serif;
}

label {
  display: flex;
  flex-direction: column;
  font-size: 16px;
  margin-bottom: 15px;
  font-family: Arial, sans-serif;
}

input {
  margin-top: 5px;
  padding: 10px;
  border: 2px solid #7675fa;
  border-radius: 5px;
  font-size: 14px;
}

/* Поля для имени и фамилии */
.name-fields {
  display: flex;
  justify-content: space-between;
  gap: 10px;
}

.name-fields label {
  flex: 1;
}

/* Кнопка регистрации */
.register-button {
  padding: 10px;
  margin-top: 20px;
  background-color: #7675fa;
  color: white;
  border: none;
  border-radius: 10px;
  font-size: 16px;
  cursor: pointer;
  font-family: 'Bahnschrift', sans-serif;
  transition: background-color 0.3s;
}

.register-button:hover {
  background-color: #1e90ff;
}
</style>
