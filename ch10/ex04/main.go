package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type packageInfo struct {
	ImportPath string   `json:"ImportPath"`
	Name       string   `json:"Name"`
	Deps       []string `json:"Deps"`
}

func main() {
	if len(os.Args) == 1 {
		log.Fatalln("package specification is required")
	}
	initialList, err := getPacakgeList(os.Args[1:]...)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	allList, err := getPacakgeList("...")
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	showDependentPackages(initialList, allList)
}
func getPacakgeList(packages ...string) ([]*packageInfo, error) {
	args := []string{"list", "-e", "-json"}
	args = append(args, packages...)
	cmd := exec.Command("go", args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	defer stdout.Close()

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}
	go func() {
		io.Copy(ioutil.Discard, stderr)
		stderr.Close()
	}()
	decoder := json.NewDecoder(stdout)
	var pInfos []*packageInfo
	for {
		var pInfo packageInfo
		err := decoder.Decode(&pInfo)
		if err != nil {
			if err != io.EOF {
				log.Printf("%v\n", err)
			}
			return pInfos, nil
		}
		pInfos = append(pInfos, &pInfo)
	}
}
func showDependentPackages(initial, all []*packageInfo) {
	initialNames := make([]string, 0, len(initial))
	for _, pInfo := range initial {
		initialNames = append(initialNames, pInfo.ImportPath)
	}
	fmt.Printf("Specified Packages : %s\n", strings.Join(initialNames, " "))
	var deps []string
	for _, pInfo := range all {
		if !isDependent(pInfo, initialNames) {
			continue
		}
		deps = append(deps, pInfo.ImportPath)
	}
	if len(deps) != 0 {
		fmt.Printf("Dependet Packages :%s\n", strings.Join(deps, " "))
	}
}

func isDependent(pInfo *packageInfo, names []string) bool {
topLoop:
	for _, name := range names {
		for _, deps := range pInfo.Deps {
			if deps == name {
				continue topLoop
			}
		}
		return false
	}
	return true
}
