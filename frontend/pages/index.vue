<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      <div v-for="book in books" :key="book.id">
        <p>{{ book.id }}</p>
        <p>{{ book.name }}</p>
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
      // const url = "https://present-api.aircatalog-dev.jp/api/v1/facilities"
      // const url = "http://localhost:8080/books"
      const response = await ctx.$axios.get("/books")
      // const response = await ctx.$axios.get("/books")
      console.log(response, 'response')
      // const books = response
      return {
        // books,
        // email: '',
        // password: '',
      }
    } catch(err) {
      console.log(err)
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
