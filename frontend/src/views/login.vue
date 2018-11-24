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
      :particlesNumber="40"
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
  import * as axios from "axios";
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
                if (this.input.username == this.$parent.mockAccount.username && this.input.password == this.$parent.mockAccount.password) {
                    // TODO: - change the mock
                    this.$emit("authenticated", true);
                    this.$router.replace({ name: "secure" });
                } else {
                    this.$emit("authenticated", false);
                    this.$router.replace({ name: "insecure" });
                }
            
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
    opacity: 0.95;
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