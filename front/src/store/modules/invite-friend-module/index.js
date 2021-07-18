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
        console.log("GET FRIENDS");
        try {
            let res = await axios.get(API.getFriends);
            console.log(res);
            context.commit('setFriends', res.data);
        } catch (e) {
            console.log(e);
        }
    },
    async getInvites(context,) {
        console.log("GET Invites");
        try {
            let res = await axios.get(API.getInvites);
            console.log(res);
            context.commit('setInvites', res.data);
        } catch (e) {
            console.log(e);
        }
    },
    async createInvite(context, payload) {
        console.log("Create Invites");
        try {
            let res = await axios.post(API.getInvites, payload);
            console.log(res);
        } catch (e) {
            console.log(e);
        }
    },
    async deleteInvite(context, payload) {
        console.log("Delete Invites");
        console.log(payload)
        try {
            let res = await axios.delete(API.getInvites, {
                data: payload
            });
            console.log(res);
        } catch (e) {
            console.log(e);
        }
    },
    async addFriend(context, payload) {
        console.log("add Friend");
        try {
            let res = await axios.post(API.getFriends, payload);
            console.log(res);
        } catch (e) {
            console.log(e);
        }
    },
    async deleteFriend(context, payload) {
        console.log("Delete Friend");
        try {
            let res = await axios.delete(API.getFriends,  {
                data: payload
            });
            console.log(res);
        } catch (e) {
            console.log(e);
        }
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