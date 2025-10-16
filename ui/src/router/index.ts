import { createRouter, createWebHistory } from 'vue-router'
import BaseLayout from '@/layout/index.vue'
import remainingRouter from "./modules/remaining";

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const modules: Record<string, any> = import.meta.glob(
  [
    "./modules/**/*.ts",
    "!./modules/**/remaining.ts",
  ],
  {
    eager: true
  }
);

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const routes:any[] = []

Object.keys(modules).forEach(key => {
  routes.push(modules[key].default);
});

const constantMenus = routes.concat(...remainingRouter);

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes:constantMenus
})

export default router
