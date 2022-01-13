<template>
  <el-container direction="vertical" id="app">
    <Header></Header>
    <el-scrollbar>
      <el-backtop target=".el-scrollbar .el-scrollbar__wrap"/>
      <el-main>
        <el-row>
          <!-- 空白边栏 -->
          <el-col :xs="0" :sm="0" :md="0" :lg="0" :xl="2"></el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="20">
            <el-row :gutter="10">
              <el-col :xs="24" :sm="24" :md="24" :lg="20" :xl="20">
                <router-view/>
              </el-col>
              <el-col :xs="24" :sm="24" :md="24" :lg="4" :xl="4">
                <!--<el-affix :offset="89">-->
                <el-card v-adhereTop:[direction]="{ value:20, parent: '.el-scrollbar .el-scrollbar__wrap' }" >
                  <el-empty description="正在施工中"></el-empty>
                </el-card>
                <!--</el-affix>-->
              </el-col>
            </el-row>
          </el-col>
          <!-- 空白边栏 -->
          <el-col :xs="0" :sm="0" :md="0" :lg="0" :xl="2"></el-col>
        </el-row>
      </el-main>
      <Footer></Footer>
    </el-scrollbar>
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

<script lang="ts" setup>
import {v4 as uuid} from "uuid"
import {useStore} from "vuex";
import {onMounted} from "vue";

const store = useStore()
const direction = 'top'
const generateClientId = () => {
  store.commit('setClientId', uuid())
}

onMounted(generateClientId)

</script>

<style>
html,body {
  margin: 0;
  overflow: hidden;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  background-color: #f6f6f7;
  font-size: 16px;
  color: #2c3e50;
  height: 100vh;
  overflow-x: hidden;
}


</style>
