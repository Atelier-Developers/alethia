<template>
  <div style="height: 100%;position: absolute;width: 100%; padding: 50px">
    <v-row>
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
      <v-col cols="8" style="height: 100%" class="px-10">
        <v-row v-for="message in messages" :key="message.id" class="my-5">
          <v-spacer v-if="+userId===+message.user_id"/>
          <v-col cols="4" class="message" @click="selectMessage(message)">
            <div v-if="message.reply_id > 0" class="mb-3 elevation-3" style="padding: 3px;">
              <h6>reply: {{ message.replied_message_username }}</h6>
             <span style="font-size: 12px">{{ message.replied_message_body }}</span>
            </div>
            <div>
              {{ message.body }}
            </div>
          </v-col>
        </v-row>
      </v-col>
    </v-row>
    <v-row v-if="selectedConv !== null">
      <v-spacer/>
      <v-col cols="8">
        <v-row v-if="selectedMessage.id > 0" dense>
          <v-col cols="12">
            <h6>Reply:</h6>
            <div>
              {{ selectedMessage.body }}
            </div>
          </v-col>
        </v-row>
        <v-row dense>
          <v-col cols="10">
            <v-text-field
                label="message"
                v-model="newMessage"
            />
          </v-col>
          <v-col cols="2">
            <v-btn @click="sendMessage" color="primary">Send</v-btn>
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
    }
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
      if (this.selectedMessage.id === -1) {
        this.createMessage({
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
.conv {
  background-color: white;
  border-bottom: 1px solid black;
}

.conv.conv-active {
  background-color: gray;
  color: white;
}

.message {
  background-color: white;
  border-radius: 3px;
}
</style>