<template>
  <div>
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
                  <p>
                    <span class="font-weight-bold">Full Name:</span> {{ user.first_name }} {{ user.last_name }}
                  </p>
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12">
                  <p>
                    <span class="font-weight-bold">Intro:</span> {{ user.intro }}
                  </p>
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12">
                  <p>
                    <span class="font-weight-bold">About:</span> {{ user.about }}
                  </p>
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12">
                  <p>
                    <span class="font-weight-bold">Accomplishments:</span> {{ user.accomplishments }}
                  </p>
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12">
                  <p>
                    <span class="font-weight-bold">Additional Info:</span> {{ user.additional_info }}
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
                <v-col cols="6" v-for="skill in skills" :key="skill.id">
                  <Skill :skill="skill"/>
                </v-col>
              </v-row>
            </v-container>
          </v-card>
        </v-col>
        <v-col cols="12">
          <v-card>
            <v-card-title>
              <h2>
                Languages
              </h2>
              <v-spacer/>
              <v-icon color="black" @click="dialogLang=true">mdi-plus</v-icon>
            </v-card-title>
            <v-divider/>
            <v-container>
              <v-row>
                <v-col cols="12">
                  <v-list>
                    <v-list-item v-for="lang in languages" :key="lang.id">
                      {{ lang.language }}
                    </v-list-item>
                  </v-list>
                </v-col>
              </v-row>
            </v-container>
          </v-card>
        </v-col>
        <v-col cols="12">
          <v-card>
            <v-card-title>
              <h2>
                Backgrounds
              </h2>
            </v-card-title>
            <v-divider/>
            <v-container>
              <v-row>
                <v-col cols="6" v-for="back in backgrounds" :key="back.id">
                  <Background :back="back"/>
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
  </div>
</template>

<script>
import {mapActions, mapGetters} from "vuex";

export default {
  name: "User",
  computed: {
    ...mapGetters("UserModules",
        ["user",
          "skills",
          "languages",
          "backgrounds",
        ]),
  },
  methods: {
    ...mapActions("UserModules",
        ["getUserById",
          "getSkills",
          "getLanguages",
          "getBackgrounds"]),
  },
  mounted() {
    this.getUserById({id: +this.$route.params.id})
    this.getBackgrounds(this.user.id)
    this.getSkills(this.user.id)
    this.getLanguages(this.user.id)
  }
}
</script>

<style scoped>

</style>