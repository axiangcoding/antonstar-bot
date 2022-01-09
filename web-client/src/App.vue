<template>
  <el-container direction="vertical" id="app">
    <Header></Header>
    <el-main>
      <router-view/>
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
  color: #2c3e50;
}


</style>
