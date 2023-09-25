// eslint-disable-next-line max-classes-per-file
import * as openapi from '@/openapi';
import API from '@/api';
import { Marked } from '@ts-stack/markdown';

export default {
  actions: {
    async getFeedPosts(ctx: any) {
      try {
        const resp = await API.viewFeed();

        ctx.commit('saveFeedPosts', {
          subPosts: resp.data.subPosts,
          noSubPosts: resp.data.noSubPosts,
        });
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },

    async getProfilePosts(ctx: any, login: string) {
      try {
        const resp = await API.viewProfilePosts(
          login,
        );

        ctx.commit('savePosts', resp.data.posts);
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },

    async changeReaction(ctx: any, payload: { postID: string, typeID: string }) {
      try {
        await API.react(
          payload.postID,
          new class implements openapi.ReactRequest {
            typeID = payload.typeID;
          }(),
        );

        const resp = await API.getPost(
          payload.postID,
        );

        ctx.commit('savePost', resp.data.post);
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },

    async changePermsPost(ctx: any, postID: string) {
      try {
        await API.changePostPerms(
          postID,
        );

        const resp = await API.getPost(
          postID,
        );

        ctx.commit('savePost', resp.data.post);
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },

    async deletePost(ctx: any, postID: string) {
      try {
        await API.deletePost(
          postID,
        );

        ctx.commit('deletePost', postID);
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },

    async publishPost(ctx: any, payload: { content: string, perms: boolean }) {
      try {
        const resp = await API.publish(
          new class implements openapi.PublishRequest {
            content = payload.content;

            perms = payload.perms;
          }(),
        );

        ctx.commit('addPost', resp.data.post);
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },
  },

  mutations: {
    saveFeedPosts(state: any, payload: { subPosts: Array<openapi.Post>,
      noSubPosts: Array<openapi.Post> }) {
      state.posts = payload.subPosts.concat(payload.noSubPosts);
    },

    savePosts(state: any, posts: Array<openapi.Post>) {
      state.posts = posts;
    },

    savePost(state: any, post: openapi.Post) {
      const index = state.posts.findIndex((el: openapi.Post) => el.id === post.id);
      if (index !== -1) {
        state.posts[index] = post;
      }
    },

    deletePost(state: any, postID: string) {
      state.posts = state.posts.filter((post: openapi.Post) => (post.id !== postID));
    },

    addPost(state: any, post: openapi.Post) {
      state.posts.unshift(post);
    },
  },

  state: {
    posts: Array<openapi.Post>(),
  },

  getters: {
    allPosts(state: any) {
      return state.posts.map((post: openapi.Post) => {
        // eslint-disable-next-line no-param-reassign
        post.content = Marked.parse(post.content);
        return post;
      });
    },

    getPostByID(state: any, postID: string) {
      for (let i = 0; i < state.posts.length; i += 1) {
        if (state.posts[i].id === postID) {
          return state.posts[i];
        }
      }

      return null;
    },
  },
};
