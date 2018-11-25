<template>
  <section v-bind:style="{ background: 'url(' + this.$store.state.backgroundImagePath + ') no-repeat center center fixed' }"
           id="main-section" class="hero is-fullheight has-background">
    <div id="app">
      <div id="nav">
        <router-link
          v-if="this.$store.state.userIsLoggedIn"
          to="/login"
          v-on:click.native="logout();"
          replace
        >
          <button class="button" @click="logout();">Logout</button>
        </router-link>
        <router-link
          v-else-if="this.$route.path==='/insecure'"
          to="/login"
          replace
        >
          <button class="button" @click="back();">Try again</button>
        </router-link>
      </div>
      <router-view />
    </div>
  </section>

</template>

<script>
  export default {
    name: "App",
    data() {
      return {};
    },
    methods: {
      back () {
        this.$router.replace({ name: "signup" });
      },
      logout() {
        this.$store.state.userIsLoggedIn = false
        this.$router.replace({ name: "login" });
      }
    }
  };
</script>

<style lang="scss">
  // Import Bulma's core
  @import "~bulma/sass/utilities/_all";
  // Set your colors
  $primary: #8c67ef;
  $primary-invert: findColorInvert($primary);
  $twitter: #4099ff;
  $twitter-invert: findColorInvert($twitter);
  // Setup $colors to use as bulma classes (e.g. 'is-twitter')
  $colors: (
    "white": (
      $white,
      $black
    ),
    "black": (
      $black,
      $white
    ),
    "light": (
      $light,
      $light-invert
    ),
    "dark": (
      $dark,
      $dark-invert
    ),
    "primary": (
      $primary,
      $primary-invert
    ),
    "info": (
      $info,
      $info-invert
    ),
    "success": (
      $success,
      $success-invert
    ),
    "warning": (
      $warning,
      $warning-invert
    ),
    "danger": (
      $danger,
      $danger-invert
    ),
    "twitter": (
      $twitter,
      $twitter-invert
    )
  );
  html {
    height: 100%;
  }
  body {
    /* The image used */
    background-image: url("assets/background.jpg");
    background-size: cover;
    /* Full height */
    height: 100%;
  }
  // Links
  $link: $primary;
  $link-invert: $primary-invert;
  $link-focus-border: $primary;
  // Import Bulma and Buefy styles
  @import "~bulma";
  @import "~buefy/src/scss/buefy";
  .has-background {
    -webkit-background-size: cover;
    -moz-background-size: cover;
    -o-background-size: cover;
    background-size: cover;
  }
</style>