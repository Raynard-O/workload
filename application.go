package main

import "Proto/router"

func main()  {
	e := router.New()
	// listening on port 5000
	e.Logger.Fatal(e.Start(":5000"))
}

