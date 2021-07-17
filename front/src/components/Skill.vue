<template>
  <v-list-item>
    <v-list-item-content>
      <v-list-item-title> {{ skill.title }}</v-list-item-title>
    </v-list-item-content>
  </v-list-item>
</template>

<script>
export default {
  name: "Skill",
  props: ['skill']
}
</script>

<style scoped>

</style>