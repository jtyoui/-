<template>
  <Header/>
  <el-container>
    <el-aside class="aside">
      <div v-for="option in message" :key="option">
        <div class="head">
          <el-row style="height: 55px">
            <el-icon :size="20" class="icon">
              <component :is="option.icon"/>
            </el-icon>
            &emsp;
            <h3><a :href="option.link">{{ option.name }}</a></h3>
          </el-row>
          <hr/>
        </div>
      </div>
    </el-aside>
    <el-container>
      <el-header class="title">
        www.我和僵尸有个约会.com
      </el-header>

      <el-main>

        <el-row>
          <div class="row">
            <vue3-video-play v-bind="options"/>
            <span># 我和僵尸有个约会1——马小玲爱上我</span>
          </div>
          <div class="row">
            <vue3-video-play v-bind="options"/>
            <span># 我和僵尸有个约会1——马小玲爱上我</span>
          </div>
          <div class="row">
            <vue3-video-play v-bind="options"/>
            <span># 我和僵尸有个约会1——马小玲爱上我</span>
          </div>
        </el-row>

        <el-row class="top">
          <div class="row">
            <vue3-video-play v-bind="options"/>
            <span># 我和僵尸有个约会1——马小玲爱上我</span>
          </div>
          <div class="row">
            <vue3-video-play v-bind="options"/>
            <span># 我和僵尸有个约会1——马小玲爱上我</span>
          </div>
          <div class="row">
            <vue3-video-play v-bind="options"/>
            <span># 我和僵尸有个约会1——马小玲爱上我</span>
          </div>
        </el-row>


        <el-row class="top">
          <div class="row">
            <vue3-video-play v-bind="options"/>
            <span># 我和僵尸有个约会1——马小玲爱上我</span>
          </div>
          <div class="row">
            <vue3-video-play v-bind="options"/>
            <span># 我和僵尸有个约会1——马小玲爱上我</span>
          </div>
          <div class="row">
            <vue3-video-play v-bind="options"/>
            <span># 我和僵尸有个约会1——马小玲爱上我</span>
          </div>
        </el-row>


      </el-main>
    </el-container>
  </el-container>
</template>

<script setup>
import Header from "./Header.vue"
import {reactive} from "vue";


const message = reactive([
  {"name": "僵约一", "icon": "star-filled", "link": ""},
  {"name": "僵约二", "icon": "ship", "link": ""},
  {"name": "僵约三", "icon": "apple", "link": ""},
  {"name": "精彩短片", "icon": "video-camera", "link": ""},
  {"name": "剪辑短片", "icon": "video-camera-filled", "link": ""},
  {"name": "网友提供", "icon": "data-line", "link": ""},
  {"name": "开幕片段", "icon": "platform", "link": ""},
  {"name": "各种音乐", "icon": "service", "link": ""},
  {"name": "精彩图片", "icon": "picture", "link": ""},
  {"name": "交流论坛", "icon": "share", "link": ""},
  {"name": "个人信息", "icon": "user", "link": ""},
])

const options = reactive({
  width: "369px",
  height: "207px",
  title: "我和僵尸有个约会开头",
  src: "https://jtyoui.oss-accelerate-overseas.aliyuncs.com/video/login.mp4",
  muted: true,
  webFullScreen: false,
  autoPlay: true,
  loop: true,
  control: true,
  volume: 0.3,
  controlBtns: [
    "fullScreen",
    "audioTrack",
    "quality",
    "speedRate",
    "volume",
  ]
})

</script>


<style scoped>

.aside {
  width: 200px;
  border-right: rgba(115, 115, 101, 0.19) solid 5px;
}

.head {
  width: 150px;
  margin: 0 auto;
}

.icon {
  float: left;
  margin-top: 25px;
}

a {
  text-decoration: none;
}

.title {
  height: 60px;
  background-color: #ee8b8b;
  font-size: 30px;
  text-align: center;
}

.row {
  margin: 0 auto;
}

.top {
  margin-top: 20px;
}
</style>