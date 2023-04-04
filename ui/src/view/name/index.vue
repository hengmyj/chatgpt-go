<template>
  <van-nav-bar title="名字查询" />
  <van-form @submit="onSubmit">
  <van-cell-group inset>
    <van-field
      v-model="firsname"
      name="firsname"
      label="姓"
      placeholder="姓"
      :rules="[{ required: true, message: '请填写姓' }]"
    />
    <van-field name="sex" label="性别">
      <template #input>
        <van-radio-group v-model="sex" direction="horizontal">
          <van-radio name="男">弟弟</van-radio>
          <van-radio name="女">妹妹</van-radio>
        </van-radio-group>
      </template>
    </van-field>
     
  <van-field
    v-model="value"
    is-link
    readonly
    name="calendar"
    label="日历"
    placeholder="点击选择日期"
    @click="showCalendar = true"
  />
<van-calendar v-model:show="showCalendar" @confirm="onConfirm" />

  </van-cell-group>

  <div style="margin: 16px;">
    <van-button round block type="primary" native-type="submit">
      提交
    </van-button>
  </div>



</van-form>
<van-cell-group inset>

<van-list
  v-model:loading="loading"
  :finished="finished"
  finished-text="没有更多了"
  @load="onLoad"
>
  <van-cell v-for="item in list" :key="item" :title="item" />
</van-list>
</van-cell-group>


<van-notify v-model:show="show">
</van-notify>

</template>
  
  <script>
  import { ref } from 'vue';
  import { showNotify,Notify,showLoadingToast, closeToast } from 'vant';
  import axios from 'axios'

export default {
  setup() {
    const firsname = ref('李');
    const sex = ref('男');
    const result = ref('');
    const showCalendar = ref(false);
    const list = ref([]);
    const show = ref(false);
    const loading = ref(false);
    const finished = ref(false);


    const onSubmit = (val) => {
      console.log('submit', val);
      list.value.length = 0;
      // 加载状态
      loading.value = false;

      const toast = showLoadingToast({
        duration: 0,
        forbidClick: true,
        message: '倒计时 120 秒',
      });

      let second = 120;
      const timer = setInterval(() => {
        second--;
        if (second) {
          toast.message = `倒计时 ${second} 秒`;
        } else {
          clearInterval(timer);
          closeToast();
        }
      }, 1000);

      axios
      .get('https://name-api.wisdom-os.top/v1/name/create?firstname='+val["firsname"]+'&sex='+val["sex"])
      .then((response) => {
        if(response.data.data){
            var str = response.data.data.split("\n")
            str.forEach(e => {
              list.value.push(e);
            });
        }
        console.log(response)
        // 主动关闭
        clearInterval(timer);
        closeToast();
      })
     
    };

    const onConfirm = (date) => {
      result.value = `${date.getMonth() + 1}/${date.getDate()}`;
      showCalendar.value = false;
    };

   
    const onLoad = () => {
      // 异步更新数据
      // setTimeout 仅做示例，真实场景中一般为 ajax 请求
      setTimeout(() => {
        for (let i = 0; i < 10; i++) {
          list.value.push(list.value.length + 1);
        }

        // 加载状态结束
        loading.value = false;

        // 数据全部加载完成
        if (list.value.length >= 40) {
          finished.value = true;
        }
      }, 1000);
    };
    




    return {
      show,
      firsname,sex,
      onSubmit,
      onConfirm,result,showCalendar,
      list,
      // onLoad,
      loading,
      finished,
    };
  },
};
  </script>
  
  <style lang="less">
 .wrapper {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100%;
  }

  .block {
    width: 120px;
    height: 120px;
    background-color: #fff;
  }
  </style>
  