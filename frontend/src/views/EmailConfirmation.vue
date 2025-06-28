<script setup>
import { useI18n } from 'vue-i18n'
import { ref, onMounted } from "vue";
import { useRoute } from 'vue-router';

const { t } = useI18n()
const route = useRoute();

const email = ref("");
const confirmationCode = ref("");
const errorMessage = ref("");
const successMessage = ref("");

const disableSendButton = ref(false);

const wrongEmailType = ref("");
const wrongEmailMessage = ref("");

const wrongCodeType = ref("");
const wrongCodeMessage = ref("");

// Get email from query parameter if available
onMounted(() => {
  if (route.query.email) {
    email.value = route.query.email;
    checkForm();
  }
});

function checkForm() {
  const emailValid = checkEmail();
  const codeValid = checkCode();

  if (emailValid && codeValid) {
    disableSendButton.value = false;
    return;
  }

  disableSendButton.value = true;
}

function checkEmail() {
  const regex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  if (regex.test(email.value)) {
    wrongEmailType.value = "is-success";
    wrongEmailMessage.value = "";
    return true;
  }
  wrongEmailType.value = "is-danger";
  wrongEmailMessage.value = t('page.registration.wrongemail');
  return false;
}

function checkCode() {
  if (confirmationCode.value.length > 0) {
    wrongCodeType.value = "is-success";
    wrongCodeMessage.value = "";
    return true;
  }
  wrongCodeType.value = "is-danger";
  wrongCodeMessage.value = "Please enter confirmation code";
  return false;
}

async function confirm() {
  disableSendButton.value = true;
  errorMessage.value = "";
  successMessage.value = "";

  try {
    // Here you would call your API to confirm the email
    // For example:
    // const response = await confirmEmail(email.value, confirmationCode.value);

    // Simulating API call for now
    await new Promise(resolve => setTimeout(resolve, 1000));

    // Success message
    successMessage.value = "Email confirmed successfully!";
  } catch (error) {
    // Error handling
    errorMessage.value = error.message || "Failed to confirm email";
  } finally {
    disableSendButton.value = false;
  }
}
</script>
<template>
  <section class="hero is-link">
    <div class="hero-body">
      <p class="title">{{ t("pageheaders.emailconfirmation") || "Email Confirmation" }}</p>
    </div>
  </section>
  <div class="box">
    <div class="columns is-mobile">
      <div class="column is-half is-offset-one-quarter">
        <p class="bd-notification is-primary">
          <b-field label="Email"
                   v-on:keyup="checkForm"
                   :type="wrongEmailType"
                   :message="wrongEmailMessage">
            <b-input v-model="email"></b-input>
          </b-field>
          <b-field label="Confirmation Code"
                   v-on:keyup="checkForm"
                   :type="wrongCodeType"
                   :message="wrongCodeMessage">
            <b-input v-model="confirmationCode"></b-input>
          </b-field>
          <div v-if="errorMessage" class="notification is-danger">
            {{errorMessage}}
          </div>
          <div v-if="successMessage" class="notification is-success">
            {{successMessage}}
          </div>
          <b-button
              v-on:click="confirm"
              :disabled="disableSendButton"
              type="is-primary"
              expanded
          >{{ t("page.email-confirmation.send") }}
          </b-button>
        </p>
      </div>
    </div>
  </div>
</template>

<style>

</style>
