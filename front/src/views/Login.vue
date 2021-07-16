<template>
  <v-container>
    <v-row justify="center">
      <v-col cols="4">
        <v-card>
          <v-card-title>
            Login
          </v-card-title>
          <v-container>
            <v-row>
              <v-col cols="12">
                <v-text-field
                    label="username"
                    v-model="user.username"
                />
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="12">
                <v-text-field
                    label="password"
                    type="password"
                    v-model="user.password"
                />
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="3">
                <v-btn
                    @click="sendLogin()"
                >
                  LOG IN
                </v-btn>
              </v-col>
            </v-row>
            <v-row v-if="status !== ''">
              <v-col cols="12">
                  <v-alert
                      border="left"
                      dense
                      outlined
                      type="error"
                  >{{this.status}}</v-alert>
              </v-col>
            </v-row>
          </v-container>
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
    }
  }
}
</script>

<style scoped>
.status {
  color: red;
}
</style>