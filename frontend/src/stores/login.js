import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useLoginStore = defineStore(
    'login',
    () => {
    const jwt = ref("")

    function login(value) {
      jwt.value = value
    }
    function logout() {
      jwt.value = value
    }
    function isLoggedIn() {
      return jwt.value != ""
    }

    return {
      jwt,
      login,
      logout,
      isLoggedIn,
    }
  }
)
