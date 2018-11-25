<template>
  <section>
    <div id="loggedIn" class="container is-fluid">
      <h1 class="title" style="text-align: center">User Control Panel</h1>

      <div>
        <button class="button is-primary" @click="addThisDevice();">
          Add this device
        </button>
      </div>

      <section>
          My devices: {{ myDevices }} <br />
          Other connected users: {{ connectedUsers }} <br />
      </section>

      <div>
          <b-field label="Device to remove">
            <b-input
                v-model="input.toRemove"
                placeholder="MAC address"
                style="width: 50%"
            ></b-input>
          </b-field>

          <button class="button is-danger" @click="removeThisDevice();">
            Remove this device
          </button>
      </div>

      <div style="margin-top: 16px">
          <button class="button is-primary" @click="getConnectedUsers();">
            Refresh the list of connected users
          </button>
      </div>

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
  name: "Secure",
  data() {
    return {
        myDevices: [],
        connectedUsers: [],
        input: {
            toRemove: ""
        }
    };
  },
  methods: {
    addThisDevice() {
      this.$http
        .get(this.$store.state.server + "/add-device", {
          headers: { Authorization: "Bearer " + this.$store.state.jwt }
        })
        .then(function(response) {
          console.log(JSON.stringify(response));
        })
        .catch(function(error) {
          this.$dialog.alert("Sorry");
        });

        this.getAllDevices()
    },
    removeThisDevice() {
      if (this.input.toRemove != "") {
        let devices = [this.input.toRemove]
        this.$http.post(this.$store.state.server + "/delete-device", {
          devices
        }, {
          headers: { Authorization: "Bearer " + this.$store.state.jwt }
        }).then(function(response) {
            console.log(JSON.stringify(response));
        }).catch(function(error) {
          // Nothing
        });
      } else {
        this.$dialog.alert("You have to chose which device to delete");
      }

      this.getAllDevices()
    },
    getAllDevices() {
      let self = this;
      this.$http.get(this.$store.state.server + "/devices", {
          headers: { Authorization: "Bearer " + this.$store.state.jwt }
        }).then(function(response) {
          console.log(response.data);
          this.myDevices = response.data.devices;
        }).catch(function(error) {
            // Nothing
        });
    },
    getConnectedUsers() {
      console.log("Getting connected users");
      this.$http
        .get(this.$store.state.server + "/connected-users", {
          headers: { Authorization: "Bearer " + this.$store.state.jwt }
        })
        .then(function(response) {
            this.$dialog.alert(response)
          connectedUsers = response.data;
        })
        .catch(function(error) {
            // Nothing
        });
    }
  },
  created:function() {
    this.getAllDevices();
    this.getConnectedUsers();
  }
};
</script>

<style scoped>
#loggedIn {
  background-color: #ffffff;
  border: 1px solid #cccccc;
  opacity: 0.9;
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