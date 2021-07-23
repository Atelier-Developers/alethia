<template>
  <div>
    <v-dialog v-model="dialogReply" max-width="750">
      <v-card>
        <v-card-title>
          <h3>Reply: </h3>
        </v-card-title>
        <v-divider/>
        <v-container>
          <v-row>
            <v-col cols="12">
              <v-textarea
                  label="your Reply"
                  v-model="newReply.text"
                  outlined
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-btn @click="sendReply()" color="primary">
                Reply
              </v-btn>
            </v-col>
          </v-row>
        </v-container>
      </v-card>
    </v-dialog>
    <v-card>
      <v-card-title>
        {{ comment.commenter_username }}:
      </v-card-title>
      <v-card-subtitle>
        {{ isoToDate(comment.created) }}
      </v-card-subtitle>
      <v-divider/>
      <v-card v-if="comment.replied_comment_id.Valid" class="my-2 mx-16">
        <v-card-title>
          {{ other.commenter_username }}:
        </v-card-title>
        <v-card-subtitle>
          {{ isoToDate(other.created) }}
        </v-card-subtitle>
        <v-container>
          <v-row class="my-3">
            <v-col cols="12">
              <div class="post-content">
                {{ other.text }}
              </div>
            </v-col>
          </v-row>
        </v-container>
      </v-card>
      <v-container>
        <v-row class="my-3">
          <v-col cols="12">
            <div class="post-content" style="white-space: pre-line;color: black">
              {{ comment.text }}
            </div>
          </v-col>
        </v-row>
      </v-container>
      <v-divider/>
      <v-container>
        <v-row>
          <v-col cols="12">
            <v-btn @click="dialogReply=true" elevation="0" icon>
              <v-icon>mdi-reply</v-icon>
            </v-btn>
            {{comment.reply_count}}
            <v-btn @click="sendLikeComment()" elevation="0" icon>
              <v-icon v-if="comment.is_liked_by_this_user" color="red">mdi-heart</v-icon>
              <v-icon v-else>mdi-heart-outline</v-icon>
            </v-btn>
            {{comment.like_count}}

          </v-col>
        </v-row>
      </v-container>
    </v-card>
  </div>
</template>

<script>
import {mapActions} from "vuex";

export default {
  name: "Comment",
  props: ["comment", 'other'],
  data: () => ({
    newReply: {
      text: '',
    },
    dialogReply: false
  }),
  methods: {
    ...mapActions("PostModules", ['replyComment', 'getCommentLikes', "likeComment"]),
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
    sendReply() {
      this.replyComment({
        text: this.newReply.text,
        post_id: this.comment.post_id,
        replied_comment_id: this.comment.id
      }).then(() => {
        this.dialogReply = false;
      });
    },
    sendLikeComment() {
      if (this.comment.is_liked_by_this_user) {
        return //todo unlike
      } else {
        this.likeComment({comment_id: this.comment.id})
        this.comment.is_liked_by_this_user = true;
      }
    }
  }
}
</script>

<style scoped>

</style>