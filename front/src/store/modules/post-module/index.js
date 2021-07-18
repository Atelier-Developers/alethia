import {API} from '@/api/api'
import axios from "axios";

const state = {
    post: {},
    repost: {},
    comments: [],
    likes: [],
    reposts: {},
    commentLikes: []
};

const mutations = {
    setPost(state, item) {
        state.post = item;
    },
    setLikes(state, item) {
        state.likes = item;
    },
    setComments(state, item) {
        state.comments = item;
    },
    setReposts(state, item) {
        state.reposts = item
    },
    setCommentLikes(state, item) {
        state.commentLikes = item;
    },
    setRepost(state, item) {
        state.repost = item
    }
};

const actions = {
    async getPost(context, payload) {
        let res = await axios.get(API.getPost + "/" + payload);
        console.log(res);
        context.commit("setPost", res.data);
    },
    async getLikes(context, payload) {
        let res = await axios.get(API.getPost + "/" + payload + '/likes');
        context.commit("setLikes", res.data);
    },
    async getComments(context, payload) {
        let res = await axios.get(API.getPost + "/" + payload + '/comments');
        context.commit("setComments", res.data);
    },
    async likePost(context, payload) {
        await axios.post(API.likePost, payload);
    },
    async commentPost(context, payload) {
        await axios.post(API.commentPost, payload);
    },
    async replyComment(context, payload) {
        await axios.post(API.replyComment, payload);
    },
    async likeComment(context, payload) {
        await axios.post(API.likeComment, payload);
    },
    async addRepost(context, payload) {
        await axios.post(API.repost, payload);
    },
    async getReposts(context, payload) {
        let res = await axios.get(API.getPost + "/" + payload + '/reposts');
        context.commit('setReposts', res.data)
    },
    async getRepost(context, payload) {
        if (+payload === 0)
            return;
        let res = await axios.get(API.getPost + "/" + payload);
        context.commit('setRepost', res.data)
    }
};

const getters = {
    post: (state) => state.post,
    likes: (state) => state.likes,
    comments: (state) => state.comments,
    reposts: (state) => state.reposts,
    repost: (state) => state.repost,
    commentLikes: (state) => state.commentLikes,
};


export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
}