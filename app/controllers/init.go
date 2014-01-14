package controllers

import "github.com/robfig/revel"

func init() {
	revel.OnAppStart(Init)
	revel.InterceptMethod((*CassandraController).Begin, revel.BEFORE)
	revel.InterceptMethod((*CassandraController).Finish, revel.AFTER)
}
