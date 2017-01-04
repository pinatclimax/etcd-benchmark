package main

import (
	"errors"
	"fmt"

	"sync"

	"climax.com/mqtt.sa/etcd"
)

var macPrefix = "11:11:11:"
var panelTopic = "panel"
var userTopic = "user"
var ffffffNum = 16777215

func main() {

	var wg sync.WaitGroup
	wg.Add(100000)
	for i := 1; i <= 100000; i++ {
		go writePanelMac(i, &wg)
	}

	wg.Wait()
}

func writePanelMac(num int, wg *sync.WaitGroup) {
	mac, err := numberToMac(num)
	etcd.Upsert("/mqtt/panel/"+mac, "undefined")

	if err != nil {
		fmt.Println(err)
	}

	wg.Done()
}

func numberToMac(num int) (string, error) {
	if num > 16777215 {
		return "", errors.New("number is greater than 16777215")
	}
	hexnum := fmt.Sprintf("%06x", num)
	postmac := fmt.Sprintf("%s:%s:%s", hexnum[0:2], hexnum[2:4], hexnum[4:6])
	return "11:11:11:" + postmac, nil
}
