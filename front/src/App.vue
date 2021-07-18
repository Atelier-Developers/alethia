<template>
  <v-app>
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
        path: 'conversations',
      },
      {
        icon: 'mdi-account',
        path: 'profile'
      }

    ]
  }),
  computed: {
    ...mapGetters("AuthModules", ["isAuthenticated"]),
  },
  methods: {
    ...mapActions("AuthModules", ['logout']),
    goto(path) {
      this.$router.push(path)
    },
    sendLogout() {
      this.logout();
      this.goto('login');
    }
  }
};
</script>
