<template>
  <v-container>
    <v-row justify="center">
      <v-col cols="12" sm="6" md="4">
        <v-card style="border-radius: 7px;" flat>
          <v-card-title>
            Create new Linkedin account
          </v-card-title>
          <v-card-subtitle>
            Fill out the information below to create your Linkedin account
          </v-card-subtitle>
          <v-divider/>
          <v-card-text>
            <v-form @submit.prevent="sendSignup">
              <v-text-field
                  outlined
                  prepend-icon="mdi-account-circle-outline"
                  label="Username"
                  v-model="user.username"
              />
              <v-text-field
                  label="Password"
                  type="password"
                  prepend-icon="mdi-lock-outline"
                  outlined
                  v-model="user.password"
              />
              <v-text-field
                  outlined
                  prepend-icon="mdi-rename-box"
                  label="First Name"
                  v-model="user.first_name"
              />
              <v-text-field
                  outlined
                  prepend-icon="mdi-rename-box"
                  label="Last Name"
                  v-model="user.last_name"
              />
              <v-text-field
                  outlined
                  prepend-icon="mdi-information-variant"
                  label="About"
                  v-model="user.about"
              />
              <v-text-field
                  outlined
                  prepend-icon="mdi-trophy-outline"
                  label="Accomplishments"
                  v-model="user.accomplishments"
              />
              <v-text-field
                  prepend-icon="mdi-file-outline"
                  outlined
                  label="Intro"
                  v-model="user.intro"
              />
              <v-text-field
                  prepend-icon="mdi-newspaper-variant-outline"
                  outlined
                  label="Additional Info"
                  v-model="user.additional_info"
              />
              <v-menu
                  ref="menu1"
                  v-model="menu1"
                  :close-on-content-click="false"
                  transition="scale-transition"
                  offset-y
                  max-width="290px"
                  min-width="auto"
              >
                <template v-slot:activator="{ on, attrs }">
                  <v-text-field
                      v-model="user.birth_date"
                      label="Date"
                      outlined
                      readonly
                      prepend-icon="mdi-calendar"
                      v-bind="attrs"
                      @blur="date = parseDate(user.birth_date)"
                      v-on="on"
                  ></v-text-field>
                </template>
                <v-date-picker
                    v-model="user.birth_date"
                    color="primary"
                    @input="menu1 = false"
                ></v-date-picker>
              </v-menu>
              <v-row class="pa-3">
                <v-spacer/>
                <v-btn
                    color="primary"
                    dark
                    type="submit"
                >
                  signup
                </v-btn>
              </v-row>
            </v-form>
            <!--TODO Add validator and submit button-->
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import {mapActions} from "vuex";

export default {
  name: "Signup",
  data: () => ({
    user: {
      about: "",
      accomplishments: "",
      additional_info: "",
      birth_date: '',
      first_name: "",
      intro: "",
      last_name: "",
      password: "",
      username: ''
    },
    date: null,
    menu1: false,
  }),
  methods: {
    ...mapActions('SignupModules', ["signup"]),
    parseDate(date) {
      let x = new Date(date)
      console.log(x.toUTCString())
      console.log(x.toISOString())
      return date;
    },
    sendSignup() {
      // TODO FIx this birth_date
      this.user.birth_date = (new Date(this.user.birth_date)).toISOString();
      this.signup(this.user);
      this.$router.replace({name: 'Login'})
    }
  }
}
</script>

<style scoped>

</style>