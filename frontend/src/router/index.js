import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Vue from "vue";
// import Router from "vue-router";
// import AppHeader from "../layout/AppHeader";
import AppHeaderAtom from "../layout/AppHeaderAtom.vue";
import LoginAtom from "../views/LoginAtom.vue";
import AppFooter from "../layout/AppFooter.vue";
import Components from "../views/Components.vue";
import Main from "../views/Main.vue";
import Landing from "../views/Landing.vue";
import Download from "../views/Download.vue";
import RegisterAtom from "../views/RegisterAtom.vue";
import Profile from "../views/Profile.vue";
import Policy from "../views/Policy.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "main",
      components: {
        header: AppHeaderAtom,
        default: Main,
        footer: AppFooter
      }
    },
    // {
    //   path: '/',
    //   name: 'home',
    //   component: HomeView,
    // },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue'),
    },
  ],
})

export default router
