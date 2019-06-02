package main

import (
	"./gohunspell"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main()  {

	file, err := os.Open("./gohunspell/include/dictionaries/fr_FR.aff")
	if err != nil {
		os.Exit(0)
	}
	aff, e := ioutil.ReadAll(file)
	if e != nil {
		os.Exit(0)
	}
	file.Close()
	file, err = os.Open("./gohunspell/include/dictionaries/fr_FR.dic")
	if err != nil {
		os.Exit(0)
	}
	dic, e2 := ioutil.ReadAll(file)
	file.Close()
	if e2 != nil {
		os.Exit(0)
	}

	spell := Gohunspell.NewGohunspell(aff, dic)
	t1 := time.Now() // get current time
	count, list := spell.Stem("allons")
	elapsed := time.Since(t1)
	fmt.Println(count)
	fmt.Println(list)
	fmt.Println("App elapsed: ", elapsed)
}
