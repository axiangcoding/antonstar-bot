<template>
  <el-footer class="footer">
    <el-space :size="size" spacer="|">
      <div v-for="i in 2" :key="i">

        <el-link type="info"> button {{ i }}</el-link>

      </div>
    </el-space>
  </el-footer>
</template>

<script lang="ts" setup>
import {ref} from 'vue'

const size = ref(10)
</script>

<style lang="scss" scoped>
.footer {
  background: rgba(0, 0, 0, .01);
}
</style>
