<template>
  <div>
    <v-row class="mx-5">
      <v-col cols="4">
        <v-row v-for="conv in conversations" :key="conv.conversation_id">
          <v-col cols="12"
                 v-bind:class="+selectedConv === +conv.conversation_id ? ['conv', 'conv-active']:['conv']"
                 @click="sendConv(conv.conversation_id)"
          >
            {{ conv.username }}
          </v-col>
        </v-row>
      </v-col>
      <v-col cols="8" v-if="selectedConv !== null">
        <v-row class="grey elevation-5" style="border-radius: 10px; overflow: hidden">
          <v-col cols="12"
                 style="max-height: 80vh; height:100%;overflow: scroll; overflow-x: hidden;"
                 ref="char-box"
                 class="grey">
            <v-row dense v-for="message in messages" :key="message.id">
              <v-spacer v-if="+userId===+message.user_id"/>
              <v-col cols="10" md="4" @click="selectMessage(message)">
                <v-card flat style="border-radius: 7px">
                  <template v-if="message.reply_id > 0">
                    <v-card-title class="grey lighten-2">
                      <v-icon left style="transform: scaleX(-1)">
                        mdi-reply
                      </v-icon>
                      {{ message.replied_message_username }}
                    </v-card-title>
                    <v-card-subtitle class="grey lighten-2">
                      {{ message.replied_message_body }}
                    </v-card-subtitle>
                  </template>
                  <v-card-title>
                    {{ message.body }}
                  </v-card-title>
                  <v-card-subtitle>
                    {{ isoToDate(message.created) }}
                  </v-card-subtitle>
                </v-card>
              </v-col>
            </v-row>
          </v-col>
          <v-col cols="12">
            <v-form @submit.prevent="sendMessage">
              <v-row v-if="selectedConv !== null">
                <v-col v-if="selectedMessage.id > 0" class="white lighten-1" cols="12">
                  <h6>Reply</h6>
                  <div>
                    {{ selectedMessage.body }}
                  </div>
                </v-col>
                <v-text-field
                    label="message"
                    v-model="newMessage"
                    hide-details
                    background-color="white"
                    style="border-radius: 0"
                    solo
                />
                <v-btn type="submit" style="display: none;"/>
              </v-row>
            </v-form>
          </v-col>
        </v-row>
      </v-col>

    </v-row>
  </div>
</template>

<script>
import {mapActions, mapGetters} from "vuex";

export default {
  name: "Conversation",
  data: () => ({
    newMessage: '',
    selectedConv: null,
    selectedMessage: {
      id: -1,
      body: ''
    },
  }),
  computed: {
    ...mapGetters("ConversationModules", ["conversations", "messages"]),
    ...mapGetters('AuthModules', ['userId']),
  },
  methods: {
    ...mapActions("ConversationModules", [
      "getConversations",
      "getMessages",
      "updateConversation",
      "deleteConversation",
      "createConversation",
      "createMessage"
    ]),
    async sendConv(id) {
      await this.getMessages(id)
      this.selectedConv = id;
      if (this.$refs["char-box"]) {
        this.$refs["char-box"].scroll(0, this.$refs["char-box"].scrollHeight)
      }
    },
    async sendMessage() {
      if (this.selectedMessage.id === -1) {
        await this.createMessage({
          conversation_id: +this.selectedConv,
          message_body: this.newMessage,
        })
      } else {
        await this.createMessage({
          reply_id: +this.selectedMessage.id,
          conversation_id: +this.selectedConv,
          message_body: this.newMessage,
        })
      }
      await this.getMessages(this.selectedConv)
      this.newMessage = ''
      this.$refs["char-box"].scroll(0, this.$refs["char-box"].scrollHeight)
    },
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
    selectMessage(message) {
      if (message.id === this.selectedMessage.id)
        this.selectedMessage = {
          id: -1,
          body: ''
        }
      else
        this.selectedMessage = message
    }
  },
  mounted() {
    this.getConversations();
  }
}
</script>

<style scoped>

/* Designing for scroll-bar */
::-webkit-scrollbar {
  width: 6px;

}

/* Track */
::-webkit-scrollbar-track {
  background: gainsboro;
  border-radius: 5px;
}

/* Handle */
::-webkit-scrollbar-thumb {
  background: #aca3a3;
  border-radius: 5px;
}

/* Handle on hover */
::-webkit-scrollbar-thumb:hover {
  background: #555;
}
</style>