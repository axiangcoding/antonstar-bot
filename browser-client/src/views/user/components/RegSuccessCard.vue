<template>
  <n-modal
      v-model:show="show"
      :mask-closable="false"
      preset="dialog"
      title="注册成功！"
      content="欢迎来到安东星！3s后将自动跳转到登录界面"
      type="success"
      :closable="false"
  />
</template>

<script lang="ts" setup>
const props = defineProps({
  show: Boolean
})

</script>

<style scoped>

</style>
