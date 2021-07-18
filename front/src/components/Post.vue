<template>
  <div>
    <v-dialog v-model="dialogLikes" max-width="500">
      <v-card flat style="border-radius: 7px;">
        <v-card-title>
          <h3>Likes: </h3>
        </v-card-title>
        <v-divider/>
        <v-container>
          <v-row v-for="like in likes" :key="like.id">
            <v-col cols="12">
              <h4>{{ like.username }}</h4>
              <v-divider/>
            </v-col>
          </v-row>
        </v-container>
      </v-card>
    </v-dialog>
    <v-dialog v-model="dialogRepost" max-width="750">
      <v-card>
        <v-card-title>
          <h3>Repost: </h3>
        </v-card-title>
        <v-divider/>
        <v-container>
          <v-row>
            <v-col cols="12">
              <v-textarea
                  label="Repost content"
                  v-model="newRepost"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-card>
                <v-card-title>
                  <h4>
                    {{ post.poster_username }}
                  </h4>
                </v-card-title>
                <v-divider/>
                <v-container>
                  <v-row>
                    <v-col cols="12">
                      <div class="post-content">
                        {{ post.description }}
                      </div>
                    </v-col>
                  </v-row>
                </v-container>
              </v-card>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-btn color="success" @click="sendRepost()">
                Repost
              </v-btn>
            </v-col>
          </v-row>
        </v-container>
      </v-card>
    </v-dialog>
    <v-card>
      <v-card-title>
        <h4>
          {{ post.poster_username }}
        </h4>
        <v-spacer/>
        <v-btn @click="dialogRepost=true" elevation="0" icon>
          <v-icon>mdi-repeat</v-icon>
        </v-btn>
        <v-btn @click="dialogLikes=true" elevation="0" icon>
          {{ reposts === null ? 0 : reposts.length }}
        </v-btn>
        <v-btn @click="sendLike()" elevation="0" icon>
          <v-icon v-if="liked" color="red">mdi-heart</v-icon>
          <v-icon v-else>mdi-heart-outline</v-icon>
        </v-btn>
        <v-btn @click="dialogLikes=true" elevation="0" icon>
          {{ likes === null ? 0 : likes.length }}
        </v-btn>
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
        <v-row>
          <v-col cols="12">
            <div class="post-content">
              {{ post.description }}
            </div>
          </v-col>
        </v-row>
      </v-container>
      <v-container v-if="+post.repost_id!== 0">
        <v-row>
          <v-col cols="12">
            <v-card>
                <v-card-title>
                  <h4>
                    {{ repost.poster_username }}
                  </h4>
                  <v-spacer/>
                </v-card-title>
                <v-card-subtitle>
                  <v-row>
                    <v-col cols="12">
                      {{ isoToDate(repost.created) }}
                    </v-col>
                  </v-row>
                </v-card-subtitle>
                <v-divider/>
                <v-container>
                  <v-row class="my-3">
                    <v-col cols="12">
                      <div class="post-content">
                        {{ repost.description }}
                      </div>
                    </v-col>
                  </v-row>
                </v-container>
            </v-card>

          </v-col>
        </v-row>
      </v-container>
    </v-card>
  </div>
</template>

<script>
import {mapActions} from "vuex";

export default {
  name: "Post",
  props: {
    post: Object,
    reposts: Object,
    repost: Object,
    liked: Boolean,
    likes: Array
  },
  data: () => ({
    dialogLikes: false,
    dialogRepost: false,
    newRepost: '',
  }),
  methods: {
    ...mapActions('PostModules', ['likePost', "addRepost"]),
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
    sendLike() {
      if (this.liked) {
        return; //TODO UNLIKE?1
      } else {
        this.likePost({
          post_id: this.post.id
        })
      }
    },
    sendRepost() {
      this.addRepost({
        description: this.newRepost,
        repost_id: this.post.id
      }).then(() => {
        this.dialogRepost = false;
      });
    },


  }
}
</script>

<style scoped>

</style>