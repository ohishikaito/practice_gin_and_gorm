export const state = () => ({
  uid: null,
})

export const getters = {
  getUid: (state) => state.uid,
}

export const mutations = {
  setUid(state, uid) {
    state.uid = uid;
    localStorage.setItem('uid', uid)
  },
  unsetUid(state) {
    state.uid = null;
    localStorage.removeItem('uid')
  },
};

export const actions = {
  nuxtClientInit ({ commit }) {
    const uid = localStorage.getItem('uid');
    if (uid) {
      commit('setUid', uid);
    }
  },
  setUid({ commit }, payload) {
    commit('setUid', payload.uid)
  },
  unsetUid({ commit }) {
    commit('unsetUid')
  },
}
