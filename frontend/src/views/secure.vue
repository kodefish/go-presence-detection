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
                    checkable>

                    <template slot="bottom-left">
                        <b>Total checked</b>: {{ checkedRows.length }}
                    </template>
                </b-table>
            </b-tab-item>

            <b-tab-item label="Selected devices">
                <pre>{{ checkedRows }}</pre>
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
    const myDevices = [
      { mac: 1, connected: "Nah" },
      { mac: 2, connected: "Yeh" }
    ];
    const connectedUsers = [
      { name: ")'; DROP TABLE Users; --" },
      { name: "H4X0R" }
    ];
    return {
      myDevices,
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
      connectedUsers,
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
      if (false /* TODO: Check that it isn't added already */) {
        this.$dialog.alert({
          message: "Looks like this device is already added",
          confirmText: "Oh right, I totes forgot!"
        });
      } else {
        // TODO: Get MAC address
        // TODO: Add the stuff to the DB
        this.$toast.open({
          duration: 5000,
          message:
            "Your device has been added to our network. Tanks for your trust, we totally won't give any of your info to the NSA, we swear!"
        });
      }
    },
    removeThisDevice() {
      if (this.input.mac != "") {
        if (false /* TODO: if it isn't in the list */) {
          this.$dialog.alert({
            message: "You don't have any such device",
            confirmText: "Oh right, I totes forgot!"
          });
        } else {
          // TODO: Remove the stuff from the DB
          this.$toast.open({
            duration: 5000,
            message:
              "Your device has been removed from our network. You can count on us to *totally* stop tracking you now!"
          });
        }
      } else {
        this.$dialog.alert({
          message:
            "At least type something. It's like you're not even trying...",
          confirmText: "K, K, I'll do it"
        });
      }
    },
    getAllDevices() {
      console.log("Getting devices");
      this.$http
        .get(this.$store.state.server + "/devices", {
          headers: { Authorization: "Bearer " + this.$store.state.jwt }
        })
        .then(function(response) {
          console.log(JSON.stringify(response));
        })
        .catch(function(error) {
          this.$dialog.alert("Shit happens");
        });
    },
    wait(ms) {
      var start = new Date().getTime();
      var end = start;
      while (end < start + ms) {
        end = new Date().getTime();
      }
    },
    getConnectedUsers() {
      while (true) {
        console.log("Getting connected users");
        this.$http
          .get(this.$store.state.server + "/connected-users", {
            headers: { Authorization: "Bearer " + this.$store.state.jwt }
          })
          .then(function(response) {
            console.log(JSON.stringify(response));
          })
          .catch(function(error) {
            this.$dialog.alert("Shit happens");
          });

        wait(10000); // 10 seconds
      }
    }
  },
  beforeMount() {
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