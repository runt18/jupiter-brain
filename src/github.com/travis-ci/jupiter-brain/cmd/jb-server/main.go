package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/travis-ci/jupiter-brain/server"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Jupiter Brain API server"
	app.Author = "Travis CI"
	app.Email = "contact+jupiter-brain@travis-ci.org"
	app.Version = VersionString
	app.Compiled = GeneratedTime()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "addr",
			Usage: "host:port to listen to",
			Value: func() string {
				v := ":" + os.Getenv("PORT")
				if v == ":" {
					v = ":42161"
				}
				return v
			}(),
			EnvVar: "JUPITER_BRAIN_ADDR",
		},
		cli.StringFlag{
			Name:   "auth-token",
			Usage:  "authentication token for the api server",
			EnvVar: "JUPITER_BRAIN_AUTH_TOKEN",
		},
		cli.StringFlag{
			Name:   "vsphere-api-url",
			Usage:  "URL to vSphere API",
			EnvVar: "JUPITER_BRAIN_VSPHERE_API_URL,VSPHERE_API_URL",
		},
		cli.StringFlag{
			Name:   "vsphere-base-path",
			Usage:  "path to folder of base VMs in vSphere inventory",
			EnvVar: "JUPITER_BRAIN_VSPHERE_BASE_PATH",
		},
		cli.StringFlag{
			Name:   "vsphere-vm-path",
			Usage:  "path to folder where VMs will be put in vSphere inventory",
			EnvVar: "JUPITER_BRAIN_VSPHERE_VM_PATH",
		},
		cli.StringFlag{
			Name:   "vsphere-cluster-path",
			Usage:  "path to compute cluster that VMs will be booted in",
			EnvVar: "JUPITER_BRAIN_VSPHERE_CLUSTER_PATH",
		},
		cli.StringFlag{
			Name:   "database-url",
			Usage:  "URL to the PostgreSQL database",
			EnvVar: "JUPITER_BRAIN_DATABASE_URL,DATABASE_URL",
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "enable debug logging",
			EnvVar: "JUPITER_BRAIN_DEBUG,DEBUG",
		},
		cli.StringFlag{
			Name:   "sentry-dsn",
			Usage:  "Sentry DSN to send errors to",
			EnvVar: "JUPITER_BRAIN_SENTRY_DSN,SENTRY_DSN",
		},
	}
	app.Action = runServer

	app.RunAndExitOnError()
}

func runServer(c *cli.Context) {
	logrus.SetFormatter(&logrus.TextFormatter{DisableColors: true})

	server.Main(&server.Config{
		Addr:      c.String("addr"),
		AuthToken: c.String("auth-token"),
		Debug:     c.Bool("debug"),
		SentryDSN: c.String("sentry-dsn"),

		VSphereURL:         c.String("vsphere-api-url"),
		VSphereBasePath:    c.String("vsphere-base-path"),
		VSphereVMPath:      c.String("vsphere-vm-path"),
		VSphereClusterPath: c.String("vsphere-cluster-path"),

		DatabaseURL: c.String("database-url"),
	})
}
