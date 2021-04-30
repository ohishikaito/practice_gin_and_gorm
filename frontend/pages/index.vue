<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      <div v-for="book in books" :key="book.id">
        <p>{{ book.ID }}</p>
        <p>{{ book.Title }}</p>
      </div>
      <!-- <v-text-field v-model="email"></v-text-field>
      <v-text-field v-model="password"></v-text-field> -->
      <v-btn @click="signUp">ボタン</v-btn>
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
      console.log(response, 'response')
      // ctx.$axios.get("books/index"),
      // await Promise.all([
      //   ctx.$axios.get("books/index"),
      //   ctx.$axios.get("books"),
      //   ctx.$axios.get("/books"),
      //   ctx.$axios.get("books/"),
      //   ctx.$axios.get("/books/"),
      // ])
      const books = response.data.data
      console.log(books)
      return {
        books,
        email: '',
        password: '',
      }
    } catch(err) {
      console.error(err)
    }
  },
  components: {
    Logo,
    VuetifyLogo
  },
  methods: {
    async signUp() {
      firebase.auth().createUserWithEmailAndPassword(this.email, this.pasword)
    }
  }
}
</script>
