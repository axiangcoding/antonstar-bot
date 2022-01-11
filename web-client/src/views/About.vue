<template>
  <div class="about">
    <h1>about page</h1>
    {{ getTest }}
    <el-button @click="Test">setUserID</el-button>
  </div>
</template>
<script lang="ts">
import {defineComponent} from "vue";
import {mapGetters} from 'vuex';

export default defineComponent({
  name: "About",
  computed: {
    ...mapGetters(['getTest'])
  },
  methods: {}
});
</script>

<script lang="ts" setup>
import {useStore} from "vuex";

const store = useStore()
const Test = () => {
  store.commit('setUserId', 100)
}
</script>
