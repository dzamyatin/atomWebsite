import Vue from "vue";
import Router from "vue-router";
import AppHeader from "./layout/AppHeader";
import AppHeaderAtom from "./layout/AppHeaderAtom";
import LoginAtom from "./views/LoginAtom";
import AppFooter from "./layout/AppFooter";
import Components from "./views/Components.vue";
import Main from "./views/Main.vue";
import Landing from "./views/Landing.vue";
import Download from "./views/Download.vue";
import RegisterAtom from "./views/RegisterAtom.vue";
import Profile from "./views/Profile.vue";
import Policy from "./views/Policy.vue";

Vue.use(Router);

export default new Router({
  linkExactActiveClass: "active",
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
    {
      path: "/login",
      name: "login",
      components: {
        header: AppHeaderAtom,
        default: LoginAtom,
        footer: AppFooter
      }
    },
    {
      path: "/register",
      name: "register",
      components: {
        header: AppHeaderAtom,
        default: RegisterAtom,
        footer: AppFooter
      }
    },
    {
      path: "/policy",
      name: "policy",
      components: {
        header: AppHeaderAtom,
        default: Policy,
        footer: AppFooter
      }
    },
    {
      path: "/components",
      name: "components",
      components: {
        header: AppHeader,
        default: Components,
        footer: AppFooter
      }
    },
    {
      path: "/download",
      name: "download",
      components: {
        header: AppHeaderAtom,
        default: Download,
        footer: AppFooter
      }
    },
    {
      path: "/landing",
      name: "landing",
      components: {
        header: AppHeader,
        default: Landing,
        footer: AppFooter
      }
    },
    {
      path: "/profile",
      name: "profile",
      components: {
        header: AppHeader,
        default: Profile,
        footer: AppFooter
      }
    }
  ],
  scrollBehavior: to => {
    if (to.hash) {
      return { selector: to.hash };
    } else {
      return { x: 0, y: 0 };
    }
  }
});
