package echomessage

import (
	"fmt"
	"log"

	"github.com/heckdevice/goactorframework-corelib"
	"github.com/heckdevice/goactorframework-examples/samples/common"
)

const (
	// ActorType - actor type for this actor
	ActorType = "GreetingActor"
	// MessageTypeHI - supported messageType for this actor
	MessageTypeHI = "HI"
	// MessageTypeBYE - supported messageType for this actor
	MessageTypeBYE = "BYE"
)

// InitActor - Initialises this actor by registering its different message handlers and spawing the actor using the Default actor system
func InitActor() {
	greetingActor := core.Actor{ActorType: ActorType}
	err := core.GetDefaultActorSystem().RegisterActor(&greetingActor, MessageTypeHI, greetHI)
	if err != nil {
		log.Panic(fmt.Sprintf("Error while registering actor %v. Details : %v", greetingActor.ActorType, err.Error()))
	}
	greetingActor.RegisterMessageHandler(MessageTypeBYE, greetBye)
	greetingActor.RegisterMessageHandler(common.ConsolePrint, consolePrint)
	go greetingActor.SpawnActor()
}

func greetHI(message core.Message) {
	fmt.Print(fmt.Sprintf("Hi there %v, i got %v", message.Sender.ActorType, message.Payload))
}

func greetBye(message core.Message) {
	fmt.Print(fmt.Sprintf("Adios %v !!!", message.Sender.ActorType))
}

func consolePrint(message core.Message) {
	fmt.Print(fmt.Sprintf("Echo : %v", message))
}
