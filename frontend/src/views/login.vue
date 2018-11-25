<template>
  <section>
    <div id="login" class="container">
      <h1 class="title">Login</h1>
      <section>
        <b-field label="Username">
          <b-input
            v-model="input.username"
            placeholder="Username"
            style="width: 50%; margin-left: 25%"
          ></b-input>
        </b-field>

        <b-field label="Password">
          <b-input
            type="password"
            v-model="input.password"
            placeholder="Password"
            password-reveal
            style="width: 50%; margin-left: 25%"
          ></b-input>
        </b-field>

        <button class="button is-primary" @click="login();">
          Login
        </button>

        <button class="button is-primary" @click="signup();">
          Sign up instead
        </button>
      </section>
    </div>
    <vue-particles>
      color="#FFFFFF"
      :particleOpacity="0.7"
      :particlesNumber="200"
      shapeType="circle"
      :particleSize="4"
      linesColor="#dedede"
      :linesWidth="1"
      :lineLinked="true"
      :lineOpacity="0.4"
      :linesDistance="150"
      :moveSpeed="2"
      :hoverEffect="false"
      :clickEffect="false"
    &gt;
    </vue-particles>
  </section>
</template>

<script>
export default {
  name: "Login",
  data() {
    return {
      filedata: 0,
      input: {
        username: "",
        password: ""
      }
    };
  },
  methods: {
    login() {
      if (this.input.username !== "" && this.input.password !== "") {
        this.$http
          .post(
            this.$store.state.server + "/get-token",
            {
              username: this.input.username,
              password: this.input.password
            }/*,
            { emulateJSON: true }*/
          )
          .then(function(response) {
            this.$store.state.userIsLoggedIn = true;
            this.$store.state.jwt = response.body.token;
            this.$router.replace({ name: "secure" });
          })
          .catch(function(error) {
            this.$router.replace({ name: "insecure" });
          });
      } else {
        this.$dialog.alert({
          message:
            "At least type something. It's like you're not even trying...",
          confirmText: "K, K, I'll do it"
        });
      }
    },
    signup() {
      this.$router.replace({ name: "signup" });
    }
  }
};
</script>

<style scoped>
  #login {
    border: 1px solid #cccccc;
    background-color: #ffffff;
    opacity: 0.9;
    display: inline-block;
    margin-left: 25%;
    margin-top: 1%;
    text-align: center;
    margin-right: auto;
    width: 50%;
    padding: 20px;
    position: fixed;
    overflow-y: scroll;
    border-radius: 5px;
  }
  .button {
    margin-top: 16px;
  }
</style>