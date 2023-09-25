// eslint-disable-next-line max-classes-per-file
import * as openapi from '@/openapi';
import API from '@/api';
import { Marked } from '@ts-stack/markdown';

export default {
  actions: {
    async getComments(ctx: any, id: string) {
      try {
        const resp1 = await API.viewComments(id);

        const resp2 = await API.getPost(
          id,
        );

        ctx.commit('savePost', resp2.data.post);

        ctx.commit('saveComments', resp1.data.comments);
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },

    async comment(ctx: any, payload: { postID: string, content: string }) {
      try {
        const resp = await API.comment(
          payload.postID,
          new class implements openapi.CommentRequest {
            content = payload.content;
          }(),
        );

        ctx.commit('saveComment', resp.data.comment);
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },

    async uncomment(ctx: any, payload: { postID: string, commID: string }) {
      try {
        console.error(payload);
        await API.uncomment(
          payload.postID,
          new class implements openapi.UncommentRequest {
            commID = payload.commID;
          }(),
        );

        ctx.commit('deleteComment', payload.commID);
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },
  },

  mutations: {
    saveComments(state: any, comments: Array<openapi.Comment>) {
      state.comments = comments;
    },

    saveComment(state: any, comment: openapi.Comment) {
      state.comments.unshift(comment);
    },

    savePost(state: any, post: openapi.Post) {
      state.post = post;
    },

    deleteComment(state: any, commID: string) {
      state.comments = state.comments.filter((comm: openapi.Comment) => (comm.id !== commID));
    },
  },

  state: {
    post: {} as openapi.Post,
    comments: Array<openapi.Comment>(),
  },

  getters: {
    allComments(state: any) {
      return state.comments.map((comm: openapi.Comment) => {
        // eslint-disable-next-line no-param-reassign
        comm.content = Marked.parse(comm.content);
        return comm;
      });
    },

    getPost(state: any) {
      // eslint-disable-next-line no-param-reassign
      state.post.content = Marked.parse(state.post.content);
      return state.post;
    },
  },
};
