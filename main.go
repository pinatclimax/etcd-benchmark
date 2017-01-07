package main

import (
	"errors"
	"fmt"
	"runtime"

	"sync"

	"time"

	"os"

	"strconv"
	"climax.com/etcd-benchmark/etcd"

)

var macPrefix = "11:11:11:"
var panelTopic = "panel"
var userTopic = "user"
var ffffffNum = 16777215

func main() {

	runtime.GOMAXPROCS(4)

	testLimit, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	etcdOper := os.Args[2]
	
	var wg sync.WaitGroup
	wg.Add(testLimit)
	for i := 1; i <= testLimit; i++ {
		time.Sleep(1000)
		switch etcdOper {
			case "select":
			readPanelMac(i, &wg)
			
			case "upsert":
			writePanelMac(i, &wg)

			default:
			readPanelMac(i, &wg)

		}
		
	}

	wg.Wait()
}

func readPanelMac(num int, wg *sync.WaitGroup) {
	mac, err := numberToMac(num)
	etcd.Select("/mqtt/panel/" + mac)

	if err != nil {
		fmt.Println(err)
	}

	wg.Done()
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
	fmt.Println(num, "11:11:11:"+postmac)
	return "11:11:11:" + postmac, nil
}
