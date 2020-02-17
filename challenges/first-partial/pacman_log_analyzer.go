package main

import (
	//	"bufio"
	//	"io"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type pack struct {
	packageName string
	installdate string
	LastUpdate  string
	numUpdates  int
	remDate     string
}

type packmanager struct {
	insPackages   int
	rmdPackages   int
	upgPackages   int
	currInstalled int
	packagesMap   map[string]*pack
}

func (packageManager *packmanager) Analyzecommand(command []string) {

	switch command[3] {

	case "installed":
		var temporaryPackage = pack{command[4], strings.TrimRight(string(command[0][1:]+" "+command[1]), "]"), strings.TrimRight(string(command[0][1:]+" "+command[1]), "]"), 0, "_"}
		packageManager.packagesMap[command[4]] = &temporaryPackage
		packageManager.insPackages++
		packageManager.currInstalled++
	case "upgraded":
		packageManager.packagesMap[command[4]].LastUpdate = strings.TrimRight(string(command[0][1:]+" "+command[1]), "]")
		if packageManager.packagesMap[command[4]].numUpdates == 0 {
			packageManager.upgPackages++
		}
		packageManager.packagesMap[command[4]].numUpdates++
	case "reinstalled":
		packageManager.packagesMap[command[4]].remDate = "_"
		packageManager.currInstalled++
		packageManager.rmdPackages--
	case "removed":
		packageManager.packagesMap[command[4]].remDate = strings.TrimRight(string(command[0][1:]+" "+command[1]), "]")
		packageManager.currInstalled--
		packageManager.rmdPackages++
	}

}
func (packageManager *packmanager) CreateReport(filepath string) {
	t, err2 := os.Create(filepath)
	CheckError(err2)
	fmt.Fprintln(t, "Pacman Packages Report")
	fmt.Fprintln(t, "----------------------")
	fmt.Fprintln(t, "- Installed packages : ", packageManager.insPackages)
	fmt.Fprintln(t, "- Removed packages   : ", packageManager.rmdPackages)
	fmt.Fprintln(t, "- Upgraded packages  : ", packageManager.upgPackages)
	fmt.Fprintln(t, "- Current installed  : ", packageManager.currInstalled)
	fmt.Fprintln(t, "  ")
	fmt.Fprintln(t, "List of packages")
	fmt.Fprintln(t, "----------------")
	for _, s := range packageManager.packagesMap {
		fmt.Fprintln(t, "- Package Name        : "+s.packageName)
		fmt.Fprintln(t, "  - Install date      : ", s.installdate)
		fmt.Fprintln(t, "  - Last update date  : ", s.LastUpdate)
		fmt.Fprintln(t, "  - How many updates  : ", s.numUpdates)
		fmt.Fprintln(t, "  - Removal date      : ", s.remDate)
		fmt.Fprintln(t, "  ")

	}
	t.Close()

}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	fmt.Println("Pacman Log Analyzer")
	if len(os.Args) < 2 {
		fmt.Println("You must send at least one pacman log file to analize")
		fmt.Println("usage: ./pacman_log_analizer <logfile>")
		os.Exit(1)
	}
	dat, err := ioutil.ReadFile(os.Args[1])
	CheckError(err)
	logs := strings.Split(string(dat), "\n")
	var packageManager = &packmanager{0, 0, 0, 0, make(map[string]*pack)}
	for _, commandLine := range logs {
		command := strings.Split(string(commandLine), " ")
		packageManager.Analyzecommand(command)
	}
	packageManager.CreateReport("packages_report.txt")
}
