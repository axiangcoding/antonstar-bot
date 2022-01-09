<template>
  <div class="about">
    <h1>about page</h1>
    {{getTest}}
  </div>
</template>
<script lang="ts">
import { defineComponent } from "vue";
import {mapGetters} from 'vuex';

export default defineComponent({
  name: "About",
  computed:{
    ...mapGetters(['getTest'])
  },
  methods:{}
});
</script>
