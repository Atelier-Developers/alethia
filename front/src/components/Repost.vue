<template>
  <v-card flat outlined>
    <div @click="gotoPost()" style="cursor: pointer">
      <v-card-title>
        <h4 class="pr-2">
          {{ post.poster_firstname }} {{ post.poster_lastname }}
        </h4>
        <h4 class="font-weight-light">
          @{{ post.poster_username }}
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
      <v-card-text>
              <div class="post-content" style="white-space: pre-line;">
                {{ post.description }}
              </div>
      </v-card-text>
    </div>
  </v-card>
</template>

<script>
import {mapActions, mapGetters} from "vuex";

export default {
  name: "repost",
  props: ['repost_id'],
  computed: {
    ...mapGetters("PostModules", ["post"]),
  },
  methods: {
    ...mapActions("PostModules", ["getPost"]),
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
    gotoPost() {
      this.$router.push("post/" + this.post.id)
    }
  },
  mounted() {
    this.getPost(this.repost_id)
  }
}
</script>

<style scoped>

</style>