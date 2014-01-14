package controllers

import (
	r "github.com/robfig/revel"
	c "tux21b.org/v1/gocql"
)

var (
	Cl *c.ClusterConfig
)

func Init() {
	cluster := c.NewCluster("127.0.0.1")
	cluster.Keyspace = "mykeyspace"
	cluster.Consistency = c.One
	Cl = cluster
}

type CassandraController struct {
	*r.Controller
	Sess *c.Session
}

func (c *CassandraController) Begin() r.Result {
	session, err := Cl.CreateSession()
	if err != nil {
		panic(err)
	}
	c.Sess = session
	return nil
}

func (c *CassandraController) Finish() r.Result {
	c.Sess.Close()
	return nil
}
