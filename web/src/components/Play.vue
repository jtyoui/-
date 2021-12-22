<template>
  <Header/>
  <el-container>
    <el-main>
      <div class="video">
        <vue3-video-play v-bind="options.video"/>
        <div class="top">
          <el-icon :size="20">
            <share/>
          </el-icon>
          <span class="tittle">&emsp;{{ options.video.title }}</span>
          <el-button icon="star" class="start" type="success">20</el-button>
        </div>
      </div>
    </el-main>

    <el-aside width="30%">
      <div class="box">
        <el-menu :default-openeds="[0]">
          <el-sub-menu :index="0">
            <template #title>
              <el-icon>
                <Menu/>
              </el-icon>
              <span>{{ options.menu }}</span>
            </template>
            <div class="list" style="overflow: auto">
              <div v-for="(video,i) in options.list" :key="i" class="infinite-list-item">
                <el-menu-item :index="i" @click="play(video)">
                  <el-tooltip :content="video.title" placement="top" show-after="800">
                    <div class="name">
                      P{{ i + 1 }}&emsp;{{ video.title }}
                    </div>
                  </el-tooltip>
                  <div class="time">{{ video.time }}</div>
                </el-menu-item>
              </div>
            </div>
          </el-sub-menu>
        </el-menu>
      </div>
    </el-aside>
  </el-container>
</template>

<script setup>
import Header from "./Header.vue"
import {onMounted, reactive} from "vue";
import {Menu} from '@element-plus/icons'
import axios from "axios";

const options = reactive({
  video: {
    width: "984px",
    height: "700px",
    title: "",
    src: "",
    type: "m3u8",
    muted: false,
    webFullScreen: false,
    autoPlay: false,
    loop: false,
    control: true,
    volume: 0.3
  },
  list: [],
  menu: ""
})

onMounted(() => {
  axios.get("http://localhost:12480/api/play?video_id=" + "1").then((resp) => {
    if (resp.status === 200 && resp.data.code === 200) {
      const data = resp.data.data
      options.video.src = data.url
      options.video.title = data.title
      options.menu = data.name
      options.list = data.list
    }
  })
  console.log(options)
})

function play(self) {
  console.log(self)
  options.video.title = self.title
  options.video.src = self.url
}

</script>

<style scoped>
.video {
  margin: 0 15%;
}

.box {
  width: 300px;
  border: #dcd9d9 solid 2px;
  margin: 0 auto;
}

.tittle {
  font-size: 20px;
  color: red;
}

.start {
  float: right;
}

.top {
  margin-top: 20px;
}

.time {
  width: 20px;
  margin-left: 30px;
}

.list {
  height: 700px;
  list-style: none;
  width: 100%;
}

.name {
  width: 200px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>