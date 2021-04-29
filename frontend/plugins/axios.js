export default (ctx) => {
  ctx.$axios.onRequest((config) => {
    console.log(config)
    return config
  })
  ctx.$axios.onResponse((response) => {
    console.log(config)
    // if (response.headers['access-token']) {
    //   ctx.store.dispatch('auth', { headers: response.headers })
    // }
    // if (!response.data || typeof response.data !== 'object') {
    //   return response
    // }
    // response.data = camelcaseKeys(response.data, { deep: true })
    return response
  })
  // ctx.$axios.onError(error => {
  //   if (error.response.status === 401) {
  //     ctx.store.dispatch('unAuthorized');
  //     ctx.redirect('/auth/sign-in');
  //   }
  // })
}
