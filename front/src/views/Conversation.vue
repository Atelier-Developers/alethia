<template>
  <v-container style="height: 100%">
    <v-row>
      <v-col cols="4">
        <v-row v-for="conv in conversations" :key="conv.conversation_id">
          <v-col cols="12" class="conv" @click="sendConv(conv.conversation_id)">
            {{ conv.username }}
          </v-col>
        </v-row>
      </v-col>
      <v-col cols="8" style="height: 100%" class="px-10">
        <v-row v-for="message in messages" :key="message.id" class="my-5">
          <v-spacer v-if="+userId===+message.user_id"/>
          <v-col cols="4" class="message" @click="selectMessage(message)">
            {{ message.body }}
          </v-col>
        </v-row>
      </v-col>
    </v-row>
    <div>
      {{selectedMessage.body}}
    </div>
    <v-text-field
        label="message"
        v-model="newMessage"
    />
    <v-btn @click="sendMessage" color="primary">Send</v-btn>
  </v-container>
</template>

<script>
import {mapActions, mapGetters} from "vuex";

export default {
  name: "Conversation",
  data: () => ({
    newMessage: '',
    selectedConv: null,
    selectedMessage: {}
  }),
  computed: {
    ...mapGetters("ConversationModules", ["conversations", "messages"]),
    ...mapGetters('AuthModules', ['userId'])
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
    sendConv(id) {
      this.getMessages(id).then(() => {
        this.selectedConv = id;
      });
    },
    sendMessage() {
      if (this.selectedMessage === null) {
        this.createMessage({
          reply_id: 1,
          conversation_id: +this.selectedConv,
          message_body: this.newMessage,
        })
      } else {
        this.createMessage({
          reply_id: +this.selectedMessage.id,
          conversation_id: +this.selectedConv,
          message_body: this.newMessage,
        })
      }
    },
    selectMessage(message) {
      if (message.id === this.selectedMessage.id)
        this.selectedMessage = {}
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
.conv {
  background-color: white;
  border-bottom: 1px solid black;
}

.message {
  background-color: white;
  border-radius: 3px;
}
</style>