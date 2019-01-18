/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"
	"strconv"

	"github.com/conreality/conreality.go/sdk"
	"github.com/spf13/cobra"
)

// UnitJoinCmd describes and implements the `concli unit join` command
var UnitJoinCmd = &cobra.Command{
	Use:   "join UNIT-NAME",
	Short: "Join a unit",
	Long:  "Conreality Command-Line Interface (CLI): Join Unit",
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

		err = session.JoinUnit(ctx, sdk.UnitID(unitID))
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	UnitCmd.AddCommand(UnitJoinCmd)
}
