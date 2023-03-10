package entity

import "context"

type Context struct {
	context.Context
	Response interface{}
}
