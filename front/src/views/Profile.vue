<template>
  <div>
    <v-dialog v-model="dialogLang">
      <v-card>
        <v-card-title>
          <h3>Add Language</h3>
        </v-card-title>
        <v-divider/>
        <v-container>
          <v-row>
            <v-col cols="2" v-for="lang in allLanguages" :key="lang.id">
              <v-checkbox
                  v-model="selectedLangs"
                  :label="lang.language"
                  :value="lang.id"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-btn @click="sendLangs()">
                Add
              </v-btn>
            </v-col>
          </v-row>
        </v-container>
      </v-card>
    </v-dialog>
    <v-dialog v-model="dialogSkill">
      <v-card>
        <v-card-title>
          <h3>Add Skill</h3>
        </v-card-title>
        <v-divider/>
        <v-container>
          <v-row>
            <v-col cols="2" v-for="skill in allSkills" :key="skill.id">
              <v-checkbox
                  v-model="selectedSkills"
                  :label="skill.title"
                  :value="skill.id"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-btn @click="sendSkills()">
                Add
              </v-btn>
            </v-col>
          </v-row>
        </v-container>
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
    <v-container>
      <v-row>
        <v-col cols="12">
          <v-card>
            <v-card-title>
              <h2>
                {{ user.username }}
              </h2>
            </v-card-title>
            <v-divider/>
            <v-container>
              <v-row>
                <v-col cols="12">
                  <h3>info</h3>
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12">
                  <p>
                    {{ user.info }}
                  </p>
                </v-col>
              </v-row>
            </v-container>
          </v-card>
        </v-col>
        <v-col cols="12">
          <v-card>
            <v-card-title>
              <h2>
                Skills
              </h2>
              <v-spacer/>
              <v-icon color="black" @click="dialogSkill=true">mdi-plus</v-icon>
            </v-card-title>
            <v-divider/>
            <v-container>
              <v-row>
                <v-col cols="6" v-for="skill in skills" :key="skill.id">
                  <Skill :skill="skill"/>
                </v-col>
              </v-row>
            </v-container>
          </v-card>
        </v-col>
        <v-col cols="12">
          <v-card>
            <v-card-title>
              <h2>
                Languages
              </h2>
              <v-spacer/>
              <v-icon color="black" @click="dialogLang=true">mdi-plus</v-icon>
            </v-card-title>
            <v-divider/>
            <v-container>
              <v-row>
                <v-col cols="12">
                  <v-list>
                    <v-list-item v-for="lang in languages" :key="lang.id">
                      {{ lang.language }}
                    </v-list-item>
                  </v-list>
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
    selectedLangs: [],
    selectedSkills: [],
    newBack: {
      start_date: '',
      end_date: '',
      title: '',
      description: '',
      category: '',
    },
    menu1: false,
    menu2: false,

  }),
  computed: {
    ...mapGetters("UserModules", ["user",
      "skills",
      "languages",
      "backgrounds",
      "allLanguages",
      "allSkills"])
  },
  methods: {
    ...mapActions('UserModules', ["getUser",
      "getBackgrounds",
      "getSkills",
      "getLanguages",
      "addLanguage",
      "addSkill",
      "addBackground",
      "getAllLanguages",
      'getAllSkills']),
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
    }
  },
  mounted() {
    this.getUser();
    this.getBackgrounds()
    this.getSkills().then(() => {
      this.selectedSkills = this.skills.map((x) => x.id)
    })
    this.getLanguages().then(() => {
      this.selectedLangs = this.languages.map((x) => x.id)
    })
    this.getAllLanguages()
    this.getAllSkills()
  }
}
</script>

<style scoped>

</style>