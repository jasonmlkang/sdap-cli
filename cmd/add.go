package cmd

import (
	// "fmt"
	"github.com/spf13/cobra"

	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
)

var addCmd = &cobra.Command{
	Use:   "add <granule 1> <granule 2> ...",
	Short: "Add granule to SDAP",
	Run: Add,
}

func init() {
	addCmd.Flags().StringP("sshUsername", "u", "", "SSH username to SDAP cluster server")
	addCmd.Flags().StringP("sshKeyPath", "k", "", "Path to SSH private key used for this connection (including filename)")
	addCmd.Flags().StringP("sshHostKeyPath", "p", "", "Path to SSH host key to SDAP cluster server (including filename)")
	ingestionCmd.AddCommand(addCmd)
}

func Add(cmd *cobra.Command, args []string) {
	key, err := ioutil.ReadFile("<path-to->")
	if err != nil {
		log.Fatalf("Unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: "kang",
		Auth: []ssh.AuthMethod{
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer),
			ssh.Password("<password>"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", "<url>:22", config)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer client.Close()
}
