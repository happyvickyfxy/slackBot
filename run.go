package main

func main() {
	users, err := getAllUsers()
	if err != nil {
		panic(err)
	}

	user, err := findUserInList(users, "Xiaoyu Fang")
	if err != nil {
		panic(err)
	}

	//err = sendMessage(user, "You are a pretty princess!")
	//if err != nil {
	//	panic(err)
	//}

	//var mode string
	//fmt.Println("Please choose the mode: handle or interactive")
	//fmt.Scan(&mode)
	//
	//if mode == "handle" {
	//	go rtm.ManageConnection()
	//	handleMessage()
	//} else if mode == "interactive" {
		sendMessage(user, "Am I beautiful??")
	//} else {
	//	fmt.Println("no mode selected")
	//}
}
