// replication-manager - Replication Manager Monitoring and CLI for MariaDB and MySQL
// Authors: Guillaume Lefranc <guillaume@signal18.io>
//          Stephane Varoqui  <stephane.varoqui@mariadb.com>
// This source code is licensed under the GNU General Public License, version 3.
// Redistribution/Reuse of this code is permitted under the GNU v3 license, as
// an additional term, ALL code must carry the original Author(s) credit in comment form.
// See LICENSE in this directory for the integral text.

package main

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	source      string
	destination string
	cleanall    = false
)

func init() {

	rootCmd.AddCommand(provisionCmd)
	provisionCmd.Flags().StringVar(&source, "source", "", "Source server")
	provisionCmd.Flags().StringVar(&destination, "destination", "", "Source server")
	//bootstrapCmd.Flags().StringVar(&conf.PrefMaster, "prefmaster", "", "Preferred server for master initialization")
	//bootstrapCmd.Flags().StringVar(&conf.MasterConn, "master-connection", "", "Connection name to use for multisource replication")
	//bootstrapCmd.Flags().IntVar(&conf.MasterConnectRetry, "master-connect-retry", 10, "Specifies how many seconds to wait between slave connect retries to master")

}

var provisionCmd = &cobra.Command{
	Use:   "provision",
	Short: "Provision a replica server",
	Long: `The provision command is used to create a new replication server
using mysqldump or xtrabackup`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("I'm dedicated this dead code to Kolbe from MariaDB :) \n\nIt should work soon via calling the rejoin dump or custum script from the api")
		/*	dbHost, dbPort := misc.SplitHostPort(source)
			destHost, destPort := misc.SplitHostPort(destination)
			dbUser, dbPass = misc.SplitPair(confs[cfgGroupIndex].User)
			hostArg := fmt.Sprintf("--host=%s", dbHost)
			portArg := fmt.Sprintf("--port=%s", dbPort)
			userArg := fmt.Sprintf("--user=%s", dbUser)
			passArg := fmt.Sprintf("--password=%s", dbPass)
			desthostArg := fmt.Sprintf("--host=%s", destHost)
			destportArg := fmt.Sprintf("--port=%s", destPort)
			dumpCmd := exec.Command("/usr/bin/mysqldump", "--opt", "--single-transaction", "--all-databases", hostArg, portArg, userArg, passArg)
			clientCmd := exec.Command("/usr/bin/mysql", desthostArg, destportArg, userArg, passArg)
			var err error
			clientCmd.Stdin, err = dumpCmd.StdoutPipe()
			if err != nil {
				log.Fatal("Error opening pipe:", err)
			}
			if err := dumpCmd.Start(); err != nil {
				log.Fatal("Error starting dump:", err, dumpCmd.Path, dumpCmd.Args)
			}
			if err := clientCmd.Run(); err != nil {
				log.Fatal("Error starting client:", err, clientCmd.Path, clientCmd.Args)
			}*/
	},
}
