package DemoService

import (
	"blog_service/App/Struct/DemoStruct"
	"context"
	"fmt"
	"testing"
)

func TestGetIndex(t *testing.T) {
	fmt.Println(1111)

	srv := NewDemoService()

	request := DemoStruct.IdxRequest{Id: 0}

	response, err := srv.GetIndex(context.Background(), request)
	if err != nil {
		return
	}

	fmt.Printf("ret: %+v\n", response)
}
