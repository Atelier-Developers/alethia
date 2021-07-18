<template>
  <div>
    <v-dialog v-model="dialogLang" max-width="750px">
      <v-card flat style="border-radius: 7px;">
        <v-card-title>
          <v-icon left>
            mdi-translate
          </v-icon>
          <h4>
            Languages
          </h4>
        </v-card-title>
        <v-divider/>
        <v-card-text>
          <v-form @submit.prevent="sendLangs">
            <v-row dense>
              <v-col cols="12" sm="4" v-for="lang in allLanguages" :key="lang.id">
                <v-checkbox
                    v-model="selectedLangs"
                    :label="lang.language"
                    :value="lang.id"
                />
              </v-col>
              <v-col cols="12">
                <v-row class="px-3 py-2">
                  <v-spacer/>
                  <v-btn
                      color="primary"
                      dark
                      type="submit"
                  >
                    Add
                  </v-btn>
                </v-row>
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>
      </v-card>
    </v-dialog>
    <v-dialog v-model="dialogSkill" max-width="750px">
      <v-card flat style="border-radius: 7px;">
        <v-card-title>
          <v-icon left>
            mdi-swim
          </v-icon>
          <h4>
            Skills
          </h4>
        </v-card-title>
        <v-divider/>
        <v-card-text>
          <v-form @submit.prevent="sendSkills">
            <v-row dense>
              <v-col cols="12" sm="4" v-for="skill in allSkills" :key="skill.id">
                <v-checkbox
                    v-model="selectedSkills"
                    :label="skill.title"
                    :value="skill.id"
                />
              </v-col>
              <v-col cols="12">
                <v-row class="px-3 py-2">
                  <v-spacer/>
                  <v-btn
                      color="primary"
                      dark
                      type="submit">
                    Add
                  </v-btn>
                </v-row>
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>
      </v-card>
    </v-dialog>
    <v-dialog v-model="dialogBack" max-width="750px">
      <v-card flat style="border-radius: 7px;">
        <v-card-title>
          <v-icon left>
            mdi-note-plus-outline
          </v-icon>
          <h4>
            Background
          </h4>
        </v-card-title>
        <v-card-text>
          <v-form
              ref="back-form"
              @submit.prevent="sendBackground">
            <v-row dense>
              <v-col cols="12" md="6">
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
                        v-model="newBack.start_date"
                        label="Start Date"
                        outlined
                        readonly
                        prepend-icon="mdi-calendar"
                        v-bind="attrs"
                        v-on="on"
                    ></v-text-field>
                  </template>
                  <v-date-picker
                      v-model="newBack.start_date"
                      @input="menu1 = false"
                  ></v-date-picker>
                </v-menu>

              </v-col>
              <v-col cols="12" md="6">
                <v-menu
                    ref="menu2"
                    v-model="menu2"
                    :close-on-content-click="false"
                    transition="scale-transition"
                    offset-y
                    max-width="290px"
                    min-width="auto"
                >
                  <template v-slot:activator="{ on, attrs }">
                    <v-text-field
                        v-model="newBack.end_date"
                        label="End Date"
                        outlined
                        prepend-icon="mdi-calendar"
                        readonly
                        v-bind="attrs"
                        v-on="on"
                    ></v-text-field>
                  </template>
                  <v-date-picker
                      v-model="newBack.end_date"
                      @input="menu2 = false"
                  ></v-date-picker>
                </v-menu>
              </v-col>
              <v-col cols="12">
                <v-select
                    label="Category"
                    outlined
                    :items="['work_experience', 'xx']"
                    v-model="newBack.category"
                />
              </v-col>
              <v-col cols="12">
                <v-text-field
                    label="Location"
                    outlined
                    v-model="newBack.location"
                />
              </v-col>
              <v-col cols="12">
                <v-text-field
                    outlined
                    label="Description"
                    v-model="newBack.description"
                />
              </v-col>
              <v-col cols="12">
                <v-text-field
                    label="Title"
                    outlined
                    v-model="newBack.title"
                />
              </v-col>
              <v-col cols="12">
                <v-row class="px-3 py-2">
                  <v-spacer/>
                  <v-btn
                      color="primary"
                      dark
                      type="submit">
                    Add
                  </v-btn>
                </v-row>
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>
      </v-card>
    </v-dialog>
    <v-dialog v-model="dialogUser" max-width="750px">
      <v-card flat style="border-radius: 7px;">
        <v-card-title>
          <h4>Edit Info</h4>
        </v-card-title>
        <v-divider/>
        <v-card-text>
          <v-form class="mt-5" @submit.prevent="sendEditUser">
            <v-row dense>
              <v-col cols="6">
                <v-menu
                    ref="menu3"
                    v-model="menu3"
                    :close-on-content-click="false"
                    transition="scale-transition"
                    offset-y
                    max-width="290px"
                    min-width="auto"
                >
                  <template v-slot:activator="{ on, attrs }">
                    <v-text-field
                        v-model="newUser.birth_date"
                        label="Start Date"
                        outlined
                        readonly
                        prepend-icon="mdi-calendar"
                        v-bind="attrs"
                        v-on="on"
                    ></v-text-field>
                  </template>
                  <v-date-picker
                      v-model="newUser.birth_date"
                      @input="menu3 = false"
                  ></v-date-picker>
                </v-menu>
              </v-col>
              <v-col cols="12">
                <v-text-field
                    v-model="newUser.intro"
                    outlined
                    label="Intro"
                />
              </v-col>
              <v-col cols="12">
                <v-text-field
                    outlined
                    v-model="newUser.about"
                    label="About"
                />
              </v-col>
              <v-col cols="12">
                <v-text-field
                    v-model="newUser.accomplishments"
                    label="Accomplishments"
                    outlined
                />
              </v-col>
              <v-col cols="12">
                <v-text-field
                    v-model="newUser.additional_info"
                    label="Additional Info"
                    outlined
                />
              </v-col>
              <v-col cols="12">
                <v-row>
                  <v-spacer/>
                  <v-btn type="submit" color="primary">
                    Edit
                  </v-btn>
                </v-row>
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>
      </v-card>
    </v-dialog>
    <v-container>
      <v-row>
        <v-col cols="12">
          <v-card flat style="border-radius: 7px">
            <v-card-title>
              <h2 class="pr-2">
                {{ user.first_name }} {{ user.last_name }}
              </h2>
              <h4 class="font-weight-light">
                @{{ user.username }}
              </h4>
              <v-spacer/>
              <v-btn @click="dialogUser=true" icon>
                <v-icon>mdi-pencil</v-icon>
              </v-btn>
            </v-card-title>
            <v-card-subtitle>
              {{ user.about }}
            </v-card-subtitle>
            <v-divider/>
            <v-card-text>
              {{ user.intro }}
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="12">
          <v-card flat style="border-radius: 7px">
            <v-card-title>
              <v-icon left>
                mdi-trophy-outline
              </v-icon>
              <h4>
                Accomplishments
              </h4>
            </v-card-title>
            <v-divider/>
            <v-card-text>
              {{ user.accomplishments }}
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="12">
          <v-card flat style="border-radius: 7px">
            <v-card-title>
              <v-icon left>
                mdi-information-outline
              </v-icon>
              <h4>
                Additional Info
              </h4>
            </v-card-title>
            <v-divider/>
            <v-card-text>
              {{ user.additional_info }}
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="12">
          <v-card flat style="border-radius: 7px">
            <v-card-title>
              <v-icon left>
                mdi-swim
              </v-icon>
              <h4>
                Skills
              </h4>
              <v-spacer/>
              <v-btn @click="dialogSkill=true" elevation="0" icon>
                <v-icon>mdi-plus</v-icon>
              </v-btn>
            </v-card-title>
            <v-divider/>
            <v-container>
              <v-row dense>
                <v-col cols="6" v-for="skill in skills" :key="skill.id">
                  <Skill :skill="skill"/>
                </v-col>
              </v-row>
            </v-container>
          </v-card>
        </v-col>
        <v-col cols="12">
          <v-card flat style="border-radius: 7px">
            <v-card-title>
              <v-icon left>
                mdi-translate
              </v-icon>
              <h4>
                Languages
              </h4>
              <v-spacer/>
              <v-btn @click="dialogLang=true" elevation="0" icon>
                <v-icon>mdi-plus</v-icon>
              </v-btn>
            </v-card-title>
            <v-divider/>
            <v-container>
              <v-row dense>
                <v-col cols="6" v-for="lang in languages" :key="lang.id">
                  <v-list-item>
                    <v-list-item-content>
                      <v-list-item-title>
                        {{ lang.language }}
                      </v-list-item-title>
                    </v-list-item-content>
                  </v-list-item>
                </v-col>
              </v-row>
            </v-container>
          </v-card>
        </v-col>
        <v-col cols="12">
          <v-card flat style="border-radius: 7px">
            <v-card-title>
              <h2>
                Backgrounds
              </h2>
              <v-spacer/>
              <v-icon @click="dialogBack=true">mdi-plus</v-icon>
            </v-card-title>
            <v-divider/>
            <v-container>
              <v-row>
                <v-col cols="12" md="6" v-for="back in backgrounds" :key="back.id">
                  <Background @delete="sendDeleteBackground" :deletable="true" :back="back"/>
                </v-col>
              </v-row>
            </v-container>
          </v-card>
        </v-col>
        <v-col cols="12">
          <v-card color="grey lighten-2">
            <v-card-title style="background-color: white">
              <h2>
                Featured Posts
              </h2>
            </v-card-title>
            <v-divider/>
            <v-container>
              <v-row v-for="post in user.posts" :key="post.id">
                <v-col cols="12">
                  <Post :post="post" :username="user.username"/>
                </v-col>
              </v-row>
            </v-container>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import Skill from "../components/Skill";
import {mapActions, mapGetters} from "vuex";
import Post from "../components/Post";
import Background from "../components/Background";

export default {
  name: "Profile",
  components: {Background, Post, Skill},
  data: () => ({
    dialogLang: false,
    dialogSkill: false,
    dialogBack: false,
    dialogUser: false,
    selectedLangs: [],
    selectedSkills: [],
    newBack: {
      start_date: '',
      end_date: '',
      title: '',
      description: '',
      category: '',
    },
    newUser: {},
    menu1: false,
    menu2: false,
    menu3: false,

  }),
  computed: {
    ...mapGetters("UserModules",
        ["user",
          "skills",
          "languages",
          "backgrounds",
          "allLanguages",
          "allSkills"])
  },
  methods: {
    ...mapActions('UserModules',
        ["getUser",
          "getBackgrounds",
          "getSkills",
          "getLanguages",
          "updateUserLanguage",
          "updateUserSkill",
          "addBackground",
          "deleteBackground",
          "getAllLanguages",
          'getAllSkills',
          "editUser"]),
    async sendLangs() {
      await this.updateUserLanguage(this.selectedLangs)
      await this.getLanguages(this.user.user_id)
      this.selectedLangs = this.languages.map((x) => x.id)
      this.dialogLang = false
    },
    async sendSkills() {
      await this.updateUserSkill(this.selectedSkills)
      await this.getSkills(this.user.user_id)
      this.selectedSkills = this.skills.map((x) => x.id)
      this.dialogSkill = false
    },
    async sendBackground() {
      const payload = {...this.newBack}
      payload.start_date = (new Date(payload.start_date)).toISOString()
      if (payload.end_date) {
        payload.end_date = (new Date(payload.end_date)).toISOString()
      } else {
        delete payload.end_date
      }
      await this.addBackground(payload)
      await this.getBackgrounds(this.user.user_id)
      this.$refs["back-form"].reset()
      this.dialogBack = false
    },
    async sendDeleteBackground(id) {
      await this.deleteBackground(id)
      await this.getBackgrounds(this.user.user_id)
    },
    parseDate(date) {
      return date;
    },
    sendEditUser() {
      this.editUser({
        about: this.newUser.about,
        accomplishments: this.newUser.accomplishments,
        additional_info: this.newUser.additional_info,
        birth_date: this.newUser.birth_date,
        intro: this.newUser.intro,
      });
      this.dialogUser = false;
    }
  },
  async mounted() {
    await this.getUser()
    this.newUser = this.user;
    await this.getBackgrounds(this.user.user_id)
    await this.getSkills(this.user.user_id)
    this.selectedSkills = this.skills.map((x) => x.id)
    await this.getLanguages(this.user.user_id)
    this.selectedLangs = this.languages.map((x) => x.id)
    await this.getAllLanguages()
    await this.getAllSkills()
  }
}
</script>

<style scoped>

</style>