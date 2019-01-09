package main

import (
	"github.com/DyingLight123/RedisInfluxdb"
	"time"
)

func main() {
	/*pause := make(chan string, 1)
	defer close(pause)*/

	x := &RedisInfluxdb.RedisInfluxdb{"127.0.0.1:6379", "", "data",
	"http://127.0.0.1:8086", "admin", "admin", "test",
	"redis"}

	/*err := x.RedisWriteToInfluxdb()
	if err != nil {
		log.Fatal(err)
	}*/

	/*field, err := x.GetRedis()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(field)
	results, err := x.GetInfluxdb()
	if err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(*result)
	}*/

	go x.RefreshRedis(100)
	time.Sleep(20 * time.Second)
	x.PauseRedis()
	//time.Sleep(time.Second)
	go x.RefreshRedis(200)
	time.Sleep(20 * time.Second)
	x.PauseRedis()
	//time.Sleep(time.Second)
	go x.RefreshRedis(100)
	time.Sleep(10 * time.Second)
	x.PauseRedis()
	time.Sleep(time.Second)

	/*go x.RefreshRedis1(200, pause)
	time.Sleep(21 * time.Second)
	//fmt.Println("pause")
	pause <- "continue"
	go x.RefreshRedis1(300, pause)
	time.Sleep(21 * time.Second)
	pause <- "continue"
	go x.RefreshRedis1(100, pause)
	time.Sleep(11 * time.Second)
	pause <- "continue"
	time.Sleep(time.Second)
*/

	/*snmpifo, err := x.GetSnmpIfo("localhost", "161", ".1.3.6.1.2.1.2.2.1.10")
	if err != nil {
		log.Fatal(err)
	}
	for a, b := range *snmpifo {
		fmt.Println(a + ":", b)
	}*/

	return
}
