<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      <div v-for="book in books" :key="book.id">
        <nuxt-link :to="`books/${book.ID}`">
          <p>id: {{ book.ID }}</p>
        </nuxt-link>
        <p>{{ book.Title }}</p>
      </div>
      <v-text-field v-model="email"></v-text-field>
      <v-text-field v-model="password"></v-text-field>
      <v-btn @click="signUp()">signUp</v-btn>
      <nuxt-link to="/sign_in">
        <v-btn>signIn</v-btn>
      </nuxt-link>
      <template v-if="authenticated">
        <p>ログイン中</p>
        <v-btn @click="signOut()">signOut</v-btn>
      </template>
    </v-col>
  </v-row>
</template>

<script>
import Logo from '~/components/Logo.vue'
import VuetifyLogo from '~/components/VuetifyLogo.vue'
import firebase from '~/plugins/firebase'

export default {
  async asyncData(ctx) {
    try {
      const response = await ctx.$axios.get("books/index")
      const books = response.data
      return {
        books,
        email: 'ohishikaito@gmail.com',
        password: 'adaadaada',
      }
    } catch(err) {
      console.error(err)
    }
  },
  components: {
    Logo,
    VuetifyLogo
  },
  computed: {
    authenticated() {
      return this.$store.getters['getUid']
    }
  },
  methods: {
    async signUp() {
      try {
      const response = await firebase.auth().createUserWithEmailAndPassword(this.email, this.password)
      const user = response.user
      const req = {
        Email: user.email,
        LastName: "大石",
        FirstName: "海渡",
      }
      await this.$axios.post("users/create", req)
      // ログイン状態を保持するため
      this.$store.dispatch('setUid', { uid: user.uid })
      } catch(err) {
        console.error(err)
        console.error(err.response)
      }
    },
    async signOut() {
      this.$store.dispatch('unsetUid')
    },
  }
}
</script>
