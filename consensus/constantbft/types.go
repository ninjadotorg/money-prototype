package constantbft

import (
	libp2p "github.com/libp2p/go-libp2p-peer"
	"github.com/ninjadotorg/constant/wire"
)

type ChainInfo struct {
	CurrentCommittee        []string
	CandidateListMerkleHash string
	ChainsHeight            []int
}

type swapSig struct {
	Validator string
	SwapSig   string
}

type serverInterface interface {
	// list functions callback which are assigned from Server struct
	GetPeerIDsFromPublicKey(string) []libp2p.ID
	PushMessageToAll(wire.Message) error
	PushMessageToPeer(wire.Message, libp2p.ID) error
	PushMessageToShard(wire.Message, byte) error
	PushMessageToBeacon(wire.Message) error
	PushMessageToPbk(wire.Message, string) error
	UpdateConsensusState(role string, userPbk string, currentShard *byte, beaconCommittee []string, shardCommittee map[byte][]string)
}
