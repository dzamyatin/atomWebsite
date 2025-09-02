<script setup>
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
</script>

<template>
  <main>
    <section class="hero is-danger is-medium">
      <div class="hero-body">
        <div class="container has-text-centered">
          <h1 class="title is-1">404</h1>
          <p class="subtitle">{{ t("page.notfound.message") }}</p>
          <b-button 
            size="is-medium"
            icon-left="home"
            tag="router-link" 
            :to="{ path: '/' }"
          >
            {{ t("page.notfound.backhome") }}
          </b-button>
        </div>
      </div>
    </section>
  </main>
</template>