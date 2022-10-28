package Helpers

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

type TestInterface interface {
	A() string
	B() int64
}

type SystemTools struct {
	Logger *logrus.Logger
	Conf   *ini.File
	Funcs  Funcs
}

func (s SystemTools) A() {
	fmt.Println("A()")
}
func (s SystemTools) B() {
	fmt.Println("A()")
}

func NewSystemTools() *SystemTools {
	var logger = logrus.New()
	logger.Hooks.Add(NewContextHook())

	conf, err := ini.Load("env.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	s := SystemTools{
		Logger: logger,
		Conf:   conf,
	}

	return &s
}
