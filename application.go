package main

import "Proto/router"

func main()  {
	e := router.New()
	e.Logger.Fatal(e.Start(":5000"))
}

