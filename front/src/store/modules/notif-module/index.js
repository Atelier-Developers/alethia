const state = {
    notifs: [],
};

const mutations = {
    setNotifs(state, item) {
        state.notifs = item;
    }
};

const actions = {
    getNotifs(context,) {
        let notifs = [
            {
                type: "BD",
                date: '12/12/12 11:11',
                user: {
                    username: 'haj zahra',
                }
            },
            {
                type: "VP",
                date: '12/12/12 11:11',
                user: {
                    username: 'ayatollah maryam',
                }
            },
            {
                type: "LP",
                date: '12/12/12 11:11',
                user: {
                    username: 'haj zahra',
                },
                post: {
                    user: {
                        username: "Ali",
                    },
                    text: 'امروز خیلی هوا کیری بود. افتاب به صورت افقی تا ما تحت ادم میرفت. من خیلی شاخم خلاصه.',
                    date: '11/1/1 12:12',
                }
            },
            {
                type: "CM",
                date: '12/12/12 11:11',
                user: {
                    username: 'ayatollah maryam',
                },
                post: {
                    user: {
                        username: "Ali",
                    },
                    text: 'امروز خیلی هوا کیری بود. افتاب به صورت افقی تا ما تحت ادم میرفت. من خیلی شاخم خلاصه.',
                    date: '11/1/1 12:12',
                },
                comment:{
                    username: "mamad",
                    text: 'to ba in rezome kirit afqanestanam nemiri guzu',
                },
            },
            {
                type: "LCM",
                date: '12/12/12 11:11',
                user: {
                    username: 'ayatollah maryam',
                },
                comment:{
                    username: "mamad",
                    text: 'to ba in rezome kirit afqanestanam nemiri guzu',
                },
            },
            {
                type: "RCM",
                date: '12/12/12 11:11',
                user: {
                    username: 'ayatollah maryam',
                },
                comment:{
                    username: "mamad",
                    text: 'to ba in rezome kirit afqanestanam nemiri guzu',
                },
                reply:{
                    username: "mamad",
                    text: 'to ba in rezome kirit afqanestanam nemiri guzu',
                },
            },
            {
                type: "END",
                date: '12/12/12 11:11',
                user: {
                    username: 'haj zahra',
                },
                skill:{
                    text: 'suck zadan'
                }
            },
            {
                type: "FCW",
                date: '12/12/12 11:11',
                user: {
                    username: 'haj zahra',
                },
                new_work: "NewYork city's MASJED"
            }
        ];
        context.commit('setNotifs', notifs);
    },
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
