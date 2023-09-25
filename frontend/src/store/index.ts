import Vuex from 'vuex';

import posts from './modules/posts';
import profile from './modules/profile';
import users from './modules/users';
import user from './modules/user';
import notification from './modules/notification';
import comments from './modules/comments';

// Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    posts,
    profile,
    users,
    user,
    notification,
    comments,
  },
});
