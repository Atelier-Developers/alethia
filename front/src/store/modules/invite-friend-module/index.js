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
        let res = await axios.get(API.getFriends);
        context.commit('setFriends', res.data);
    },
    async getInvites(context,) {
        const res = await axios.get(API.getInvites);
        context.commit('setInvites', res.data);
    },
    async createInvite(context, payload) {
        await axios.post(API.getInvites, payload);
    },
    async deleteInvite(context, payload) {
        await axios.delete(API.getInvites, {
            data: payload
        });
    },
    async addFriend(context, payload) {
        await axios.post(API.getFriends, payload);
    },
    async deleteFriend(context, payload) {
        await axios.delete(API.getFriends, {
            data: payload
        });
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