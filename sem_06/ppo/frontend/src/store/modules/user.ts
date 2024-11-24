// eslint-disable-next-line max-classes-per-file
import * as openapi from '@/openapi/api';
import API from '@/api';
import { LoginRequest, RegisterRequest } from '@/openapi/api';
import Cookies from 'cookies-ts';
import { ActionContext } from 'vuex';
import {
  UserState, State,
} from '@/store/states';

export default {
  actions: {
    async userInfo(ctx: ActionContext<UserState, State>) {
      try {
        const resp = await API.userAPI.getCurrentUser();

        ctx.commit('saveUser', resp.data.user);
        // eslint-disable-next-line
      } catch (err: any) {
        await ctx.dispatch('addError', err.response.data.message);
      }
    },

    async login(ctx: ActionContext<UserState, State>, payload: {login: string, password: string}) {
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

    async register(ctx: ActionContext<UserState, State>, payload: {login: string, password: string,
      email: string,
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
};
