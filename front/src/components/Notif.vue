<template>
  <div>
    <!--        BD-->
    <v-card v-if="notif.type === 'BD'">
      <v-card-title class="notif-header">
        {{ notif.creator.username }}'s Birthday!
      </v-card-title>
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
    <v-card v-else-if="notif.type === 'VP'">
      <v-card-title class="notif-header">
        {{ notif.creator.username }}'s viewed ur profile!
      </v-card-title>
      <v-divider/>
      <v-container>
        <v-row>
          <v-col cols="12">
            <div class="notif-content">
              {{ notif.creator.username }}'s viewed ur profile at {{ notif.date }}
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-card>

    <!--        LP-->
    <v-card v-else-if="notif.type === 'LP'">
      <v-card-title class="notif-header">
        {{ notif.creator.username }} has Liked ur post!
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
    <v-card v-else-if="notif.type === 'CM'">
      <v-card-title class="notif-header">
        {{ notif.creator.username }} has commented under ur post!
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
                      {{ notif.post.poster_username }}
                    </v-col>
                  </v-row>
                </v-card-title>
                <v-card-subtitle>
                  <v-row>
                    <v-col cols="12">
                      {{ notif.isoToDate(notif.post.created) }}
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
              <Comment :comment="notif.comment" :dense="true"/>
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-card>

    <!--        LCM-->
    <v-card v-else-if="notif.type === 'LCM'">
      <v-card-title class="notif-header">
        {{ notif.creator.username }} has Liked ur comment!
      </v-card-title>
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
    <v-card v-else-if="notif.type === 'RCM'">
      <v-card-title class="notif-header">
        {{ notif.creator.username }} has replied ur comment!
      </v-card-title>
      <v-divider/>
      <v-container>
        <v-row>
          <v-col cols="12">
            <div class="notif-content">
              <LittleComment :dense="true" :comment="notif.comment"/>
            </div>
            <div class="my-3">
              has been replied by {{ notif.creator.username }}:
            </div>
            <div class="notif-content">
              <Comment :dense="true" :comment="notif.replied_comment"/>
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-card>

    <!--        END-->
    <v-card v-else-if="notif.type === 'END'">
      <v-card-title class="notif-header">
        {{ notif.creator.username }} has endorsed ur skill!
      </v-card-title>
      <v-divider/>
      <v-container>
        <v-row>
          <v-col cols="12">
            <div class="notif-content">
              ur {{ notif.skill.text }} has been
              endorsed by {{ notif.creator.username }}:
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-card>

    <!--        FCW-->
    <v-card v-else-if="notif.type === 'FCW'">
      <v-card-title class="notif-header">
        {{ notif.creator.username }} has changed his company!
      </v-card-title>
      <v-divider/>
      <v-container>
        <v-row>
          <v-col cols="12">
            <div class="notif-content">
              {{ notif.creator.username }} has changed his company
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
import LittleComment from "./LittleComment";

export default {
  name: "Notif",
  components: {LittleComment, Comment},
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