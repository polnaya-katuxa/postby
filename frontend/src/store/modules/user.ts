// eslint-disable-next-line max-classes-per-file
import * as openapi from '@/openapi';
import API from '@/api';
import { LoginRequest, RegisterRequest } from '@/openapi';
import Cookies from 'cookies-ts';

export default {
  actions: {
    async userInfo(ctx: any) {
      try {
        const resp = await API.userInfo();

        ctx.commit('saveUser', resp.data.user);
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },

    async login(ctx: any, payload: {login: string, password: string}) {
      try {
        const resp = await API.login(
          new class implements LoginRequest {
            login = payload.login;

            password = payload.password;
          }(),
        );

        const cookies = new Cookies();
        cookies.set('user-token', resp.data.token);
        window.location.replace('/');
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },

    async register(ctx: any, payload: {login: string, password: string, email: string,
      picture: string, description: string}) {
      try {
        const resp = await API.register(
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
        window.location.replace('/');
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },
  },

  mutations: {
    saveUser(state: any, user: openapi.User) {
      state.user = user;
    },
  },

  state: {
    user: {} as openapi.User,
  },

  getters: {
    user(state: any) {
      return state.user;
    },

    isAdmin(state: any) {
      return state.user.isAdmin;
    },

    curLogin(state: any) {
      return state.user.login;
    },
  },
};
