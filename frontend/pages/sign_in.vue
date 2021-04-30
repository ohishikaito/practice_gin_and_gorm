<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      <v-btn @click="signIn()">signIn</v-btn>
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
      return {
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
  methods: {
    async signIn() {
      try {
        const response = await firebase.auth().signInWithEmailAndPassword(this.email, this.password)
        const user = response.user
        this.$store.dispatch('setUid', { uid: user.uid })
        this.$router.push('/')
      } catch(err) {
        console.error(err)
        console.error(err.response)
      }
    },
  }
}
</script>
