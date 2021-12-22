<template>
  <Header/>
  <el-container>
    <el-aside class="aside">
      <div v-for="option in message" :key="option">
        <el-button type="primary" class="head" @click="load(option.index)" :icon="option.icon">
          {{ option.name }}
        </el-button>
        <hr/>
      </div>
    </el-aside>

    <el-container style="height: 800px">
      <el-header class="title">
        www.我和僵尸有个约会.com
      </el-header>

      <el-main>
        <el-row>
          <div v-for="video in videos.list" :key="video">
            <div class="row">
              <el-alert :title="video.title " type="success" center show-icon :closable="false"></el-alert>
              <vue3-video-play width="400px" height="225px" :src="video.url" type="m3u8" :control="false"
                               @play="play(video.id)">
              </vue3-video-play>
            </div>
          </div>
        </el-row>
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup>
import Header from "./Header.vue"
import {onMounted, reactive} from "vue";
import {useRouter} from 'vue-router'
import axios from "axios";

const router = useRouter()

const message = reactive([
  {"name": "僵约一", "icon": "star-filled", "index": 1},
  {"name": "僵约二", "icon": "ship", "index": 2},
  {"name": "僵约三", "icon": "apple", "index": 3},
  {"name": "精彩短片", "icon": "video-camera", "index": 4},
  {"name": "剪辑短片", "icon": "video-camera-filled", "index": 5},
  {"name": "网友提供", "icon": "data-line", "index": 6},
  {"name": "开幕片段", "icon": "platform", "index": 7},
  {"name": "各种音乐", "icon": "service", "index": 8},
  {"name": "精彩图片", "icon": "picture", "index": 9},
  {"name": "交流论坛", "icon": "share", "index": 10},
  {"name": "个人信息", "icon": "user", "index": 11},
])

const videos = reactive({
  "list": []
})

onMounted(() => {
  load(3)
})

function load(type) {
  const params = {
    "type": type,
    "page": 0,
    "size": 10
  }
  axios.post("http://localhost:12480/api/videoList", params).then((resp) => {
    if (resp.status === 200 && resp.data.code === 200) {
      videos.list = resp.data.data
    }
  })
}


function play(id) {
  router.push({path: "/play", query: {"id": id}})
}
</script>


<style scoped>

.aside {
  width: 180px;
  height: 800px;
  border-right: rgba(115, 115, 101, 0.19) solid 5px;
}

.head {
  width: 150px;
  height: 60px;
  margin: 2px auto 0;
}

.title {
  height: 60px;
  background-color: #ee8b8b;
  font-size: 30px;
  text-align: center;
}

.row {
  width: 400px;
  height: 225px;
  margin: 0 50px 50px 50px;
  cursor: pointer;
}

</style>