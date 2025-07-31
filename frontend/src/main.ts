import { createApp } from "vue";
import App from "./App.vue";
import "./style.css";
import { router } from "./router";

import "@wailsio/runtime";

import { createPinia } from "pinia";
const pinia = createPinia();

createApp(App).use(router).use(pinia).mount("#app");
