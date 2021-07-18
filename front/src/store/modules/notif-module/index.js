import {API} from "@/api/api";
import axios from "axios";


const state = {
    notifs: [],
};

const mutations = {
    setNotifs(state, item) {
        state.notifs = item;
    }
};

const actions = {
    async getNotifs(context,) {
        console.log("GET NOTIFS");
        let res = await axios.get(API.notif);
        console.log(res);
        context.commit('setNotifs', res.data);
    },
    async viewProfile(context, payload){
        console.log("VIEW PROFILE NOTIF");
        let res = await axios.post(API.notif + '/profile', payload)
        console.log(res);
    }
};

const getters = {
    notifs: (state) => state.notifs,
};


export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
}
