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
    async viewProfile(context, payload) {
        console.log("VIEW PROFILE NOTIF");
        let res = await axios.post(API.notif + '/profile', payload)
        console.log(res);
    }
};

const getters = {
    notifs: (state) => {
        if (state.notifs.length === 0)
            return [];
        let all = [];
        if (state.notifs.birthday !== null) {
            state.notifs.birthday.forEach((x) => x.type = 'BD')
            all.push(...state.notifs.birthday)
        }
        if (state.notifs.change_work !== null) {
            state.notifs.change_work.forEach((x) => x.type = 'FCW')
            all.push(...state.notifs.change_work)
        }
        if (state.notifs.comment !== null) {
            state.notifs.comment.forEach((x) => x.type = 'CM')
            all.push(...state.notifs.comment)
        }
        if (state.notifs.endorse !== null) {
            state.notifs.endorse.forEach((x) => x.type = 'END')
            all.push(...state.notifs.endorse)
        }
        if (state.notifs.like_comment !== null) {
            state.notifs.like_comment.forEach((x) => x.type = 'LCM')
            all.push(...state.notifs.like_comment)
        }
        if (state.notifs.like_post !== null) {
            state.notifs.like_post.forEach((x) => x.type = 'LP')
            all.push(...state.notifs.like_post)
        }
        if (state.notifs.reply_comment !== null) {
            state.notifs.reply_comment.forEach((x) => x.type = 'RCM')
            all.push(...state.notifs.reply_comment)
        }
        if (state.notifs.view_profile !== null) {
            state.notifs.view_profile.forEach((x) => x.type = 'VP')
            all.push(...state.notifs.view_profile)
        }
        all.sort(function (a, b) {
            return (a.created < b.created) ? 1 : ((a.created > b.created) ? -1 : 0);
        });
        return all;
    },
};


export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
}
