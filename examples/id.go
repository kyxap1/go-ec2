package main

import "github.com/mitchellh/goamz/aws"
import "github.com/segmentio/go-ec2"
import "encoding/json"
import "os"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	auth, err := aws.EnvAuth()
	check(err)

	hosts := ec2.New(auth, aws.USWest2)
	check(err)

	nodes, err := hosts.Id("i-2c6ceb18", "i-3792b43c")
	check(err)

	b, err := json.MarshalIndent(nodes, "", "  ")
	check(err)

	os.Stdout.Write(b)
}
