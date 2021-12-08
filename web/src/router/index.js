import {createRouter, createWebHistory} from 'vue-router'
import Login from "../components/Login.vue"
import Register from "../components/Register.vue"
import Index from "../components/Index.vue"
import Video from "../components/Video.vue"

const routes = [{
    path: '/login', component: Login
}, {
    path: '/register', component: Register
}, {
    path: '/video', component: Video
}, {
    path: '/', component: Index
}]

const Router = createRouter({
    history: createWebHistory(), routes
})
export default Router
