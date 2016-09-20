package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func setup() {
	cd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current director:", err)
		os.Exit(1)
	}

	out, err := exec.Command(cd+"/bin/cf", "--version").Output()
	if err != nil {
		fmt.Println("Error running cf command1", err)
		os.Exit(1)
	}

	out, err = exec.Command(cd+"/bin/cf", "login", "-a", "api.bosh-lite.com", "-u", "admin", "-p", "admin", "--skip-ssl-validation").Output()
	if err != nil {
		fmt.Println("Error running cf command2", err)
		os.Exit(1)
	}

	out, err = exec.Command(cd+"/bin/cf", "curl", "v2/apps").Output()
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
	var health_check_timeout int
	for _, app := range apps.Resources {
		if app.Entity.Env.AppID == "this_is_not_a_test" {
			guid = app.Metadata.Guid
			health_check_timeout = app.Entity.HealthCheckTimeout
			break
		}
	}

	if health_check_timeout == 0 {
		out, err = exec.Command(cd+"/bin/cf", "curl", "v2/apps/"+guid, "-X", "PUT", "-d", `'{"health_check_timeout":2}'`).Output()
		if err != nil {
			fmt.Println("Error running cf command8", err)
			os.Exit(1)
		}

		os.Exit(0)
	}
}
