<template>
  <v-card flat outlined>
    <v-card-title>
      {{ comment.commenter_username }}:
    </v-card-title>
    <v-card-subtitle>
      {{ isoToDate(comment.created) }}
    </v-card-subtitle>
    <v-divider/>
    <v-container>
      <v-row class="my-3">
        <v-col cols="12">
          <div class="post-content">
            {{ comment.text }}
          </div>
        </v-col>
      </v-row>
    </v-container>
  </v-card>
</template>

<script>
export default {
  name: "LittleComment",
  props: ['comment'],
  methods:{
    isoToDate(iso) {
      let date = new Date(iso);
      let year = date.getFullYear();
      let month = date.getMonth() + 1;
      let dt = date.getDate();
      let h = date.getHours();
      let m = date.getMinutes();

      if (m < 10)
        m = '0' + m
      if (h < 10)
        h = '0' + h

      if (dt < 10) {
        dt = '0' + dt;
      }
      if (month < 10) {
        month = '0' + month;
      }

      return year + '-' + month + '-' + dt + ' ' + h + ":" + m;
    }
  }
}
</script>

<style scoped>

</style>