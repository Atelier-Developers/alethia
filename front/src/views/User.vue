<template>
  <div>
    <v-dialog max-width="700px" v-model="dialogInvite">
      <v-card flat style="border-radius: 7px">
        <v-card-title>
          <h4>Invitation</h4>
        </v-card-title>
        <v-card-text>
          <v-form @submit.prevent="sendCreateInvite">
            <v-row>
              <v-col cols="12">
                <v-textarea
                    outlined
                    lable="Invite Message"
                    v-model="inviteReq.body"
                />
              </v-col>
              <v-col cols="12">
                <v-row>
                  <v-spacer/>
                  <v-btn class="mx-2" color="primary" type="submit">
                    Send
                  </v-btn>
                </v-row>
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>
      </v-card>
    </v-dialog>
    <v-container>
      <v-row>
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
              <v-btn v-if="!user.is_invited_by_this_user && user.is_friends_with_this_user === 0" class="mr-5"
                     @click="dialogInvite=true" color="primary">
                INVITE
                <v-icon right>
                  mdi-plus
                </v-icon>
              </v-btn>
              <v-btn @click="startConversation()" color="secondary" v-if="!user.has_conversation_with_this_user">
                Message
                <v-icon right>
                  mdi-message
                </v-icon>
              </v-btn>
            </v-card-title>
            <v-card-subtitle>
              Lives in {{ user.location }}
            </v-card-subtitle>
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
            </v-card-title>
            <v-divider/>
            <v-container>
              <v-row dense>
                <v-col cols="6" v-for="skill in skills" :key="skill.id">
                  <Skill :skill="skill" :user_id="user_id"/>
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
    user_id: 0
  }),
  computed: {
    ...mapGetters("UserModules",
        ["user",
          "skills",
          "languages",
          "backgrounds",
        ]),
    ...mapGetters('AuthModules', ['userId'])
  },
  methods: {
    ...mapActions("UserModules",
        ["getUserById",
          "getSkills",
          "getLanguages",
          "getBackgrounds"]),
    ...mapActions("InviteFriendModules", ["createInvite"]),
    ...mapActions("ConversationModules", ["createConversation"]),
    ...mapActions("NotifModules", ["viewProfile"]),
    async sendCreateInvite() {
      this.inviteReq.receiver_id = +this.$route.params.id;
      await this.createInvite(this.inviteReq);
      await this.getUserById({
        id: +this.$route.params.id
      })
      this.dialogInvite = false
    },
    async startConversation() {
      let id = +this.$route.params.id
      await this.createConversation({user_id_2: id})
      await this.$router.push({name: "Conversation"})
    }
  },
  async mounted() {
    let id = +this.$route.params.id
    this.user_id = id;
    await this.getUserById({id: id})
    await this.getBackgrounds(id)
    await this.getSkills(id)
    await this.getLanguages(id)
    if (id !== +this.userId)
      await this.viewProfile({id: id})
  },
  watch: {
    '$route.params.id': async () => {
      let id = +this.$route.params.id
      this.user_id = id;
      await this.getUserById({id: id})
      await this.getBackgrounds(id)
      await this.getSkills(id)
      await this.getLanguages(id)
      if (id !== +this.userId)
        await this.viewProfile({id: id})
    }
  }
}
</script>

<style scoped>

</style>