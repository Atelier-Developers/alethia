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
                    @click="sendLangs()">
                  Add
                </v-btn>
              </v-row>
            </v-col>
          </v-row>
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
                    @click="sendSkills()">
                  Add
                </v-btn>
              </v-row>
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>
    </v-dialog>
    <v-dialog v-model="dialogBack">
      <v-card>
        <v-card-title>
          <h3>Add Background</h3>
        </v-card-title>
        <v-divider/>
        <v-container>
          <v-row>
            <v-col cols="6">
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
                      hint="MM/DD/YYYY format"
                      persistent-hint
                      prepend-icon="mdi-calendar"
                      v-bind="attrs"
                      @blur="date = parseDate(newBack.start_date)"
                      v-on="on"
                  ></v-text-field>
                </template>
                <v-date-picker
                    v-model="newBack.start_date"
                    no-title
                    @input="menu1 = false"
                ></v-date-picker>
              </v-menu>

            </v-col>
            <v-col cols="6">
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
                      hint="MM/DD/YYYY format"
                      persistent-hint
                      prepend-icon="mdi-calendar"
                      v-bind="attrs"
                      @blur="date = parseDate(newBack.end_date)"
                      v-on="on"
                  ></v-text-field>
                </template>
                <v-date-picker
                    v-model="newBack.end_date"
                    no-title
                    @input="menu2 = false"
                ></v-date-picker>
              </v-menu>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-select
                  label="Category"
                  :items="['work_experience', 'xx']"
                  v-model="newBack.category"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-text-field
                  label="Location"
                  v-model="newBack.location"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-text-field
                  label="Description"
                  v-model="newBack.description"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-text-field
                  label="Title"
                  v-model="newBack.title"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-btn @click="sendBackground()">
                Add
              </v-btn>
            </v-col>
          </v-row>
        </v-container>
      </v-card>
    </v-dialog>
    <v-dialog v-model="dialogUser">
      <v-card>
        <v-card-title>
          <h3>Edit Info</h3>
        </v-card-title>
        <v-divider/>
        <v-container>
          <v-row>
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
                      hint="MM/DD/YYYY format"
                      persistent-hint
                      prepend-icon="mdi-calendar"
                      v-bind="attrs"
                      @blur="date = parseDate(newUser.birth_date)"
                      v-on="on"
                  ></v-text-field>
                </template>
                <v-date-picker
                    v-model="newUser.birth_date"
                    no-title
                    @input="menu3 = false"
                ></v-date-picker>
              </v-menu>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-text-field
                  v-model="newUser.intro"
                  label="Intro"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-text-field
                  v-model="newUser.about"
                  label="About"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-text-field
                  v-model="newUser.accomplishments"
                  label="Accomplishments"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-text-field
                  v-model="newUser.additional_info"
                  label="Additional Info"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-btn @click="sendEditUser()">
                Change
              </v-btn>
            </v-col>
          </v-row>
        </v-container>
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
              <v-btn @click="dialogUser=true" elevation="0" icon>
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
        <!--        TODO remove 409 conflict when adding new skill-->
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
          <v-card>
            <v-card-title>
              <h2>
                Backgrounds
              </h2>
              <v-spacer/>
              <v-icon color="black" @click="dialogBack=true">mdi-plus</v-icon>
            </v-card-title>
            <v-divider/>
            <v-container>
              <v-row>
                <v-col cols="6" v-for="back in backgrounds" :key="back.id">
                  <Background :back="back"/>
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
          "addLanguage",
          "addSkill",
          "addBackground",
          "getAllLanguages",
          'getAllSkills',
          "editUser"]),
    sendLangs() {
      console.log(this.selectedLangs);
      this.addLanguage(this.selectedLangs)
      this.dialogLang = false
    },
    sendSkills() {
      console.log(this.selectedSkills);
      this.addSkill(this.selectedSkills)
      this.dialogSkill = false
    },
    sendBackground() {
      this.newBack.start_date = (new Date(this.newBack.start_date)).toISOString()
      this.newBack.end_date = (new Date(this.newBack.end_date)).toISOString()
      this.addBackground(this.newBack)
      this.dialogBack = false
    },
    parseDate(date) {
      let x = new Date(date)
      console.log(x.toUTCString())
      console.log(x.toISOString())
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
  mounted() {
    this.getUser().then(() => {
      this.newUser = this.user;
      this.getBackgrounds(this.user.user_id)
      this.getSkills(this.user.user_id).then(() => {
        this.selectedSkills = this.skills.map((x) => x.id)
      })
      this.getLanguages(this.user.user_id).then(() => {
        this.selectedLangs = this.languages.map((x) => x.id)
      })
    })

    this.getAllLanguages()
    this.getAllSkills()
  }
}
</script>

<style scoped>

</style>