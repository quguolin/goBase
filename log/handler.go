package main

import "context"

type Field struct {
	Key string
	Value string
}

func KVString(key string,value string)Field  {
	return Field{
		Key:key,
		Value:value,
	}
}

type Handler interface {
	Log(context.Context,...Field)(error)
	Close()error
}
