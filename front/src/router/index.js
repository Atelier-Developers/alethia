import Vue from 'vue'
import VueRouter from 'vue-router'
import store from '../store/index'

Vue.use(VueRouter)

const routes = [
    {
        path: "/",
        redirect: {
            name: "Home"
        }
    },
    {
        path: "/home",
        name: "Home",
        component: () => import("@/views/Home"),
    },
    {
        path: "/notifications",
        name: "Notifications",
        component: () => import("@/views/Notifs"),
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import("@/views/Login"),
    },
    {
        path: '/signup',
        name: 'Signup',
        component: () => import("@/views/Signup"),
    },
    {
        path: '/profile',
        name: 'Profile',
        component: () => import("@/views/Profile")
    },
    {
        path: '/user/:id',
        name: 'User',
        component: () => import('@/views/User')
    },
    {
        path: '/post/:id',
        name: 'Post',
        component: () => import('@/views/PostPage')
    },
    {
        path: '/network',
        name: 'Network',
        component: () => import("@/views/Network")
    },
    {
        path: '/conversation',
        name: 'Conversation',
        component: () => import("@/views/Conversation")
    }
    // {
    //     path: '*',
    //     name: 'Error404',
    //     component: () => import("../views/Error404")
    // },
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})


router.beforeEach((to, from, next) => {
    if (to.name === 'Login' || to.name === 'Signup') next()
    else if (!store.getters["AuthModules/isAuthenticated"]) {
        next({name: "Login"});
    } else {
        next()
    }
})


export default router