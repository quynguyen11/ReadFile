package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v1"
)

// Go provides a `flag` package supporting basic
// command-line flag parsing. We'll use this package to
// implement our example command-line program.

func main() {

	// File, err := os.Open("application.license.yaml")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	byteValue, err := ioutil.ReadFile("application.license.yaml")
	if err != nil {
		fmt.Println(err)
	}

	/*
	Nếu đã dùng ioutil sao không đọc file thẳng luôn mà phải mở bằng os ? lý do ?

	byteValue, err := 	ioutil.ReadFile("application.license.yaml")
	if err != nil {
		fmt.Println(err)
	}
	*/
	setting := &Application{}
	
	/*
	Em thường bỏ qua biến err - không bắt lỗi. vậy nếu xảy ra lỗi thì chương trình sẽ như thế nào ?
	*/
	yaml.Unmarshal(byteValue, setting)

	a := CheckExpiration(setting.ExpirationDate)
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
	/* 
	Hai file license tạo ra và đọc vào định dạng dữ liệu khác nhau. như vậy lúc sử dụng phải tạo file đọc vào bằng tay sao ?
	Nếu không nhớ định dạng hoặc người khác sử dụng pakage này thì chuyện gì sẽ xảy ra ?
	*/

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

	/*
	Mục đích của pakage này là giới hạn thời gian sử dụng chương trình. tức là người sử dụng không thể sử dụng nếu hết hạn.
	vì vậy file license hoặc giá trị của nó phải được mã hóa để người dùng không thể thay đổi để bypass./ nếu không package này vô dụng.
	Em thấy string base64 bị đọc và thay đổi 1 cách dễ dàng.

	Hacker chính là hành động khai thác lỗ hổng nhằm bypass chương trình. cách làm như vậy ai cũng có thể qua dễ dàng.
	*/

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
