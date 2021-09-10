package twopc

import (
	"fmt"
	"github.com/RosettaFlow/Carrier-Go/common"
	"github.com/RosettaFlow/Carrier-Go/common/rlputil"
	"github.com/RosettaFlow/Carrier-Go/common/timeutils"
	ctypes "github.com/RosettaFlow/Carrier-Go/consensus/twopc/types"
	apipb "github.com/RosettaFlow/Carrier-Go/lib/common"
	pb "github.com/RosettaFlow/Carrier-Go/lib/consensus/twopc"
	"github.com/RosettaFlow/Carrier-Go/types"
	"testing"
	"time"
)

func genHash() (h common.Hash) {
	proposalHash := rlputil.RlpHash([]interface{}{
		"t.config.Option.NodeID",
		uint64(timeutils.UnixMsec()),
		"task.TaskId()",
		"task.TaskData().TaskName",
		5678,
		uint64(time.Now().Nanosecond()),
	})

	return proposalHash
}
func TestUpdateStateToDatabase(t *testing.T) {
	hash := genHash()
	//region proposalSet
	proposalSetMap := make(map[common.Hash]*ctypes.ProposalState, 0)
	sender:=&apipb.TaskOrganization{
		PartyId:    "2",
		NodeName:   "NodeName",
		NodeId:     "NodeId",
		IdentityId: "IdentityId",
	}
	proposalSetInfo := ctypes.NewProposalState(hash,"task001",sender)
	proposalSetMap[hash]=proposalSetInfo
	UpdateStateToDatabase(proposalSetMap)
	// endregion

	taskRoleMap := make(map[apipb.TaskRole]uint32, 0)
	taskRoleMap[22] = 112

	//region prepareVotes
	prepareVoteMap := make(map[string]*types.PrepareVote, 0)
	prepareVoteMap["prepareVoteMap"] = &types.PrepareVote{
		MsgOption: &types.MsgOption{
			ProposalId:      hash,
			SenderRole:      2,
			SenderPartyId:   "SenderPartyId",
			ReceiverRole:    3,
			ReceiverPartyId: "ReceiverPartyId",
			Owner: &apipb.TaskOrganization{
				PartyId:    "2",
				NodeName:   "NodeName",
				NodeId:     "NodeId",
				IdentityId: "IdentityId",
			},
		},
		PeerInfo: &types.PrepareVoteResource{
			Id:      "PeerInfo_id",
			Ip:      "192.178.22.222",
			Port:    "7777",
			PartyId: "PartyIdTest",
		},
		VoteOption: 2,
		CreateAt:   4444,
		Sign:       []byte("test sign"),
	}
	prepareVoteStateInfo := &prepareVoteState{
		votes:      prepareVoteMap,
		yesVotes:   taskRoleMap,
		voteStatus: taskRoleMap,
	}
	prepareVotes := make(map[common.Hash]*prepareVoteState, 0)
	prepareVotes[hash] = prepareVoteStateInfo
	UpdateStateToDatabase(prepareVotes)
	//endregion

	//region confirmVotes
	votesMap := make(map[string]*types.ConfirmVote, 0)
	votesMap["vote1"] = &types.ConfirmVote{
		MsgOption: &types.MsgOption{
			ProposalId:      hash,
			SenderRole:      2,
			SenderPartyId:   "SenderPartyId",
			ReceiverRole:    3,
			ReceiverPartyId: "ReceiverPartyId",
			Owner: &apipb.TaskOrganization{
				PartyId:    "2",
				NodeName:   "NodeName",
				NodeId:     "NodeId",
				IdentityId: "IdentityId",
			},
		},
		VoteOption: 2,
		CreateAt:   4444,
		Sign:       []byte("test sign"),
	}
	confirmVoteStateInfo := &confirmVoteState{
		votes:      votesMap,
		yesVotes:   taskRoleMap,
		voteStatus: taskRoleMap,
	}
	confirmVoteState := make(map[common.Hash]*confirmVoteState, 0)
	confirmVoteState[hash] = confirmVoteStateInfo
	UpdateStateToDatabase(confirmVoteState)
	//endregion

	//region proposalPeerInfoCache
	proposalPeerInfoCacheDB := make(map[common.Hash]*pb.ConfirmTaskPeerInfo, 0)
	confirmTaskPeerInfo := &pb.ConfirmTaskPeerInfo{
		OwnerPeerInfo: &pb.TaskPeerInfo{
			Ip:      []byte("192.168.1.111"),
			Port:    []byte("9988"),
			PartyId: []byte("5"),
		},
		DataSupplierPeerInfoList: []*pb.TaskPeerInfo{
			{
				Ip:      []byte("192.168.1.112"),
				Port:    []byte("9987"),
				PartyId: []byte("4"),
			},
		},
		PowerSupplierPeerInfoList: []*pb.TaskPeerInfo{
			{
				Ip:      []byte("192.168.1.113"),
				Port:    []byte("9986"),
				PartyId: []byte("3"),
			},
		},
		ResultReceiverPeerInfoList: []*pb.TaskPeerInfo{
			{
				Ip:      []byte("192.168.1.112"),
				Port:    []byte("9985"),
				PartyId: []byte("2"),
			},
		},
	}
	proposalPeerInfoCacheDB[hash] = confirmTaskPeerInfo
	UpdateStateToDatabase(proposalPeerInfoCacheDB)
	//endregion
}

func TestConsensusStateFromDatabase(t *testing.T) {
	result := ConsensusStateFromDatabase()
	fmt.Println("TestConsensusStateFromDatabase result is:", result)
}
