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
                </v-col>
              </v-row>
            </v-container>
          </v-card>
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

export default {
  name: "Home",
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
    ...mapActions('PostModules', ["likePost"]),
    sendPost() {
      this.addPost(this.newPost);
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
  mounted() {
    this.getPosts().then(() => {
      console.log(this.posts.map((x) => [x.type, x]))
      console.log(this.posts.map((x) => x.type))
    });
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