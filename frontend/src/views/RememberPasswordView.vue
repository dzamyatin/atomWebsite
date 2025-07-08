<script setup>
import { useI18n } from 'vue-i18n'
import {BNavbarItem} from "buefy";
import {useLoginStore} from './../stores/login.js'
import {ref} from "vue";
import router from "@/router/index.js";
import {rememberPassword} from "./../client/client"

const { t } = useI18n()
const store = useLoginStore()

const email = ref('')
const phone = ref('')
const isLoading = ref(false)
const errorMessage = ref('')
const successMessage = ref('')

async function submitForm() {
  errorMessage.value = ''
  successMessage.value = ''

  // Validate that at least one field is filled
  if (!email.value && !phone.value) {
    errorMessage.value = t('errors.provide_email_or_phone')
    return
  }

  isLoading.value = true

  try {
    const result = await rememberPassword(email.value, phone.value)

    if (result.error) {
      errorMessage.value = result.error.message || t('errors.unknown_error')
    } else {
      successMessage.value = t('success.password_reset_sent')
      // Clear form after successful submission
      email.value = ''
      phone.value = ''
    }
  } catch (error) {
    errorMessage.value = error.message || t('errors.unknown_error')
  } finally {
    isLoading.value = false
  }
}
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
        <div class="bd-notification is-primary">
          <!-- Success message -->
          <div v-if="successMessage" class="notification is-success">
            {{ successMessage }}
          </div>

          <!-- Error message -->
          <div v-if="errorMessage" class="notification is-danger">
            {{ errorMessage }}
          </div>

          <form @submit.prevent="submitForm">
            <div class="field">
              <label class="label">{{ t('fields.email') }}</label>
              <div class="control">
                <input 
                  class="input" 
                  type="email" 
                  v-model="email" 
                  :placeholder="t('placeholders.email')"
                >
              </div>
            </div>

            <div class="field">
              <label class="label">{{ t('fields.phone') }}</label>
              <div class="control">
                <input 
                  class="input" 
                  type="tel" 
                  v-model="phone" 
                  :placeholder="t('placeholders.phone')"
                >
              </div>
              <p class="help">{{ t('help.provide_email_or_phone') }}</p>
            </div>

            <div class="field">
              <div class="control">
                <button 
                  class="button is-primary" 
                  type="submit" 
                  :class="{ 'is-loading': isLoading }" 
                  :disabled="isLoading"
                >
                  {{ t('buttons.reset_password') }}
                </button>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<style>

</style>
