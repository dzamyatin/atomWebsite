<script setup>
import {useI18n} from 'vue-i18n'
import {useLoginStore} from './../stores/login.js'
import {ref} from "vue";
import router from "@/router/index.js";
import {register} from "./../client/client"

const {t} = useI18n()
const store = useLoginStore()

</script>
<template>
  <section class="hero is-link">
    <div class="hero-body">
      <p class="title">{{ t("pageheaders.changepassword") }}</p>
    </div>
  </section>
  <div class="box">
    <div class="columns is-mobile">
      <div class="column is-half is-offset-one-quarter">
        <p class="bd-notification is-primary">
          <div class="tabs">
            <ul>
              <li><a>{{ t("page.changepassword.bypass") }}</a></li>
              <li><a>{{ t("page.changepassword.bycode") }}</a></li>
            </ul>
          </div>
        </p>
      </div>
    </div>
  </div>

</template>

<style>

</style>
