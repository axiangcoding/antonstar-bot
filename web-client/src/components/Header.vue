<template>
  <v-app-bar
      app
      color="white"
      flat
  >
    <v-img
        lazy-src=""
        height="100%"
        width="250"
        src="https://picsum.photos/id/11/500/3002"
        alt="项目的logo"
    ></v-img>

    <v-spacer class="px-4"></v-spacer>
    <v-tabs
        align-with-title
        class="ml-n9"
        v-model="selected"
    >
      <v-tab
          v-for="link in links"
          :key="link.id"
          class="text-lg-button font-weight-bold"
          @click="$router.push({name: link.routeName})"
      >
        {{ link.name }}
      </v-tab>
    </v-tabs>
    <v-menu offset-y nudge-top="-8" bottom>
      <template v-slot:activator="{ on, attrs }">
        <v-avatar
            class="hidden-sm-and-down"
            color="grey darken-1 shrink"
            size="40"
            v-bind="attrs"
            v-on="on"
        ></v-avatar>
      </template>
      <v-list>
        <v-list-item
            v-for="(item, index) in items"
            :key="index"
        >
          <v-list-item-title>{{ item.title }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>
  </v-app-bar>
</template>

<script>
export default {
  name: "Header",
  data: () => ({
    links: [
      {id: 1, name: '首页', routeName: 'Home'},
      {id: 2, name: '攻略教程'},
      {id: 3, name: '玩家社区'},
      {id: 4, name: '资源下载'},
      {id: 5, name: '游戏数据'},
      {id: 6, name: '封神榜'},
      {id: 7, name: '关于我们', routeName: 'About'},

    ],
    selected: 2,
    items: [
      {title: 'Click Me'},
      {title: 'Click Me'},
      {title: 'Click Me'},
      {title: 'Click Me 2'},
    ],
  }),
}
</script>

<style scoped>

</style>
