<template>
  <div>
    <!--        BD-->
    <v-card v-if="notif.type === 'BD'" flat>
      <v-card-title class="notif-header">
        {{ notif.creator.username }}'s Birthday!
      </v-card-title>
      <v-card-subtitle>
        {{isoToDate(notif.created)}}
      </v-card-subtitle>
      <v-divider/>
      <v-container>
        <v-row>
          <v-col cols="12">
            <div class="notif-content">
              today is {{ notif.creator.username }}'s Birthday!!
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-card>

    <!--        VP-->
    <v-card v-else-if="notif.type === 'VP'" flat>
      <v-card-title class="notif-header">
        {{ notif.creator.username }}'s viewed ur profile!
      </v-card-title>
      <v-card-subtitle>
        {{isoToDate(notif.created)}}
      </v-card-subtitle>
      <v-divider/>
      <v-container>
        <v-row>
          <v-col cols="12">
            <div class="notif-content">
              {{ notif.creator.username }}'s viewed ur profile at {{ isoToDate(notif.created) }}
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-card>

    <!--        LP-->
    <v-card v-else-if="notif.type === 'LP'" flat>
      <v-card-title class="notif-header">
        {{ notif.creator.username }} has Liked ur post!
      </v-card-title>
      <v-card-subtitle>
        {{isoToDate(notif.created)}}
      </v-card-subtitle>
      <v-divider/>
      <v-container>
        <v-row>
          <v-col cols="12">
            <div class="notif-content">
              <v-card flat outlined>
                <v-card-title class="notif-header">
                  <v-row>
                    <v-col cols="12">
                      {{ notif.post.poster_username }}
                    </v-col>
                  </v-row>
                </v-card-title>
                <v-card-subtitle>
                  <v-row>
                    <v-col cols="12">
                      {{ isoToDate(notif.post.created) }}
                    </v-col>
                  </v-row>
                </v-card-subtitle>
                <v-divider/>
                <v-container>
                  <v-row>
                    <v-col cols="12">
                      <div class="post-content">
                        {{ notif.post.description }}
                      </div>
                    </v-col>
                  </v-row>
                </v-container>
              </v-card>
            </div>
            <div class="my-3">
              has been liked by {{ notif.creator.username }}
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-card>

    <!--        CM-->
    <v-card v-else-if="notif.type === 'CM'" flat>
      <v-card-title class="notif-header">
        {{ notif.creator.username }} has commented under ur post!
      </v-card-title>
      <v-card-subtitle>
        {{ isoToDate(notif.created) }}
      </v-card-subtitle>
      <v-divider/>
      <v-container>
        <v-row>
          <v-col cols="12">
            <div class="notif-content">
              <v-card flat outlined>
                <v-card-title class="notif-header">
                  <v-row>
                    <v-col cols="12">
                      {{ notif.post.poster_username }}
                    </v-col>
                  </v-row>
                </v-card-title>
                <v-card-subtitle>
                  <v-row>
                    <v-col cols="12">
                      {{ isoToDate(notif.post.created) }}
                    </v-col>
                  </v-row>
                </v-card-subtitle>
                <v-divider/>
                <v-container>
                  <v-row>
                    <v-col cols="12">
                      <div class="post-content">
                        {{ notif.post.description }}
                      </div>
                    </v-col>
                  </v-row>
                </v-container>
              </v-card>
            </div>
            <div class="my-3">
              has been commented by {{ notif.creator.username }}:
            </div>
            <div class="notif-content">
              <LittleComment :comment="notif.comment" :dense="true"/>
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-card>

    <!--        LCM-->
    <v-card v-else-if="notif.type === 'LCM'" flat>
      <v-card-title class="notif-header">
        {{ notif.creator.username }} has Liked ur comment!
      </v-card-title>
      <v-card-subtitle>
        {{isoToDate(notif.created)}}
      </v-card-subtitle>
      <v-divider/>

      <v-container>
        <v-row>
          <v-col cols="12">
            <div class="notif-content">
              <LittleComment :dense="true" :comment="notif.comment"/>
            </div>
            <div class="my-3">
              has been liked by {{ notif.creator.username }}
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-card>

    <!--        RCM-->
    <v-card v-else-if="notif.type === 'RCM'" flat>
      <v-card-title class="notif-header">
        {{ notif.creator.username }} has replied ur comment!
      </v-card-title>
      <v-card-subtitle>
        {{isoToDate(notif.created)}}
      </v-card-subtitle>
      <v-divider/>
      <v-container>
        <v-row>
          <v-col cols="12">
            <div class="notif-content">
              <LittleComment :dense="true" :comment="notif.replied_comment"/>
            </div>
            <div class="my-3">
              has been replied by {{ notif.creator.username }}:
            </div>
            <div class="notif-content">
              <LittleComment :dense="true" :comment="notif.comment"/>
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-card>

    <!--        END-->
    <v-card v-else-if="notif.type === 'END'" flat>
      <v-card-title class="notif-header">
        {{ notif.creator.username }} has endorsed ur skill!
      </v-card-title>
      <v-card-subtitle>
        {{isoToDate(notif.created)}}
      </v-card-subtitle>
      <v-divider/>
      <v-container>
        <v-row>
          <v-col cols="12">
            <div class="notif-content">
              ur {{ notif.user_skill.title }} has been
              endorsed by {{ notif.creator.username }}
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-card>

    <!--        FCW-->
    <v-card v-else-if="notif.type === 'FCW'" flat>
      <v-card-title class="notif-header">
        {{ notif.creator.username }} has changed his company!
      </v-card-title>
      <v-card-subtitle>
        {{isoToDate(notif.created)}}
      </v-card-subtitle>
      <v-divider/>
      <v-container>
        <v-row>
          <v-col cols="12">
            <div class="notif-content">
              {{ notif.creator.username }} has changed his company
              to {{ notif.background_history.location }}
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-card>
  </div>
</template>

<script>
import LittleComment from "./LittleComment";

export default {
  name: "Notif",
  components: {LittleComment},
  props: ['notif'],
  methods: {
    isoToDate(iso) {
      let date = new Date(iso);
      let year = date.getFullYear();
      let month = date.getMonth() + 1;
      let dt = date.getDate();
      let h = date.getHours();
      let m = date.getMinutes();

      if (m < 10)
        m = '0' + m
      if (h < 10)
        h = '0' + h

      if (dt < 10) {
        dt = '0' + dt;
      }
      if (month < 10) {
        month = '0' + month;
      }

      return year + '-' + month + '-' + dt + ' ' + h + ":" + m;
    }
  }
}
</script>

<style scoped>

</style>