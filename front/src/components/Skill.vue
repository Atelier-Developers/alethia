<template>
  <v-list-item>
    <v-list-item-content>
      <v-list-item-title> {{ skill.title }}</v-list-item-title>
    </v-list-item-content>
    <v-list-item-action>
      <div>
        <v-btn icon @click="likeSkill()">
          <v-icon v-bind:color="skill.is_endorsed_by_this_user? 'red': 'grey lighten-1'">mdi-heart</v-icon>
        </v-btn>
        {{skill.endorse_count}}
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
        this.skill.is_endorsed_by_this_user = !this.skill.liked;
        this.skill.endorse_count += 1;
      })
    }
  }
}
</script>

<style scoped>

</style>