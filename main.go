package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
	"uploadFile/crypto"

	"github.com/boltdb/bolt"
	"gopkg.in/yaml.v1"
)

const (

	//RSAPublicKey  is
	RSAPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzjJOT5mxrUslWOHJGAe6g0fDAaky5K7t23rDscUeh4viE2ayH4rBW++Yzup6ON6UPqDDY2QBYI9EGvX7EawU5Fo/WwwQUzlcKrzV5lhB2Sb38rk6qjsq+1GgS6JJPYKiF8/Ib3yXgpS5jXo5dZ4N4LpKuU5oSKzhvccS/7h9nv0mcA0PcOGYU4fsmQKXZnbgTzNakc++FPB3XwQJi167lkd71gtxdVrtpAhr2tj9qYLnlXibjXauoK+VKTuZIdNWKrhLI6bYLpIcD7qv6pCQQ6I96pKBz9yPCUWsdD7qOXXtqHtWalzID3Bh5RemK8wCXmhHKmX7LG0gSzWOrQhbnQIDAQAB"

	//RSAPrivateKey  is
	RSAPrivateKey = "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDOMk5PmbGtSyVY4ckYB7qDR8MBqTLkru3besOxxR6Hi+ITZrIfisFb75jO6no43pQ+oMNjZAFgj0Qa9fsRrBTkWj9bDBBTOVwqvNXmWEHZJvfyuTqqOyr7UaBLokk9gqIXz8hvfJeClLmNejl1ng3gukq5TmhIrOG9xxL/uH2e/SZwDQ9w4ZhTh+yZApdmduBPM1qRz74U8HdfBAmLXruWR3vWC3F1Wu2kCGva2P2pgueVeJuNdq6gr5UpO5kh01YquEsjptgukhwPuq/qkJBDoj3qkoHP3I8JRax0Puo5de2oe1ZqXMgPcGHlF6YrzAJeaEcqZfssbSBLNY6tCFudAgMBAAECggEBALmA6JMUYpTVFjRwbMoBqfqRhQ7pea/i8HqUZ5p3DJBFeI6bxYQ7ANoFJCSDLpIbLKNrXlz8i4CjY1IeGlI7zk/iIT14DPfSkuigGo+wtwz10fu8SNr9gB25wcxJpDqCW/GwLVKfdG81/fkaDLeUCkgzUSaFM6yuXCiwAJevEtUEq2VDSUHXQwdY6FpNS75A/mRsTNWCmqJBDJVfPCdzXlBjdWqeTL4BKaLKMM90gbnH0P6hoJK6Mh+NnenazPu/6q8ICJJ9xUlVzOqdsdfQJ5/gubZCBWpbjMx48yv++1hB02r5MdPwpZTE4Ej3adu0J0E8WoTVMdqFpJBuKHipnsECgYEA/Gwx3RigPZWYAPZHkb4WRd2YnbQ1FBLSOqmC8aIOit3gGc+SVvh4pqdJlg+vE2DsWpRgRwYUc0l+9nDGSTxHXsG9LmEpL76rh8LvUbro1hdki+SWWBh+WEEACXUjdcXlARr1YH1XXNbLoIi3d+d607mgT5RIfehY+cXP0cGlny8CgYEA0R5mWmg7OEpa+e7Gzk8NUoSX0ISdTakTRAuG/eZ3jiRsXA8kCgoOxmEvA8+WNw5mSo/qadZlDx1kLai4Ul6CSDWPidpKs7rhjN1i9/h8lpqGQDMo2frAIwSMOHnsHQaUI3Lcdzk4TLM21J0WzqlWGg6RzF7RUfyhUbSgMZnRXvMCgYEAuGVFS9FIhZR5NQK0J5hn5uPJMDNLrv1MzAO2n1OWMgWBRvmmWpgqcvuzusZ8S7i7EDRh3KBpYgqnj9m0UB2TuXnn/DCICNPNtGBHuTnEC1mNXtA+r948tbXOFBqZK9jDwLnz1Gfb4PscR4p4FERqKq7omBmnlyqbjOLfPMisd4sCgYA4Cadv7qJ/8Rz0ANJxkqmFRVbRX9gvaXFqOJSSEWJUStpmyP/lWNCgxIYuxUABvPAYZvxwJC2soTmpKp9KI11SMFgonsCJ7Thn4SOWQ5ZPXVVAevUlhJZcS+mvcKyfEpY78Zm2sTSvCQ9WZkooUyRpkyHq3DXHfKVWGcsbv5ZllQKBgHpZBFU7J+h5BmWZdUdJyIZE5+4ZJomf8kf76vfe8kNMUg3kIsFv/LTa95dLwbk8lIMVbk1aIJXimyCgDHHbZ5ySwrG47dnFPt77MhJBZ/7nBVczgyleA80v/fx87MOQUc+5kOAJSNX9Cy2WIXOhHvL3KUoSepP2yDTM/aiVpnXo"
)

var (
	q []byte
	w []byte
)

func main() {

	// not use BoltDB. Create lincense.yaml
	b := CreateLicense("2020-08-01T11:45:26.371Z", "2020-12-12T11:45:26.371Z")
	print(b)

	// use BoltDB. Create lincense.yaml
	c := CreateLicenseDB("2020-08-01T11:45:26.371Z", "2020-12-12T11:45:26.371Z")
	print(c)

	// Read lincense.yaml
	byteValue, err := ioutil.ReadFile("license.yaml")
	if err != nil {
		fmt.Println(err)
	}

	setting := &Application{}

	err = yaml.Unmarshal(byteValue, setting)
	if err != nil {
		fmt.Println(err)
		return
	}

	a := CheckExpiration(setting.ExpirationDate)
	// a := CheckExpirationDB(setting.ExpirationDate)
	if a == false {

		fmt.Println("Bad")
		return

	}

	{
		yamlFile, err := ioutil.ReadFile("test.yaml")
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successfully Opened test.yaml")

		st := &Config{}

		yaml.Unmarshal(yamlFile, st)

		fmt.Println(st.HostService)
	}

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

/*-------------------- NOT USE BoltDB --------------------- */

/*
CheckExpiration object
*/
func CheckExpiration(timeExp string) bool {

	sDec, err := base64.StdEncoding.DecodeString(timeExp)
	if err != nil {
		fmt.Println(err)
		return false
	}

	keyDec, err := base64.StdEncoding.DecodeString(RSAPrivateKey)
	if err != nil {
		fmt.Println(err)
		return false
	}

	cre, err := crypto.RSADecryptOAEP(sDec, keyDec)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(string(cre))

	now := time.Now()
	t, err := time.Parse(time.RFC3339, string(cre))
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

	keyDec, err := base64.StdEncoding.DecodeString(RSAPublicKey)
	if err != nil {

		return err
	}

	a := []byte(cr)
	b := []byte(exp)

	cre, err := crypto.RSAEncryptOAEP(keyDec, a)
	if err != nil {

		return err
	}

	expi, err := crypto.RSAEncryptOAEP(keyDec, b)
	if err != nil {

		return err
	}

	created := base64.StdEncoding.EncodeToString(cre)
	expiration := base64.StdEncoding.EncodeToString(expi)

	defaut := &Application{
		DateCreated:    created,
		ExpirationDate: expiration,
	}

	body, _ := yaml.Marshal(defaut)

	err = ioutil.WriteFile("license.yaml", body, 0664)
	if err != nil {

		return err
	}
	return nil

}

/*-------------------- USE BoltDB --------------------- */

/*
CheckExpirationDB object
*/
func CheckExpirationDB(timeExp string) bool {

	db, err := bolt.Open("k.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
		return false
	}

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte("RSAPrivateKey"))

		w = v
		return nil
	})
	defer db.Close()

	sDec, err := base64.StdEncoding.DecodeString(timeExp)
	if err != nil {
		fmt.Println(err)
		return false
	}

	keyDec, err := base64.StdEncoding.DecodeString(string(w))
	if err != nil {
		fmt.Println(err)
		return false
	}

	cre, err := crypto.RSADecryptOAEP(sDec, keyDec)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(string(cre))

	now := time.Now()
	t, err := time.Parse(time.RFC3339, string(cre))
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
CreateLicenseDB object
*/
func CreateLicenseDB(cr string, exp string) error {

	db, err := bolt.Open("k.db", 0600, nil)
	if err != nil {

		return err
	}

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte("RSAPublicKey"))

		q = v
		return nil
	})
	defer db.Close()
	keyDec, err := base64.StdEncoding.DecodeString(string(q))
	if err != nil {

		return err
	}

	a := []byte(cr)
	b := []byte(exp)

	cre, err := crypto.RSAEncryptOAEP(keyDec, a)
	if err != nil {

		return err
	}

	expi, err := crypto.RSAEncryptOAEP(keyDec, b)
	if err != nil {

		return err
	}

	created := base64.StdEncoding.EncodeToString(cre)
	expiration := base64.StdEncoding.EncodeToString(expi)

	defaut := &Application{
		DateCreated:    created,
		ExpirationDate: expiration,
	}

	body, _ := json.Marshal(defaut)

	err = ioutil.WriteFile("license.json", body, 0664)
	if err != nil {

		return err
	}
	return nil

}
