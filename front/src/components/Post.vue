<template>
  <div>
    <v-dialog v-model="dialogRepost" max-width="750">
      <v-card>
        <v-card-title>

        </v-card-title>
        <v-card-text>
          <v-form @submit.prevent="sendRepost">
            <v-textarea
                label="Share your thoughts"
                v-model="newRepost"
                outlined
            />
            <v-row>
              <v-col cols="12">
                <v-card flat outlined>
                  <v-card-title>
                    <h4>
                      {{ post.poster_username }}
                    </h4>
                  </v-card-title>
                  <v-divider/>
                  <v-card-text style="white-space: pre-line;">
                    {{ post.description }}
                  </v-card-text>
                </v-card>
              </v-col>
            </v-row>
            <v-row>
              <v-spacer/>
              <v-btn color="primary" type="submit">
                Repost
              </v-btn>
            </v-row>
          </v-form>
        </v-card-text>
      </v-card>
    </v-dialog>
    <v-card flat>
      <v-card-title>
        <h4>
          {{ post.poster_username }}
        </h4>
        <v-spacer/>
        <v-btn @click="dialogRepost=true" elevation="0" icon class="mr-3">
          <v-icon left>mdi-repeat</v-icon>
          {{ reposts === null ? 0 : reposts.length }}
        </v-btn>
        <v-btn @click="sendLike" elevation="0" icon>
          <v-icon v-if="liked" color="red" left>mdi-heart</v-icon>
          <v-icon v-else left>mdi-heart-outline</v-icon>
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
            <div class="post-content" style="white-space: pre-line;">
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
                    <div class="post-content" style="white-space: pre-line;">
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
    reposts: {},
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
    async sendLike() {
      await this.likePost({
        post_id: this.post.id
      })
      this.$emit('updatePost')
    },
    async sendRepost() {
      await this.addRepost({
        description: this.newRepost,
        repost_id: this.post.id
      })
      this.$emit('updatePost')
      this.dialogRepost = false;
    },


  }
}
</script>

<style scoped>

</style>