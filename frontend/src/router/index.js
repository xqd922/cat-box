import { createRouter, createWebHashHistory } from "vue-router";

import { useLoginStore } from "@/stores/login";

import Layout from "@/views/Layout/index.vue";
import Home from "@/views/Home/index.vue";
import Login from "@/views/Login/index.vue";

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      component: Layout,
      children: [
        {
          path: "",
          name: "home",
          component: Home,
          meta: {
            isAuth: true,
          },
        },
      ],
    },
    {
      path: "/login",
      name: "login",
      component: Login,
    },
  ],
});

router.beforeEach(async (to, from, next) => {
  if (to.meta.isAuth) {
    const loginStore = useLoginStore();
    if (!loginStore.token) {
      next("login");
      return;
    }
  }
  next();
});

export default router;
