<script setup>
import { useI18n } from 'vue-i18n'
import {BNavbarItem} from "buefy";
import {useLoginStore} from './../stores/login.js'
import {ref} from "vue";
import router from "@/router/index.js";
import {login} from "./../client/client"

const { t } = useI18n()
const store = useLoginStore()

const email = ref("")
const password = ref("")
const errorMessage = ref("")

const disableLoginButton = ref(false)

const wrongEmailType = ref("")
const wrongEmailMessage = ref("")

const wrongPasswordType = ref("")
const wrongPasswordMessage = ref("")

function checkLogin() {
  const emailValid = checkEmail()
  const passwordValid = checkPassword()

  if (emailValid && passwordValid) {
    disableLoginButton.value = false
    return
  }

  disableLoginButton.value = true
}

function checkEmail() {
  const regex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  if (regex.test(email.value)) {
    wrongEmailType.value = "is-success"
    wrongEmailMessage.value = ""

    return true
  }
  wrongEmailType.value = "is-danger"
  wrongEmailMessage.value = t('page.registration.wrongemail')

  return false
}

function checkPassword() {
  if (password.value.length > 0) {
    wrongPasswordType.value = "is-success"
    wrongPasswordMessage.value = ""

    return true
  }

  wrongPasswordType.value = "is-danger"
  wrongPasswordMessage.value = t('page.registration.wrongpassword')

  return false
}

async function loginUser() {
  disableLoginButton.value = true

  let response = await login(
      email.value,
      password.value,
  )

  if (response.error != null) {
    disableLoginButton.value = false
    errorMessage.value = response.response?.body?.message
    return
  }

  router.push('/profile')
  store.login(response.data.token)
}
</script>
<template>
  <section class="hero is-link">
    <div class="hero-body">
      <p class="title">{{t("pageheaders.login")}}</p>
    </div>
  </section>
  <div class="box">
    <div class="columns is-mobile">
      <div class="column is-half is-offset-one-quarter">
        <p class="bd-notification is-primary">
          <b-field :label="t('page.registration.email')"
                   v-on:keyup="checkLogin"
                   :type="wrongEmailType"
                   :message="wrongEmailMessage">
            <b-input v-model="email"></b-input>
          </b-field>
          <b-field :label="t('page.registration.password')"
                   v-on:keyup="checkLogin"
                   :type="wrongPasswordType"
                   :message="wrongPasswordMessage">
            <b-input v-model="password" type="password"></b-input>
          </b-field>
          <div v-if="errorMessage" class="notification is-danger">
            {{errorMessage}}
          </div>
          <p>
            <b-navbar-item tag="router-link" :to="{ path: '/remember-password' }">
              <a>{{ t("page.login.remember-password") }}</a>
            </b-navbar-item>
          </p>
          <b-button
              v-on:click="loginUser"
              :disabled="disableLoginButton"
              type="is-primary"
              expanded
          >{{ t("pageheaders.login") }}
          </b-button>
        </p>
      </div>
    </div>
  </div>
</template>

<style>

</style>
