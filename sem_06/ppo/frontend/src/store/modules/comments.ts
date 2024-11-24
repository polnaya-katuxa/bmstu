// eslint-disable-next-line max-classes-per-file
import * as openapi from '@/openapi/api';
import API from '@/api/index';
import { Marked } from '@ts-stack/markdown';
import { ActionContext } from 'vuex';
import {
  CommentsState, State,
} from '@/store/states';

export default {
  actions: {
    async getComments(ctx: ActionContext<CommentsState, State>, payload: { id: string, page: number,
      num: number }) {
      try {
        const resp1 = await API.commentAPI.getComments(payload.id, payload.page, payload.num);

        const resp2 = await API.postAPI.getPost(payload.id);

        ctx.commit('savePost', resp2.data.post);

        ctx.commit('saveComments', { comments: resp1.data.comments, total: resp1.data.total });
        // eslint-disable-next-line
      } catch (err: any) {
        await ctx.dispatch('addError', err.response.data.message);
      }
    },

    async comment(ctx: ActionContext<CommentsState, State>, payload: { postID: string,
      content: string }) {
      try {
        const resp = await API.commentAPI.comment(
          payload.postID,
          new class implements openapi.CommentRequest {
            content = payload.content;
          }(),
        );

        ctx.commit('saveComment', resp.data.comment);
        // eslint-disable-next-line
      } catch (err: any) {
        await ctx.dispatch('addError', err.response.data.message);
      }
    },

    async uncomment(ctx: ActionContext<CommentsState, State>, payload: { postID: string,
      commID: string }) {
      try {
        await API.commentAPI.uncomment(
          payload.postID,
          payload.commID,
        );

        ctx.commit('deleteComment', payload.commID);
        // eslint-disable-next-line
      } catch (err: any) {
        await ctx.dispatch('addError', err.response.data.message);
      }
    },

    clearComments(ctx: ActionContext<CommentsState, State>) {
      ctx.commit('clearCommentsMut');
    },

    incPageComments(ctx: ActionContext<CommentsState, State>) {
      ctx.commit('incComments');
    },
  },

  mutations: {
    clearCommentsMut(state: CommentsState) {
      state.totalComments = 0;
      state.post = {} as openapi.Post;
      state.comments = [];
      state.page = 1;
    },

    saveComments(state: CommentsState, payload: {comments: Array<openapi.Comment>, total: number}) {
      state.comments.push(...payload.comments);
      state.totalComments = payload.total;
    },

    saveComment(state: CommentsState, comment: openapi.Comment) {
      state.comments.unshift(comment);
    },

    savePost(state: CommentsState, post: openapi.Post) {
      state.post = post;
    },

    deleteComment(state: CommentsState, commID: string) {
      state.comments = state.comments.filter((comm: openapi.Comment) => (comm.id !== commID));
    },

    incComments(state: CommentsState) {
      state.page += 1;
    },
  },

  state: (): CommentsState => ({
    post: {} as openapi.Post,
    comments: Array<openapi.Comment>(),
    totalComments: 0,
    page: 1,
    num: 10,
  }),

  getters: {
    totalComments(state: CommentsState) {
      return state.totalComments;
    },

    allComments(state: CommentsState) {
      return state.comments.map((comm: openapi.Comment) => {
        // eslint-disable-next-line no-param-reassign
        comm.content = Marked.parse(comm.content);
        return comm;
      });
    },

    getPost(state: CommentsState) {
      // eslint-disable-next-line no-param-reassign
      state.post.content = Marked.parse(state.post.content);
      return state.post;
    },

    getPageComm(state: CommentsState) {
      return state.page;
    },

    getNumComm(state: CommentsState) {
      return state.num;
    },
  },
};
