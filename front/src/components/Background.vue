<template>
  <v-card>
    <v-card-title>
      {{ back.category }}
    </v-card-title>
    <v-divider/>
    <v-list>
      <v-list-item>
        <span class="font-weight-bold mr-5">Start Date:</span>{{ isoToDate(back.start_date) }}
      </v-list-item>
      <v-list-item>
        <span class="font-weight-bold mr-5">End Date:</span> {{ isoToDate(back.end_date) }}
      </v-list-item>
      <v-list-item>
        <span class="font-weight-bold mr-5">Title: </span>{{ back.title }}
      </v-list-item>
      <v-list-item>
        <span class="font-weight-bold mr-5">Location: </span>{{ back.location }}
      </v-list-item>
      <v-list-item>
        <span class="font-weight-bold mr-5">Description: </span>{{ back.description }}
      </v-list-item>
    </v-list>
  </v-card>
</template>

<script>
export default {
  name: "Background",
  props: ['back'],
  methods: {
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