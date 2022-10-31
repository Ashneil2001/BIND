package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
)

// func GetCredential() (string, string) {
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Enter Username: ")
// 	username, _ := reader.ReadString('\n')

// 	fmt.Print("Enter Password: ")
// 	bytePassword, err := terminal.ReadPassword(0)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	fmt.Println(err)
// 	password := string(bytePassword)

// 	return strings.TrimSpace(username), strings.TrimSpace(password)
// }

func main() {
	fmt.Println(os.Args[1], os.Args[2])
	if len(os.Args) != 3 {
		fmt.Printf("Usage : %s <hostname> <port> \n", os.Args[0])
		os.Exit(0)
	}

	hostname := os.Args[1]
	port := os.Args[2]

	config := &ssh.ClientConfig{
		User: "arv",
		Auth: []ssh.AuthMethod{ssh.Password("824200")},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	fmt.Println("\nConnecting to ", hostname, port)

	hostaddress := strings.Join([]string{hostname, port}, ":")
	client, err := ssh.Dial("tcp", hostaddress, config)
	if err != nil {
		panic(err.Error())
	}

	for {
		session, err := client.NewSession()
		if err != nil {
			panic(err.Error())
		}
		defer session.Close()

		fmt.Println("To exit this program, hit Control-C")
		fmt.Printf("Enter command to execute on %s : ", hostname)

		// fmt.Scanf is unable to accept command with parameters
		// see solution at
		// https://www.socketloop.com/tutorials/golang-accept-input-from-user-with-fmt-scanf-skipped-white-spaces-and-how-to-fix-it
		//fmt.Scanf("%s", &cmd)

		commandReader := bufio.NewReader(os.Stdin)
		cmd, _ := commandReader.ReadString('\n')
		//log.Printf(cmd)
		fmt.Println("Executing command ", cmd)

		// capture standard output
		// will NOT be able to handle refreshing output such as TOP command
		// executing top command will result in panic
		var buff bytes.Buffer
		session.Stdout = &buff
		if err := session.Run(cmd); err != nil {
			panic(err.Error())
		}

		fmt.Println(buff.String())

	}
}
