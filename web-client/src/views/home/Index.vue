<template>
  <div class="home">
    <el-card class="card-main" :body-style="bodyStyle">
      <el-carousel height="400px">
        <el-carousel-item v-for="item in carousel_list" :key="item.id">
          <!--<h2>{{ item.content }}</h2>-->
          <el-image style="width: 100%;height: 400px"
                    :src="item.url" fit="cover">

          </el-image>
        </el-carousel-item>
      </el-carousel>
      <el-empty description="正在施工中"></el-empty>
    </el-card>
  </div>
</template>

<script lang="ts" setup>

import service from "../../util/request";
import {onMounted, computed, ref} from "vue";
import {useStore} from "vuex";

const carousel_list = ref([
  {
    id: 1,
    content: '安东星内测啦，目前仅战绩查询可用！敬请期待！',
    to: '',
    url: 'https://s2.loli.net/2022/01/14/2wIoaU193gQYn6A.png'
  },
  {
    id: 2,
    content: '【广告位招租】 ',
    to: '',
    url: 'https://s2.loli.net/2022/01/14/CHkQW3msRbT5ayS.png'
  },
  {
    id: 3,
    content: '【广告位招租】 ',
    to: '',
    url: 'https://s2.loli.net/2022/01/14/O681mVtkdKDRxZz.png'
  }
])

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

.card-main {

  margin: 0;
  padding: 0;
}

:deep(.el-carousel__item) h2 {
  color: #475669;
  opacity: 0.75;
  margin: 0;
  text-align: center;
  line-height: 200px;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}
</style>
