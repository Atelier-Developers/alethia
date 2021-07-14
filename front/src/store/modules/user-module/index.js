const state = {
    user: {},
};

const mutations = {
    setUser(state, item) {
        state.user = item;
    }
};

const actions = {
    getUser(context,) {
        let user = {
            username: "ahmad dul kaj",
            info: "hichvaqt tu surakhayi ke motmaen nisti saf hastn foru nakon",
            skills:[
                {
                    name: "surakh shenas",
                    n_endorse: 153,
                },
                {
                    name: "surakh va kon",
                    n_endorse: 587,
                },
                {
                    name: "ahrom",
                    n_endorse: 52817,
                },
            ],
            posts: [
                {
                    text: 'امروز خیلی هوا کیری بود. افتاب به صورت افقی تا ما تحت ادم میرفت. من خیلی شاخم خلاصه.',
                    date: '11/1/1 12:12',
                    // media: ''
                },
                {
                    text: 'امروز خیلی هوا کیری بود. افتاب به صورت افقی تا ما تحت ادم میرفت. من خیلت افقی تا ما تحت ادم میرفت. من خیلت افقی تا ما تحت ادم میرفت. من خیلت افقی تا ما تحت ادم میرفت. من خیل ی شاخم خلاصه.',
                    date: '11/1/1 12:12',
                    // media: ''
                },

            ]
        };
        context.commit('setUser', user);
    },
};

const getters = {
    user: (state) => state.user,
};


export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
}
