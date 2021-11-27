import { createApp } from 'vue';
import PrimeVue from 'primevue/config';
import Button from 'primevue/button';
import Chip from 'primevue/chip';
import Panel from 'primevue/panel';

import 'primevue/resources/themes/saga-blue/theme.css';
import 'primevue/resources/primevue.min.css';
import 'primeicons/primeicons.css';
import 'primeflex/primeflex.css';

import App from './App.vue';

const app = createApp(App);
app.use(PrimeVue);
app.mount('#app');
app.component('Button', Button);
app.component('Chip', Chip);
app.component('Panel', Panel);
