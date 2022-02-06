package propose

import (
	"context"
	"errors"
	"github.com/0xPolygon/polygon-edge/command/helper"
	"github.com/0xPolygon/polygon-edge/command/output"
	ibftOp "github.com/0xPolygon/polygon-edge/consensus/ibft/proto"
	"github.com/0xPolygon/polygon-edge/types"
)

const (
	voteFlag    = "vote"
	addressFlag = "addr"
)

const (
	authVote = "auth"
	dropVote = "drop"
)

var (
	errInvalidVoteType      = errors.New("invalid vote type")
	errInvalidAddressFormat = errors.New("invalid address format")
)

var (
	params = &proposeParams{}
)

type proposeParams struct {
	addressRaw string

	vote    string
	address types.Address
}

func (p *proposeParams) getRequiredFlags() []string {
	return []string{
		voteFlag,
		addressFlag,
	}
}

func (p *proposeParams) validateFlags() error {
	if !isValidVote(p.vote) {
		return errInvalidVoteType
	}

	p.address = types.Address{}
	if err := p.address.UnmarshalText([]byte(p.addressRaw)); err != nil {
		return errInvalidAddressFormat
	}

	return nil
}

func isValidVote(vote string) bool {
	return vote != authVote && vote != dropVote
}

func (p *proposeParams) proposeCandidate(grpcAddress string) error {
	ibftClient, err := helper.GetIBFTOperatorClientConnection(grpcAddress)
	if err != nil {
		return err
	}

	if _, err := ibftClient.Propose(
		context.Background(),
		&ibftOp.Candidate{
			Address: p.address.String(),
			Auth:    p.vote == authVote,
		},
	); err != nil {
		return err
	}

	return nil
}

func (p *proposeParams) getResult() output.CommandResult {
	return &IBFTProposeResult{
		Address: p.address.String(),
		Vote:    p.vote,
	}
}
