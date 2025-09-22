package main

import (
	"os/exec"
)

type Commands []*exec.Cmd

func (c Commands) Execute() error {
	for _, v := range c {
		err := v.Run()
		if err != nil {
			return err
		}
	}

	return nil
}

func (c Commands) ExecuteInDirectory(dir string) error {
	for _, v := range c {
		v.Dir = dir
		err := v.Run()
		if err != nil {
			return err
		}
	}

	return nil
}
