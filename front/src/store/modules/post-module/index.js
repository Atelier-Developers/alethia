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
        console.log("GET POST");
        try {
            let res = await axios.get(API.getPost + "/" + payload);
            console.log(res);
            context.commit("setPost", res.data);
        } catch (e) {
            console.log(e);
        }
    },
    async getLikes(context, payload) {
        console.log("GET Likes");
        try {
            let res = await axios.get(API.getPost + "/" + payload + '/likes');
            console.log(res);
            context.commit("setLikes", res.data);
        } catch (e) {
            console.log(e);
        }
    },
    async getComments(context, payload) {
        console.log("GET POST");
        try {
            let res = await axios.get(API.getPost + "/" + payload + '/comments');
            console.log(res);
            context.commit("setComments", res.data);
        } catch (e) {
            console.log(e);
        }
    },
    async likePost(context, payload) {
        console.log("Like POST");
        try {
            let res = await axios.post(API.likePost, payload);
            console.log(res);
        } catch (e) {
            console.log(e);
        }
    },
    async commentPost(context, payload) {
        console.log("Comment POST");
        try {
            let res = await axios.post(API.commentPost, payload);
            console.log(res);
        } catch (e) {
            console.log(e);
        }
    },
    async replyComment(context, payload) {
        console.log("Comment POST");
        try {
            let res = await axios.post(API.replyComment, payload);
            console.log(res);
        } catch (e) {
            console.log(e);
        }
    },
    async likeComment(context, payload) {
        console.log("Comment POST");
        try {
            let res = await axios.post(API.likeComment, payload);
            console.log(res);
        } catch (e) {
            console.log(e);
        }
    },
    async addRepost(context, payload) {
        console.log("Add Repost");
        try {
            let res = await axios.post(API.repost, payload);
            console.log(res);
        } catch (e) {
            console.log(e);
        }
    },
    async getReposts(context, payload) {
        console.log("Get Repost");
        try {
            let res = await axios.get(API.getPost + "/" + payload + '/reposts');
            console.log(res);
            context.commit('setReposts', res.data)
        } catch (e) {
            console.log(e);
        }
    },
    async getRepost(context, payload) {
        console.log("Get Repost");
        if (+payload===0)
            return;
        try {
            let res = await axios.get(API.getPost + "/" + payload);
            console.log(res);
            context.commit('setRepost', res.data)
        } catch (e) {
            console.log(e);
        }
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