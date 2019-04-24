package ste_mail

import (
	"fmt"
	"net/smtp"
	"log"
	""
)
type mailconfig struct {
	SMTP_server string
	SMTP_port int
	sender string
}


func readFile(path string) []byte {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return dat
}

func send(recpient string, body string) {
	var config mailconfig
	data := readFile("config.json")
	if err := json.Unmarshal(data, &config); err != nil {
        log.Fatal(err)
    }
	c, err := smtp.Dial(fmt.Sprintf("%s:%d", config.SMTP_server, config.SMTP_port))
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Mail(config.sender); err != nil {
		log.Fatal(err)
	}

	if err := c.Rcpt(recpient); err != nil {
		log.Fatal(err)
	}
	wc, err = c.Data()
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintf(wc, body)
	if err != nil {
		log.Fatal(err)
	}
	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = c.Quit()
	if err != nil {
		log.Fatal(err)
	}
}