<script setup>
import { useI18n } from 'vue-i18n'
import {BNavbarItem} from "buefy";

const { t } = useI18n()
</script>

<template>
  <main>
    <section class="hero is-primary is-medium">
      <!-- Hero head: will stick at the top -->
      <div class="hero-head">
      </div>

      <!-- Hero content: will be in the middle -->
      <div class="hero-body">
        <div class="container has-text-centered">
          <p class="title">ATOM VPN</p>
          <p class="subtitle">{{ t("home.subtitle") }}</p>
          <b-button size="is-large"
                    icon-left="download-circle-outline"
                    tag="router-link" :to="{ path: '/download' }"
          >
            {{ t("home.download") }}
          </b-button>
        </div>
      </div>

      <!-- Hero footer: will stick at the bottom -->
      <div class="hero-foot">
      </div>
    </section>
  </main>
</template>
