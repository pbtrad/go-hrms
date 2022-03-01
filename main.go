package main

type MongoInstance struct {
	Client
	Db
}

var mg MongoInstance

const dbName = "fiber-hrms"
const mongoURI = ""
