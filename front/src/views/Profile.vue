<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <v-card>
          <v-card-title>
            <h2>
              {{ user.username }}
            </h2>
          </v-card-title>
          <v-divider/>
          <v-container>
            <v-row>
              <v-col cols="12">
                <h3>info</h3>
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="12">
                <p>
                  {{ user.info }}
                </p>
              </v-col>
            </v-row>
          </v-container>
        </v-card>
      </v-col>
      <v-col cols="12">
        <v-card>
          <v-card-title>
            <h2>
              Skills
            </h2>
          </v-card-title>
          <v-divider/>
          <v-container>
            <v-row>
              <v-col cols="6" v-for="skill in user.skills" :key="skill.id">
                <Skill :skill="skill"/>
              </v-col>
            </v-row>
          </v-container>
        </v-card>
      </v-col>
      <v-col cols="12">
        <v-card color="grey lighten-2">
          <v-card-title style="background-color: white">
            <h2>
              Featured Posts
            </h2>
          </v-card-title>
          <v-divider/>
          <v-container>
            <v-row v-for="post in user.posts" :key="post.id">
              <v-col cols="12">
                <Post :post="post" :username="user.username"/>
              </v-col>
            </v-row>
          </v-container>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import Skill from "../components/Skill";
import {mapActions, mapGetters} from "vuex";
import Post from "../components/Post";

export default {
  name: "Profile",
  components: {Post, Skill},
  computed: {
    ...mapGetters('UserModules', ["user"])
  },
  methods: {
    ...mapActions('UserModules', ["getUser"])
  },
  mounted() {
    this.getUser();
  }
}
</script>

<style scoped>

</style>