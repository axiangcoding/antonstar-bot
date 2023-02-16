<template>

  <div id="footer">
    Powered by
    <n-image
        preview-disabled
        width="60"
        :src="logo"
    />

  </div>


</template>

<script setup lang="ts">
import logo from '@/assets/logo/logo_text_color.png'
import {onMounted, ref} from "vue";
import api from "@/api";

onMounted(() => {
  getSystemInfo()
})

const version = ref('')

function getSystemInfo() {

  api.get('/api/v1/system/info').then((resp) => {
    version.value = resp.data.data.version
    console.log(resp)
  })
}

</script>


<style scoped>
#footer {
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f5f5f5;
  margin: auto auto;
  height: 70px;
}
</style>