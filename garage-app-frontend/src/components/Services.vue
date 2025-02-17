<script setup>
</script>

<template>
  <div class="container">

    <div class="service">
      <h2>Ремонт ДВС</h2>
      <div style="display: flex; align-items: center; justify-content: center; flex-direction: row;">
        <img src="../assets/services/engine.png" alt="Двигатель">
        <div class="info">
          <ul>
            <li>Диагностика и ремонт неисправностей двигателя</li>
            <li>Капитальный ремонт двигателя</li>
            <li>Ремонт системы охлаждения</li>
          </ul>
        </div>
      </div>
    </div>

    <div class="service">
      <h2>Ремонт КПП</h2>
      <div style="display: flex; align-items: center; justify-content: center; flex-direction: row;">
        <img src="../assets/services/gearbox.png" alt="Коробка передач">
        <div class="info">
          <ul>
            <li>Ремонт механических коробок передач</li>
            <li>Ремонт автоматических трансмиссий</li>
          </ul>
        </div>
      </div>
    </div>

    <div class="service">
      <h2>Автоэлектрика</h2>
      <div style="display: flex; align-items: center; justify-content: center; flex-direction: row;">
        <img src="../assets/services/electric.png" alt="Электрика">
        <div class="info">
          <ul>
            <li>Диагностика электрооборудования</li>
            <li>Замена аккумулятора</li>
            <li>Ремонт стартера и генератора</li>
          </ul>
        </div>
      </div>
    </div>

    <div class="service">
      <h2>Подвеска и ходовая часть</h2>
      <div style="display: flex; align-items: center; justify-content: center; flex-direction: row;">
        <img src="../assets/services/amort.png" alt="Электрика">
        <div class="info">
          <ul>
            <li>Диагностика подвески</li>
            <li>Замена амортизаторов, стоек, пружин</li>
            <li>Регулировка сход-развала</li>
          </ul>
        </div>
      </div>
    </div>

    <div class="service">
      <h2>Тормозная система</h2>
      <div style="display: flex; align-items: center; justify-content: center; flex-direction: row;">
        <img src="../assets/services/brakes.png" alt="Электрика">
        <div class="info">
          <ul>
            <li>Замена тормозных колодок и дисков</li>
            <li>Ремонт или замена суппортов</li>
            <li>Замена тормозной жидкости</li>
          </ul>
        </div>
      </div>
    </div>

    <div class="service">
      <h2>Кузовной ремонт</h2>
      <div style="display: flex; align-items: center; justify-content: center; flex-direction: row;">
        <img src="../assets/services/body.png" alt="Электрика">
        <div class="info">
          <ul>
            <li>Ремонт вмятин и повреждений кузова</li>
            <li>Ремонт и замена стекол</li>
            <li>Антикоррозийная обработка</li>
          </ul>
        </div>
      </div>
    </div>






  </div>
</template>

<style scoped>


/* Контейнер для всех услуг */
.container {
  display: flex;
  flex-wrap: wrap; /* Перенос строк, если блоки не помещаются */
  gap: 20px; /* Расстояние между блоками */
  justify-content: center; /* Центрирование блоков */
  align-content: flex-start;
  padding: 20px;
  min-height: 100vh;
}

/* Блок одной услуги */
.service {
  display: flex;
  flex-direction: column;
  background-color: white;
  border-radius: 15px;
  border: 2px solid #7675fa; /* Цвет рамки */
  width: 400px; /* Ширина блока */
  height: 200px; /* Высота блока */
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1); /* Тень для блока */
  align-items: center;
  padding: 30px;
}

/* Информация в блоке */
.info {
  margin-left: 20px;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

/* Стили для заголовка */
h2 {
  font-size: 24px;
  color: #222;
  margin-bottom: 10px;
  font-family: 'Bahnschrift', sans-serif;
}

/* Стили для списка */
ul {
  padding-left: 20px; /* Отступ слева для списков */
}

li {
  font-size: 16px;
  line-height: 1.5;
  color: #555;
  margin-bottom: 5px;
  font-family: Arial, sans-serif;
  font-weight: bold;
}

/* Картинки */
.service img {
  width: 100px;
  height: 100px;
}


</style>
