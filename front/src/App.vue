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
          <v-row align="center" dense>
            <v-col cols="12" style="display: flex">
              <v-checkbox
                  v-model="searchOpt.language.use"
                  hide-details
                  class="shrink mr-2 mb-10"
              ></v-checkbox>
              <v-text-field
                  :disabled="!searchOpt.language.use"
                  label="Language"
                  outlined
                  v-model="searchOpt.language.text"
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row align="center" dense>
            <v-col cols="12" style="display: flex">
              <v-checkbox
                  v-model="searchOpt.location.use"
                  hide-details
                  class="shrink mr-2 mb-10"
              ></v-checkbox>
              <v-text-field
                  :disabled="!searchOpt.location.use"
                  label="Location"
                  outlined
                  v-model="searchOpt.location.text"
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row align="center" dense>
            <v-col cols="12" style="display: flex">
              <v-checkbox
                  v-model="searchOpt.company.use"
                  hide-details
                  class="shrink mr-2 mb-10"
              ></v-checkbox>
              <v-text-field
                  :disabled="!searchOpt.company.use"
                  label="Company"
                  outlined
                  v-model="searchOpt.company.text"
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row dense>
            <v-spacer/>
            <v-col cols="4">
              <v-btn @click="doSearch()" color="primary">
                Search
                <v-icon>mdi-magnify</v-icon>
              </v-btn>
            </v-col>
          </v-row>
        </v-container>
        <v-divider v-if="users.length > 0"/>
        <v-container>
          <v-row
              v-for="user in users"
              :key="user.id"
              @click="goto('user/'+user.user_id);dialogSearch=false;"
              style="border-bottom: 1px solid rgba(0, 0, 0, 0.12); cursor: pointer"
          >
            <v-col cols="4">
              {{ user.username }}
            </v-col>
            <v-col cols="4" v-if="user.is_friends_with_this_user === 1">
              Your Friend
            </v-col>
            <v-col cols="4" v-else>
              Connections: {{ user.mutual_connection }}
            </v-col>
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
            <v-btn icon @click="dialogSearch=true; clearUsers()">
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

    ],
    searchOpt: {
      location: {
        use: false,
        text: '',
      },
      language: {
        use: false,
        text: '',
      },
      company: {
        use: false,
        text: '',
      },
    }
  }),
  computed: {
    ...mapGetters("AuthModules", ["isAuthenticated"]),
    ...mapGetters("UserModules", ['users'])
  },
  methods: {
    ...mapActions("AuthModules", ['logout']),
    ...mapActions("UserModules", ['getUsersByUsername', 'clearUsers']),
    goto(path) {
      if (this.$route.path !== `/${path}`) {
        this.$router.push(`/${path}`)
      }
    },
    sendLogout() {
      this.logout();
      this.goto('login');
    },
    doSearch() {
      this.getUsersByUsername({
        username: this.username,
        location: this.searchOpt.location.use ? this.searchOpt.location.text : '',
        company: this.searchOpt.company.use ? this.searchOpt.company.text : '',
        language: this.searchOpt.language.use ? this.searchOpt.language.text : '',
      });
    }
  }
};
</script>
