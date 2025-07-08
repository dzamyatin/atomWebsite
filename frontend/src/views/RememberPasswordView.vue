<script setup>
import { useI18n } from 'vue-i18n'
import {BNavbarItem} from "buefy";
import {useLoginStore} from './../stores/login.js'
import {ref} from "vue";
import router from "@/router/index.js";
import {login} from "./../client/client"

const { t } = useI18n()
const store = useLoginStore()

</script>
<template>
  <section class="hero is-link">
    <div class="hero-body">
      <p class="title">{{t("pageheaders.remember-password")}}</p>
    </div>
  </section>
  <div class="box">
    <div class="columns is-mobile">
      <div class="column is-half is-offset-one-quarter">
        <p class="bd-notification is-primary">

        </p>
      </div>
    </div>
  </div>
</template>

<style>

</style>
