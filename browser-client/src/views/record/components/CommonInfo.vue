<template>
  <n-card bordered>
    <n-space>
      <n-badge value="已被封禁！">
        <n-avatar
            round
            :size="60"
            src="https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png"
        />
      </n-badge>

    </n-space>
  </n-card>
</template>

<script lang="ts">
export default {
  name: "CommonInfo"
}
</script>

<script lang="ts" setup>
import {NCard} from "naive-ui";
</script>

<style scoped>

</style>
