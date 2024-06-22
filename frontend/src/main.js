// Vuetify
import "vuetify/styles";
import "@mdi/font/css/materialdesignicons.css";

import { createApp } from "vue";
import { createPinia } from "pinia";

// Vuetify
import { createVuetify } from "vuetify";
const vuetify = createVuetify();

import App from "./App.vue";
import router from "./router";

const app = createApp(App);

app.use(vuetify);
app.use(createPinia());
app.use(router);

app.mount("#app");
