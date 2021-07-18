<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <v-card flat style="border-radius: 7px;">
          <v-card-title>
            <h4>
              Invites
            </h4>
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
</template>

<script>
import {mapActions, mapGetters} from "vuex";

export default {
  name: "Network",
  data: () => ({}),
  computed: {
    ...mapGetters("InviteFriendModules", ["friends", "invites"]),
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
          "deleteInvite"]),
    async sendAcceptInvite(id) {
      await this.addFriend({user_id: id})
      await this.getFriends();
      await this.getInvites();
    },
    async sendDeleteFriend(id) {
      await this.deleteFriend({user_id: id})
      await this.getFriends();
    },
    async sendDeleteInvite(id) {
      await this.deleteInvite({user_id: id})
      await this.getInvites();
    }

  },
  mounted() {
    this.getInvites();
    this.getFriends();
  }
}
</script>

<style scoped>

</style>