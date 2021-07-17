<template>
  <div>
    <v-dialog v-model="dialogComment" max-width="750px">
      <v-card>
        <v-card-title>
          <h3>Add New Comment</h3>
        </v-card-title>
        <v-divider/>
        <v-container>
          <v-row>
            <v-col cols="12">
              <v-textarea
                  v-model="newComment.text"
                  label="What do you think"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-btn @click="sendComment()">
                Comment
              </v-btn>
            </v-col>
          </v-row>
        </v-container>
      </v-card>

    </v-dialog>
    <v-container>
      <v-row>
        <v-col cols="12">
          <Post :post="post" :liked="liked" :likes="likes"/>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12">
          <v-card>
            <v-card-title>
              <h4>
                Comments:
              </h4>
              <v-spacer/>
              <v-btn @click="dialogComment=true" elevation="0" icon>
                <v-icon>mdi-plus</v-icon>
              </v-btn>
            </v-card-title>
            <v-divider/>
            <v-container>
              <v-row v-for="comment in comments" :key="comment.id">
                <v-col cols="12" v-if="comment.replied_comment_id.Valid">
                  <Comment :comment="comment" :other="getRepliedComment(comment.replied_comment_id.Int64)"/>
                </v-col>
                <v-col cols="12" v-else>
                  <Comment :comment="comment"/>
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
import Post from "../components/Post";
import Comment from "../components/Comment";

export default {
  name: "PostPage",
  components: {Comment, Post},
  data: () => ({
    dialogComment: false,
    newComment: {
      text: '',
    },
    liked: false,
  }),
  computed: {
    ...mapGetters('PostModules', ["post", "likes", "comments", "reposts"]),
    ...mapGetters("AuthModules", ['userId'])
  },
  methods: {
    ...mapActions('PostModules',
        [
          "getPost",
          "getComments",
          "getLikes",
          "likePost",
          "likeComment",
          "replyComment",
          "commentPost",
          "addRepost",
          "getRepost"
        ]),
    sendComment() {
      let id = +this.$route.params.id;
      this.commentPost({
        text: this.newComment.text,
        post_id: id
      })
    },
    checkAmILike() {
      let my_id = this.userId;
      for (let l of this.likes)
        if (+l.user_id === +my_id)
          return true;
      return false;
    },
    getRepliedComment(id) {
      for (let c of this.comments) {
        if (+c.id === +id)
          return c;
      }
      return null;
    }
  },
  mounted() {
    let id = +this.$route.params.id;
    this.getPost(id);
    this.getLikes(id).then(() => {
      this.liked = this.checkAmILike()
    });
    this.getComments(id);
    this.getRepost(id);

  }
}
</script>

<style scoped>

</style>