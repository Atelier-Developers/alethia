<template>
  <div>
    <!--        BD-->
    <v-card v-if="notif.type === 'BD'">
      <v-card-title class="notif-header">
        {{ notif.user.username }}'s Birthday!
      </v-card-title>
      <v-divider/>
      <v-container>
        <v-row>
          <v-col cols="12">
            <div class="notif-content">
              today is {{ notif.user.username }}'s Birthday!!
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-card>

    <!--        VP-->
    <v-card v-else-if="notif.type === 'VP'">
      <v-card-title class="notif-header">
        {{ notif.user.username }}'s viewed ur profile!
      </v-card-title>
      <v-divider/>
      <v-container>
        <v-row>
          <v-col cols="12">
            <div class="notif-content">
              {{ notif.user.username }}'s viewed ur profile at {{ notif.date }}
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-card>

    <!--        LP-->
    <v-card v-else-if="notif.type === 'LP'">
      <v-card-title class="notif-header">
        {{ notif.user.username }} has Liked ur post!
      </v-card-title>
      <v-divider/>
      <v-container>
        <v-row>
          <v-col cols="12">
            <div class="notif-content">
              <v-card>
                <v-card-title class="notif-header">
                  <v-row>
                    <v-col cols="12">
                      {{ notif.post.user.username }}
                    </v-col>
                  </v-row>
                </v-card-title>
                <v-card-subtitle>
                  <v-row>
                    <v-col cols="12">
                      {{ notif.post.date }}
                    </v-col>
                  </v-row>
                </v-card-subtitle>
                <v-divider/>
                <v-container>
                  <v-row>
                    <v-col cols="12">
                      <div class="post-content">
                        {{ notif.post.text }}
                      </div>
                    </v-col>
                  </v-row>
                </v-container>
              </v-card>
            </div>
            has been liked by {{ notif.user.username }}
          </v-col>
        </v-row>
      </v-container>
    </v-card>

    <!--        CM-->
    <v-card v-else-if="notif.type === 'CM'">
      <v-card-title class="notif-header">
        {{ notif.user.username }} has commented under ur post!
      </v-card-title>
      <v-divider/>
      <v-container>
        <v-row>
          <v-col cols="12">
            <div class="notif-content">
              <v-card>
                <v-card-title class="notif-header">
                  <v-row>
                    <v-col cols="12">
                      {{ notif.post.user.username }}
                    </v-col>
                  </v-row>
                </v-card-title>
                <v-card-subtitle>
                  <v-row>
                    <v-col cols="12">
                      {{ notif.post.date }}
                    </v-col>
                  </v-row>
                </v-card-subtitle>
                <v-divider/>
                <v-container>
                  <v-row>
                    <v-col cols="12">
                      <div class="post-content">
                        {{ notif.post.text }}
                      </div>
                    </v-col>
                  </v-row>
                </v-container>
              </v-card>
            </div>
            has been commented by {{ notif.user.username }}:
            <div class="notif-content">
              <comment :comment="notif.comment"/>
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-card>

    <!--        LCM-->
    <v-card v-else-if="notif.type === 'LCM'">
      <v-card-title class="notif-header">
        {{ notif.user.username }} has Liked ur comment!
      </v-card-title>
      <v-divider/>
      <v-container>
        <v-row>
          <v-col cols="12">
            <div class="notif-content">
              <comment :comment="notif.comment"/>
            </div>
            has been liked by {{ notif.user.username }}:
          </v-col>
        </v-row>
      </v-container>
    </v-card>

    <!--        RCM-->
    <v-card v-else-if="notif.type === 'RCM'">
      <v-card-title class="notif-header">
        {{ notif.user.username }} has replied ur comment!
      </v-card-title>
      <v-divider/>
      <v-container>
        <v-row>
          <v-col cols="12">
            <div class="notif-content">
              <comment :comment="notif.comment"/>
            </div>
            has been replied by {{ notif.user.username }}:
            <div class="notif-content">
              <comment :comment="notif.reply"/>
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-card>

    <!--        END-->
    <v-card v-else-if="notif.type === 'END'">
      <v-card-title class="notif-header">
        {{ notif.user.username }} has endorsed ur skill!
      </v-card-title>
      <v-divider/>
      <v-container>
        <v-row>
          <v-col cols="12">
            <div class="notif-content">
              ur {{ notif.skill.text }} has been
              endorsed by {{ notif.user.username }}:
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-card>

    <!--        FCW-->
    <v-card v-else-if="notif.type === 'FCW'">
      <v-card-title class="notif-header">
        {{ notif.user.username }} has changed his company!
      </v-card-title>
      <v-divider/>
      <v-container>
        <v-row>
          <v-col cols="12">
            <div class="notif-content">
              {{ notif.user.username }} has changed his company
              to {{ notif.new_work }}
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-card>
  </div>
</template>

<script>
import Comment from "./Comment";
export default {
  name: "Notif",
  components: {Comment},
  props: ['notif']
}
</script>

<style scoped>

</style>