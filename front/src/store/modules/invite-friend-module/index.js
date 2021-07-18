import {API} from '@/api/api'
import axios from "axios";

const state = {
    invites: [],
    friends: [],
};

const mutations = {
    setFriends(state, item) {
        state.friends = item;
    },
    setInvites(state, item) {
        state.invites = item;
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
};

const getters = {
    invites: (state) => state.invites,
    friends: (state) => state.friends,
};


export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
}