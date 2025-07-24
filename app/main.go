package main

import (
	"fmt"
	"log"
	"net/http"
	"xray-panel/app/controllers"
	"xray-panel/app/services"
)

func main() {
	xrayService := services.NewXrayService()
	err := xrayService.Start()
	if err != nil {
		log.Fatalf("Error starting Xray: %v", err)
	}

	inboundController := controllers.NewInboundController()
	clientController := controllers.NewClientController()

	http.HandleFunc("/inbounds", inboundController.GetInbounds)
	http.HandleFunc("/inbounds/add", inboundController.AddInbound)
	http.HandleFunc("/inbounds/update", inboundController.UpdateInbound)
	http.HandleFunc("/inbounds/delete", inboundController.DeleteInbound)

	http.HandleFunc("/clients", clientController.GetClients)
	http.HandleFunc("/clients/add", clientController.AddClient)
	http.HandleFunc("/clients/update", clientController.UpdateClient)
	http.HandleFunc("/clients/delete", clientController.DeleteClient)

	serverController := controllers.NewServerController()
	http.HandleFunc("/servers", serverController.GetServers)
	http.HandleFunc("/servers/add", serverController.AddServer)
	http.HandleFunc("/servers/update", serverController.UpdateServer)
	http.HandleFunc("/servers/delete", serverController.DeleteServer)

	tunnelController := controllers.NewTunnelController()
	http.HandleFunc("/tunnels", tunnelController.GetTunnels)
	http.HandleFunc("/tunnels/add", tunnelController.AddTunnel)
	http.HandleFunc("/tunnels/update", tunnelController.UpdateTunnel)
	http.HandleFunc("/tunnels/delete", tunnelController.DeleteTunnel)

	subscriptionController := controllers.NewSubscriptionController()
	http.HandleFunc("/sub/{token}", subscriptionController.GetSubscription)

	userController := controllers.NewUserController()
	http.HandleFunc("/user", userController.GetUser)
	http.HandleFunc("/user/update", userController.UpdateUser)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
