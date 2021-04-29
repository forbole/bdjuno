package types

// StakingPoolRow represents a single row inside the staking_pool table
type StakingPoolRow struct {
	BondedTokens    int64 `db:"bonded_tokens"`
	NotBondedTokens int64 `db:"not_bonded_tokens"`
}

// NewStakingPoolRow allows to easily create a new StakingPoolRow
func NewStakingPoolRow(bondedTokens, notBondedTokens int64) StakingPoolRow {
	return StakingPoolRow{
		BondedTokens:    bondedTokens,
		NotBondedTokens: notBondedTokens,
	}
}

// Equal allows to tells whether r and as represent the same rows
func (r StakingPoolRow) Equal(s StakingPoolRow) bool {
	return r.BondedTokens == s.BondedTokens &&
		r.NotBondedTokens == s.NotBondedTokens
}
