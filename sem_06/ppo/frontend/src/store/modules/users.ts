// eslint-disable-next-line max-classes-per-file
import * as openapi from '@/openapi/api';
import API from '@/api/index';
import { ActionContext } from 'vuex';
import {
  UsersState, State,
} from '@/store/states';

export default {
  actions: {
    async viewUsers(ctx: ActionContext<UsersState, State>, payload: {page: number, num: number}) {
      try {
        const resp = await API.userAPI.getUsers(payload.page, payload.num);

        ctx.commit('saveUsers', { users: resp.data.users, total: resp.data.total });
        // eslint-disable-next-line
      } catch (err: any) {
        await ctx.dispatch('addError', err.response.data.message);
      }
    },

    async deleteUser(ctx: ActionContext<UsersState, State>, login: string) {
      try {
        await API.userAPI.deleteUser(login);

        ctx.commit('deleteUser', login);
        // eslint-disable-next-line
      } catch (err: any) {
        await ctx.dispatch('addError', err.response.data.message);
      }
    },

    clearUsers(ctx: ActionContext<UsersState, State>) {
      ctx.commit('clearUsersMut');
    },

    incPageUsers(ctx: ActionContext<UsersState, State>) {
      ctx.commit('incUsers');
    },
  },

  mutations: {
    clearUsersMut(state: UsersState) {
      state.users = [];
      state.page = 1;
      state.totalUsers = 0;
    },

    saveUsers(state: UsersState, payload: {users: Array<openapi.User>, total: number}) {
      state.users.push(...payload.users);
      state.totalUsers = payload.total;
    },

    deleteUser(state: UsersState, login: string) {
      state.users = state.users.filter((user: openapi.User) => (user.login !== login));
    },

    incUsers(state: UsersState) {
      state.page += 1;
    },
  },

  state: (): UsersState => ({
    users: Array<openapi.User>(),
    totalUsers: 0,
    page: 1,
    num: 10,
  }),

  getters: {
    totalUsers(state: UsersState) {
      return state.totalUsers;
    },

    users(state: UsersState) {
      return state.users;
    },

    getPageUsers(state: UsersState) {
      return state.page;
    },

    getNumUsers(state: UsersState) {
      return state.num;
    },
  },
};
