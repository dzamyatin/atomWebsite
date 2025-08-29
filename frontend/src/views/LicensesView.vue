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
      <p class="title">{{ t("pageheaders.licenses") }}</p>
    </div>
  </section>

      <div class="box">
        <div class="columns is-mobile">
          <div class="column is-half is-offset-one-quarter">
            <div class="bd-notification is-primary">
              <table class="table">
                <thead>
                <tr>
                  <th>Идентификатор</th>
                  <th>Название</th>
                  <th>Валиден до</th>
                  <th>Дата создания</th>
                  <th>Ключ</th>
                </tr>
                </thead>
                <tbody>
                <tr>
                  <th>1</th>
                  <td>
                    10 Mbit/sec. Месяц. Германия.
                  </td>
                  <td>2025.01.01</td>
                  <td>2026.01.01</td>
                  <td>
                    <textarea>asdasfmioqwmdoXASmdfo,asdasfg</textarea>
                  </td>
                  <td>
                    <b-button>Копировать</b-button>
                  </td>
                </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>


</template>

<style>

</style>
