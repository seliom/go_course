//--Summary:
//  Write a program to display server status.
//
//--Requirements:


package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

//* Create a function to print server status, including:
//  - Number of servers
//  - Number of servers for each status (Online, Offline, Maintenance, Retired)

func serverStatus(serverMap map[string] string){
	fmt.Println("There are", len(serverMap),"servers:")
	statusMap := make(map[string]int)
	for _, status := range serverMap{
		if status == "Online"{
			statusMap["Online"] += 1
		} else if status == "Offline"{
			statusMap["Offline"] += 1
		} else if status == "Maintenance"{
			statusMap["Maintenance"] += 1
		} else if status == "Retired"{
			statusMap["Retired"] += 1
		} else{
			panic("unhandled server status")
		}
	}
	fmt.Println("", statusMap["Online"], "Online servers\n",
				statusMap["Offline"], "Offline servers\n",
				statusMap["Maintenance"], "Maintenance servers\n", 
				statusMap["Retired"], "Retired servers")
	
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverMap := make(map[string]string)

	for _, server := range servers{
		serverMap[server] = "Online"
	}

	serverStatus(serverMap)
	serverMap["darkstar"] = "Retired"
	serverMap["aiur"] = "Offline"
	serverStatus(serverMap)

	for _, server := range servers{
		serverMap[server] = "Maintenance"
	}

	serverStatus(serverMap)
	
	//* Store the existing slice of servers in a map
	//* Default all of the servers to `Online`
	//* Perform the following status changes and display server info:
	//  - display server info
	//  - change `darkstar` to `Retired`
	//  - change `aiur` to `Offline`
	//  - display server info
	//  - change all servers to `Maintenance`
	//  - display server info
}
