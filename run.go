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

	err = sendMessage(user, "You are a pretty princess!")
	if err != nil {
		panic(err)
	}

	go rtm.ManageConnection()
	handleMessage()
}
