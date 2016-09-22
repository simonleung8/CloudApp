package main

import (
	"encoding/json"
	"os"
	"os/exec"
)

func setup() {
	cd, err := os.Getwd()
	catch(err, "Error getting current directory")

	catch(exec.Command(cd+"/bin/cf", "--version").Run(), "cf --version")

	catch(exec.Command(cd+"/bin/cf", "login", "-a", "api.bosh-lite.com", "-u", "admin", "-p", "admin", "--skip-ssl-validation").Run(), "cf login")

	out, err := exec.Command(cd+"/bin/cf", "curl", "v2/apps").Output()
	catch(err, "Error running cf command7")

	var apps apps_json
	catch(json.Unmarshal(out, &apps), "Error unmarshaling")

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
		catch(err, "Error running cf command8")

		os.Exit(0)
	}
}

func catch(err error, comment string) {
	if err != nil {
		panic(err.Error() + " - " + comment)
	}
}
