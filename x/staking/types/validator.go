package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// Validator represents a single validator.
// This is defined as an interface so that we can use the SDK types
// as well as database types properly.
type Validator interface {
	GetConsAddr() string
	GetConsPubKey() string
	GetOperator() string
	GetSelfDelegateAddress() string
	GetMaxChangeRate() *sdk.Dec
	GetMaxRate() *sdk.Dec
	GetHeight() int64
}

// validator allows to easily implement the Validator interface
type validator struct {
	ConsensusAddr       string
	ConsPubKey          string
	OperatorAddr        string
	SelfDelegateAddress string
	MaxChangeRate       *sdk.Dec
	MaxRate             *sdk.Dec
	Height              int64
}

// NewValidator allows to build a new Validator implementation having the given data
func NewValidator(
	consAddr string, opAddr string, consPubKey string,
	selfDelegateAddress string, maxChangeRate *sdk.Dec,
	maxRate *sdk.Dec, height int64,
) Validator {
	return validator{
		ConsensusAddr:       consAddr,
		ConsPubKey:          consPubKey,
		OperatorAddr:        opAddr,
		SelfDelegateAddress: selfDelegateAddress,
		MaxChangeRate:       maxChangeRate,
		MaxRate:             maxRate,
		Height:              height,
	}
}

// GetConsAddr implements the Validator interface
func (v validator) GetConsAddr() string {
	return v.ConsensusAddr
}

// GetConsPubKey implements the Validator interface
func (v validator) GetConsPubKey() string {
	return v.ConsPubKey
}

func (v validator) GetOperator() string {
	return v.OperatorAddr
}

func (v validator) GetSelfDelegateAddress() string {
	return v.SelfDelegateAddress
}

func (v validator) GetMaxChangeRate() *sdk.Dec {
	return v.MaxChangeRate
}

func (v validator) GetMaxRate() *sdk.Dec {
	return v.MaxRate
}

func (v validator) GetHeight() int64 {
	return v.Height
}

// _________________________________________________________

// ValidatorDescription contains the description of a validator
// and timestamp do the description get changed
type ValidatorDescription struct {
	OperatorAddress string
	Description     stakingtypes.Description
	AvatarURL       string
	Height          int64
}

// NewValidatorDescription return a new ValidatorDescription object
func NewValidatorDescription(
	opAddr string, description stakingtypes.Description, avatarURL string, height int64,
) ValidatorDescription {
	return ValidatorDescription{
		OperatorAddress: opAddr,
		Description:     description,
		AvatarURL:       avatarURL,
		Height:          height,
	}
}

// ----------------------------------------------------------------------------------------------------------

// ValidatorCommission contains the data of a validator commission at a given height
type ValidatorCommission struct {
	ValAddress        string
	Commission        *sdk.Dec
	MinSelfDelegation *sdk.Int
	Height            int64
}

// NewValidatorCommission return a new validator commission instance
func NewValidatorCommission(
	valAddress string, rate *sdk.Dec, minSelfDelegation *sdk.Int, height int64,
) ValidatorCommission {
	return ValidatorCommission{
		ValAddress:        valAddress,
		Commission:        rate,
		MinSelfDelegation: minSelfDelegation,
		Height:            height,
	}
}

//--------------------------------------------

// ValidatorVotingPower represents the voting power of a validator at a specific block height
type ValidatorVotingPower struct {
	ConsensusAddress string
	VotingPower      int64
	Height           int64
}

// NewValidatorVotingPower creates a new ValidatorVotingPower
func NewValidatorVotingPower(address string, votingPower int64, height int64) ValidatorVotingPower {
	return ValidatorVotingPower{
		ConsensusAddress: address,
		VotingPower:      votingPower,
		Height:           height,
	}
}

//--------------------------------------------------------

// ValidatorStatus represents the current state for the specified validator at the specific height
type ValidatorStatus struct {
	ConsensusAddress string
	ConsensusPubKey  string
	Status           int
	Jailed           bool
	Height           int64
}

// NewValidatorStatus creates a new ValidatorVotingPower
func NewValidatorStatus(address, pubKey string, status int, jailed bool, height int64) ValidatorStatus {
	return ValidatorStatus{
		ConsensusAddress: address,
		ConsensusPubKey:  pubKey,
		Status:           status,
		Jailed:           jailed,
		Height:           height,
	}
}

//---------------------------------------------------------------

// DoubleSignEvidence represent a double sign evidence on each tendermint block
type DoubleSignEvidence struct {
	VoteA  DoubleSignVote
	VoteB  DoubleSignVote
	Height int64
}

// NewDoubleSignEvidence return a new DoubleSignEvidence object
func NewDoubleSignEvidence(height int64, voteA DoubleSignVote, voteB DoubleSignVote) DoubleSignEvidence {
	return DoubleSignEvidence{
		VoteA:  voteA,
		VoteB:  voteB,
		Height: height,
	}
}

// DoubleSignVote represents a double vote which is included inside a DoubleSignEvidence
type DoubleSignVote struct {
	BlockID          string
	ValidatorAddress string
	Signature        string
	Type             int
	Height           int64
	Round            int32
	ValidatorIndex   int32
}

// NewDoubleSignVote allows to create a new DoubleSignVote instance
func NewDoubleSignVote(
	roundType int,
	height int64,
	round int32,
	blockID string,
	validatorAddress string,
	validatorIndex int32,
	signature string,
) DoubleSignVote {
	return DoubleSignVote{
		Type:             roundType,
		Height:           height,
		Round:            round,
		BlockID:          blockID,
		ValidatorAddress: validatorAddress,
		ValidatorIndex:   validatorIndex,
		Signature:        signature,
	}
}
