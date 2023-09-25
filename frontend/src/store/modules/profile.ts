// eslint-disable-next-line max-classes-per-file
import * as openapi from '@/openapi';
import API from '@/api';

export default {
  actions: {
    async viewProfile(ctx: any, login: string) {
      try {
        const resp = await API.viewProfileUser(login);

        ctx.commit('saveProfile', {
          profile: resp.data.user,
          subscribed: resp.data.subscribed,
          self: resp.data.self,
        });
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },

    async subscribe(ctx: any) {
      try {
        const resp = await API.subscribe(ctx.state.profile.id);

        ctx.commit('subscribe', resp.data.subscribed);

        ctx.dispatch('viewProfile', ctx.state.profile.login);
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },
  },

  mutations: {
    saveProfile(state: any, payload: { profile: openapi.User,
      subscribed: boolean, self: boolean }) {
      state.profile = payload.profile;
      state.subscribed = payload.subscribed;
      state.self = payload.self;
    },

    subscribe(state: any, subscribed: boolean) {
      state.subscribed = subscribed;
    },
  },

  state: {
    profile: {} as openapi.User,
    subscribed: false,
    self: false,
  },

  getters: {
    profile(state: any) {
      return state.profile;
    },

    subscribed(state: any) {
      return state.subscribed;
    },

    self(state: any) {
      return state.self;
    },
  },
};
