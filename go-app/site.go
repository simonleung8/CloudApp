package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	cd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current director:", err)
		os.Exit(1)
	}

	fmt.Println("Current dir:", cd+"/bin")
	err = exec.Command("chmod", "+x", cd+"/bin").Run()
	if err != nil {
		fmt.Println("Error setting bin/ permission", err)
		os.Exit(1)
	}

	err = exec.Command("chmod", "+x", cd+"/bin/cf").Run()
	if err != nil {
		fmt.Println("Error setting cf permission", err)
		os.Exit(1)
	}

	out, err := exec.Command(cd+"/bin/cf", "--version").Output()
	fmt.Fprintln(res, "Output is:", string(out))
	if err != nil {
		fmt.Println("Error running cf command1", err)
		os.Exit(1)
	}

	out, err = exec.Command(cd+"/bin/cf", "api", "api.ng.bluemix.net").Output()
	fmt.Fprintln(res, "Output is:", string(out))
	if err != nil {
		fmt.Println("Error running cf command2", err)
		os.Exit(1)
	}

	out, err = exec.Command(cd+"/bin/cf", "auth", os.Getenv("user"), os.Getenv("pass")).Output()
	fmt.Fprintln(res, "Output is:", string(out))
	if err != nil {
		fmt.Println("Error running cf command3", err)
		os.Exit(1)
	}

}
