package annotation

import (
	"fmt"
	"github.com/choutsugi/freedom/examples/pb/player"
	"google.golang.org/protobuf/proto"
	"reflect"
	"testing"
)

func TestAnnotation(t *testing.T) {
	anno := Annotation{Prototypes: map[uint16]reflect.Type{}}

	anno.Register(01, &player.CS_Login{})

	data, _ := proto.Marshal(&player.CS_Login{
		Username: "tsugi",
		Password: "123456",
	})

	msg := anno.UnMarshal(01, data)
	fmt.Println(msg)
}
