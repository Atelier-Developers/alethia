<template>
  <v-list-item>
    <v-list-item-content>
      <v-list-item-title> {{ skill.title }}</v-list-item-title>
    </v-list-item-content>
    <v-list-item-action>
      <div>
        <v-btn icon @click="likeSkill()">
          <v-icon v-bind:color="skill.liked? 'red': 'grey lighten-1'">mdi-heart</v-icon>
        </v-btn>
      </div>
    </v-list-item-action>
  </v-list-item>
</template>

<script>
import {mapActions} from "vuex";

export default {
  name: "Skill",
  props: ['skill', 'user_id'],
  methods: {
    ...mapActions("UserModules", ["endorseSkill"]),
    likeSkill() {
      this.endorseSkill({
        skill_id: this.skill.id,
        user_id: this.user_id
      }).then(() => {
        this.skill.liked = !this.skill.liked
      })
    }
  }
}
</script>

<style scoped>

</style>