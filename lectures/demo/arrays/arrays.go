package main

import "fmt"

type Room struct{
	roomNumber int
	isCleaned bool
}

func checkCleanliness (rooms [4]Room) {
	for i := 0; i < len(rooms); i++{
		room := rooms[i]
		if room.isCleaned{
			fmt.Println(room.roomNumber, "is clean!")
		} else{
			fmt.Println(room.roomNumber, "is dirty!")
		}
	}
}
func main() {
	motel := [...]Room{
		{1, false}, 
		{2, false}, 
		{3, true}, 
		{4, false}}
	fmt.Println(motel)
	checkCleanliness(motel)


}
