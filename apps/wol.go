package apps

import (
	"log"

	"github.com/linde12/gowol"
)

func WOL(mac string) {

	if packet, err := gowol.NewMagicPacket(mac); err == nil {
		packet.Send("255.255.255.255")          // send to broadcast
		packet.SendPort("255.255.255.255", "7") // specify receiving port
	} else {
        log.Fatalf("Error running wol command: %v\n", err)
	}

	log.Printf("wol: command executed successfully!\n\n")
}