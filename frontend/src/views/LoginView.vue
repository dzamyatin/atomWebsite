<script setup>
import { useI18n } from 'vue-i18n'
import {BNavbarItem} from "buefy";

const { t } = useI18n()
</script>
<template>
  <section class="hero is-link">
    <div class="hero-body">
      <p class="title">{{t("pageheaders.login")}}</p>
    </div>
  </section>
</template>

<style>

</style>
