// eslint-disable-next-line max-classes-per-file
import * as openapi from '@/openapi';
import API from '@/api';

export default {
  actions: {
    async viewUsers(ctx: any) {
      try {
        const resp = await API.viewUsers();

        ctx.commit('saveUsers', resp.data.users);
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },

    async deleteUser(ctx: any, login: string) {
      try {
        await API.deleteUser(
          login,
        );

        ctx.commit('deleteUser', login);
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },
  },

  mutations: {
    saveUsers(state: any, users: Array<openapi.User>) {
      state.users = users;
    },

    deleteUser(state: any, login: string) {
      state.users = state.users.filter((user: openapi.User) => (user.login !== login));
    },
  },

  state: {
    users: Array<openapi.User>(),
  },

  getters: {
    users(state: any) {
      return state.users;
    },
  },
};
