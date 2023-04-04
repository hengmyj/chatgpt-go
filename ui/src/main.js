import { createApp } from 'vue';
import App from './App.vue';
import { Form, Field, CellGroup,List,Popup,Toast } from 'vant';
import { router } from './router';

const app = createApp(App);
app.use(router);
app.use(Form);
app.use(Field);
app.use(CellGroup);
app.use(List);
app.use(Popup);
app.use(Toast);
app.mount('#app');
