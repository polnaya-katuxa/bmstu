// eslint-disable-next-line max-classes-per-file
import * as openapi from '@/openapi/api';
import API from '@/api';
import { Marked } from '@ts-stack/markdown';
import { ActionContext } from 'vuex';
import {
  PostsState, State,
} from '@/store/states';

export default {
  actions: {
    async getFeedPosts(ctx: ActionContext<PostsState, State>, payload: { page: number,
      num: number }) {
      try {
        const resp = await API.postAPI.getPosts('', 'feed', payload.page, payload.num);

        ctx.commit('saveFeedPosts', { posts: resp.data.posts, total: resp.data.total });
        // eslint-disable-next-line
      } catch (err: any) {
        await ctx.dispatch('addError', err.response.data.message);
      }
    },

    async getProfilePosts(ctx: ActionContext<PostsState, State>, payload: { login: string,
      page: number, num: number }) {
      try {
        const resp = await API.postAPI.getPosts(payload.login, '', payload.page, payload.num);

        ctx.commit('savePosts', { posts: resp.data.posts, total: resp.data.total });
        // eslint-disable-next-line
      } catch (err: any) {
        await ctx.dispatch('addError', err.response.data.message);
      }
    },

    async changeReaction(ctx: ActionContext<PostsState, State>, payload: { postID: string,
      typeID: string }) {
      try {
        await API.reactionAPI.react(
          payload.postID,
          new class implements openapi.ReactRequest {
            typeID = payload.typeID;
          }(),
        );

        const resp = await API.postAPI.getPost(
          payload.postID,
        );

        ctx.commit('savePost', resp.data.post);
        // eslint-disable-next-line
      } catch (err: any) {
        await ctx.dispatch('addError', err.response.data.message);
      }
    },

    async changePermsPost(ctx: ActionContext<PostsState, State>, postID: string) {
      try {
        await API.postAPI.changePostPerms(
          postID,
        );

        const resp = await API.postAPI.getPost(
          postID,
        );

        ctx.commit('savePost', resp.data.post);
        // eslint-disable-next-line
      } catch (err: any) {
        await ctx.dispatch('addError', err.response.data.message);
      }
    },

    async deletePost(ctx: ActionContext<PostsState, State>, postID: string) {
      try {
        await API.postAPI.deletePost(
          postID,
        );

        ctx.commit('deletePost', postID);
        // eslint-disable-next-line
      } catch (err: any) {
        await ctx.dispatch('addError', err.response.data.message);
      }
    },

    async publishPost(ctx: ActionContext<PostsState, State>, payload: { content: string,
      perms: boolean }) {
      try {
        const resp = await API.postAPI.publishPost(
          new class implements openapi.PublishRequest {
            content = payload.content;

            perms = payload.perms;
          }(),
        );

        ctx.commit('addPost', resp.data.post);
        // eslint-disable-next-line
      } catch (err: any) {
        await ctx.dispatch('addError', err.response.data.message);
      }
    },

    clearPosts(ctx: ActionContext<PostsState, State>) {
      ctx.commit('clearPostsMut');
    },

    incPagePosts(ctx: ActionContext<PostsState, State>) {
      ctx.commit('incPosts');
    },
  },

  mutations: {
    clearPostsMut(state: PostsState) {
      state.posts = [];
      state.page = 1;
      state.totalPosts = 0;
    },

    saveFeedPosts(state: PostsState, payload: {posts: Array<openapi.Post>, total: number}) {
      state.posts.push(...payload.posts);
      state.totalPosts = payload.total;
    },

    savePosts(state: PostsState, payload: {posts: Array<openapi.Post>, total: number}) {
      state.posts.push(...payload.posts);
      state.totalPosts = payload.total;
    },

    savePost(state: PostsState, post: openapi.Post) {
      const index = state.posts.findIndex((el: openapi.Post) => el.id === post.id);
      if (index !== -1) {
        state.posts[index] = post;
      }
    },

    deletePost(state: PostsState, postID: string) {
      state.posts = state.posts.filter((post: openapi.Post) => (post.id !== postID));
    },

    addPost(state: PostsState, post: openapi.Post) {
      state.posts.unshift(post);
    },

    incPosts(state: PostsState) {
      state.page += 1;
    },
  },

  state: (): PostsState => ({
    posts: Array<openapi.Post>(),
    totalPosts: 0,
    page: 1,
    num: 10,
  }),

  getters: {
    totalPosts(state: PostsState) {
      return state.totalPosts;
    },

    allPosts(state: PostsState) {
      return state.posts.map((post: openapi.Post) => {
        // eslint-disable-next-line no-param-reassign
        post.content = Marked.parse(post.content);
        return post;
      });
    },

    getPostByID(state: PostsState, postID: string) {
      for (let i = 0; i < state.posts.length; i += 1) {
        if (state.posts[i].id === postID) {
          return state.posts[i];
        }
      }

      return null;
    },

    getPagePosts(state: PostsState) {
      return state.page;
    },

    getNumPosts(state: PostsState) {
      return state.num;
    },
  },
};
