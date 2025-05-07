import {ref, computed} from 'vue'
import {defineStore} from 'pinia'

export const useLoginStore = defineStore(
    'login',
    () => {
        const jwt = ref("")
        const isLoggedIn = ref(false)

        function login(value) {
            jwt.value = value
            isLoggedIn.value = true
        }

        function logout() {
            jwt.value = value
            isLoggedIn.value = false
        }


        return {
            jwt,
            isLoggedIn,
            login,
            logout,
        }
    }
)
