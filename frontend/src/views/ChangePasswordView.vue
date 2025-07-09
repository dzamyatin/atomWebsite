<script setup>
import {useI18n} from 'vue-i18n'
import {useLoginStore} from './../stores/login.js'
import {ref} from "vue";
import router from "@/router/index.js";
import {changePassword} from "./../client/client"
import {useRoute} from "vue-router";

const {t} = useI18n()
const route = useRoute();
const store = useLoginStore()

// Form fields
const oldPassword = ref('')
const newPassword = ref('')
const email = ref('')
const phone = ref('')
const code = ref('')

// UI state
const activeTab = ref('bypass') // 'bypass' or 'bycode'
const isLoading = ref(false)
const errorMessage = ref('')
const successMessage = ref('')

// Change active tab
function setActiveTab(tab) {
  activeTab.value = tab
  // Clear form and messages when switching tabs
  clearForm()
}

// Clear form fields and messages
function clearForm() {
  oldPassword.value = ''
  newPassword.value = ''
  email.value = ''
  phone.value = ''
  code.value = ''
  errorMessage.value = ''
  successMessage.value = ''
}

// Submit form based on active tab
async function submitForm() {
  errorMessage.value = ''
  successMessage.value = ''

  if (activeTab.value === 'bypass') {
    // Validate bypass tab fields
    if (!oldPassword.value || !newPassword.value) {
      errorMessage.value = t('errors.provide_both_passwords')
      return
    }
  } else {
    // Validate bycode tab fields
    if (!newPassword.value) {
      errorMessage.value = t('errors.provide_new_password')
      return
    }
    if (!code.value) {
      errorMessage.value = t('errors.provide_verification_code')
      return
    }
    if (!email.value && !phone.value) {
      errorMessage.value = t('errors.provide_email_or_phone')
      return
    }
  }

  isLoading.value = true

  try {
    // Call API with different parameters based on active tab
    let result
    if (activeTab.value === 'bypass') {
      result = await changePassword(newPassword.value, oldPassword.value, '', '', '')
    } else {
      result = await changePassword(newPassword.value, '', email.value, phone.value, code.value)
    }

    if (result.error) {
      errorMessage.value = result.error.message || t('errors.unknown_error')
    } else {
      successMessage.value = t('success.password_changed')
      clearForm()

      router.push('/login')
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
      <p class="title">{{ t("pageheaders.changepassword") }}</p>
    </div>
  </section>
  <div class="box">
    <div class="columns is-mobile">
      <div class="column is-half is-offset-one-quarter">
        <div class="bd-notification is-primary">
          <!-- Tabs -->
          <div class="tabs">
            <ul>
              <li :class="{ 'is-active': activeTab === 'bypass' }">
                <a @click="setActiveTab('bypass')">{{ t("page.changepassword.bypass") }}</a>
              </li>
              <li :class="{ 'is-active': activeTab === 'bycode' }">
                <a @click="setActiveTab('bycode')">{{ t("page.changepassword.bycode") }}</a>
              </li>
            </ul>
          </div>

          <!-- Success message -->
          <div v-if="successMessage" class="notification is-success">
            {{ successMessage }}
          </div>

          <!-- Error message -->
          <div v-if="errorMessage" class="notification is-danger">
            {{ errorMessage }}
          </div>

          <!-- Form for changing password by old password -->
          <form v-if="activeTab === 'bypass'" @submit.prevent="submitForm">
            <div class="field">
              <label class="label">{{ t('fields.oldPassword') }}</label>
              <div class="control">
                <input 
                  class="input" 
                  type="password" 
                  v-model="oldPassword" 
                  :placeholder="t('placeholders.oldPassword')"
                >
              </div>
            </div>

            <div class="field">
              <label class="label">{{ t('fields.newPassword') }}</label>
              <div class="control">
                <input 
                  class="input" 
                  type="password" 
                  v-model="newPassword" 
                  :placeholder="t('placeholders.newPassword')"
                >
              </div>
            </div>

            <div class="field">
              <div class="control">
                <button 
                  class="button is-primary" 
                  type="submit" 
                  :class="{ 'is-loading': isLoading }" 
                  :disabled="isLoading"
                >
                  {{ t('buttons.change_password') }}
                </button>
              </div>
            </div>
          </form>

          <!-- Form for changing password by code -->
          <form v-if="activeTab === 'bycode'" @submit.prevent="submitForm">
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
              <label class="label">{{ t('fields.code') }}</label>
              <div class="control">
                <input 
                  class="input" 
                  type="text" 
                  v-model="code" 
                  :placeholder="t('placeholders.code')"
                >
              </div>
            </div>

            <div class="field">
              <label class="label">{{ t('fields.newPassword') }}</label>
              <div class="control">
                <input 
                  class="input" 
                  type="password" 
                  v-model="newPassword" 
                  :placeholder="t('placeholders.newPassword')"
                >
              </div>
            </div>

            <div class="field">
              <div class="control">
                <button 
                  class="button is-primary" 
                  type="submit" 
                  :class="{ 'is-loading': isLoading }" 
                  :disabled="isLoading"
                >
                  {{ t('buttons.change_password') }}
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
