import { ActionContext, createStore } from 'vuex';
import { expect, test } from 'vitest';
import API from '@/api';
import * as openapi from '@/openapi/api';
import { ProfileState } from '@/store/states';

// eslint-disable-next-line
const ProfileVuexStore = (initialState: any) => createStore({
  actions: {
    async viewProfile(ctx: ActionContext<ProfileState, ProfileState>, login: string) {
      try {
        const resp = await API.userAPI.getUser(login);

        ctx.commit('saveProfile', {
          profile: resp.data.user,
          subscribed: resp.data.subscribed,
          self: resp.data.self,
        });
        // eslint-disable-next-line
      } catch (err: any) {
        await ctx.dispatch('addError', err);
      }
    },

    async subscribe(ctx: ActionContext<ProfileState, ProfileState>) {
      try {
        const resp = await API.subscriberAPI.subscribe(ctx.state.profile.id);

        ctx.commit('subscribe', resp.data.subscribed);

        await ctx.dispatch('viewProfile', ctx.state.profile.login);
        // eslint-disable-next-line
      } catch (err: any) {
        await ctx.dispatch('addError', err);
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
    ...initialState,
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
});

test('view profile', () => {
  const store = ProfileVuexStore({ });
  store.dispatch('viewProfile', 'muhomorfus').then(() => {
    expect(store.state.profile).toBe([{
      id: "'e95ab7b2-636e-447f-9f87-04072e4b3b9d",
      login: 'muhomorfus',
      picture: 'string',
      description: 'string',
      balance: 0,
      mail: 'string',
      isAdmin: false,
    },
    ]);
    expect(store.state.subscribed).toBe(false);
    expect(store.state.self).toBe(false);
  });
});

test('subscribe', () => {
  const store = ProfileVuexStore({ profile: { id: 'e95ab7b2-636e-447f-9f87-04072e4b3b9d', login: 'muhomorfus' } });
  store.dispatch('subscribe').then(() => {
    expect(store.state.subscribed).toBe(true);
  });
});
