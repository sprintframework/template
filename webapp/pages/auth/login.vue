<template>
  <section class="section">
    <div class="container">
      <div class="columns">
        <div class="column is-4 is-offset-4">
          <h2 class="title has-text-centered">Login</h2>

          <Notification v-if="error" :message="error" @close="error=null"/>

          <form method="post" class="box" @submit.prevent="login">

            <div class="field">
              <label class="label required">Login</label>

              <div class="control">
                <input
                  v-model="username"
                  type="text"
                  class="input"
                  name="username"
                >
              </div>
            </div>

            <div class="field">
              <label class="label required">Password</label>

              <div class="control">
                <input
                  v-model="password"
                  type="password"
                  class="input"
                  name="password"
                >
              </div>
            </div>

            <div class="control">
              <button type="submit" class="button is-dark is-fullwidth">Log In</button>
            </div>
          </form>

          <div class="has-text-centered" style="margin-top: 20px">
            <p>
              Don't have an account? <nuxt-link to="/auth/register">Register</nuxt-link>
            </p>
          </div>

          <div class="has-text-centered" style="margin-top: 20px">
            <p>
              Forgot password? <nuxt-link to="/auth/restore_password">Restore</nuxt-link>
            </p>
          </div>

        </div>
      </div>
    </div>
  </section>
</template>

<script>
import Notification from '~/components/Notification';

export default {

  components: {
    Notification,
  },

  middleware: 'guest',

  data() {
    return {
      username: '',
      password: '',
      error: null,
    };
  },

  methods: {
    async login() {
      try {
        await this.$auth.loginWith('local', {
          data: {
            login: this.username,
            password: this.password,
          },
        });

        this.$router.push('/');
      } catch (e) {
        this.error = e.response.data.message;
      }
    },
  },
};
</script>

<style>
  .required:after {
    content:" *";
    color: red;
  }
</style>

