<script setup>
  import {useI18n} from 'vue-i18n'
  import {useLoginStore} from './../stores/login.js'
  import {ref, onMounted} from "vue";
  import router from "@/router/index.js";
  import {getCurrentUser} from "./../client/client"
  import {BButton} from "buefy";

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

  function confirmEmail() {
    console.log("Confirming with email:", email.value)
  }

  function confirmPhone() {
    console.log("Confirming with and phone:", phone.value)
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
</template>

<style>

</style>
