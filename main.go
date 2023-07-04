package main

import "HackDayBackend/global"

func main() {
	if err := global.Set; err != nil {
		panic(err)
	}
}
