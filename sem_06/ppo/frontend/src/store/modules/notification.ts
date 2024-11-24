import { ActionContext } from 'vuex';
import {
  Notification, State,
} from '@/store/states';

export default {
  actions: {
    async addError(ctx: ActionContext<Notification, State>, error: string) {
      ctx.commit('addError', error);
    },

    async deleteError(ctx: ActionContext<Notification, State>, error: string) {
      ctx.commit('deleteError', error);
    },
  },

  mutations: {
    addError(state: Notification, error: string) {
      state.errors.unshift(error);
      console.error(state.errors);
    },

    deleteError(state: Notification, errorDel: string) {
      state.errors = state.errors.filter((error: string) => (error !== errorDel));
    },
  },

  state: (): Notification => ({
    errors: Array<string>(),
  }),

  getters: {
    errors(state: Notification) {
      return state.errors;
    },
  },
};
