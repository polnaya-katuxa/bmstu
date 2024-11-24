import { createStore, ActionContext } from 'vuex';
import { expect, test } from 'vitest';
import API from '@/api';
import * as openapi from '@/openapi/api';
import { Marked } from '@ts-stack/markdown';
import { CommentsState } from '@/store/states';

const CommentsVuexStore = () => createStore({
  actions: {
    async getComments(ctx: ActionContext<CommentsState, CommentsState>, payload: { id: string,
      page: number, num: number }) {
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

    async comment(ctx: ActionContext<CommentsState, CommentsState>, payload: { postID: string,
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

    async uncomment(ctx: ActionContext<CommentsState, CommentsState>, payload: { postID: string,
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

    clearComments(ctx: ActionContext<CommentsState, CommentsState>) {
      ctx.commit('clearCommentsMut');
    },

    incPageComments(ctx: ActionContext<CommentsState, CommentsState>) {
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
});

test('view comments', () => {
  const store = CommentsVuexStore();
  store.dispatch('getComments', { id: 'e95ab7b2-636e-447f-9f87-04072e4b3b9d', page: 1, num: 10 }).then(() => {
    expect(store.state.comments).toBe([{
      id: "'e95ab7b2-636e-447f-9f87-04072e4b3b9d",
      content: 'aaaaa',
      pubTime: 'string',
      commentator: {
        id: "'e95ab7b2-636e-447f-9f87-04072e4b3b9d",
        login: 'string',
        picture: 'string',
        description: 'string',
        balance: 0,
        mail: 'string',
        isAdmin: false,
      },
      postID: 'string',
    },
    ]);
    expect(store.state.totalComments).toBe(1);
  });
});

test('comment', () => {
  const store = CommentsVuexStore();
  store.dispatch('comment', { postID: 'e95ab7b2-636e-447f-9f87-04072e4b3b9d', content: 'aaaaa' }).then(() => {
    expect(store.state.comments).toBe([{
      id: "'e95ab7b2-636e-447f-9f87-04072e4b3b9d",
      content: 'aaaaa',
      pubTime: 'string',
      commentator: {
        id: "'e95ab7b2-636e-447f-9f87-04072e4b3b9d",
        login: 'string',
        picture: 'string',
        description: 'string',
        balance: 0,
        mail: 'string',
        isAdmin: false,
      },
      postID: 'string',
    },
    ]);
    expect(store.state.totalComments).toBe(0);
  });
});

test('uncomment', () => {
  const store = CommentsVuexStore();
  store.dispatch('uncomment', {
    postID: 'e95ab7b2-636e-447f-9f87-04072e4b3b9d',
    commID: 'e95ab7b2-636e-447f-9f87-04072e4b3b9d',
  }).then(() => {
    expect(store.state.comments).toBe([]);
    expect(store.state.totalComments).toBe(0);
  });
});
