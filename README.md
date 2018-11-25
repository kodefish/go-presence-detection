# Sensorless Presence Detection System
## Description
Lauzhack was the perfect platform for us to build our sensorless presece detection system. Having thought about it for quite some time, it was time to finally build it.
SDPS is meant to be an easy to use RESTful API one can use to monitor devices in their localnet. One can map certain devices to their owners, and so uniqueley via the MAC address.
To do so, our system scans the local network and keeps track of the MAC addresses. The users authenticate themselves to the API via a JWT webtoken, which then let's us match their identity to a MAC address. 

## Frontend
As a proof of concept, we built two front-ends, a web-based interface along with a mobile one. The web interface is built using VueJS and the Beaufy framework. The mobile one is a native Android app.

## Backend
The backend is built in Go, and relies on mongoDB database for persistence. It exposes the following end-points:

* `/register` register an account
* `/get-token` returns a JWT token that is used to identify you later on
* `/add-device` adds the current device to your pool
* `/devices` returns the devices in your pool
* `/delete-device` remove the current device from the pool
* `/connected-users` shows a list of users currently connected users

Since the project was mostly done for localnets, we did not focus much on security. The reasoning being that the intended use of the project is at home, with devices under your control.

## Future Work
Now that we have an api that allows for easy tracking of persons, many applications can be built upon it, such as:

* Shared Music box which adapts to current present person's tastes
* Home automation/IoT without any external sensors for presence detection
* Social media bot

We will save the ideas for our next hackathon!

## Repo structure
* detection         - Contains logic for scanning the localnet and saving IPs and MAC addresses
* crpyto            - Contains helper functions to hash user passwords
* api/ - contains all api related stuff
    + controller.go - Defines handler methods for endpoints
    + model.go      - User models
    + db.go         - Methods to interact with db
    + router.go     - Defines routes and endpoints
* README.md         - Readme for documentation
* main.go           - Entry point for api
* frontend          - Contains VueJS frontend
* android           - Android app
