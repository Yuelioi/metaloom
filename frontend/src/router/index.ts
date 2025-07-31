// router/index.ts
import { createRouter, createWebHistory } from "vue-router";
import Home from "@/pages/Home.vue";
import Previewer from "@/pages/Previewer.vue";

const routes = [
  { path: "/", component: Home },
  { path: "/previewer", component: Previewer },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
