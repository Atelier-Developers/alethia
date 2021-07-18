import {API} from '@/api/api'
import axios from "axios";

const state = {
    invites: [],
    friends: [],
    mutual: [],
};

const mutations = {
    setFriends(state, item) {
        state.friends = item;
    },
    setInvites(state, item) {
        state.invites = item;
    },
    setMutual(state, item) {
        state.mutual = item;
    }

};

const actions = {
    async getFriends(context,) {
        console.log("GET FRIENDS");
        let res = await axios.get(API.getFriends);
        console.log(res);
        context.commit('setFriends', res.data);
    },
    async getInvites(context,) {
        console.log("GET Invites");
        let res = await axios.get(API.getInvites);
        console.log(res);
        context.commit('setInvites', res.data);
    },
    async createInvite(context, payload) {
        console.log("Create Invites");
        let res = await axios.post(API.getInvites, payload);
        console.log(res);
    },
    async deleteInvite(context, payload) {
        console.log("Delete Invites");
        let res = await axios.delete(API.getInvites, {
            data: payload
        });
        console.log(res);
    },
    async addFriend(context, payload) {
        console.log("add Friend");
        let res = await axios.post(API.getFriends, payload);
        console.log(res);
    },
    async deleteFriend(context, payload) {
        console.log("Delete Friend");
        let res = await axios.delete(API.getFriends, {
            data: payload
        });
        console.log(res);
    },
    async getMutualConnections(context) {
        console.log("get mutual")
        let res = await axios.get(API.mutual)
        console.log(res);
        context.commit('setMutual', res.data);
    }
};

const getters = {
    invites: (state) => state.invites,
    friends: (state) => state.friends,
    mutual: (state) => state.mutual,
};


export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
}