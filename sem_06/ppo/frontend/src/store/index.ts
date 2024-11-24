import Vuex from 'vuex';

import profile from '@/store/modules/profile';
import posts from '@/store/modules/posts';
import users from '@/store/modules/users';
import user from '@/store/modules/user';
import notification from '@/store/modules/notification';
import comments from '@/store/modules/comments';
import {
  State,
} from './states';

export default new Vuex.Store<State>({
  modules: {
    posts,
    profile,
    users,
    user,
    notification,
    comments,
  },
});
