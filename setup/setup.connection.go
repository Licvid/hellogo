package setup

import (
	"bufio"
	"fmt"
	"os"

	"../docker"
	"../utils/color"
)

type connector struct {
}

func Pin() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(color.BCyan("1. Please type your connection PIN code or Email: "))
	connectionKey, _ := reader.ReadString('\n')
	fmt.Println()
	fmt.Println(connectionKey)

	docker.InstallCore()
}
