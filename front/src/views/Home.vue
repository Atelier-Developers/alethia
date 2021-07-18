<template>
  <div>
    <v-container>
      <v-row justify="center">
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
      <v-row v-for="post in posts" :key="post.id">
        <v-col cols="12">
          <v-card>
            <div @click="gotoPost(post)" style="cursor: pointer">
              <v-card-subtitle v-if="post.type==='LP'">
                <h4> {{ post.like_username }} has Liked this Post:</h4>
              </v-card-subtitle>
              <v-card-subtitle v-else-if="post.type==='CP'">
                <h4> {{ post.commenter_username }} has Commented this Post:</h4>
              </v-card-subtitle>
              <v-card-title>
                <h4>
                  {{ post.poster_username }}
                </h4>
                <v-spacer/>
              </v-card-title>
              <v-card-subtitle>
                <v-row>
                  <v-col cols="12">
                    {{ isoToDate(post.created) }}
                  </v-col>
                </v-row>
              </v-card-subtitle>
              <v-divider/>
              <v-container>
                <v-row class="my-3">
                  <v-col cols="12">
                    <div class="post-content">
                      {{ post.description }}
                    </div>
                  </v-col>
                </v-row>
              </v-container>
            </div>
            <v-container v-if="+post.repost_id!== 0">
              <v-row>
                <v-col cols="12">
                  <Repost :repost_id="post.repost_id"/>
                </v-col>
              </v-row>
            </v-container>
            <v-container>
              <v-row>
                <v-col cols="12">
                  <v-btn elevation="0" icon>
                    <v-icon>mdi-reply</v-icon>
                  </v-btn>
                  {{ post.comment_count }}
                  <v-btn @click="sendLike(post)" elevation="0" icon>
                    <v-icon v-if="post.is_liked_by_this_user" color="red">mdi-heart</v-icon>
                    <v-icon v-else>mdi-heart-outline</v-icon>
                  </v-btn>
                  {{ post.like_count }}
                  <v-btn elevation="0" icon>
                    <v-icon>mdi-repeat</v-icon>
                  </v-btn>
                  {{ post.repost_count }}
                </v-col>
              </v-row>
            </v-container>
          </v-card>
        </v-col>
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
      text: '',
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
      await this.getPosts();
    },
    sendLike(post) {
      if (this.liked) {
        return; //TODO UNLIKE?1
      } else {
        this.likePost({
          post_id: post.id
        });
        post.is_liked_by_this_user = true;
      }
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