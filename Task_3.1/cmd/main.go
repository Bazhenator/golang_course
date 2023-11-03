package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	calc "task_3.1/pkg"
)

func main() {
	var n int
	var err error

	switch len(os.Args) {
	case 2:
		n, err = strconv.Atoi(os.Args[1])
	case 3:
		n, err = strconv.Atoi(os.Args[2])
	default:
		log.Fatal("invalid amount of cml args")
	}
	if err != nil {
		err = errors.New("bad input. please enter the integer as last cml arg")
		log.Fatal(err)
	}

	logFlag := flag.Bool("log", false, "enable logging")
	flag.Parse()
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	//if *logFlag {
	//	logger.SetOutput(os.Stdout)
	//} else {
	//	logger.SetOutput(io.Discard)
	//}

	logger.Info(calc.Calculate(int64(n), *logFlag))
}
