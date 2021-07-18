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
        let res = await axios.get(API.getUser);
        context.commit('setUser', res.data);
    },
    async getUserById(context, payload) {
        let res = await axios.post(API.getUser, payload);
        context.commit('setUser', res.data);
    },
    async getSkills(context, payload) {
        let res = await axios.get(API.getSkills + '/' + payload);
        context.commit('setSkills', res.data);
    },
    async getLanguages(context, payload) {
        let res = await axios.get(API.getLanguages + '/' + payload);
        context.commit('setLanguages', res.data);
    },
    async getBackgrounds(context, payload) {
        const res = await axios.get(API.getBackground + '/' + payload);
        context.commit('setBackgrounds', res.data);
    },
    async updateUserLanguage(context, payload) {
        const mappedLanguage = context.getters.languages.map((x) => x.id)

        for (const p of payload) {
            if (!mappedLanguage.includes(p)) {
                await axios.post(API.userLanguage, {
                    language_id: p
                });
            }
        }

        for (const p of mappedLanguage) {
            if (!payload.includes(p)) {
                await axios.delete(`${API.userLanguage}/${p}`, {
                    language_id: p
                });
            }
        }
    },
    async addBackground(context, payload) {
        await axios.post(API.userBackground, payload);
    },
    async deleteBackground(context, id) {
        await axios.delete(`${API.userBackground}/${id}`);
    },
    async updateUserSkill(context, payload) {
        const mappedSkills = context.getters.skills.map((x) => x.id)
        for (const p of payload) {
            if (!mappedSkills.includes(p)) {
                await axios.post(API.userSkill, {
                    skill_id: p
                });
            }
        }
        for (const p of mappedSkills) {
            if (!payload.includes(p)) {
                await axios.delete(`${API.userSkill}/${p}`, {
                    skill_id: p
                });
            }
        }
    },
    async getAllLanguages(context,) {
        const res = await axios.get(API.getAllLanguages);
        context.commit("setAllLanguages", res.data);
    },
    async getAllSkills(context,) {
        let res = await axios.get(API.getAllSkills);
        console.log(res)
        context.commit("setAllSkills", res.data);
    },
    async editUser(context, payload) {
        await axios.put(API.editUser, payload);
    },

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
