<template>
  <el-footer class="footer" height="auto">
    <el-space direction="vertical">
      <el-space :size="10" spacer="|">
        <div v-for="item in btnList" :key="item.id">
          <el-link type="info" @click="routeTo(item.to)">{{ item.name }}</el-link>
        </div>
      </el-space>
      <el-link type="info">备案号</el-link>
    </el-space>
  </el-footer>
</template>

<script lang="ts" setup>
import {ref} from 'vue'
import {useRouter} from "vue-router";

const btnList = ref([
  {id: 1, name: '加入我们', to: 'About'},
  {id: 2, name: '意见反馈', to: ''}
])

const route = useRouter()
const routeTo = (name: string) => {
  route.push({name: name})
}
</script>

<style lang="scss" scoped>
.footer {
  //background: rgba(0, 0, 0, .01);
  background-color: #f6f6f7;
}
</style>
