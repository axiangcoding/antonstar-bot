<template>
  <div class="about">
    <h1>This is an about page</h1>
    <div>{{ res }}</div>
  </div>
</template>

<script>
import {HTTP} from "@/plugins/http";

export default {
  data: () => ({
    res: ''
  }),
  mounted() {
    this.testGet()
  },
  methods: {
    testGet() {
      HTTP.post('posts').then(res => {
        this.res = res
      })
    }
  }
}
</script>
