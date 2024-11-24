import { ActionContext, createStore } from 'vuex';
import { expect, test } from 'vitest';
import API from '@/api';
import * as openapi from '@/openapi/api';
import { UsersState } from '@/store/states';

// eslint-disable-next-line
const UserVuexStore = (initialState: any) => createStore({
  actions: {
    async viewUsers(ctx: ActionContext<UsersState, UsersState>, payload: {page: number,
      num: number}) {
      try {
        const resp = await API.userAPI.getUsers(payload.page, payload.num);

        ctx.commit('saveUsers', { users: resp.data.users, total: resp.data.total });
        // eslint-disable-next-line
      } catch (err: any) {
        await ctx.dispatch('addError', err.response.data.message);
      }
    },

    async deleteUser(ctx: ActionContext<UsersState, UsersState>, login: string) {
      try {
        await API.userAPI.deleteUser(login);

        ctx.commit('deleteUser', login);
        // eslint-disable-next-line
      } catch (err: any) {
        await ctx.dispatch('addError', err.response.data.message);
      }
    },

    clearUsers(ctx: ActionContext<UsersState, UsersState>) {
      ctx.commit('clearUsersMut');
    },

    incPageUsers(ctx: ActionContext<UsersState, UsersState>) {
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
});

test('view users', () => {
  const store = UserVuexStore({ });
  store.dispatch('viewUsers', { page: 1, num: 10 }).then(() => {
    expect(store.state.users).toBe([{
      id: 'e95ab7b2-636e-447f-9f87-04072e4b3b9d',
      login: 'muhomorfus',
      picture: 'a.png',
      description: 'aaaa',
      balance: 0,
      mail: 'a@a.ru',
      isAdmin: false,
    },
    ]);
    expect(store.state.totalUsers).toBe(1);
  });
});

test('delete user', () => {
  const store = UserVuexStore({
    users: [{
      id: 'e95ab7b2-636e-447f-9f87-04072e4b3b9d',
      login: 'muhomorfus',
      picture: 'a.png',
      description: 'aaaa',
      balance: 0,
      mail: 'a@a.ru',
      isAdmin: false,
    },
    ],
    totalUsers: 1,
  });
  store.dispatch('deleteUser', 'muhomorfus').then(() => {
    expect(store.state.users).toBe([]);
    expect(store.state.totalUsers).toBe(1);
  });
});
