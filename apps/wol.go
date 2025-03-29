package apps

import (
	"log"
	"os/exec"
)

func WOL(mac string) {
	// Define the command and arguments
	cmd := exec.Command("wol", "wake", mac)

    // Capture the combined output (stdout and stderr)
    output, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatalf("Error running wol command: %v\nOutput: %s", err, string(output))
    }

	log.Printf("wol: command executed successfully!\n\n%s\n", string(output))
}