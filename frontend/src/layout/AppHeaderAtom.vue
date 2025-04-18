<template>
    <header class="header-global">
        <base-nav class="navbar-main" transparent type="" effect="light" expand>
            <router-link slot="brand" class="navbar-brand mr-lg-5" to="/">
                <img src="/img/atomvpn2.png" alt="logo">
                ATOM VPN
            </router-link>

            <div class="row" slot="content-header" slot-scope="{closeMenu}">
                <div class="col-6 collapse-brand">
                    <a href="https://demos.creative-tim.com/vue-argon-design-system/documentation/">
                        <img src="/img/brand/blue.png">
                    </a>
                </div>
                <div class="col-6 collapse-close">
                    <close-button @click="closeMenu"></close-button>
                </div>
            </div>

            <ul class="navbar-nav navbar-nav-hover align-items-lg-center">
            </ul>
            <ul class="navbar-nav align-items-lg-center ml-lg-auto">
                <li class="nav-item">
                  <router-link slot="brand" class="" to="/login">
                    <base-button class="btn-3" type="primary" icon="ni ni-atom">Вход</base-button>
                  </router-link>
                </li>
                <li class="nav-item">
                  <router-link to="/download">
                    <a href="#" target="_blank" rel="noopener" class="btn btn-neutral btn-icon">
                        <span class="btn-inner--icon">
                          <i class="fa fa-cloud-download mr-2"></i>
                        </span>
                        <span class="nav-link-inner--text">Скачать</span>
                    </a>
                  </router-link>
                </li>
            </ul>
        </base-nav>
    </header>
</template>
<script>
import BaseNav from "@/components/BaseNav.vue";
import BaseDropdown from "@/components/BaseDropdown.vue";
import CloseButton from "@/components/CloseButton.vue";

export default {
  components: {
    BaseNav,
    CloseButton,
    BaseDropdown
  }
};
</script>
<style>
</style>
