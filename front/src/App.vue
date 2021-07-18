<template>
  <v-app>
    <v-dialog v-model="dialogSearch" max-width="500">
      <v-card>
        <v-card-title>
          Search User
        </v-card-title>
        <v-divider/>
        <v-container>
          <v-row dense>
            <v-col cols="12">
              <v-text-field
                  label="Username"
                  v-model="username"
                  outlined
              />
            </v-col>
          </v-row>
          <v-row dense>
          <v-col cols="4">
              <v-btn  @click="doSearch()">
                Search <v-icon>mdi-magnify</v-icon>
              </v-btn>
            </v-col>
          </v-row>
        </v-container>
        <v-divider/>
        <v-container>
          <v-row>

          </v-row>
        </v-container>
      </v-card>
    </v-dialog>
    <div>
      <v-app-bar
          color="grey lighten-5"
          dense
          flat
          fixed
      >
        <v-toolbar-title>
          <v-icon x-large left color="primary">mdi-linkedin</v-icon>
        </v-toolbar-title>

        <v-spacer></v-spacer>
        <div v-if="isAuthenticated" style="display: flex">
          <div>
            <v-btn icon @click="dialogSearch=true">
              <v-icon>mdi-magnify</v-icon>
            </v-btn>
          </div>
          <div v-for="item in appBarItems" :key="item.path">
            <v-btn icon @click="goto(item.path)">
              <v-icon>{{ item.icon }}</v-icon>
            </v-btn>
          </div>
        </div>
        <div v-if="isAuthenticated">
          <v-btn icon @click="sendLogout()">
            <v-icon>mdi-power</v-icon>
          </v-btn>
        </div>
        <div v-else>
          <v-btn icon @click="goto('login')">
            <v-icon>mdi-login</v-icon>
          </v-btn>
        </div>
      </v-app-bar>
    </div>

    <v-main class="pt-16 grey lighten-3">
      <router-view></router-view>
    </v-main>
  </v-app>
</template>

<script>

import {mapActions, mapGetters} from "vuex";

export default {
  name: 'App',
  data: () => ({
    dialogSearch: false,
    username: '',
    appBarItems: [
      {
        icon: 'mdi-home',
        path: 'home',
      },
      {
        icon: 'mdi-account-multiple-plus',
        path: 'network',
      },
      {
        icon: 'mdi-bell',
        path: 'notifications',
      },
      {
        icon: 'mdi-comment',
        path: 'conversation',
      },
      {
        icon: 'mdi-account',
        path: 'profile'
      }

    ]
  }),
  computed: {
    ...mapGetters("AuthModules", ["isAuthenticated"]),
    ...mapGetters("UserModules", ['users'])
  },
  methods: {
    ...mapActions("AuthModules", ['logout']),
    ...mapActions("UserModules", ['getUsersByUsername']),
    goto(path) {
      if (this.$route.path !== `/${path}`) {
        this.$router.push(`/${path}`)
      }
    },
    sendLogout() {
      this.logout();
      this.goto('login');
    },
    doSearch(){
      this.getUsersByUsername(this.username);
    }
  }
};
</script>
