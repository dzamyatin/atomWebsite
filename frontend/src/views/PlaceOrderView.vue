<script setup>
  import {useI18n} from 'vue-i18n'
  import {useLoginStore} from './../stores/login.js'
  import {ref, reactive} from "vue";
  import router from "@/router/index.js";
  import {register} from "./../client/client"
  import {BButton, BTable, BInput, BRadio} from "buefy";

  const {t} = useI18n()
  const store = useLoginStore()

  if (!store.isLoggedIn) {
    router.push('/login')
  }

  // Sample products data - in a real app, this would come from an API
  const products = reactive([
    { id: 1, name: '10 Mbit/sec. Месяц. Германия.', price: 1000, quantity: 1 },
    { id: 2, name: '50 Mbit/sec. Месяц. Франция.', price: 1500, quantity: 1 },
    { id: 3, name: '100 Mbit/sec. Месяц. США.', price: 2000, quantity: 1 }
  ]);

  // Available payment methods
  const paymentMethods = ref([
    { id: 1, name: 'Credit Card' },
    { id: 2, name: 'PayPal' },
    { id: 3, name: 'Bank Transfer' }
  ]);

  const selectedPaymentMethod = ref(1);

  // Function to validate quantity input (only positive integers)
  const validateQuantity = (value) => {
    const parsedValue = parseInt(value);
    return !isNaN(parsedValue) && parsedValue > 0;
  };

  // Function to update quantity
  const updateQuantity = (product, event) => {
    const value = event.target.value;
    if (validateQuantity(value)) {
      product.quantity = parseInt(value);
    } else {
      // Reset to 1 if invalid
      product.quantity = 1;
    }
  };

  // Calculate total price
  const calculateTotal = () => {
    return products.reduce((total, product) => {
      return total + (product.price * product.quantity);
    }, 0);
  };

  // Handle order confirmation
  const confirmOrder = () => {
    // In a real app, this would submit the order to an API
    console.log('Order confirmed', {
      products,
      paymentMethod: paymentMethods.value.find(method => method.id === selectedPaymentMethod.value)
    });
    // Navigate to payment page
    router.push('/pay');
  };
</script>

<template>
  <section class="hero is-link">
    <div class="hero-body">
      <p class="title">{{ t("pageheaders.placeorder") }}</p>
    </div>
  </section>

  <div class="box">
    <div class="columns">
      <div class="column is-8 is-offset-2">
        <!-- Products Table -->
        <h3 class="title is-4">{{ t("page.placeorder.productName") }}</h3>
        <table class="table is-fullwidth">
          <thead>
            <tr>
              <th>{{ t("page.placeorder.productName") }}</th>
              <th>{{ t("page.placeorder.quantity") }}</th>
              <th>{{ t("page.placeorder.price") }}</th>
              <th>{{ t("page.placeorder.total") }}</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="product in products" :key="product.id">
              <td>{{ product.name }}</td>
              <td>
                <input 
                  class="input" 
                  type="number" 
                  min="1" 
                  :value="product.quantity" 
                  @input="updateQuantity(product, $event)"
                />
              </td>
              <td>{{ product.price }} ₽</td>
              <td>{{ product.price * product.quantity }} ₽</td>
            </tr>
          </tbody>
          <tfoot>
            <tr>
              <td colspan="3" class="has-text-right"><strong>{{ t("page.placeorder.total") }}:</strong></td>
              <td><strong>{{ calculateTotal() }} ₽</strong></td>
            </tr>
          </tfoot>
        </table>

        <!-- Payment Methods -->
        <h3 class="title is-4 mt-5">{{ t("page.placeorder.paymentMethods") }}</h3>
        <div class="field">
          <div v-for="method in paymentMethods" :key="method.id" class="control">
            <label class="radio">
              <input 
                type="radio" 
                :value="method.id" 
                v-model="selectedPaymentMethod"
              >
              {{ method.name }}
            </label>
          </div>
        </div>

        <!-- Confirm Button -->
        <div class="field mt-5">
          <div class="control">
            <button class="button is-primary" @click="confirmOrder">
              {{ t("page.placeorder.confirmOrder") }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
.mt-5 {
  margin-top: 1.5rem;
}
</style>
