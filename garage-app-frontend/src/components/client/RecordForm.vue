<script setup>
import { ref } from 'vue';

const carBrand = ref('');
const carModel = ref('');
const virtualGarage = ref('');
const issue = ref('');
const details = ref('');
const date = ref('');
const time = ref('');

const submitForm = () => {
  console.log({
    carBrand: carBrand.value,
    carModel: carModel.value,
    virtualGarage: virtualGarage.value,
    issue: issue.value,
    details: details.value,
    date: date.value,
    time: time.value,
  });
  alert('Форма отправлена!');
};


</script>

<template>
  <div class="form-container">
    <h1>Запись на техническое обслуживание</h1>
    <div class="form-grid">
      <!-- Укажите ваш автомобиль -->
      <div class="form-block">
        <h2>Укажите ваш автомобиль</h2>
        <label>
          Марка:
          <input type="text" v-model="carBrand" />
        </label>
        <label>
          Модель:
          <input type="text" v-model="carModel" />
        </label>
        <p>или</p>
        <label>
          Виртуальный гараж:
          <select v-model="virtualGarage" :disabled="true">
            <option disabled value="">Выберите</option>
            <option>Автомобиль 1</option>
            <option>Автомобиль 2</option>
          </select>
        </label>
      </div>

      <!-- Опишите вашу проблему -->
      <div class="form-block">
        <h2>Опишите вашу проблему</h2>
        <label>
          Выберите вариант из списка:
          <select v-model="issue">
            <option disabled value="">Выберите проблему</option>
            <option>Проблема с двигателем</option>
            <option>Проблема с КПП</option>
            <option>Электрика</option>
          </select>
        </label>
        <label>
          Расскажите нам подробнее:
          <textarea v-model="details" placeholder="Необязательное поле"></textarea>
        </label>
      </div>

      <!-- Выберите дату и время -->
      <div class="form-block">
        <h2>Выберите дату и время</h2>
        <label>
          Дата:
          <input type="date" v-model="date" />
        </label>
        <label>
          Время:
          <input type="time" v-model="time" />
        </label>
        <button class="record-button" @click="submitForm">Оформить запись</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* Контейнер для всей формы */
.form-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
  font-family: 'Bahnschrift', sans-serif;
}

h1 {
  font-size: 36px;
  margin-bottom: 30px;
  color: #222;
  text-align: center;
}

/* Сетка формы */
.form-grid {
  display: flex;
  justify-content: center;
  gap: 30px;
  width: 100%;
  max-width: 1200px;
}

/* Блок формы */
.form-block {
  display: flex;
  flex-direction: column;
  background-color: white;
  border: 2px solid #7675fa;
  border-radius: 10px;
  padding: 20px;
  width: 350px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

h2 {
  font-size: 24px;
  margin-bottom: 15px;
}

label {
  display: flex;
  flex-direction: column;
  font-size: 16px;
  margin-bottom: 10px;
}

input,
select,
textarea {
  margin-top: 5px;
  padding: 8px;
  border: 2px solid #7675fa;
  border-radius: 5px;
  font-size: 14px;
}

textarea {
  resize: none;
  height: 80px;
}

/* Кнопка записи */
.record-button {
  padding: 10px 20px;
  background-color: #7675fa;
  color: white;
  border: none;
  border-radius: 10px;
  font-size: 16px;
  cursor: pointer;
  font-family: 'Bahnschrift', sans-serif;
  transition: background-color 0.3s;
  margin-top: 20px;
  align-self: flex-start;
}

.record-button:hover {
  background-color: #1e90ff; /* Синий цвет при наведении */
}
</style>
