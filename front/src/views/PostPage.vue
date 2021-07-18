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
          <div @click="gotoPost(repost.id)" style="cursor: pointer">
            <Post @updatePost="getAllPostDetail" :post="post" :reposts="reposts" :repost="repost" :liked="liked"
                  :likes="likes"/>
          </div>
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
      await this.getRepost(id);
      await this.getReposts(id)
    }
  },
  async mounted() {
    await this.getAllPostDetail();
  }
}
</script>

<style scoped>

</style>