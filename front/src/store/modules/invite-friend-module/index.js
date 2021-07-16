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
    getFriends(context,) {
        console.log("GET FRIENDS");
        try {
            let res = axios.get(API.getFriends);
            console.log(res);
            context.commit('setFriends', res.data);
        } catch (e) {
            console.log(e);
        }
    },
    getInvites(context,) {
        console.log("GET Invites");
        try {
            let res = axios.get(API.getInvites);
            console.log(res);
            context.commit('setInvites', res.data);
        } catch (e) {
            console.log(e);
        }
    },
    createInvite(context, payload){
        console.log("Create Invites");
        try {
            let res = axios.post(API.getInvites, payload);
            console.log(res);
        } catch (e) {
            console.log(e);
        }
    },
    deleteInvite(context, payload){
        console.log("Delete Invites");
        try {
            let res = axios.delete(API.getInvites, payload);
            console.log(res);
        } catch (e) {
            console.log(e);
        }
    },
    addFriend(context, payload){
        console.log("Delete Invites");
        try {
            let res = axios.post(API.getFriends, payload);
            console.log(res);
        } catch (e) {
            console.log(e);
        }
    },
    deleteFriend(context, payload){
        console.log("Delete Invites");
        try {
            let res = axios.delete(API.getFriends, payload);
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