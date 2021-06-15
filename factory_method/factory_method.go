package main

import (
	"fmt"
)

// Product Interface
type ITransport interface {
	getName() string
	getTrack() string
	getContainerName() string
	deliver()
}

// Concrete Product
type transport struct {
	name          string
	track         string
	containerName string
}

func (t *transport) getName() string {
	return t.name
}

func (t *transport) getTrack() string {
	return t.track
}

func (t *transport) getContainerName() string {
	return t.containerName
}

func (t *transport) deliver() {
	fmt.Printf("Deliver with %s by %s i a %s", t.name, t.track, t.containerName)
}

// Concrete Product
type truck struct {
	transport
}

func newTruck() ITransport {
	return &truck{
		transport: transport{
			name:          "truck",
			track:         "land",
			containerName: "box",
		},
	}
}

// Concrete Product
type ship struct {
	transport
}

func newShip() ITransport {
	return &ship{
		transport: transport{
			name:          "ship",
			track:         "sea",
			containerName: "container",
		},
	}
}

// Factory/Creator
func getTransport(transportName string) (ITransport, error) {
	switch transportName {
	case "truck":
		return newTruck(), nil
	case "ship":
		return newShip(), nil
	default:
		return nil, fmt.Errorf("transport %s unavailable", transportName)
	}
}

// Client
func main() {
	transport, err := getTransport("truck")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("========Transport 1=======")
		fmt.Println("Transport name: ", transport.getName())
		fmt.Println("Transport track: ", transport.getTrack())
		fmt.Println("Transport container: ", transport.getContainerName())
		transport.deliver()
	}

	fmt.Println()

	transport2, err := getTransport("ship")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("========Transport 2=======")
		fmt.Println("Transport name: ", transport2.getName())
		fmt.Println("Transport track: ", transport2.getTrack())
		fmt.Println("Transport container: ", transport2.getContainerName())
		transport2.deliver()
	}
}
