package drand

import (
	"bytes"
	"encoding/hex"
	"errors"

	"github.com/BurntSushi/toml"

	"github.com/drand/drand/v2/common/log"
	"github.com/drand/drand/v2/internal/net"
)

func generateJoiningProposal(l log.Logger, beaconID string, joining []string) (string, error) {
	return generateProposal(l, beaconID, joining, []string{}, []string{})
}

func generateProposal(l log.Logger, beaconID string, joining, remaining, leaving []string) (string, error) {
	if len(joining) == 0 && len(remaining) == 0 && len(leaving) == 0 {
		return "", errors.New("you must fill at least one of the ")
	}

	extractParticipants := func(addrs []string) ([]*TomlParticipant, error) {
		var participants []*TomlParticipant
		for _, addr := range addrs {
			client, err := net.NewControlClient(l, addr)
			if err != nil {
				return nil, err
			}
			res, err := client.PublicKey(beaconID)
			if err != nil {
				return nil, err
			}
			participants = append(participants, &TomlParticipant{
				Address:   res.Addr,
				Key:       hex.EncodeToString(res.PubKey),
				Signature: hex.EncodeToString(res.Signature),
			})
		}
		return participants, nil
	}

	joiners, err := extractParticipants(joining)
	if err != nil {
		return "", err
	}
	remainers, err := extractParticipants(remaining)
	if err != nil {
		return "", err
	}
	leavers, err := extractParticipants(leaving)
	if err != nil {
		return "", err
	}

	outputFile := ProposalFileFormat{
		Joining:   joiners,
		Leaving:   leavers,
		Remaining: remainers,
	}

	b := bytes.NewBufferString("")
	err = toml.NewEncoder(b).Encode(outputFile)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}
