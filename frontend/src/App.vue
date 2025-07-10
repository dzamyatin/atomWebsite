<script setup>
import { RouterLink, RouterView } from 'vue-router'
import { BNavbar, BNavbarDropdown, BNavbarItem } from 'buefy'
import { useI18n } from 'vue-i18n'
import NavProfile from "@/components/NavProfile.vue";
import SocialNetworks from "./components/SocialNetworks.vue";

const { t } = useI18n()
</script>

<template>
  <div>
    <header>
      <b-navbar>
        <template #brand>
          <b-navbar-item tag="router-link" :to="{ path: '/' }">
            <img
                src="/atomvpn.png"
                :alt="t('header.logoalt')"
            >
          </b-navbar-item>
        </template>
        <template #start>
            <b-navbar-item href="/download">
              {{ t("nav.download") }}
            </b-navbar-item>
        </template>

        <template #end>
          <b-navbar-item tag="div">
            <NavProfile></NavProfile>
          </b-navbar-item>
        </template>
      </b-navbar>
    </header>
  </div>
  <div>
    <RouterView />
  </div>
  <footer class="footer">
    <div class="content has-text-centered">
      <p>
        {{t("footer.text")}}
        <strong>Android, Linux, Windows</strong>
        <br/>
        <a href="/download"> {{t("footer.download")}}</a>
      </p>
    </div>
    <div class="content has-text-centered">
      <SocialNetworks/>
    </div>
  </footer>
</template>

<style scoped>

</style>
