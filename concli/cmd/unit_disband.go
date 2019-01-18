/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"
	"strconv"

	"github.com/conreality/conreality.go/sdk"
	"github.com/spf13/cobra"
)

// UnitDisbandCmd describes and implements the `concli unit disband` command
var UnitDisbandCmd = &cobra.Command{
	Use:   "disband UNIT-NAME",
	Short: "Disband a unit",
	Long:  "Conreality Command-Line Interface (CLI): Disband Unit",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		client, err := sdk.Connect(masterURL)
		if err != nil {
			panic(err)
		}
		defer client.Disconnect()

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		session, err := client.Authenticate(ctx, playerNick)
		if err != nil {
			panic(err)
		}
		defer session.Close()

		unitID, err := strconv.Atoi(args[0]) // TODO: support unit names as well
		if err != nil {
			panic(err)
		}

		err = session.DisbandUnit(ctx, sdk.UnitID(unitID))
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	UnitCmd.AddCommand(UnitDisbandCmd)
}
