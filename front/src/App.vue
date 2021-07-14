<template>
  <v-app>
    <div>
      <v-app-bar
          color="deep-purple accent-4"
          dense
          dark
          fixed
      >
        <v-app-bar-nav-icon></v-app-bar-nav-icon>

        <v-toolbar-title>Page title</v-toolbar-title>

        <v-spacer></v-spacer>

        <v-btn icon>
          <v-icon>mdi-heart</v-icon>
        </v-btn>

        <v-btn icon>
          <v-icon>mdi-magnify</v-icon>
        </v-btn>

        <v-menu
            left
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
            <v-list-item
                v-for="n in 5"
                :key="n"
                @click="() => {}"
            >
              <v-list-item-title>Option {{ n }}</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
      </v-app-bar>
    </div>

    <v-main class="mt-10">
      <v-container>
        <v-row >
          <v-col cols="3"></v-col>
          <v-col cols="6">
            <Home/>
          </v-col>
          <v-col cols="3">
            <Notifs style="position: fixed;"/>
          </v-col>
        </v-row>
      </v-container>
    </v-main>

  </v-app>
</template>

<script>
import Home from "@/views/Home";
import Notifs from "@/views/Notifs";

export default {
  name: 'App',

  components: {
    Notifs,
    Home
  },

  data: () => ({
    //
  }),
};
</script>
