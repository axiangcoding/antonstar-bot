<template>
  <el-container direction="vertical" id="app">
    <Header></Header>
    <el-main>
      <el-row>
        <!-- 空白边栏 -->
        <el-col :span="3"></el-col>
        <el-col :span="18">
          <el-row :gutter="6">
            <el-col :span="20">
              <router-view/>
            </el-col>
            <el-col :span="4">
              <el-affix :offset="88">
              <el-card>
                <el-empty description="正在施工中"></el-empty>
              </el-card>
              </el-affix>
            </el-col>
          </el-row>
        </el-col>
        <!-- 空白边栏 -->
        <el-col :span="3"></el-col>
      </el-row>
    </el-main>
    <Footer></Footer>
  </el-container>
</template>

<script lang="ts">
import Header from './components/Header.vue'
import Footer from './components/Footer.vue'
import {defineComponent} from 'vue'

export default defineComponent({
  name: 'App',
  components: {Header, Footer}
})
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  background-color: #f6f6f7;
  color: #2c3e50;
  height: 98vh;
}


</style>
