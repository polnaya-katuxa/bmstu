import { createApp } from 'vue';
import { Marked, Renderer } from '@ts-stack/markdown';
import App from './App.vue';

import router from './router';
import store from './store';

import 'bootstrap/dist/css/bootstrap.css';
import './assets/style.css';

import 'bootstrap/dist/js/bootstrap';
// eslint-disable-next-line
import 'vue-popperjs/dist/vue-popper';

Marked.setOptions({
  renderer: new Renderer(),
  gfm: true,
  tables: true,
  breaks: false,
  pedantic: false,
  sanitize: false,
  smartLists: true,
  smartypants: false,
});

createApp(App).use(router).use(store).mount('#app');
