package main

import (
	"consumer/shared"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go shared.ContainerItem.Service.Consume(&wg)
	wg.Wait()
}
