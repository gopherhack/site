package controllers

import (
	mgo "gopkg.in/mgo.v2"
)

type Controller struct {
	DB *mgo.Session
}

const (
	// Key (Should come from somewhere else).
	Key = "secret"
)
