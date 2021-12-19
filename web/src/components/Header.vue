<template>
  <div class="head">
    <div class="header">

      <div class="icon">
        <el-icon :size="35">
          <House/>
        </el-icon>
      </div>

      <div class="search">
        <el-input clearable v-model="search" placeholder="请输入要搜索的视频" :prefix-icon="Search"/>
      </div>

      <div class="name">
        <span class="text" @click="index">首页</span>
        <span class="text">联系我<del class="small">们</del></span>
        <span class="text">关于网站</span>
      </div>
    </div>
  </div>
  <hr/>
</template>

<script setup>
import {House, Search} from '@element-plus/icons'
import {onMounted, ref} from "vue";
import {useRouter} from "vue-router";

const search = ref("")
const route = useRouter()

onMounted(() => {
  const vampire = localStorage.getItem("vampire")
  if (!(vampire && vampire.trim().length)) {
    route.push('/login')
  }
})

function index() {
  route.push('/')
}
</script>

<style scoped>
.head {
  height: 80px;
  background-color: rgba(255, 255, 255);
}

.header {
  padding-top: 20px;
}

.icon {
  width: 10%;
  float: left;
  padding-left: 5%;
}

.search {
  width: 50%;
  float: left;
}

.name {
  width: 30%;
  font-size: 20px;
  float: right;
}

.text {
  padding-left: 15%;
  cursor: pointer;
}

.small {
  font-size: 15px;
}
</style>