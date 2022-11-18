import { createApp } from 'vue';
import App from './App.vue';
import './normalize.css';
import router from './router';
import './style.scss';

const app = createApp(App);

app.use(router);
app.mount('#app');

export default app;
