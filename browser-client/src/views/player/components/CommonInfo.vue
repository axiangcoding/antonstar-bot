<template>
  <n-card bordered>

    <n-space align="center" vertical>
      <n-avatar
          v-if="data"
          round
          :size="80"
          src="https://pic.imgdb.cn/item/6242d7ae27f86abb2a2fda29.png"
      />
      <n-button
          text
          tag="a"
          :href="'https://warthunder.com/'+data.clan_url"
          target="_blank"
          type="primary"
          v-if="data.clan"
      >
        <n-gradient-text
            :size="18"
            gradient="linear-gradient(90deg, red 0%, green 50%, blue 100%)"
        >{{ data.clan }}
        </n-gradient-text>
      </n-button>
      <n-gradient-text
          :size="18"
          v-else
          type="error"
      >【未加入联队】
      </n-gradient-text>
      <n-gradient-text size="16" :gradient="{from: '#555555', to: '#aaaaaa'}">【{{ data.title }}】
      </n-gradient-text>
      <n-gradient-text size="18" :gradient="{from: '#555555', to: '#000000'}"> {{ data.nick }}</n-gradient-text>
      <n-tag size="large" type="info"> 等级 {{ data.level }}</n-tag>
      <n-gradient-text size="14" :gradient="{from: '#555555', to: '#aaaaaa'}">注册于 {{ data.register_date }}
      </n-gradient-text>
    </n-space>
    <n-alert title="已被官方封禁" type="error" v-if="data.banned">
      🍾🍾🍾好似，开香槟喽！
    </n-alert>
  </n-card>
</template>

<script lang="ts" setup>
defineProps({
  data: {}
})
</script>

<style scoped>

</style>
