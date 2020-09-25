package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/wisdommatt/grpc/proto/person"
)

func personGrpc() {
	wisdom := &person.Person{
		Name: "Wisdom",
		Age:  20,
		SocialFollowers: &person.SocialFollowers{
			Youtube: 10000,
			Twitter: 3500,
		},
	}

	data, err := proto.Marshal(wisdom)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

	newWisdom := &person.Person{}
	err = proto.Unmarshal(data, newWisdom)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(newWisdom.GetName())
}
