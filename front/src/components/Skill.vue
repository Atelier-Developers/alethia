<template>
  <v-container>
    <v-row>
      <v-col cols="8">
        {{skill.title}}
      </v-col>
      <v-col cols="4">
<!--        {{skill.n_endorse}}-->
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "Skill",
  props: ['skill']
}
</script>

<style scoped>

</style>