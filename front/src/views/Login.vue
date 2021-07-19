<template>
  <v-container>
    <v-row justify="center">
      <v-col cols="12" sm="6" md="4">
        <v-card style="border-radius: 7px;" flat>
          <v-card-title style="justify-content: center">
            Sign in to Linkedin
          </v-card-title>
          <v-card-subtitle style="text-align: center">
            See your growth and get consulting support
          </v-card-subtitle>
          <v-divider/>
          <v-card-text>
            <v-form @submit.prevent="sendLogin">
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
              <template v-if="status !== ''">
                <v-alert
                    border="left"
                    dense
                    outlined
                    type="error"
                >{{ this.status }}
                </v-alert>
              </template>
              <v-row class="pa-3">
                <v-col cols="4" class="mt-2" >
                  <a @click="gotoSignup()">Register</a>
                </v-col>
                <v-spacer/>
                <v-col cols="6" style="display: flex">
                  <v-spacer/>
                  <v-btn
                      color="primary"
                      dark
                      type="submit"
                  >
                    Sign in
                  </v-btn>
                </v-col>
              </v-row>
            </v-form>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import {mapActions} from "vuex";

export default {
  name: "Login",
  data: () => ({
    user: {
      username: '',
      password: ''
    },
    status: '',
  }),
  methods: {
    ...mapActions('AuthModules', ['login']),
    sendLogin() {
      this.status = '';
      this.login(this.user).then((r) => {
        if (r) {
          this.$router.replace({name: "Home"})
        } else {
          this.status = "User or Password is invalid"
        }
      });
    },
    gotoSignup() {
      this.$router.push('signup')
    }
  }
}
</script>

<style scoped>
.status {
  color: red;
}
</style>