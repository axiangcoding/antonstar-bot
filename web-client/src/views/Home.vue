<template>
  <div class="home">
    {{ data }}
  </div>
</template>

<script lang="ts" setup>

import service from "../util/request";
import {onMounted, ref} from "vue";

const data = ref()

const testGet = () => {
  service.get('/v1/demo/get', {
    params: {
      "param1": "test1",
      "param2": "test2"
    }
  }).then(res => {
        data.value = res
      }
  )
}
onMounted(testGet)


</script>
