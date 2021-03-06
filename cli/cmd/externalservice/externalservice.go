package externalservice

import (
	"github.com/rancher/rio/cli/pkg/builder"
	"github.com/rancher/rio/cli/pkg/clicontext"
	"github.com/rancher/rio/cli/pkg/table"
	"github.com/urfave/cli"
)

func ExternalService(app *cli.App) cli.Command {
	create := builder.Command(&Create{},
		"Create external services",
		app.Name+" external create [EXTERNAL_SERVICE] [(IP)(FQDN)(SERVICE/APP/ROUTER)]",
		"To create an externalservice by pointing to FQDN, run `rio external create [NAMESPACE:]NAME foo.bar.\n"+
			"	 To create an externalservice by pointing to IPs, run `rio external create [NAMESPACE:]NAME 1.1.1.1 2.2.2.2.\n"+
			"	 To create an externalservice by pointing to service/router in another namespace, run `rio external create [NAMESPACE:]NAME [NAMESPACE:]@NAME`")
	ls := builder.Command(&Ls{},
		"List external services",
		app.Name+" external ls",
		"")
	rm := builder.Command(&Rm{},
		"Remove external services",
		app.Name+" external rm [EXTERNAL_SERVICE]",
		"Example: run `rio external rm externalservice/redis`")
	rm.Aliases = []string{"delete"}
	return cli.Command{
		Name:      "externalservices",
		Aliases:   []string{"external"},
		ShortName: "externalservice",
		Usage:     "Operation on externalservices",
		Action:    clicontext.DefaultAction(ls.Action),
		Flags:     table.WriterFlags(),
		Category:  "SUB COMMANDS",
		Subcommands: []cli.Command{
			ls,
			create,
			rm,
		},
	}
}
