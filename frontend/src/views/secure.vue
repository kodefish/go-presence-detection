<template>
  <section>
    <div id="loggedIn" class="container is-fluid">
      <h1 class="title" style="text-align: center">User Control Panel</h1>

      <div>
        <button class="button is-primary" @click="addThisDevice();">
          Add this device
        </button>
      </div>

      <div style="margin-top: 16px">
        <button class="button field is-danger" @click="checkedRows = []"
            :disabled="!checkedRows.length">
            <b-icon icon="close"></b-icon>
            <span>Remove these devices</span>
        </button>

        <b-tabs>
            <b-tab-item label="My Devices">
                <b-table
                    :data="myDevices"
                    :columns="columnsDevices"
                    :checked-rows.sync="checkedRows"
                    focusable>

                    <template slot="bottom-left">
                        <b>Total checked</b>: {{ checkedRows.length }}
                    </template>
                </b-table>
            </b-tab-item>
        </b-tabs>
      </div>

      <div style="margin-top: 16px">
        <b-tabs>
            <b-tab-item label="Connected Users">
                <b-table
                    :data="connectedUsers"
                    :columns="columnsUsers">
                </b-table>
            </b-tab-item>
        </b-tabs>
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
        checkedRows: [],
        columnsDevices: [
            {
            field: "mac",
            label: "MAC address"
            },
            {
            field: "connected",
            label: "Currently connected"
            }
        ],
        connectedUsers: [],
        columnsUsers: [
            {
            field: "name",
            label: "Name"
            }
        ],
        input: {
            mac: ""
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
    },
    removeThisDevice() {
      if (this.checkedRows.length > 0) {
        this.$http.post(this.$store.state.server + "/delete-device", {
            // TODO: give the mac addresses to delete
        }, {
          headers: { Authorization: "Bearer " + this.$store.state.jwt }
        })
        .then(function(response) {
            console.log(JSON.stringify(response));
        }).catch(function(error) {
          this.$dialog.alert("Nothing here");
        });
      } else {
        this.$dialog.alert("You have to chose which device to delete");
      }
    },
    getAllDevices() {
      let self = this;
      this.$http.get(this.$store.state.server + "/devices", {
          headers: { Authorization: "Bearer " + this.$store.state.jwt }
        }).then(function(response) {
          console.log(JSON.stringify(response));
          this.myDevices.push(response.data[0]);
        }).catch(function(error) {
          this.$dialog.alert("Nothing here");
        });
    },
    getConnectedUsers() {
      console.log("Getting connected users");
      this.$http
        .get(this.$store.state.server + "/connected-users", {
          headers: { Authorization: "Bearer " + this.$store.state.jwt }
        })
        .then(function(response) {
          var asString = JSON.stringify(response); // Of the form userName -> MAC
          var asStrings = asString.split("\n");
          var users = asStrings;
          for (var i = 0; i < asStrings.length; i++) {
            users[i] = asStrings[i].split(" ")[0];
          }
        })
        .catch(function(error) {
          this.$dialog.alert("No one here");
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