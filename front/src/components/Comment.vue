<template>
  <v-card>
    <v-card-title>
      {{ comment.username }}:
    </v-card-title>
    <v-divider/>
    <v-container>
      <v-row>
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
  name: "Comment",
  props: ["comment"]
}
</script>

<style scoped>

</style>