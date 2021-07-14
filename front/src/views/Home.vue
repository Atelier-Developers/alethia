<template>
  <v-container>
    <v-row v-for="post in posts" :key="post.id">
      <v-col cols="12">
        <Post :post="post"></Post>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>

import {mapActions, mapGetters} from "vuex";
import Post from "../components/Post";

export default {
  name: "Home",
  components: {Post},
  computed: {
    ...mapGetters('PostsModules', ["posts"])
  },
  methods: {
    ...mapActions('PostsModules', ["getPosts"])
  },
  mounted() {
    this.getPosts();
  }
}
</script>

<style scoped>

</style>