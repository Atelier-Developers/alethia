import axios from 'axios';
import {API} from "@/api/api";


const state = {
    token: localStorage.getItem("userToken") || '',
};

const mutations = {
    setToken(state, payload) {
        state.token = payload;
    },
};

const actions = {
    async login(context, payload) {
        try {
            let response = await axios.post(API.login, payload);
            context.commit("setToken", response.data.access_token);
            localStorage.setItem("userToken", state.token);
            axios.defaults.headers.common['Authorization'] = "Bearer " + state.token;
            return true;
        } catch (e) {
            console.log(e);
            return false;
        }
    },
    logout(context) {
        axios.defaults.headers.common['Authorization'] = '';
        localStorage.removeItem("userToken");
        context.commit('setToken', "");
    },
};

const getters = {
    isAuthenticated: (state) => {
        return state.token !== ""
    },
};


export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
}
