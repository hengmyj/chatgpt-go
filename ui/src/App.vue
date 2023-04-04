<template>
  <div id="app">
    <router-view />
    <van-tabbar route>
      <van-tabbar-item replace to="/home" icon="home-o">主页</van-tabbar-item>
      <van-tabbar-item replace to="/name" icon="hot-o">起名</van-tabbar-item>
      <van-tabbar-item replace to="/chattgpt" icon="chat-o">chatgpt</van-tabbar-item>
      <van-tabbar-item replace to="/user" icon="apps-o">个人中心</van-tabbar-item>
    </van-tabbar>
  </div>
</template>

<style>
body {
  font-size: 16px;
  background-color: #f8f8f8;
  -webkit-font-smoothing: antialiased;
}
</style>
