<template>

  <van-nav-bar title="ChatGPT" />
  <van-grid :column-num="3">
    <van-grid-item icon="smile-o" text="起名" to = "/name" />
    <van-grid-item icon="fire-o" text="创意写作" to = "/creativity" badge="hot" />
    <van-grid-item v-for="value in 15" :key="value" icon="photo-o" text="逐步添加中..." />
    <van-grid-item icon="shopping-cart-o" text="goods" to = "/goods" />
  </van-grid>
</template>
  
  <script>
  import { ref } from 'vue';
  import { Grid, GridItem } from 'vant';

export default {
  setup() {
    const active = ref('home');
    const value = ref('');
    const onSearch = (val) => showToast(val);
    const onClickButton = () => showToast(value.value);

    const skip = ref(1);
    return { active ,value,onSearch,onClickButton,skip};
  },
};
  </script>
  
  <style lang="less">
  
  </style>
  