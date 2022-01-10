<template>
  <div class="home">
    <el-card class="main-card" :body-style="bodyStyle">
      <el-carousel height="200px">
        <el-carousel-item v-for="item in 4" :key="item">
          <h3 class="small">{{ item }}</h3>
        </el-carousel-item>
      </el-carousel>
    </el-card>
  </div>
</template>

<script lang="ts" setup>

import service from "../util/request";
import {onMounted, computed, ref} from "vue";
import {useStore} from "vuex";

const data = ref()
const data2 = ref()
const store = useStore()

const bodyStyle = {
  padding: '10px'
}
const testGet = () => {
  service.get('/v1/demo/get', {
    headers: {
      'Authorization': store.state.count + "test"
    },
    params: {
      "param1": "test1",
      "param2": "test2"
    }
  }).then(res => {
        data.value = res
      }
  )
}
data2.value = computed(() => store.state.count)
const addClick = () => {
  store.commit('increment')
}
onMounted(testGet)


</script>

<style lang="scss" scoped>
.home {
}

.main-card {
  min-height: 2000px;
  margin: 0;
  padding: 0;
}

.el-carousel__item h3 {
  color: #475669;
  font-size: 14px;
  opacity: 0.75;
  line-height: 150px;
  margin: 0;
  text-align: center;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}
</style>
