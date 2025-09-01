<script setup>
  import {useI18n} from 'vue-i18n'
  import {useLoginStore} from './../stores/login.js'
  import {ref} from "vue";
  import router from "@/router/index.js";
  import {register} from "./../client/client"
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

  function checkFormEmail() {
    const emailValid = checkEmail()

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

              >{{ t("page.licenses.confirm") }}
              </b-button>
              <hr/>
              <b-field label="Phone"
                       v-on:keyup="checkForm"
                       :type="wrongPhoneType"
                       :message="wrongPhoneMessage">
                <b-input v-model="phone"></b-input>
              </b-field>
              <b-button
                  v-on:click="confirmPhone"
                  :disabled="disableSendPhoneButton"
                  type="is-primary"
                  icon-left=""

              >{{ t("page.licenses.confirm") }}
              </b-button>
            </div>
          </div>
        </div>
      </div>
</template>

<style>

</style>
