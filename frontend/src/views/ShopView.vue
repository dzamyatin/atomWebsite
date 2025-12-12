<script setup>
  import {useI18n} from 'vue-i18n'
  import {useLoginStore} from './../stores/login.js'
  import {ref, reactive} from "vue";
  import router from "@/router/index.js";
  import {BButton, BIcon} from "buefy";

  const {t} = useI18n()
  const store = useLoginStore()

  if (!store.isLoggedIn) {
    router.push('/login')
  }

  // Sample products data - in a real app, this would come from an API
  const products = reactive([
    { id: 1, name: '10 Mbit/sec. Месяц. Германия.', price: 1000, icon: 'server' },
    { id: 2, name: '50 Mbit/sec. Месяц. Франция.', price: 1500, icon: 'server-network' },
    { id: 3, name: '100 Mbit/sec. Месяц. США.', price: 2000, icon: 'server-plus' }
  ]);

  // Handle buy button click
  const buyProduct = (product) => {
    // In a real app, this would add the product to a cart or navigate to checkout
    console.log('Buy product', product);
    router.push('/place-order');
  };
</script>

<template>
  <section class="hero is-link">
    <div class="hero-body">
      <p class="title">{{ t("pageheaders.shop") }}</p>
    </div>
  </section>

  <div class="box">
    <div class="columns is-multiline">
      <!-- Product Tiles -->
      <div v-for="product in products" :key="product.id" class="column is-4">
        <div class="card">
          <div class="card-content">
            <div class="has-text-centered mb-4">
              <b-icon :icon="product.icon" size="is-large" type="is-primary"></b-icon>
            </div>
            <p class="title is-5">{{ product.name }}</p>
            <p class="subtitle is-6 has-text-weight-bold">{{ product.price }} ₽</p>
            <div class="has-text-centered mt-4">
              <button class="button is-primary" @click="buyProduct(product)">
                {{ t("page.licenses.buttonToBuy") }}
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
.mb-4 {
  margin-bottom: 1rem;
}
.mt-4 {
  margin-top: 1rem;
}
</style>