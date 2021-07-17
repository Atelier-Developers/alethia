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
import Skill from "../components/Skill";
import Background from "../components/Background";

export default {
  name: "User",
  components: {Background, Skill},
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