package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {

	var err error

	// tlsConfig := &tls.Config{
	// 	InsecureSkipVerify: true,
	// }

	// // config := &tls.Config{
	// 	Certificates: []tls.Certificate{cer},
	// }

	// client := fasthttp.Client{
	// 	TLSConfig: tlsConfig,
	// }

	agent := fiber.AcquireAgent()
	agent.Request().Header.SetMethod("GET")
	agent.Request().Header.Set("Accept", "application/json")
	agent.Request().Header.Set("Content-Type", "application/x-www-form-urlencoded")
	agent.Request().Header.Set("Accept-Language", "en_US")
	agent.Request().Header.Set("App-Key", "m.AppConfig.SsoAppId")
	agent.Request().Header.Set("App-Secret", " m.AppConfig.SsoAppSecret")
	agent.Request().Header.Set("User-Token", "dsfsdfsdfsf")
	agent.Request().SetRequestURI("https://sso.ptpn.id/test-api")
	// agent.InsecureSkipVerify()
	err = agent.Parse()
	if err != nil {
		fmt.Println("error cuy")
		return
	}

	statusNum, body, errs := agent.Bytes()

	fmt.Println("Status Num dari SSO : ")
	fmt.Println(statusNum)

	if len(errs) > 0 {
		fmt.Println(errs)
		return
	}

	if statusNum != fiber.StatusOK {
		fmt.Println("error cuy")
		return
	}

	fmt.Println("Sukses")
	fmt.Println(body)
}
