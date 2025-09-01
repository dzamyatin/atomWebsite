<script setup>
  import {useI18n} from 'vue-i18n'
  import {useLoginStore} from './../stores/login.js'
  import {ref, onMounted} from "vue";
  import router from "@/router/index.js";
  import {getCurrentUser, sendEmailConfirmation, sendPhoneConfirmation, confirmEmail as confirmEmailApi, confirmPhone as confirmPhoneApi} from "./../client/client"
  import {BButton, BModal, BInput, BField} from "buefy";

  const {t} = useI18n()
  const store = useLoginStore()

  if (!store.isLoggedIn) {
    router.push('/login')
  }

  // Define reactive properties
  const email = ref("")
  const phone = ref("")
  const wrongEmailType = ref("")
  const wrongEmailMessage = ref("")
  const disableSendButton = ref(false)
  const wrongPhoneType = ref("")
  const wrongPhoneMessage = ref("")
  const disableSendPhoneButton = ref(false)
  const isEmailConfirmed = ref(false)
  const isPhoneConfirmed = ref(false)

  // Modal properties
  const isModalActive = ref(false)
  const confirmationCode = ref("")
  const confirmationType = ref("") // "email" or "phone"
  const confirmationError = ref("")
  const isConfirmationLoading = ref(false)

  // Fetch current user data when component is mounted
  onMounted(async () => {
    try {
      const result = await getCurrentUser()
      if (!result.error && result.data) {
        email.value = result.data.email || ""
        phone.value = result.data.phone || ""
        isEmailConfirmed.value = result.data.confirmedEmail || false
        isPhoneConfirmed.value = result.data.confirmedPhone || false

        // Disable email confirmation button if email is already confirmed
        if (isEmailConfirmed.value) {
          disableSendButton.value = true
        } else {
          checkFormEmail() // Check email validity
        }

        // Disable phone confirmation button if phone is already confirmed
        if (isPhoneConfirmed.value) {
          disableSendPhoneButton.value = true
        }
      }
    } catch (error) {
      console.error("Error fetching user data:", error)
    }
  })

  function checkFormEmail() {
    const emailValid = checkEmail()

    isEmailConfirmed.value = false

    if (emailValid) {
      disableSendButton.value = false
      return
    }

    disableSendButton.value = true

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

  function checkFormPhone() {
    const phoneValid = checkPhone()

    isPhoneConfirmed.value = false

    if (phoneValid && !isPhoneConfirmed.value) {
      disableSendPhoneButton.value = false
      return
    }

    disableSendPhoneButton.value = true
  }

  function checkPhone() {
    // Simple phone validation - can be improved based on requirements
    if (phone.value && phone.value.length >= 10) {
      wrongPhoneType.value = "is-success"
      wrongPhoneMessage.value = ""
      return true
    }
    wrongPhoneType.value = "is-danger"
    wrongPhoneMessage.value = t('page.registration.wrongphone') || "Invalid phone number"
    return false
  }

  async function confirmEmail() {
    try {
      const result = await sendEmailConfirmation(email.value)
      if (!result.error) {
        confirmationType.value = "email"
        confirmationCode.value = ""
        confirmationError.value = ""
        isModalActive.value = true
      } else {
        console.error("Error sending email confirmation:", result.error)
      }
    } catch (error) {
      console.error("Error sending email confirmation:", error)
    }
  }

  async function confirmPhone() {
    try {
      const result = await sendPhoneConfirmation(phone.value)
      if (!result.error) {
        confirmationType.value = "phone"
        confirmationCode.value = ""
        confirmationError.value = ""
        isModalActive.value = true
      } else {
        console.error("Error sending phone confirmation:", result.error)
      }
    } catch (error) {
      console.error("Error sending phone confirmation:", error)
    }
  }

  async function submitConfirmationCode() {
    isConfirmationLoading.value = true
    confirmationError.value = ""

    try {
      let result

      if (confirmationType.value === "email") {
        result = await confirmEmailApi(email.value, confirmationCode.value)
        if (!result.error) {
          isEmailConfirmed.value = true
          disableSendButton.value = true
        }
      } else if (confirmationType.value === "phone") {
        result = await confirmPhoneApi(phone.value, confirmationCode.value)
        if (!result.error) {
          isPhoneConfirmed.value = true
          disableSendPhoneButton.value = true
        }
      }

      if (!result.error) {
        isModalActive.value = false
      } else {
        confirmationError.value = "Invalid confirmation code"
        console.error("Error confirming:", result.error)
      }
    } catch (error) {
      confirmationError.value = "An error occurred"
      console.error("Error confirming:", error)
    } finally {
      isConfirmationLoading.value = false
    }
  }

</script>
    <template>
  <section class="hero is-link">
    <div class="hero-body">
      <p class="title">{{ t("pageheaders.profile") }}</p>
    </div>
  </section>

      <div class="box">
        <div class="columns is-mobile">
          <div class="column is-half is-offset-one-quarter">
            <div class="bd-notification is-primary">
              <b-field label="Email"
                       v-on:keyup="checkFormEmail"
                       :type="wrongEmailType"
                       :message="wrongEmailMessage">
                <b-input v-model="email"></b-input>
              </b-field>
              <b-button
                  v-on:click="confirmEmail"
                  :disabled="disableSendButton"
                  type="is-primary"
                  icon-left="check"
              >
                {{ isEmailConfirmed ? t("page.profile.confirmed") : t("page.profile.confirm") }}
              </b-button>
              <hr/>
              <b-field label="Phone"
                       v-on:keyup="checkFormPhone"
                       :type="wrongPhoneType"
                       :message="wrongPhoneMessage">
                <b-input v-model="phone"></b-input>
              </b-field>
              <b-button
                  v-on:click="confirmPhone"
                  :disabled="disableSendPhoneButton"
                  type="is-primary"
                  icon-left="check"
              >
                {{ isPhoneConfirmed ? t("page.profile.confirmed") : t("page.profile.confirm") }}
              </b-button>
            </div>
          </div>
        </div>
      </div>

      <!-- Confirmation Code Modal -->
      <b-modal v-model="isModalActive" :title="confirmationType === 'email' ? t('page.profile.confirmEmail') : t('page.profile.confirmPhone')" has-modal-card>
        <div class="modal-card">
          <header class="modal-card-head">
            <p class="modal-card-title">{{ confirmationType === 'email' ? t('page.profile.confirmEmail') : t('page.profile.confirmPhone') }}</p>
          </header>
          <section class="modal-card-body">
            <b-field :label="t('page.profile.confirmationCode')" :type="confirmationError ? 'is-danger' : ''" :message="confirmationError">
              <b-input v-model="confirmationCode" type="text" placeholder="123456"></b-input>
            </b-field>
          </section>
          <footer class="modal-card-foot">
            <button class="button is-primary" @click="submitConfirmationCode" :disabled="isConfirmationLoading || !confirmationCode">
              <b-loading :is-full-page="false" :active="isConfirmationLoading"></b-loading>
              {{ t('page.profile.confirm') }}
            </button>
            <button class="button" @click="isModalActive = false">{{ t('page.profile.cancel') }}</button>
          </footer>
        </div>
      </b-modal>
    </template>

<style>

</style>
