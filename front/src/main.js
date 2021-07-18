import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import store from './store'
import router from './router'
import axios from "axios"

Vue.config.productionTip = false
Vue.prototype.$axios = axios
const token = localStorage.getItem('userToken');
if (token) {
    Vue.prototype.$axios.defaults.headers.common['Authorization'] = 'Bearer ' + token
}

new Vue({
    vuetify,
    store,
    router,
    render: h => h(App)
}).$mount('#app')
