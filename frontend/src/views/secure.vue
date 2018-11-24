<template>
  <section>
    <div id="loggedIn" class="container is-fluid">
      <h1 class="title" style="text-align: center">User Control Panel</h1>

      <div>
        <b-collapse :open="false">
            <button class="button is-primary" slot="trigger">View my devices</button>
            <div class="notification">
                <div class="content">
                    <h3>
                        My devices
                    </h3>
                    <p>
                        Gotta fill that
                    </p>
                </div>
            </div>
        </b-collapse>
      </div>

      <div style="margin-top: 16px">
        <button class="button is-primary" @click="addThisDevice();">
          Add this device
        </button>
      </div>

      <div style="margin-top: 16px">
        <div class="t">
        <b-field label="Device you want to remove">
          <b-input
            v-model="input.mac"
            placeholder="MAC address"
            style="width: 500px"
          ></b-input>
        </b-field>
        </div>
        
        <button class="button is-primary" style="margin-top: 16px" @click="removeThisDevice();">
          Remove this device
        </button>
      </div>
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
    name: "Secure",
    data() {
      return {
          input: {
              mac: ""
          }
      };
    },
    mounted() {
      this.get_user_info();
    },
    methods: {
        addThisDevice() {
            if (false /* TODO: Check that it isn't added already */) {
                this.$dialog.alert({
                    message:
                    "Looks like this device is already added",
                    confirmText: "Oh right, I totes forgot!"
                });
            } else {
                // TODO: Get MAC address
                // TODO: Add the stuff to the DB
                this.$toast.open({
                    duration: 5000,
                    message: "Your device has been added to our network. Tanks for your trust, we totally won't give any of your info to the NSA, we swear!"
                });
            }
        },
        removeThisDevice() {
            if (this.input.mac != "") {
                if (false /* TODO: if it isn't in the list */) {
                    this.$dialog.alert({
                        message:
                        "You don't have any such device",
                        confirmText: "Oh right, I totes forgot!"
                    });
                } else {
                    // TODO: Remove the stuff from the DB
                    this.$toast.open({
                        duration: 5000,
                        message: "Your device has been removed from our network. You can count on us to *totally* stop tracking you now!"
                    });
                }
            } else {
                this.$dialog.alert({
                    message:
                    "At least type something. It's like you're not even trying...",
                    confirmText: "K, K, I'll do it"
                });
            }
        }
    }
  };
</script>

<style scoped>
  #loggedIn {
    background-color: #ffffff;
    border: 1px solid #cccccc;
    opacity: 0.95;
    padding: 20px;
    margin-top: 10px;
  }
  .article-container {
    display: flex;
    flex-wrap: wrap;
    text-align: left;
  }
  .article {
    flex: 0 0 50%;
    padding: 10px;
  }
  * {
    margin: 0;
    box-sizing: border-box;
  }
</style>