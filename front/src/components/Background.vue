<template>
  <v-card >
    <v-card-title>
      {{ back.title }} at {{ back.location }}
      <v-spacer/>
      <v-btn icon v-if="deletable" @click="deleteBackground" color="red">
        <v-icon>
          mdi-delete
        </v-icon>
      </v-btn>
    </v-card-title>
    <v-card-subtitle>
      From {{ isoToDate(back.start_date) }} to
      <template v-if="back.end_date">
        {{ isoToDate(back.end_date) }}
      </template>
      <template v-else>
        Present
      </template>
    </v-card-subtitle>
    <v-card-text>
      {{ back.description }}
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  name: "Background",
  props: ['back', 'deletable'],
  methods: {
    deleteBackground() {
      this.$emit('delete', this.back.id)
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