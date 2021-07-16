import {API} from '@/api/api'
import axios from "axios";

const state = {
    user: {},
    skills: [],
    backgrounds: [],
    languages: [],
    allLanguages: [],
    allSkills: []
};

const mutations = {
    setUser(state, item) {
        state.user = item;
    },
    setSkills(state, item) {
        state.skills = item;
    },
    setBackgrounds(state, item) {
        state.backgrounds = item;
    },
    setLanguages(state, item) {
        state.languages = item;
    },
    setAllLanguages(state, item) {
        state.allLanguages = item;
    },
    setAllSkills(state, item) {
        state.allSkills = item;
    }
};

const actions = {
    async getUser(context,) {
        console.log("Get User")
        try {
            let res = await axios.get(API.getUser);
            console.log(res)
            context.commit('setUser', res.data);
        } catch (e) {
            console.log(e);
        }
    },
    async getSkills(context,) {
        console.log("Get SKILLS")
        try {
            let res = await axios.get(API.getSkills);
            console.log(res)
            context.commit('setSkills', res.data);
        } catch (e) {
            console.log(e);
        }
    },
    async getLanguages(context,) {
        console.log("Get Langs")
        try {
            let res = await axios.get(API.getLanguages);
            console.log(res)
            context.commit('setLanguages', res.data);
        } catch (e) {
            console.log(e);
        }
    },
    async getBackgrounds(context,) {
        console.log("Get Backgrounds")
        try {
            let res = await axios.get(API.getBackground);
            console.log(res)
            context.commit('setBackgrounds', res.data);
        } catch (e) {
            console.log(e);
        }
    },
    async addLanguage(context, payload) {
        console.log("Add Lang")
        for (const p of payload) {
            try {
                let res = await axios.post(API.addLanguage, {
                    language_id: p
                });
                console.log(res);
            } catch (e) {
                console.log(e);
            }
        }
    },
    async addBackground(context, payload) {
        console.log("Add Back")
        try {
            let res = await axios.post(API.addBackground, payload);
            console.log(res);
        } catch (e) {
            console.log(e);
        }
    },
    async addSkill(context, payload) {
        console.log("Add Skill")
        for (const p of payload) {
            try {
                let res = await axios.post(API.addSkill, {
                    skill_id: p
                });
                console.log(res);
            } catch (e) {
                console.log(e);
            }
        }
    },
    async getAllLanguages(context,) {
        console.log("Get All LAngs")
        try {
            let res = await axios.get(API.getAllLanguages);
            console.log(res)
            context.commit("setAllLanguages", res.data);
        } catch (e) {
            console.log(e);
        }
    },
    async getAllSkills(context,) {
        console.log("Get All Skills")
        try {
            let res = await axios.get(API.getAllSkills);
            console.log(res)
            context.commit("setAllSkills", res.data);
        } catch (e) {
            console.log(e);
        }
    }

};

const getters = {
    user: (state) => state.user,
    skills: (state) => state.skills,
    backgrounds: (state) => state.backgrounds,
    languages: (state) => state.languages,
    allLanguages: (state) => state.allLanguages,
    allSkills: (state) => state.allSkills,
};


export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
}
