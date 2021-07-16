<template>
  <v-container>
    <v-row justify="center">
      <v-col cols="4">
        <v-card>
          <v-card-title>
            Sign Up
          </v-card-title>
          <v-container>
            <v-form>
              <v-row>
                <v-col cols="12">
                  <v-text-field
                      label="Username"
                      v-model="user.username"
                  />
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12">
                  <v-text-field
                      label="Password"
                      type="password"
                      v-model="user.password"
                  />
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12">
                  <v-text-field
                      label="Fist Name"
                      v-model="user.first_name"
                  />
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12">
                  <v-text-field
                      label="Last Name"
                      v-model="user.last_name"
                  />
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12">
                  <v-text-field
                      label="about"
                      v-model="user.about"
                  />
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12">
                  <v-text-field
                      label="accomplishments"
                      v-model="user.accomplishments"
                  />
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12">
                  <v-text-field
                      label="intro"
                      v-model="user.intro"
                  />
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12">
                  <v-text-field
                      label="additionalInfo"
                      v-model="user.additional_info"
                  />
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12">
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
                          hint="MM/DD/YYYY format"
                          persistent-hint
                          prepend-icon="mdi-calendar"
                          v-bind="attrs"
                          @blur="date = parseDate(user.birth_date)"
                          v-on="on"
                      ></v-text-field>
                    </template>
                    <v-date-picker
                        v-model="user.birth_date"
                        no-title
                        @input="menu1 = false"
                    ></v-date-picker>
                  </v-menu>
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="3">
                  <v-btn
                      @click="sendSignup()"
                  >
                    LOG IN
                  </v-btn>
                </v-col>
              </v-row>
            </v-form>
          </v-container>
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
    sendSignup(){
      console.log("k");
      this.user.birth_date = (new Date(this.user.birth_date)).toISOString();
      this.signup(this.user);
      console.log("k");
    }
  }
}
</script>

<style scoped>

</style>