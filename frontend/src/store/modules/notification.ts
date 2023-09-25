export default {
  actions: {
    async addError(ctx: any, error: string) {
      ctx.commit('addError', error);
    },

    async deleteError(ctx: any, error: string) {
      ctx.commit('deleteError', error);
    },
  },

  mutations: {
    addError(state: any, error: string) {
      state.errors.unshift(error);
      console.error(state.errors);
    },

    deleteError(state: any, errorDel: string) {
      state.errors = state.errors.filter((error: string) => (error !== errorDel));
    },
  },

  state: {
    errors: Array<string>(),
  },

  getters: {
    errors(state: any) {
      return state.errors;
    },
  },
};
