<template>
  <v-card>
    <v-card-title>
      <v-row>
        <v-col cols="12">
          {{ post.user.username }}
        </v-col>
      </v-row>
    </v-card-title>
    <v-card-subtitle>
      <v-row>
        <v-col cols="12">
          {{ post.date }}
        </v-col>
      </v-row>
    </v-card-subtitle>
    <v-divider/>
    <v-container>
      <v-row>
        <v-col cols="12">
          <div class="post-content">
            {{ post.text }}
          </div>
        </v-col>
      </v-row>
    </v-container>
  </v-card>
</template>

<script>
export default {
  name: "Post",
  props:['post']
}
</script>

<style scoped>

</style>