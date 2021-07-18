<template>
  <div>
    <v-dialog v-model="dialogInvite">
      <v-card>
        <v-card-title>
          <h3>Invitation</h3>
        </v-card-title>
        <v-divider/>
        <v-container>
          <v-row>
            <v-col cols="12">
              <v-textarea
                  lable="Invite Message"
                  v-model="inviteReq.body"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-btn @click="sendCreateInvite()">
                Send
              </v-btn>
            </v-col>
          </v-row>
        </v-container>
      </v-card>
    </v-dialog>
    <v-container>
      <v-row>
        <v-col cols="12">
          <v-card>
            <v-card-title>
              <h2>
                {{ user.username }}
              </h2>
              <v-spacer/>
              <v-btn class="mr-5" @click="dialogInvite=true" color="primary">
                INVITE +
              </v-btn>
              <v-btn @click="startConversation()" color="secondary">
                Message
              </v-btn>
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
              <v-card flat style="border-radius: 7px">
                <v-card-title>
                  <h2 class="pr-2">
                    {{ user.first_name }} {{ user.last_name }}
                  </h2>
                  <h4 class="font-weight-light">
                    @{{ user.username }}
                  </h4>
                  <v-spacer/>
                  <v-btn @click="dialogUser=true" icon>
                    <v-icon>mdi-pencil</v-icon>
                  </v-btn>
                </v-card-title>
                <v-card-subtitle>
                  {{ user.about }}
                </v-card-subtitle>
                <v-divider/>
                <v-card-text>
                  {{ user.intro }}
                </v-card-text>
              </v-card>
            </v-col>
            <v-col cols="12">
              <v-card flat style="border-radius: 7px">
                <v-card-title>
                  <v-icon left>
                    mdi-trophy-outline
                  </v-icon>
                  <h4>
                    Accomplishments
                  </h4>
                </v-card-title>
                <v-divider/>
                <v-card-text>
                  {{ user.accomplishments }}
                </v-card-text>
              </v-card>
            </v-col>
            <v-col cols="12">
              <v-card flat style="border-radius: 7px">
                <v-card-title>
                  <v-icon left>
                    mdi-information-outline
                  </v-icon>
                  <h4>
                    Additional Info
                  </h4>
                </v-card-title>
                <v-divider/>
                <v-card-text>
                  {{ user.additional_info }}
                </v-card-text>
              </v-card>
            </v-col>
            <v-col cols="12">
              <v-card flat style="border-radius: 7px">
                <v-card-title>
                  <v-icon left>
                    mdi-swim
                  </v-icon>
                  <h4>
                    Skills
                  </h4>
                  <v-spacer/>
                  <v-btn @click="dialogSkill=true" elevation="0" icon>
                    <v-icon>mdi-plus</v-icon>
                  </v-btn>
                </v-card-title>
                <v-divider/>
                <v-container>
                  <v-row dense>
                    <v-col cols="6" v-for="skill in skills" :key="skill.id">
                      <Skill :skill="skill"/>
                    </v-col>
                  </v-row>
                </v-container>
              </v-card>
            </v-col>
            <v-col cols="12">
              <v-card flat style="border-radius: 7px">
                <v-card-title>
                  <v-icon left>
                    mdi-translate
                  </v-icon>
                  <h4>
                    Languages
                  </h4>
                  <v-spacer/>
                  <v-btn @click="dialogLang=true" elevation="0" icon>
                    <v-icon>mdi-plus</v-icon>
                  </v-btn>
                </v-card-title>
                <v-divider/>
                <v-container>
                  <v-row dense>
                    <v-col cols="6" v-for="lang in languages" :key="lang.id">
                      <v-list-item>
                        <v-list-item-content>
                          <v-list-item-title>
                            {{ lang.language }}
                          </v-list-item-title>
                        </v-list-item-content>
                      </v-list-item>
                    </v-col>
                  </v-row>
                </v-container>
              </v-card>
            </v-col>
            <v-col cols="12">
              <v-card flat style="border-radius: 7px">
                <v-card-title>
                  <h2>
                    Backgrounds
                  </h2>
                  <v-spacer/>
                  <v-icon @click="dialogBack=true">mdi-plus</v-icon>
                </v-card-title>
                <v-divider/>
                <v-container>
                  <v-row>
                    <v-col cols="12" md="6" v-for="back in backgrounds" :key="back.id">
                      <Background :deletable="false" :back="back"/>
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
import Skill from "../components/Skill";
import Background from "../components/Background";
import Post from "@/components/Post";

export default {
  name: "User",
  components: {Post, Background, Skill},
  data: () => ({
    inviteReq: {
      receiver_id: 0,
      body: '',
    },
    dialogInvite: false,
  }),
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
    ...mapActions("InviteFriendModules", ["createInvite"]),
    ...mapActions("ConversationModules", ["createConversation"]),
    sendCreateInvite() {
      this.inviteReq.receiver_id = +this.$route.params.id;
      this.createInvite(this.inviteReq);
      this.dialogInvite = false
    },
    startConversation() {
      let id = +this.$route.params.id
      this.createConversation({user_id_2: id}).then(() => {
        this.$router.push({name: "Conversation"})
      })
    }
  },
  mounted() {
    let id = +this.$route.params.id
    this.getUserById({id: id})
    this.getBackgrounds(id)
    this.getSkills(id)
    this.getLanguages(id)
  }
}
</script>

<style scoped>

</style>