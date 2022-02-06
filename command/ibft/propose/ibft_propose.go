package propose

import (
	"fmt"
	"github.com/0xPolygon/polygon-edge/command/output"
	"github.com/spf13/cobra"

	"github.com/0xPolygon/polygon-edge/command/helper"
)

func GetCommand() *cobra.Command {
	ibftSnapshotCmd := &cobra.Command{
		Use:     "propose",
		Short:   "Proposes a new candidate to be added or removed from the validator set",
		PreRunE: runPreRun,
		Run:     runCommand,
	}

	setFlags(ibftSnapshotCmd)
	setRequiredFlags(ibftSnapshotCmd)

	return ibftSnapshotCmd
}

func setFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(
		&params.addressRaw,
		addressFlag,
		"",
		"the address of the account to be voted for"+
			"",
	)

	cmd.Flags().StringVar(
		&params.vote,
		voteFlag,
		"",
		fmt.Sprintf(
			"requested change to the validator set. Possible values: [%s, %s]",
			authVote,
			dropVote,
		),
	)
}

func setRequiredFlags(cmd *cobra.Command) {
	for _, requiredFlag := range params.getRequiredFlags() {
		_ = cmd.MarkFlagRequired(requiredFlag)
	}
}

func runPreRun(_ *cobra.Command, _ []string) error {
	return params.validateFlags()
}

func runCommand(cmd *cobra.Command, _ []string) {
	outputter := output.InitializeOutputter(cmd)
	defer outputter.WriteOutput()

	if err := params.proposeCandidate(helper.GetGRPCAddress(cmd)); err != nil {
		outputter.SetError(err)

		return
	}

	outputter.SetCommandResult(params.getResult())
}
