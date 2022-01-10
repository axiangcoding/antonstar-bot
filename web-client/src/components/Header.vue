<template>
  <el-header class="header" height="60px">
    <el-row>
      <el-col :span="4">
        <el-image
            style="width: 200px; height: 60px"
            :src="url"
            fit="cover"
        ></el-image>
      </el-col>
      <el-col :span="19">
        <el-menu
            :default-active="activeIndex"
            class="el-menu"
            mode="horizontal"
            @select="handleSelect"
        >
          <el-menu-item v-for="item in titleList" :index="item.id">
            <span style="font-size: var(--el-font-size-medium);font-weight: bold">{{ item.name }}</span>
          </el-menu-item>
        </el-menu>
      </el-col>
      <el-col :span="1">
        <el-dropdown>
          <el-avatar size="default"
                     src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
          ></el-avatar>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item>Action 1</el-dropdown-item>
              <el-dropdown-item>Action 2</el-dropdown-item>
              <el-dropdown-item>Action 3</el-dropdown-item>
              <el-dropdown-item>Action 4</el-dropdown-item>
              <el-dropdown-item>Action 5</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </el-col>
    </el-row>
  </el-header>
</template>

<script lang="ts" setup>
import {ref} from 'vue'
import logo from "../assets/image/logo.png"
import {useRouter} from "vue-router";

const titleList = [
  {id: 'Home', name: '首页'},
  {id: '2', name: '攻略教程'},
  {id: '3', name: '玩家社区'},
  {id: '4', name: '资源下载'},
  {id: '5', name: '战绩查询'},
  {id: '6', name: '封神榜'},
  {id: 'About', name: '关于我们'},
]

const url = ref(logo)
const activeIndex = ref('1')

const router = useRouter()
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
  router.push({name: key})
}


</script>
<style lang="css" scoped>
 :deep(.el-header)  {
  padding: 0;
}

.header {
  text-align: left;
  border-bottom: 1px #ccc solid;
  background-color: #ffffff;
}

.el-menu {
  margin-top: 0;
  border-top: 0;
  border-bottom: 0;
  padding-top: 0;
}
</style>
