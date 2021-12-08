import {createApp} from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import vue3videoPlay from 'vue3-video-play'
import 'vue3-video-play/dist/style.css'
import 'element-plus/dist/index.css'
import './assets/vampire.css'
import Router from './router'
import * as ElIconModules from '@element-plus/icons'


const app = createApp(App)
for (let iconName in ElIconModules) {
    app.component(iconName, ElIconModules[iconName])
}

app.use(Router)
app.use(ElementPlus)
app.use(vue3videoPlay)
app.mount('#app')
