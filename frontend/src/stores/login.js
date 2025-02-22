import {ref, computed} from 'vue'
import {defineStore} from 'pinia'

export const useLoginStore = defineStore(
    'login',
    () => {
        // Initialize from session storage if available
        const storedJwt = sessionStorage.getItem('jwt') || ""
        const jwt = ref(storedJwt)
        const isLoggedIn = ref(storedJwt !== "")

        function login(value) {
            jwt.value = value
            isLoggedIn.value = true
            // Store in session storage
            sessionStorage.setItem('jwt', value)
        }

        function logout() {
            jwt.value = ""
            isLoggedIn.value = false
            // Remove from session storage
            sessionStorage.removeItem('jwt')
        }

        return {
            jwt,
            isLoggedIn,
            login,
            logout,
        }
    }
)
