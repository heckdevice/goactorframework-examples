package printmessage

import (
	"fmt"
	"log"

	"github.com/heckdevice/goactorframework-corelib"
	"github.com/heckdevice/goactorframework-examples/samples/common"
)

const (
	// ActorType - actor type for this actor
	ActorType = "PrintActor"
)

// InitActor - Initialises this actor by registering its different message handlers and spawing the actor using the Default actor system
func InitActor() {
	printActor := core.Actor{ActorType: ActorType}
	err := core.GetDefaultActorSystem().RegisterActor(&printActor, common.ConsolePrint, consolePrint)
	if err != nil {
		log.Panic(fmt.Sprintf("Error while registering actor %v. Details : %v", printActor.ActorType, err.Error()))
	}
	go printActor.SpawnActor()
}

func consolePrint(message core.Message) {
	fmt.Print(fmt.Sprintf("Got Message %v", message))
}
