<template>
  <div>

    <video autoplay loop class="video-bg" muted>
      <source src="https://jtyoui.oss-accelerate-overseas.aliyuncs.com/video/login.mp4" type="video/mp4">
    </video>

    <div class="layout">

      <div class="tip">
        <el-tooltip content="如果没有唯一的僵尸码请先注册" placement="top">
          <el-button size="medium" type="text" class="text">😀 欢 迎 登 陆</el-button>
        </el-tooltip>
      </div>

      <div class="token">
        <el-input clearable maxlength="16" v-model="token" placeholder="请输入唯一的僵尸码" show-password :prefix-icon="Lock"/>
      </div>

      <div class="button">
        <el-button type="primary" class="len" :icon="Avatar" :loading="loading" auto-insert-space>
          点 击 登 陆
        </el-button>
      </div>

      <div class="register">
        <el-link type="info" class="text" href="/register">👉 注册新账号→</el-link>
      </div>

    </div>
  </div>

</template>

<script setup>
import {ref} from "vue";
import {Avatar, Lock} from '@element-plus/icons'

const token = ref("")
const loading = ref(false)
</script>

<style scoped>
.layout {
  width: 25%;
  height: 500px;
  margin-left: 70%;
  margin-top: 10%;
  background-color: rgba(175, 207, 239, 0.86);
}

.tip {
  text-align: center;
  padding-top: 10%;
}

.tip .text {
  font-size: 26px;
  color: black;
}

.token {
  padding-top: 10%;
  width: 80%;
  margin: auto;
}

.button {
  min-width: 200px;
  padding-top: 10%;
  text-align: center;
}

.button .len {
  min-width: 50%;
}

.register {
  padding-top: 20%;
  padding-left: 10%;
}

.register .text {
  font-size: 20px;
  color: black;
}

</style>