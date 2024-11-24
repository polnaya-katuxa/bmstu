// eslint-disable-next-line max-classes-per-file
import { ActionContext, createStore } from 'vuex';
import { expect, test } from 'vitest';
import API from '@/api';
import * as openapi from '@/openapi/api';
import { LoginRequest, RegisterRequest } from '@/openapi/api';
import Cookies from 'cookies-ts';
import { UserState } from '@/store/states';

// eslint-disable-next-line
const UserVuexStore = (initialState: any) => createStore({
  actions: {
    async userInfo(ctx: ActionContext<UserState, UserState>) {
      try {
        const resp = await API.userAPI.getCurrentUser();

        ctx.commit('saveUser', resp.data.user);
        // eslint-disable-next-line
      } catch (err: any) {
        await ctx.dispatch('addError', err.response.data.message);
      }
    },

    async login(ctx: ActionContext<UserState, UserState>, payload: {login: string,
      password: string}) {
      try {
        const resp = await API.userAPI.login(
          new class implements LoginRequest {
            login = payload.login;

            password = payload.password;
          }(),
        );

        const cookies = new Cookies();
        cookies.set('user-token', resp.data.token);
        window.location.replace(process.env.BASE_URL || '');
        // eslint-disable-next-line
      } catch (err: any) {
        await ctx.dispatch('addError', err.response.data.message);
      }
    },

    async register(ctx: ActionContext<UserState, UserState>, payload: {login: string,
      password: string, email: string,
      picture: string, description: string}) {
      try {
        const resp = await API.userAPI.register(
          new class implements RegisterRequest {
            login = payload.login;

            password = payload.password;

            mail = payload.email;

            picture = payload.picture;

            description = payload.description;
          }(),
        );

        const cookies = new Cookies();
        cookies.set('user-token', resp.data.token);
        window.location.replace(process.env.BASE_URL || '');
        // eslint-disable-next-line
      } catch (err: any) {
        await ctx.dispatch('addError', err.response.data.message);
      }
    },
  },

  mutations: {
    saveUser(state: UserState, user: openapi.User) {
      state.user = user;
    },
  },

  state: (): UserState => ({
    user: {} as openapi.User,
  }),

  getters: {
    user(state: UserState) {
      return state.user;
    },

    isAdmin(state: UserState) {
      return state.user.isAdmin;
    },

    curLogin(state: UserState) {
      return state.user.login;
    },
  },
});

test('user info', () => {
  const store = UserVuexStore({ });
  store.dispatch('userInfo').then(() => {
    expect(store.state.user).toBe([{
      id: 'e95ab7b2-636e-447f-9f87-04072e4b3b9d',
      login: 'muhomorfus',
      picture: 'string',
      description: 'string',
      balance: 0,
      mail: 'string',
      isAdmin: false,
    },
    ]);
  });
});

test('login', () => {
  const store = UserVuexStore({ });
  store.dispatch('login', { login: 'muhomorfus', password: '12345678' }).then(() => {
    expect(store.state.user).toBe({});
    const cookies = new Cookies();
    expect(cookies.get('user-token')).toBe('e95ab7b2-636e-447f-9f87-04072e4b3b9d');
  });
});

test('register', () => {
  const store = UserVuexStore({ });
  store.dispatch('register', {
    login: 'muhomorfus',
    password: '12345678',
    email: 'string',
    picture: 'string',
    description: 'string',
  }).then(() => {
    expect(store.state.user).toBe({});
    const cookies = new Cookies();
    expect(cookies.get('user-token')).toBe('e95ab7b2-636e-447f-9f87-04072e4b3b9d');
  });
});
