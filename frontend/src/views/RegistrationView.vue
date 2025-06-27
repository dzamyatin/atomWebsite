<script setup>
import {useI18n} from 'vue-i18n'
import {useLoginStore} from './../stores/login.js'
import {ref} from "vue";
import router from "@/router/index.js";
import {register} from "./../client/client"

const {t} = useI18n()
const store = useLoginStore()

const email = ref("")
const password = ref("")
const errorMessage = ref("")

const disableRegisterButton = ref(false)

const wrongEmailType = ref("")
const wrongEmailMessage = ref("")

const wrongPasswordType = ref("")
const wrongPasswordMessage = ref("")

function checkRegistration() {
  const emailValid = checkEmail()
  const passwordValid = checkPassword()

  if (emailValid && passwordValid) {
    disableRegisterButton.value = false
    return
  }

  disableRegisterButton.value = true
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
  const regex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*()_+])[A-Za-z\d!@#$%^&*()_+]{8,}$/;
  if (regex.test(password.value)) {
    wrongPasswordType.value = "is-success"
    wrongPasswordMessage.value = ""

    return true
  }

  wrongPasswordType.value = "is-danger"
  wrongPasswordMessage.value = t('page.registration.wrongpassword')

  return false
}

async function registration() {
  disableRegisterButton.value = true

  let response = await register(
      email.value,
      password.value,
  )

  console.log("done")
  console.log(response)

  if (response.error != null) {
    disableRegisterButton.value = false
    errorMessage.value = response.response?.body?.message
    return
  }

  router.push('/profile')
  store.login("some")

  // setTimeout(function () {
  //       disableRegisterButton.value = false
  //       router.push('/profile')
  //       console.log(store.isLoggedIn)
  //       store.login("some")
  //     },
  //     1000,
  // )
}
</script>
<template>
  <section class="hero is-link">
    <div class="hero-body">
      <p class="title">{{ t("pageheaders.registration") }}</p>
    </div>
  </section>
  <div class="box">
    <div class="columns is-mobile">
      <div class="column is-half is-offset-one-quarter">
        <p class="bd-notification is-primary">
          <b-field :label="t('page.registration.email')"
                   v-on:keyup="checkRegistration"
                   :type="wrongEmailType"
                   :message="wrongEmailMessage">
            <b-input v-model="email"></b-input>
          </b-field>
          <b-field :label="t('page.registration.password')"
                   v-on:keyup="checkRegistration"
                   :type="wrongPasswordType"
                   :message="wrongPasswordMessage">
            <b-input v-model="password"></b-input>
          </b-field>
          <div v-if="errorMessage" class="notification is-danger">
            {{errorMessage}}
          </div>
          <b-button
              v-on:click="registration"
              :disabled="disableRegisterButton"
              type="is-primary"
              expanded
          >{{ t("page.registration.registration") }}
          </b-button>
        </p>
      </div>
    </div>
  </div>

</template>

<style>

</style>
