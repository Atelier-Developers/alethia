<template>
  <div>
    <v-container>
      <v-row justify="center" class="mb-3">
        <v-card flat width="100%" max-width="950px" style="border-radius: 7px">
          <v-card-title>
            Share your ideas with your network
          </v-card-title>
          <v-card-text>
            <v-textarea
                v-model="newPost.description"
                outlined
                label="Your Content"
            />
            <v-row class="px-3 py-2">
              <v-spacer/>
              <v-btn @click="sendPost" color="primary">
                Share
              </v-btn>
            </v-row>
          </v-card-text>
        </v-card>
      </v-row>
      <v-row justify="center" v-for="post in posts" :key="post.id" class="my-3">
        <v-card max-width="950px" width="100%" flat style="border-radius: 7px;">
          <v-card-subtitle v-if="post.type==='LP'">
            <h4> {{ post.like_username }} has Liked this Post:</h4>
          </v-card-subtitle>
          <v-card-subtitle v-else-if="post.type==='CP'">
            <h4> {{ post.commenter_username }} has Commented this Post:</h4>
          </v-card-subtitle>
          <v-card-title>
            <h4 class="pr-2">
              {{ post.poster_firstname }} {{ post.poster_lastname }}
            </h4>
            <h4 class="font-weight-light">
              @{{ post.poster_username }}
            </h4>
            <v-spacer/>
            <v-btn color="secondary" @click="gotoPost(post)">
              Show
              <v-icon right>
                mdi-eye
              </v-icon>
            </v-btn>
          </v-card-title>
          <v-card-subtitle>
            {{ isoToDate(post.created) }}
          </v-card-subtitle>
          <v-divider/>
          <v-card-text style="white-space: pre-line" >
            {{ post.description }}
          </v-card-text>
          <v-container v-if="+post.repost_id!== 0">
            <v-row>
              <v-col cols="12">
                <Repost :repost_id="post.repost_id"/>
              </v-col>
            </v-row>
          </v-container>
          <v-card-text>
            <v-row>
              <div class="mx-1">
                <v-icon>mdi-comment-processing-outline</v-icon>
                {{ post.comment_count }}
              </div>
              <div>
                <v-icon>mdi-repeat</v-icon>
                {{ post.repost_count }}
              </div>
              <v-spacer/>
              <div class="mx-1">
                <v-btn @click="sendLike(post)" elevation="0" icon>
                  <v-icon v-if="post.is_liked_by_this_user" color="red">mdi-heart</v-icon>
                  <v-icon v-else>mdi-heart-outline</v-icon>
                </v-btn>
                {{ post.like_count }}
              </div>
            </v-row>
          </v-card-text>
        </v-card>
      </v-row>
    </v-container>

  </div>
</template>

<script>

import {mapActions, mapGetters} from "vuex";
import Repost from "../components/Repost";

export default {
  name: "Home",
  components: {Repost},
  data: () => ({
    newPost: {
      description: '',
    }
  }),
  computed: {
    ...mapGetters('PostsModules', ["posts"])
  },
  methods: {
    ...mapActions('PostsModules', ["getPosts", "addPost"]),
    ...mapActions('PostModules', ["likePost"]),
    async sendPost() {
      await this.addPost(this.newPost);
      this.newPost.description = '';
      await this.getPosts();
    },
    async sendLike(post) {
      await this.likePost({
        post_id: post.id
      });
      post.is_liked_by_this_user = true;
      await this.getPosts();
    },
    isoToDate(iso) {
      let date = new Date(iso);
      let year = date.getFullYear();
      let month = date.getMonth() + 1;
      let dt = date.getDate();
      let h = date.getHours();
      let m = date.getMinutes();

      if (m < 10)
        m = '0' + m
      if (h < 10)
        h = '0' + h

      if (dt < 10) {
        dt = '0' + dt;
      }
      if (month < 10) {
        month = '0' + month;
      }

      return year + '-' + month + '-' + dt + ' ' + h + ":" + m;
    },
    gotoPost(post) {
      this.$router.push("post/" + post.id)
    }
  },
  async mounted() {
    await this.getPosts();
    console.log(this.posts)
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