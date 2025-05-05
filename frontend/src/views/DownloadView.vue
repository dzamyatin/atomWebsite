<script setup>
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
</script>

<template>
  <main>
    <section class="hero is-link">
      <div class="hero-body">
        <p class="title">{{t("pageheaders.download")}}</p>
      </div>
    </section>
    <div class="box">
      <div class="columns is-3">
        <div class="column">
          <b-button class="is-fullwidth" size="is-large"
                    icon-left="microsoft-windows">
            {{t("download.download")}} Windows
          </b-button>
        </div>
        <div class="column">
          <b-button class="is-fullwidth" size="is-large"
                    icon-left="penguin">
            {{t("download.download")}} Linux
          </b-button>
        </div>
        <div class="column">
          <b-button class="is-fullwidth" size="is-large"
                    icon-left="android">
            {{t("download.download")}} Android
          </b-button>
        </div>
      </div>
    </div>
  </main>
</template>
