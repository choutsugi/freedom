package annotation

import (
	"google.golang.org/protobuf/proto"
	"reflect"
)

type Annotation struct {
	Prototypes map[uint16]reflect.Type
}

func (anno *Annotation) Register(id uint16, msg proto.Message) {
	if _, ok := anno.Prototypes[id]; ok {
		return
	}
	of := reflect.TypeOf(msg)
	anno.Prototypes[id] = of.Elem()
}

func (anno *Annotation) GetPrototype(id uint16) reflect.Type {
	return anno.Prototypes[id]
}

func (anno *Annotation) UnMarshal(id uint16, data []byte) proto.Message {

	if protoType := anno.GetPrototype(id); protoType != nil {
		ret := reflect.New(protoType).Interface().(proto.Message)
		if err := proto.Unmarshal(data, ret); err != nil {
			return nil
		}
		return ret
	}
	return nil
}
