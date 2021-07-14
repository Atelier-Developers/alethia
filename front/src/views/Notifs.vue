<template>
  <v-card
      class="notifs"
      elevation="0"
  >
    <v-card-title>
    <v-row>
      <v-col cols="12">
        <h1>Notifications</h1>
      </v-col>
    </v-row>
    </v-card-title>
    <v-row v-for="notif in notifs" :key="notif.id">
      <v-col cols="12">
        <Notif :notif="notif"/>
      </v-col>
    </v-row>
  </v-card>
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
.notifs{
  max-height: 500px;
  overflow-y: scroll;
  overflow-x: hidden;
  padding: 15px;
}
.notif-header {
  font-size: 16px;
}

.notif-content {
  font-size: 13px;
  margin-top: 10px;
  margin-bottom: 10px;
}
</style>