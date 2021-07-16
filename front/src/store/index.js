import Vue from 'vue'
import Vuex from 'vuex'
import PostsModules from './modules/posts-module'
import NotifModules from './modules/notif-module'
import UserModules from './modules/user-module'
import SignupModules  from './modules/sign-up-module'
import AuthModules from './modules/auth-module'

Vue.use(Vuex)

export default new Vuex.Store({
    modules: {
        PostsModules,
        NotifModules,
        UserModules,
        SignupModules,
        AuthModules
    }
})