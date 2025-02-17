import { createApp } from 'vue';
import App from './App.vue';
import router from '@/router.js'; // Импортируем роутер

const app = createApp(App);

app.use(router); // Добавляем роутер в приложение
app.mount('#app'); // Монтируем приложение
