package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"
)

type apps_json struct {
	Resources []resources `json:"resources"`
}
type resources struct {
	Metadata metadata `json:"metadata"`
	Entity   entity   `json:"entity"`
}

type metadata struct {
	Guid string `json:"guid"`
}
type entity struct {
	Env env `json:"environment_json"`
}

type env struct {
	AppID string `json:"app_id"`
}

func main() {
	http.HandleFunc("/", start)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func start(res http.ResponseWriter, req *http.Request) {
	time.Sleep(5 * time.Second)

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

	out, err = exec.Command(cd+"/bin/cf", "api", "api.bosh-lite.com", "--skip-ssl-validation").Output()
	fmt.Fprintln(res, "Output is:", string(out))
	if err != nil {
		fmt.Println("Error running cf command2", err)
		os.Exit(1)
	}

	out, err = exec.Command(cd+"/bin/cf", "auth", "admin", "admin").Output()
	fmt.Fprintln(res, "Output is:", string(out))
	if err != nil {
		fmt.Println("Error running cf command3", err)
		os.Exit(1)
	}

	out, err = exec.Command(cd+"/bin/cf", "create-org", "temp").Output()
	fmt.Fprintln(res, "Output is:", string(out))
	if err != nil {
		fmt.Println("Error running cf command4", err)
		os.Exit(1)
	}

	out, err = exec.Command(cd+"/bin/cf", "create-space", "temp", "-o", "temp").Output()
	fmt.Fprintln(res, "Output is:", string(out))
	if err != nil {
		fmt.Println("Error running cf command5", err)
		os.Exit(1)
	}

	out, err = exec.Command(cd+"/bin/cf", "target", "-o", "temp", "-s", "temp").Output()
	fmt.Fprintln(res, "Output is:", string(out))
	if err != nil {
		fmt.Println("Error running cf command6", err)
		os.Exit(1)
	}

	out, err = exec.Command(cd+"/bin/cf", "curl", "v2/apps").Output()
	fmt.Fprintln(res, "Output is:", string(out))
	if err != nil {
		fmt.Println("Error running cf command7", err)
		os.Exit(1)
	}

	var apps apps_json
	err = json.Unmarshal(out, &apps)
	if err != nil {
		fmt.Println("Error unmarshaling", err)
		os.Exit(1)
	}

	var guid string
	for _, app := range apps.Resources {
		if app.Entity.Env.AppID == "this_is_not_a_test" {
			guid = app.Metadata.Guid
			break
		}
	}

	out, err = exec.Command(cd+"/bin/cf", "curl", "v2/apps/"+guid, "-X", "PUT", "-d", `'{"health_check_timeout":2}'`).Output()
	fmt.Fprintln(res, "Output is:", string(out))
	if err != nil {
		fmt.Println("Error running cf command8", err)
		os.Exit(1)
	}

	os.Exit(0)
}
