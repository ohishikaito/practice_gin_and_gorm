export default (ctx) => {
  ctx.$axios.onRequest((config) => {
    // const idToken = ctx.store.getters['getIdToken']
    config.headers = {
      'Authorization': `Bearer ${ctx.store.getters['getIdToken']}`
    }
    return config
  })
  ctx.$axios.onResponse((response) => {
    // console.error(config)
    // if (response.headers['access-token']) {
    //   ctx.store.dispatch('auth', { headers: response.headers })
    // }
    // if (!response.data || typeof response.data !== 'object') {
    //   return response
    // }
    // response.data = camelcaseKeys(response.data, { deep: true })
    // return response
  })
  // ctx.$axios.onError(error => {
  //   if (error.response.status === 401) {
  //     ctx.store.dispatch('unAuthorized');
  //     ctx.redirect('/auth/sign-in');
  //   }
  // })
  ctx.$axios.onError(error => {
    if (error.response.status === 401) {
      ctx.store.dispatch('unsetIdToken')
      ctx.redirect('/sign_in')
    }
  });
}
