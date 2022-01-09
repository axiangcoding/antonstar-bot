<template>
  <div class="home">
    {{ data }}
    <el-divider></el-divider>
    {{ data2 }}
    <button @click="addClick()">测试</button>
    <button @click="testGet()">测试2</button>
  </div>
</template>

<script lang="ts" setup>

import service from "../util/request";
import {onMounted, computed, ref} from "vue";
import {useStore} from "vuex";

const data = ref()
const data2 = ref()
const store = useStore()
const testGet = () => {
  service.get('/v1/demo/get', {
    headers: {
      'token': store.state.count + "test"
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
