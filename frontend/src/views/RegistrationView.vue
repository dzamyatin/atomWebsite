<script setup>
import { useI18n } from 'vue-i18n'
import { useLoginStore } from './../stores/login.js'
import {ref} from "vue";
import router from "@/router/index.js";
const { t } = useI18n()
const { store } = useLoginStore()

const email = ref("")
const password = ref("")
const disableRegister = ref(false)

function registration() {
  disableRegister.value = true

  setTimeout(function () {
      disableRegister.value = false
      router.push('/profile')
    },
    1000,
  )
}
</script>
<template>
  <section class="hero is-link">
    <div class="hero-body">
      <p class="title">{{t("pageheaders.registration")}}</p>
    </div>
  </section>
    <div class="box">
      <div class="columns is-mobile">
        <div class="column is-half is-offset-one-quarter">
          <p class="bd-notification is-primary">
            <b-field :label="t('page.registration.email')" type="is-danger" :message="t('page.registration.wrongemail')">
              <b-input v-model="email" ></b-input>
            </b-field>
            <b-field :label="t('page.registration.password')" type="is-success" :message="t('page.registration.wrongpassword')">
              <b-input v-model="password" ></b-input>
            </b-field>
            <b-button
                v-on:click="registration"
                :disabled="disableRegister.valueOf()"
                type="is-primary"
                expanded
            >{{t("page.registration.registration")}}</b-button>
          </p>
        </div>
      </div>
    </div>

</template>

<style>

</style>
