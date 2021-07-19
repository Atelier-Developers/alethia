<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <v-container>
          <v-row>
            <v-col cols="12">
                  <h1>Notifications</h1>
            </v-col>
          </v-row>
          <v-row v-for="notif in notifs" :key="([notif.id, notif.type]).toString()">
            <v-col cols="12">
              <Notif :notif="notif"/>
            </v-col>
          </v-row>
        </v-container>
      </v-col>
    </v-row>
  </v-container>

</template>

<script>
import {mapActions, mapGetters} from "vuex";
import Notif from "../components/Notif";

export default {
  name: "Notifs",
  components: {Notif},
  computed: {
    ...mapGetters('NotifModules', ["notifs"])
  },
  methods: {
    ...mapActions('NotifModules', ["getNotifs"])
  },
  mounted() {
    this.getNotifs();
  }
}
</script>

<style scoped>

</style>