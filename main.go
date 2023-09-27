package main

import (
	"project-api/database"
	"project-api/router"
)

// get all config of database & env
func init(){
  database.ConfigEnv()
  database.ConnectDB()
}

// call all of the route
func main(){
  router.Route()
}