<template>
  <v-container>
    <v-row v-for="post in posts" :key="post.id">
      <v-col cols="12">
        <v-card>
          <v-card-title>
            <v-row>
              <v-col cols="12">
                {{ post.user.username }}
              </v-col>
            </v-row>
          </v-card-title>
          <v-card-subtitle>
            <v-row>
              <v-col cols="12">
                {{ post.date }}
              </v-col>
            </v-row>
          </v-card-subtitle>
          <v-divider/>
          <v-container>
            <v-row>
              <v-col cols="12">
                <div class="post-content">
                  {{ post.text }}
                </div>
              </v-col>
            </v-row>
          </v-container>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>

import {mapActions, mapGetters} from "vuex";

export default {
  name: "Home",
  computed: {
    ...mapGetters('PostsModules', ["posts"])
  },
  methods: {
    ...mapActions('PostsModules', ["getPosts"])
  },
  mounted() {
    this.getPosts();
  }
}
</script>

<style scoped>

</style>