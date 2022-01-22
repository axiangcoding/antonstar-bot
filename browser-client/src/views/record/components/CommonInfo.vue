<template>
  <n-card bordered>
    <n-space align="center" vertical>
      <n-badge v-if="data" value="已被封禁！" :show="data.banned">
        <n-avatar
            round
            :size="80"
            src="https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png"
        />
      </n-badge>
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
      <n-gradient-text size="16" :gradient="{from: '#555555', to: '#000000'}"> {{ data.nick }}</n-gradient-text>
      <n-tag size="large" type="info"> 等级 {{ data.level }}</n-tag>
      <n-gradient-text size="14" :gradient="{from: '#555555', to: '#aaaaaa'}">注册于 {{ data.register_date }}
      </n-gradient-text>
    </n-space>
  </n-card>
</template>

<script lang="ts" setup>
defineProps({
  data: {}
})
</script>

<style scoped>

</style>
