<template>
  <v-card>
    <v-card-title>
      {{ back.title }} at {{ back.location }}
      <v-spacer/>
      <v-btn icon @click="deleteBackground">
        <v-icon>
          mdi-delete
        </v-icon>
      </v-btn>
    </v-card-title>
    <v-card-subtitle>
      From {{ isoToDate(back.start_date) }} to {{ isoToDate(back.end_date) }}
    </v-card-subtitle>
    <v-card-text>
      {{ back.description }}
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  name: "Background",
  props: ['back'],
  methods: {
    async deleteBackground() {

    },
    isoToDate(iso) {
      let date = new Date(iso);
      let year = date.getFullYear();
      let month = date.getMonth() + 1;
      let dt = date.getDate();

      if (dt < 10) {
        dt = '0' + dt;
      }
      if (month < 10) {
        month = '0' + month;
      }

      return year + '-' + month + '-' + dt
    }
  }
}
</script>

<style scoped>

</style>