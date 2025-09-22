package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

var uname string
var configFile string

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("done!")
}

func run() error {
	var err error

	configFile, err = getConfigPath()
	if err != nil {
		return err
	}

	uname, err = getUname()
	if err != nil {
		return err
	}

	if len(os.Args) < 2 {
		return push()
	}

	action := os.Args[1]

	switch action {
	default:
		return push()
	case "init":
		return initProject()
	case "fuck", "redo":
		return redo()
	}
}

func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return home + "/.config/ggit", nil

}

func getUname() (string, error) {
	contents, errExists := os.ReadFile(configFile)
	if errExists != nil {
		_, err := os.Create(configFile)
		if err != nil {
			return "", nil
		}
		username, err := promptUser("Github username:")
		if err != nil {
			return "", err
		}

		err = os.WriteFile(configFile, username, 0755)
		if err != nil {
			return "", err
		}

		fmt.Println("username set!")
		return string(username), nil
	}

	return string(contents), nil

}

func initProject() error {
	projectName, err := getArg(2, "Project Name:")
	if err != nil {
		return err
	}

	exec.Command("mkdir", "-p", projectName).Run()

	commands := Commands{
		exec.Command("go", "mod", "init", "github.com/"+uname+"/"+projectName),
		exec.Command("git", "init"),
		exec.Command("git", "remote", "add", "origin", "git@github.com:"+uname+"/"+projectName),
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	dir := wd + "/" + projectName
	return commands.ExecuteInDirectory(dir)
}

func push() error {
	commitName, err := getArg(1, "Commit Name:")
	if err != nil {
		return err
	}
	fmt.Println("commit name", commitName)

	commands := Commands{
		exec.Command("git", "add", "--all"),
		exec.Command("git", "commit", "-m", commitName),
		exec.Command("git", "push", "--set-upstream", "origin", "master"),
	}

	return commands.Execute()
}

func redo() error {
	fmt.Println("redo!")
	commands := Commands{
		exec.Command("git", "add", "--all"),
		exec.Command("git", "commit", "--amend", "--no-edit"),
		exec.Command("git", "push", "--force"),
	}

	return commands.Execute()
}

func getArg(n int, prompt string) (string, error) {
	if len(os.Args) > 2 {
		return os.Args[2], nil
	}

	arg, err := promptUser(prompt)
	if err != nil {
		return "", nil
	}

	return string(arg), nil
}

func promptUser(prompt string) ([]byte, error) {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	if err != nil {
		return nil, err
	}

	return line, nil
}
