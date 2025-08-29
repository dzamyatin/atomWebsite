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
              tst
            </div>
          </div>
        </div>
      </div>


</template>

<style>

</style>
