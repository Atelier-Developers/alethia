<template>
  <div>
    <v-dialog v-model="dialogComment" max-width="750px">
      <v-card flat style="border-radius: 7px;">
        <v-card-title>
        </v-card-title>
        <v-card-text>
          <v-textarea
              outlined
              v-model="newComment.text"
              label="What do you think"
          />
          <v-row>
            <v-spacer/>
            <v-btn @click="sendComment()" color="primary">
              Comment
            </v-btn>
          </v-row>
        </v-card-text>
      </v-card>
    </v-dialog>
    <v-container>
      <v-row>
        <v-col cols="12">
          <div  >
            <Post @updatePost="getAllPostDetail" :post="post" :reposts="reposts" :repost="repost" :liked="liked"
                  :likes="likes"/>
          </div>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12">
            <v-row>
              <v-col cols="4">
                <h1>
                  Comments:
                </h1>
              </v-col>
              <v-spacer/>
              <v-col cols="4" style="display: flex">
                <v-spacer/>
                <v-btn class="mt-2" @click="dialogComment=true" elevation="0" color="primary">
                  Add Comment
                </v-btn>
              </v-col>
            </v-row>
            <v-row v-for="comment in comments" :key="comment.id">
              <v-col cols="12" v-if="comment.replied_comment_id.Valid">
                <Comment :comment="comment" :other="getRepliedComment(comment.replied_comment_id.Int64)"/>
              </v-col>
              <v-col cols="12" v-else>
                <Comment @sendReply="getAllPostDetail" :comment="comment"/>
              </v-col>
            </v-row>
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
    ...mapGetters('PostModules', ["post", "likes", "comments", "reposts", "repost"]),
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
          "getRepost",
          "getReposts"
        ]),
    async sendComment() {
      let id = +this.$route.params.id;
      await this.commentPost({
        text: this.newComment.text,
        post_id: id
      }).then(() => {
        this.dialogComment = false;
      })
      this.dialogComment = false;
      await this.getAllPostDetail();
    },
    checkAmILike() {
      let my_id = this.userId;
      if (this.likes === null)
        return false;
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
    },
    gotoPost(id) {
      // console.log(id)
      this.$router.replace({name: 'Post', params: {id: id}});
      this.getPost(id).then(() => {
        this.getRepost(this.post.repost_id);
      });
      this.getLikes(id).then(() => {
        this.liked = this.checkAmILike()
      });
      this.getComments(id);
      this.getRepost(id);
      this.getReposts(id)
    },
    async getAllPostDetail() {
      let id = +this.$route.params.id;
      await this.getPost(id)
      await this.getRepost(this.post.repost_id);
      await this.getLikes(id)
      this.liked = this.checkAmILike()
      await this.getComments(id);
      await this.getReposts(id)
      console.log('this.reposts')
      console.log(this.reposts)
    }
  },
  async mounted() {
    await this.getAllPostDetail();
  }
}
</script>

<style scoped>

</style>