<template>
  <n-drawer v-model:show="show" placement="top" :on-update:show="updateStatus" height="500">
    <n-drawer-content title="安东星行为准则 v1.0" closable>
      <ol>
        <li>遵守中华人民共和国的各项有关法律法规.</li>
        <li>不得发布任何色情非法以及危害国家安全的言论.</li>
        <li>严禁链接有关政治, 色情, 宗教, 迷信等违法信息.</li>
        <li>互相尊重, 遵守互联网络道德, 严禁互相恶意攻击, 漫骂.</li>
        <li>本站管理员拥有一切管理权力，有权保留或删除论坛中的任意内容.</li>
        <li>由用户发布的所有信息均代表作者的个人观点, 不能代表"安东星运营方".</li>
        <li>内测期间，用户产生的所有数据，"安东星运营方"均具有在不告知用户的情况下进行清理的权力</li>
        <li>该行为准则将会适时更新</li>
      </ol>
    </n-drawer-content>
  </n-drawer>
</template>


<script lang="ts" setup>
const props = defineProps({
  show: Boolean,
})

const emit = defineEmits(['update-show'])

const updateStatus = (show: boolean) => {
  emit('update-show', show)
}
</script>


<style lang="scss" scoped>
.content-codex {
  min-width: 2000px;
}
</style>
