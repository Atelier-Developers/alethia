<template>
  <div>
    <v-dialog v-model="dialogPost">
      <v-card>
        <v-card-title>
          New Post
        </v-card-title>
        <v-divider/>
        <v-container>
          <v-row>
            <v-col cols="12">
              <v-textarea
                  v-model="newPost.description"
                  label="Your Content"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-btn @click="sendPost()">
                Post
              </v-btn>
            </v-col>
          </v-row>
        </v-container>
      </v-card>
    </v-dialog>
    <v-container>
      <v-row v-for="post in posts" :key="post.id">
        <v-col cols="12">
          <Post :post="post"></Post>
        </v-col>
      </v-row>
    </v-container>
    <v-btn
        fab
        bottom
        right
        class="btn-float"
        @click="dialogPost=true"
    >
      <v-icon black x-large>mdi-plus</v-icon>
    </v-btn>
  </div>
</template>

<script>

import {mapActions, mapGetters} from "vuex";
import Post from "../components/Post";

export default {
  name: "Home",
  components: {Post},
  data: () => ({
    dialogPost: false,
    newPost: {
      text: '',
    }
  }),
  computed: {
    ...mapGetters('PostsModules', ["posts"])
  },
  methods: {
    ...mapActions('PostsModules', ["getPosts", "addPost"]),
    sendPost() {
      this.addPost(this.newPost);
    },
  },
  mounted() {
    this.getPosts();
  }
}
</script>

<style scoped>
.btn-float {
  bottom: 0;
  right: 0;
  position: fixed;
  margin: 0 0 16px 16px;
}
</style>