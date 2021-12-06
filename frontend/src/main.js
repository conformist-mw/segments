import { createApp } from 'vue';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import router from '@/router/router';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap';

import App from './App.vue';

const app = createApp(App);

app.component('font-awesome-icon', FontAwesomeIcon);
app.config.productionTip = false;

app.use(router);
app.mount('#app');
