export const state = () => ({
  idToken: null,
})

export const getters = {
  getIdToken: (state) => state.idToken,
}

export const mutations = {
  setIdToken(state, idToken) {
    state.idToken = idToken;
    localStorage.setItem('idToken', idToken)
  },
  unsetIdToken(state) {
    state.idToken = null;
    localStorage.removeItem('idToken')
  },
};

export const actions = {
  nuxtClientInit ({ commit }) {
    const idToken = localStorage.getItem('idToken');
    if (idToken) {
      commit('setIdToken', idToken);
    }
  },
  setIdToken({ commit }, payload) {
    // const idToken = JSON.parse()
    commit('setIdToken', payload.idToken)
  },
  unsetIdToken({ commit }) {
    commit('unsetIdToken')
  },
}
