<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      <h1>{{ email }}</h1>
      <h1>{{ password }}</h1>
      <v-btn @click="signIn()">signIn</v-btn>
    </v-col>
  </v-row>
</template>

<script>
import Logo from '~/components/Logo.vue'
import VuetifyLogo from '~/components/VuetifyLogo.vue'
import firebase from '~/plugins/firebase'

export default {
  data() {
    return {
      email: 'ohishikaito@gmail.com',
      password: 'adaadaada',
    }
  },
  components: {
    Logo,
    VuetifyLogo
  },
  methods: {
    async signIn() {
      try {
        const response = await firebase.auth().signInWithEmailAndPassword(this.email, this.password)
        const user = response.user
        const idToken = await user.getIdToken(/* forceRefresh */ true)
        this.$store.dispatch('setIdToken', { idToken })
        this.$router.push('/')
      } catch(err) {
        console.error(err)
        console.error(err.response)
      }
    },
  }
}
</script>
