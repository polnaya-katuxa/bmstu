// eslint-disable-next-line max-classes-per-file
import * as openapi from '@/openapi/api';
import API from '@/api';
import { ActionContext } from 'vuex';
import {
  ProfileState, State,
} from '@/store/states';

export default {
  actions: {
    async viewProfile(ctx: ActionContext<ProfileState, State>, login: string) {
      try {
        const resp = await API.userAPI.getUser(login);

        ctx.commit('saveProfile', {
          profile: resp.data.user,
          subscribed: resp.data.subscribed,
          self: resp.data.self,
        });
        // eslint-disable-next-line
      } catch (err: any) {
        await ctx.dispatch('addError', err.response.data.message);
      }
    },

    async subscribe(ctx: ActionContext<ProfileState, State>) {
      try {
        const resp = await API.subscriberAPI.subscribe(ctx.state.profile.id);

        ctx.commit('subscribe', resp.data.subscribed);

        await ctx.dispatch('viewProfile', ctx.state.profile.login);
        // eslint-disable-next-line
      } catch (err: any) {
        await ctx.dispatch('addError', err.response.data.message);
      }
    },
  },

  mutations: {
    saveProfile(state: ProfileState, payload: { profile: openapi.User,
      subscribed: boolean, self: boolean }) {
      state.profile = payload.profile;
      state.subscribed = payload.subscribed;
      state.self = payload.self;
    },

    subscribe(state: ProfileState, subscribed: boolean) {
      state.subscribed = subscribed;
    },
  },

  state: (): ProfileState => ({
    profile: {} as openapi.User,
    subscribed: false,
    self: false,
  }),

  getters: {
    profile(state: ProfileState) {
      return state.profile;
    },

    subscribed(state: ProfileState) {
      return state.subscribed;
    },

    self(state: ProfileState) {
      return state.self;
    },
  },
};
