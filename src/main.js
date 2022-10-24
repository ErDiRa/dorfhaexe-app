import router from 'router.js';
import { createApp } from 'vue';
import App from './App.vue';
import './normalize.css';
import './style.scss';

const app = createApp(App)

app.use(router)

app.mount('#app')

