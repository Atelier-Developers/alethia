import Vue from 'vue'
import VueRouter from 'vue-router'

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


export default router