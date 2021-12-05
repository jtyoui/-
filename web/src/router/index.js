import {createRouter, createWebHistory} from 'vue-router'
import Login from "../components/Login.vue";
import Register from "../components/Register.vue";

const routes = [{
    path: '/', component: Login
}, {
    path: '/register', component: Register
}]

const Router = createRouter({
    history: createWebHistory(), routes
})
export default Router
