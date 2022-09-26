package services

import (
	"fmt"
	// "sync"
	"time"

	// "github.com/gin-gonic/gin"
)

// var wg sync.WaitGroup

func RunTimer() {
	for {
		fmt.Println("mengupdate data...")
		go UpdateFileWithNewValue()
		TestFileAfterUpdate("./status.json")
		time.Sleep(time.Second * 15);
		fmt.Println("hasil update")
	}
}