import { createRouter, createWebHistory } from "vue-router";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "public",
      component: () => import("../views/PublicView.vue"),
    },
    {
      path: "/app",
      name: "home",
      component: () => import("../views/HomeView.vue"),
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: "/login",
      name: "login",
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("../views/auth/LoginView.vue"),
      meta: {
        hiddenIfLoggedIn: true,
      },
    },
    {
      path: "/register",
      name: "register",
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("../views/auth/RegisterView.vue"),
      meta: {
        hiddenIfLoggedIn: true,
      },
    },
    {
      path: "/profile",
      name: "profile",
      component: () => import("../views/ProfileView.vue"),
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: "/channels/explore",
      name: "channels",
      component: () => import("../views/ChannelsView.vue"),
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: "/channels",
      name: "my-channels",
      component: () => import("../views/MyChannelsView.vue"),
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: "/channels/create",
      name: "create-channel",
      component: () => import("../views/CreateChannelsView.vue"),
      meta: {
        requiresAuth: true,
      },
    },
  ],
});

export default router;
