<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      <nuxt-link to="/">BACK</nuxt-link>

      <p>id: {{ book.ID }}</p>
      <p>{{ book.Title }}</p>
      <p>{{ book.CreatedAt }}</p>

    </v-col>
  </v-row>
</template>

<script>
import Logo from '~/components/Logo.vue'
import VuetifyLogo from '~/components/VuetifyLogo.vue'

export default {
  async asyncData(ctx) {
    try {
      const response = await ctx.$axios.get(`/books/${ctx.route.params.id}`, {
        headers: {'Authorization': `Bearer ${ctx.store.getters['getIdToken']}`}
      })
      console.log(response)
      const book = response.data.data
      return {
        book,
      }
    } catch(err) {
      console.error(err, 'err')
      console.error(err.response, 'err.response')
    }
  },
  components: {
    Logo,
    VuetifyLogo
  }
}
</script>
