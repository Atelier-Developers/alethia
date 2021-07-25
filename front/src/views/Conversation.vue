<template>
  <div>
    <v-row class="mx-5">
      <v-col cols="4" style="border-radius:10px 0 0 10px;overflow-y: scroll;overflow-x:hidden;"
             v-bind:class="$vuetify.theme.dark ? 'black white--text' : 'white black--text'">
        <v-row>
          <v-col cols="12">
            <v-text-field
                outlined
                label="Search"
                dense
                v-model="search"
            />
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12"
                 class="pt-6"
                 style="border-bottom: 1px solid gray!important">
            <v-row>
              <v-col cols="8" class="font-weight-bold">Filter</v-col>
              <v-col cols="2"></v-col>
              <v-col cols="2">
                <v-menu
                    right
                    bottom
                >
                  <template v-slot:activator="{ on, attrs }">
                    <v-btn
                        icon
                        v-bind="attrs"
                        v-on="on"
                    >
                      <v-icon>mdi-dots-vertical</v-icon>
                    </v-btn>
                  </template>

                  <v-list>
                    <v-list-item>
                      <v-list-item-title>
                        <v-checkbox label="UnRead" v-model="filter.read"/>
                      </v-list-item-title>
                    </v-list-item>
                    <v-list-item>
                      <v-list-item-title>
                        <v-checkbox label="Archived" v-model="filter.archived"/>
                      </v-list-item-title>
                    </v-list-item>
                  </v-list>
                </v-menu>
              </v-col>
            </v-row>
          </v-col>
        </v-row>
        <v-row v-for="conv in convs" :key="conv.conversation_id"
               v-bind:class="$vuetify.theme.dark ? 'black white--text' : 'white black--text'">
          <v-col cols="12"
                 v-bind:class="(+selectedConv === +conv.conversation_id ? 'conv py-10 grey conv-active ':'conv py-10 ') + ($vuetify.theme.dark ? 'black white--text' : 'white black--text')"
                 @click="sendConv(conv.conversation_id, conv)"
                 style="border-bottom: 1px solid gray!important">
            <v-row>
              <v-col cols="8">
                {{ conv.username }}
              </v-col>
              <v-col cols="2">
                <v-icon v-if="conv.is_read" color="green">
                  mdi-check-all
                </v-icon>
                <v-icon v-else>
                  mdi-check
                </v-icon>
              </v-col>
              <v-col cols="2">
                <v-menu
                    right
                    bottom
                >
                  <template v-slot:activator="{ on, attrs }">
                    <v-btn
                        icon
                        v-bind="attrs"
                        v-on="on"
                    >
                      <v-icon>mdi-dots-vertical</v-icon>
                    </v-btn>
                  </template>

                  <v-list>
                    <v-list-item @click="changeConv('read',conv.conversation_id,conv)">
                      <v-list-item-title>Read/Unread</v-list-item-title>
                      <v-list-item-action>
                        <v-icon v-bind:color="conv.is_read? 'green': 'red'">mdi-read</v-icon>
                      </v-list-item-action>
                    </v-list-item>
                    <v-list-item @click="changeConv('archive',conv.conversation_id,conv)">
                      <v-list-item-title>Archived/Unarchived</v-list-item-title>
                      <v-list-item-action>
                        <v-icon v-bind:color="conv.is_archived ? 'green': 'red'">mdi-archive</v-icon>
                      </v-list-item-action>
                    </v-list-item>
                    <v-list-item @click="changeConv('delete',conv.conversation_id,conv)">
                      <v-list-item-title>Delete</v-list-item-title>
                      <v-list-item-action>
                        <v-icon>mdi-delete</v-icon>
                      </v-list-item-action>
                    </v-list-item>
                  </v-list>
                </v-menu>
              </v-col>
            </v-row>
          </v-col>
        </v-row>
      </v-col>
      <v-col cols="8" v-if="selectedConv !== null">
        <v-row class="grey elevation-5" style="border-radius: 0 10px 10px 0 ; overflow: hidden"
               v-bind:class="this.$vuetify.theme.dark ? 'black white--text' : 'grey elevation-5'">
          <v-col cols="12"
                 style="max-height: 80vh; min-height: 80vh; height:100%;overflow: scroll; overflow-x: hidden;"
                 ref="char-box"
                 class="grey">
            <v-row dense v-for="message in messages" :key="message.id">
              <v-spacer v-if="+userId===+message.user_id"/>
              <v-col cols="10" md="4" @click="selectMessage(message)">
                <v-card flat style="border-radius: 7px">
                  <template v-if="message.reply_id > 0">
                    <v-card-title v-bind:class="$vuetify.theme.dark ? 'grey darken-2' : 'grey lighten-2'">
                      <v-icon left style="transform: scaleX(-1)">
                        mdi-reply
                      </v-icon>
                      {{ message.replied_message_username }}
                    </v-card-title>
                    <v-card-subtitle v-bind:class="$vuetify.theme.dark ? 'grey darken-2' : 'grey lighten-2'">
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
                <v-col v-if="selectedMessage.id > 0" v-bind:class="$vuetify.theme.dark ? 'black white--text' : 'white black--text'" cols="12">
                  <h6>Reply</h6>
                  <div>
                    {{ selectedMessage.body }}
                  </div>
                </v-col>
                <v-text-field
                    label="message"
                    v-model="newMessage"
                    hide-details
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
    theme: '',
    selectedConv: null,
    selectedMessage: {
      id: -1,
      body: ''
    },
    filter: {
      read: false,
      archived: false
    },
    search: ''
  }),
  computed: {
    ...mapGetters("ConversationModules", ["conversations", "messages"]),
    ...mapGetters('AuthModules', ['userId']),
    convs() {
      let all = [];
      for (let conv of this.conversations) {
        if ((this.filter.read && !this.filter.archived && !conv.is_read)
            || (this.filter.read && this.filter.archived && !conv.is_read && conv.is_archived)
            || (!this.filter.read && this.filter.archived && conv.is_archived)
            || (!this.filter.read && !this.filter.archived))
          if (conv.username.search(this.search) >= 0)
            all.push(conv);
      }
      return all;
    }
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
    async sendConv(id, conv) {
      await this.getMessages(id)
      this.selectedConv = id;
      if (this.$refs["char-box"]) {
        this.$refs["char-box"].scroll(0, this.$refs["char-box"].scrollHeight)
      }
      this.updateConversation({
        conversation_id: +id,
        is_archived: conv.is_archived,
        is_deleted: false,
        is_read: true
      })

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
    },
    changeConv(opt, id, conv) {
      if (opt === 'read') {
        this.updateConversation({
          conversation_id: +id,
          is_archived: conv.is_archived,
          is_deleted: false,
          is_read: !conv.is_read
        })
      } else if (opt === 'archive') {
        this.updateConversation({
          conversation_id: +id,
          is_archived: !conv.is_archived,
          is_deleted: false,
          is_read: conv.is_read
        })
      } else if (opt === 'delete') {
        this.deleteConversation({
          conversation_id: +id,
        })
      }
      this.getConversations()
    }
  },
  mounted() {
    this.getConversations();
    this.theme = this.$vuetify.theme.dark ? 'black white--text' : 'white black--text';
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