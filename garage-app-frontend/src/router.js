import { createRouter, createWebHistory } from 'vue-router';
import MainComponent from '@/components/Main.vue';
import ServicesComponent from '@/components/Services.vue';
import ReviewsComponent from '@/components/Reviews.vue';
import ContactsComponent from '@/components/Contacts.vue';
import StuffComponent from '@/components/Stuff.vue';
import RecordForm from "@/components/client/RecordForm.vue";
import Registration from "@/components/auth/Registration.vue";
import LoginClient from "@/components/auth/LoginClient.vue";
import LoginWorker from "@/components/auth/LoginWorker.vue";
import WhoAuth from "@/components/auth/WhoAuth.vue";

const routes = [
    { path: '/', component: MainComponent }, // Главная страница
    { path: '/services', component: ServicesComponent }, // Услуги
    { path: '/reviews', component: ReviewsComponent }, // Отзывы
    { path: '/contacts', component: ContactsComponent }, // Контакты
    { path: '/team', component: StuffComponent }, // Команда
    { path: '/record', component: RecordForm },
    { path: '/auth', component: WhoAuth },

    { path: '/client/registration', component: Registration },
    { path: '/client/login', component: LoginClient },
    { path: '/client/profile', component: LoginClient }, // ЛК

    { path: '/worker/login', component: LoginWorker },
];

const router = createRouter({
    history: createWebHistory(), // Используем режим истории
    routes,
});

export default router;
