<template>
  <section>
    <div id="signup" class="container">
      <h1 class="title">Sign up</h1>
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
            style="width: 50%; margin-left: 25%"
          ></b-input>
        </b-field>

        <b-field label="Password (again)">
          <b-input
            type="password"
            v-model="input.passwordAgain"
            placeholder="Same password as before"
            style="width: 50%; margin-left: 25%"
          ></b-input>
        </b-field>

        <button class="button is-primary" @click="signup();">
          Sign up
        </button>

        <button class="button is-primary" @click="login();">
          Log in instead
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
import * as axios from "axios";
export default {
  name: "Signup",
  data() {
    return {
      filedata: 0,
      input: {
        username: "",
        password: "",
        passwordAgain: ""
      }
    };
  },
  methods: {
    signup() {
      if (this.input.username !== "" && this.input.password !== "") {
        if (this.input.password != this.input.passwordAgain) {
          this.$dialog.alert("The passwords you entered are not the same");
        } else {
          this.$http
            .post(
              this.$store.state.server + "/register",
              {
                username: this.input.username,
                password: this.input.password
              }/*,
              { emulateJSON: true }*/
            )
            .then(function(response) {
              console.log(response);
              this.$store.state.userIsLoggedIn = true;
              this.$store.state.jwt = response.body.token;
              this.$router.replace({ name: "secure" });
            })
            .catch(function(error) {
              console.log(error);
              this.$dialog.alert("This username is already taken");
            });
        }
      } else {
        this.$dialog.alert("You cannot have an empty username or password");
      }

      this.input.username = "";
      this.input.password = "";
      this.input.passwordAgain = "";
    },
    login() {
      this.$router.replace({ name: "login" });
    }
  }
};
</script>

<style scoped>
#signup {
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