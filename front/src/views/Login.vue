<template>
  <v-container>
    <v-row>
      <v-col cols="4">
        <v-card>
          <v-container>
            <v-row>
              <v-col cols="12">
                <v-text-field
                    label="username"
                />
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="12">
                <v-text-field
                    label="password"
                    type="password"
                />
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="3">
                <v-btn
                    @click="login()"
                >
                  LOG IN
                </v-btn>
              </v-col>
            </v-row>
          </v-container>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "Login"
}
</script>

<style scoped>

</style>