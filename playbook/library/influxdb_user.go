package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strings"

	influxdbClient "github.com/influxdb/influxdb/client"
)

func success(response interface{}) {
	if outputJSON, err := json.Marshal(response); err != nil {
		fail(err)
	} else {
		fmt.Printf("%s", outputJSON)
	}
	os.Exit(0)
}

func fail(err error) {
	ansibleResponse := map[string]interface{}{
		"failed": true,
		"msg":    err.Error(),
	}
	s, _ := json.Marshal(ansibleResponse)
	fmt.Printf("%s\n", s)
	os.Exit(1)
}

func parseAnsibleArguments(path string) map[string]string {
	if len(os.Args) < 2 {
		fail(errors.New("No argument file receive"))
	}

	conf := make(map[string]string)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fail(err)
	}
	for _, arg := range strings.Split(string(data), " ") {
		if strings.Contains(arg, "=") {
			line := strings.Split(arg, "=")
			conf[line[0]] = line[1]
		}
	}

	return conf
}

func buildInfluxDbClient(args map[string]string) influxdbClient.Client {
	parsedURL, _ := url.Parse("http://localhost:8086")

	loginConfig := influxdbClient.Config{
		URL:      *parsedURL,
		Username: "",
		Password: "",
	}

	if loginURL, present := args["login_url"]; present {
		parsedURL, err := url.Parse(loginURL)
		if err != nil {
			fail(err)
		}
		loginConfig.URL = *parsedURL
	}

	if loginURL, present := args["login_user"]; present {
		loginConfig.Username = loginURL
	}

	if loginPassword, present := args["login_password"]; present {
		loginConfig.Password = loginPassword
	}

	client, err := influxdbClient.NewClient(loginConfig)
	if err != nil {
		fail(err)
	}

	_, _, err = client.Ping()
	if err != nil {
		fail(err)
	}

	return *client
}

func queryDB(client *influxdbClient.Client, database string, cmd string) (results []influxdbClient.Result, err error) {
	q := influxdbClient.Query{
		Command:  cmd,
		Database: database,
	}
	response, err := client.Query(q)
	if err == nil {
		if response.Error() != nil {
			return response.Results, response.Error()
		}
	}
	return response.Results, nil
}

func userExists(client *influxdbClient.Client, user string) bool {
	results, err := queryDB(client, "", "SHOW USERS")
	if err != nil {
		fail(err)
	}

	users := make(map[string]bool)
	for _, userRow := range results[0].Series[0].Values {
		users[userRow[0].(string)] = true
	}

	return users[user]
}

func noResultQuery(client *influxdbClient.Client, query string) {
	_, err := queryDB(client, "", query)
	if err != nil {
		fail(err)
	}
}

func createUser(client *influxdbClient.Client, user string, password string, privledges string, database string) {
	userExists := userExists(client, user)

	if !userExists {
		createUserQuery := fmt.Sprintf("CREATE USER %s WITH PASSWORD '%s'", user, password)
		noResultQuery(client, createUserQuery)

		var grantPrivledgesQuery string
		if database == "" {
			grantPrivledgesQuery = fmt.Sprintf("GRANT %s TO %s", privledges, user)
		} else {
			grantPrivledgesQuery = fmt.Sprintf("GRANT %s ON %s TO %s", privledges, database, user)
		}

		noResultQuery(client, grantPrivledgesQuery)
	}

	ansibleResponse := map[string]interface{}{
		"changed": !userExists,
	}

	success(ansibleResponse)
}

func main() {
	args := parseAnsibleArguments(os.Args[1])
	client := buildInfluxDbClient(args)

	if user, userPresent := args["user"]; userPresent {

		password, passwordPresent := args["password"]
		if !passwordPresent {
			fail(errors.New("Password not present."))
		}

		privledges, privledgesPresent := args["privledges"]
		if !privledgesPresent {
			privledges = "ALL"
		}

		database, databasePresent := args["database"]
		if !databasePresent {
			database = ""
		}

		createUser(&client, user, password, privledges, database)
	}

	fail(errors.New("User not present."))
}
