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
          <v-card flat style="border-radius: 7px;">
            <v-card-title>
              <h3>
                Invites
              </h3>
            </v-card-title>
            <v-divider/>
            <v-card-text>
              <v-row v-for="invite in allInvites" :key="invite.other_user_id">
                <v-col cols="8">
                  <div>
                    <h3>
                      {{ invite.other_user_name }}
                    </h3>
                  </div>
                  <p>
                    {{ invite.body }}
                  </p>
                </v-col>
                <v-spacer/>
                <div v-if="invite.am_i_sender" class="mr-5 mt-5">
                  <p class="body-1">
                    PENDING
                    <v-icon right>mdi-clock</v-icon>
                  </p>
                </div>
                <div v-else style="display: flex" class="pr-5">
                  <v-btn
                      @click="sendAcceptInvite(invite.other_user_id)"
                      color="success"
                      fab
                      icon
                  >
                    <v-icon>mdi-check</v-icon>
                  </v-btn>
                  <v-btn
                      color="error"
                      fab
                      icon
                      @click="sendDeleteInvite(invite.other_user_id)"
                  >
                    <v-icon>mdi-delete</v-icon>
                  </v-btn>
                </div>
              </v-row>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12">
          <v-card flat>
            <v-card-title>
              <h3>
                People You May Know
              </h3>
            </v-card-title>
            <v-divider/>
            <v-container fluid>
              <v-row v-for="m in mutual" :key="m.username">
                <v-col cols="8">
                  <v-row>
                    <v-col cols="12">
                      <h3>
                        {{ m.username }}
                      </h3>
                      Mutual Connections: {{ m.mutual_connection }}
                    </v-col>
                  </v-row>
                </v-col>
                <v-col cols="4" style="display: flex">
                  <v-spacer/>
                  <v-btn
                      @click="dialogInvite=true; inviteReq.receiver_id=+m.user_id"
                      color="primary"
                      class="mt-2"
                  >
                    Invite +
                  </v-btn>
                </v-col>
              </v-row>
            </v-container>
          </v-card>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12">
          <v-card flat style="border-radius: 7px;">
            <v-card-title>
              <h3>
                Friends
              </h3>
            </v-card-title>
            <v-divider/>
            <v-card-text>
              <div v-for="friend in allFriends" :key="friend.other_user_id">
                <v-row>
                  <v-col cols="8">
                    <div>
                      <h3>
                        {{ friend.other_user_name }}
                      </h3>
                    </div>
                  </v-col>
                  <v-spacer/>
                  <v-btn
                      @click="sendDeleteFriend(friend.other_user_id)"
                      color="error"
                      fab
                      icon
                  >
                    <v-icon>mdi-delete</v-icon>
                  </v-btn>
                </v-row>
              </div>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import {mapActions, mapGetters} from "vuex";

export default {
  name: "Network",
  data: () => ({
    dialogInvite: false,
    inviteReq: {
      body: '',
      receiver_id: 0
    }
  }),
  computed: {
    ...mapGetters("InviteFriendModules", ["friends", "invites", "mutual"]),
    ...mapGetters('AuthModules', ['userId']),
    allInvites() {
      if (this.invites == null)
        return [];

      let invites = []
      for (let inv of this.invites) {
        let invite = {
          other_user_id: 0,
          other_user_name: '',
          body: '',
          am_i_sender: false
        }
        if (+inv.user_id_1 === +this.userId) {
          invite.am_i_sender = true;
          invite.body = inv.body;
          invite.other_user_id = +inv.user_id_2;
          invite.other_user_name = inv.username_2;
        } else {
          invite.am_i_sender = false;
          invite.body = inv.body;
          invite.other_user_id = +inv.user_id_1;
          invite.other_user_name = inv.username_1;
        }
        invites.push(invite)
      }
      return invites;
    },
    allFriends() {
      if (this.friends == null)
        return [];
      let friends = []
      for (let inv of this.friends) {
        let friend = {
          other_user_id: 0,
          other_user_name: '',
        }
        if (+inv.user_id_1 === +this.userId) {
          friend.other_user_id = +inv.user_id_2;
          friend.other_user_name = inv.username_2;
        } else {
          friend.other_user_id = +inv.user_id_1;
          friend.other_user_name = inv.username_1;
        }
        friends.push(friend)
      }
      return friends;

    }

  },
  methods: {
    ...mapActions("InviteFriendModules",
        ["getFriends",
          "getInvites",
          "addFriend",
          "deleteFriend",
          "deleteInvite",
          "createInvite",
          "getMutualConnections"]),
    async sendAcceptInvite(id) {
      await this.addFriend({user_id: id})
      await this.getFriends();
      await this.getInvites();
      await this.getMutualConnections();
    },
    async sendDeleteFriend(id) {
      await this.deleteFriend({user_id: id})
      await this.getInvites();
      await this.getFriends();
      await this.getMutualConnections();
    },
    async sendDeleteInvite(id) {
      await this.deleteInvite({user_id: id})
      await this.getInvites();
      await this.getFriends();
      await this.getMutualConnections();
    },
    async sendCreateInvite() {
      this.createInvite(this.inviteReq)
      await this.getInvites();
      await this.getFriends();
      await this.getMutualConnections();
      this.dialogInvite = false;
    }

  },
  async mounted() {
    await this.getInvites();
    await this.getFriends();
    await this.getMutualConnections();

  }
}
</script>

<style scoped>

</style>