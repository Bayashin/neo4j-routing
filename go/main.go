package main

import (
	"fmt"

	"github.com/Bayashin/neo4j-routing/go/handler"
)

func main() {
	fmt.Println(handler.Neo4jConnect("10号館", "AITプラザ"))
}
