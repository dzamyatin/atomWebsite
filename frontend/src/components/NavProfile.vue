<script setup>

import {BNavbarItem} from "buefy";
import {useI18n} from 'vue-i18n'
import {useLoginStore} from './../stores/login.js'
import router from "@/router/index.js";
const {t} = useI18n()

const store = useLoginStore();

function logout() {
  store.logout()
  router.push('/login')
}

</script>

<template>
  <div v-if="store.isLoggedIn.valueOf()">
    <div class="buttons">
      <b-navbar-item tag="router-link" :to="{ path: '/profile' }">
        <a class="button is-primary">
          <strong>{{ t("nav.profile") }}</strong>
        </a>
      </b-navbar-item>
      <b-navbar-item tag="router-link" :to="{ path: '/licenses' }">
        <a class="button is-primary">
          <strong>{{ t("nav.license") || "Licenses" }}</strong>
        </a>
      </b-navbar-item>
      <b-navbar-item tag="router-link" :to="{ path: '/shop' }">
        <a class="button is-primary">
          <strong>{{ t("nav.shop") || "Shop" }}</strong>
        </a>
      </b-navbar-item>
      <b-navbar-item>
        <a class="button is-light" @click="logout">
          {{ t("nav.logout") || "Logout" }}
        </a>
      </b-navbar-item>
    </div>
  </div>
  <div v-else>
    <div class="buttons">
      <b-navbar-item tag="router-link" :to="{ path: '/registration' }">
        <a class="button is-primary">
          <strong>{{ t("nav.registration") }}</strong>
        </a>
      </b-navbar-item>
      <b-navbar-item tag="router-link" :to="{ path: '/login' }">
        <a class="button is-light" tag="router-link" :to="{ path: '/login' }">
          {{ t("nav.signin") }}
        </a>
      </b-navbar-item>
    </div>
  </div>
</template>

<style scoped>

</style>
