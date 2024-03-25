package main

/*
By default we have one main go routine and that only creates new routines by go using keyword.
go --> This will create a new go routine i.e child go routine.
Once the main function of main go routine gets over it autmoatically executes and doesn't care
about the child routines.
That's where Channels came into action as they are used to communicate b/w different routines and
lets main routines know about the child ones.

// Here the requests are being handled one at a time

func checkLink(link string) {
	_, err := http.Get(link) // This is the Blocking Function as the go routine cann't do anything until its gets some response

	if err != nil {
		fmt.Println(link, err)
		return
	}
	fmt.Println(link, "is up!")
}


// Here the requests are being handled parallely
func checkLinkisRunning(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
		c <- "Maybe there's an issue"
		return
	}
	fmt.Println(link, "is up!")
	c <- "Yep its up"

}

*/
