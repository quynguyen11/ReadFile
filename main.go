package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"gopkg.in/yaml.v1"
)

// Go provides a `flag` package supporting basic
// command-line flag parsing. We'll use this package to
// implement our example command-line program.

func main() {

	File, err := os.Open("application.license.yaml")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, err := ioutil.ReadAll(File)
	if err != nil {
		fmt.Println(err)
	}

	setting := &Application{}
	yaml.Unmarshal(byteValue, setting)

	a := CheckExpiration(setting.ExpirationDate)
	if a == false {

		fmt.Println("Bad")
		return

	}

	{
		yamlFile, err := os.Open("test.yaml")
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successfully Opened test.yaml")

		byteValu, _ := ioutil.ReadAll(yamlFile)

		st := &Config{}

		yaml.Unmarshal(byteValu, st)

		fmt.Println(st.HostService)
	}

	b := CreateLicense("2020-10-01T11:45:26.371Z", "2020-10-01T11:45:26.371Z")

	print(b)
	/* ---------------------------------- */
	// Open our jsonFile

}

/*
Config object
*/
type Config struct {
	HostService domain `yaml:"host_service"`
}

type domain struct {
	Bind string `yaml:"bind"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

/*
Application object
*/
type Application struct {
	DateCreated    string `yaml:"date_created" json:"date_created" `
	ExpirationDate string `yaml:"expiration_date" json:"expiration_date" `
}

/*
CheckExpiration object
*/
func CheckExpiration(timeExp string) bool {

	sDec, err := base64.StdEncoding.DecodeString(timeExp)
	if err != nil {
		fmt.Println(err)
		return false
	}

	now := time.Now()
	t, err := time.Parse(time.RFC3339, string(sDec))
	if err != nil {
		fmt.Println(err)
		return false
	}

	checktime := now.After(t)
	if checktime == true {

		return false

	}
	return true
}

/*
CreateLicense object
*/
func CreateLicense(cr string, exp string) error {

	created := base64.StdEncoding.EncodeToString([]byte(cr))

	expiration := base64.StdEncoding.EncodeToString([]byte(exp))

	defaut := &Application{
		DateCreated:    created,
		ExpirationDate: expiration,
	}

	body, _ := json.Marshal(defaut)

	// json.Unmarshal(body, &st)

	err := ioutil.WriteFile("license.json", body, 0664)
	if err != nil {

		return err
	}
	return nil

}
