package bigdipper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/forbole/bdjuno/database/types"

	"github.com/forbole/bdjuno/types"

	dbtypes "github.com/forbole/bdjuno/database/bigdipper/types"
)

func (suite *DbTestSuite) TestBigDipperDb_SaveCommunityPool() {
	// Save data
	original := sdk.NewDecCoins(sdk.NewDecCoin("uatom", sdk.NewInt(100)))
	err := suite.database.SaveCommunityPool(original, 10)
	suite.Require().NoError(err)

	// Verify data
	expected := dbtypes.NewCommunityPoolRow(types.NewDbDecCoins(original), 10)
	var rows []dbtypes.CommunityPoolRow
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM community_pool`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1, "community_pool table should contain only one row")
	suite.Require().True(expected.Equals(rows[0]))

	// ---------------------------------------------------------------------------------------------------------------

	// Try updating with lower height
	coins := sdk.NewDecCoins(sdk.NewDecCoin("uatom", sdk.NewInt(50)))
	err = suite.database.SaveCommunityPool(coins, 5)
	suite.Require().NoError(err)

	// Verify data
	expected = dbtypes.NewCommunityPoolRow(types.NewDbDecCoins(original), 10)
	rows = []dbtypes.CommunityPoolRow{}
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM community_pool`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1, "community_pool table should contain only one row")
	suite.Require().True(expected.Equals(rows[0]), "updating with lower height should not modify the data")

	// ---------------------------------------------------------------------------------------------------------------

	// Try updating with equal height
	coins = sdk.NewDecCoins(sdk.NewDecCoin("uatom", sdk.NewInt(120)))
	err = suite.database.SaveCommunityPool(coins, 10)
	suite.Require().NoError(err)

	// Verify data
	expected = dbtypes.NewCommunityPoolRow(types.NewDbDecCoins(coins), 10)
	rows = []dbtypes.CommunityPoolRow{}
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM community_pool`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1, "community_pool table should contain only one row")
	suite.Require().True(expected.Equals(rows[0]), "updating with same height should modify the data")

	// ---------------------------------------------------------------------------------------------------------------

	// Try updating with higher height
	coins = sdk.NewDecCoins(sdk.NewDecCoin("uatom", sdk.NewInt(200)))
	err = suite.database.SaveCommunityPool(coins, 11)
	suite.Require().NoError(err)

	// Verify data
	expected = dbtypes.NewCommunityPoolRow(types.NewDbDecCoins(coins), 11)
	rows = []dbtypes.CommunityPoolRow{}
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM community_pool`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1, "community_pool table should contain only one row")
	suite.Require().True(expected.Equals(rows[0]), "updating with higher height should modify the data")
}

func (suite *DbTestSuite) TestBigDipperDb_SaveValidatorCommissionAmount() {
	validator := suite.getValidator(
		"cosmosvalcons1qqqqrezrl53hujmpdch6d805ac75n220ku09rl",
		"cosmosvaloper1rcp29q3hpd246n6qak7jluqep4v006cdsc2kkl",
		"cosmosvalconspub1zcjduepq7mft6gfls57a0a42d7uhx656cckhfvtrlmw744jv4q0mvlv0dypskehfk8",
	)

	// Save the data
	original := sdk.NewDecCoins(sdk.NewDecCoin("atom", sdk.NewInt(100)))
	amount := types.NewValidatorCommissionAmount(
		validator.GetOperator(),
		validator.GetSelfDelegateAddress(),
		original,
		10,
	)
	err := suite.database.SaveValidatorCommissionAmount(amount)
	suite.Require().NoError(err)

	// Verify the data
	originalRow := dbtypes.NewValidatorCommissionAmountRow(validator.GetConsAddr(), types.NewDbDecCoins(original), 10)
	var rows []dbtypes.ValidatorCommissionAmountRow
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM validator_commission_amount`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().Equal(originalRow, rows[0])

	// ---------------------------------------------------------------------------------------------------------------

	// Try updating with lower height
	coins := sdk.NewDecCoins(sdk.NewDecCoin("atom", sdk.NewInt(120)))
	amount = types.NewValidatorCommissionAmount(validator.GetOperator(), validator.GetSelfDelegateAddress(), coins, 9)
	err = suite.database.SaveValidatorCommissionAmount(amount)
	suite.Require().NoError(err)

	// Verify the data
	rows = []dbtypes.ValidatorCommissionAmountRow{}
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM validator_commission_amount`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().Equal(originalRow, rows[0], "updating with lower height should not modify the data")

	// ---------------------------------------------------------------------------------------------------------------

	// Try updating with same height
	coins = sdk.NewDecCoins(sdk.NewDecCoin("atom", sdk.NewInt(200)))
	amount = types.NewValidatorCommissionAmount(validator.GetOperator(), validator.GetSelfDelegateAddress(), coins, 10)
	err = suite.database.SaveValidatorCommissionAmount(amount)
	suite.Require().NoError(err)

	// Verify the data
	expected := dbtypes.NewValidatorCommissionAmountRow(validator.GetConsAddr(), types.NewDbDecCoins(coins), 10)
	rows = []dbtypes.ValidatorCommissionAmountRow{}
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM validator_commission_amount`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().Equal(expected, rows[0], "updating with same height should modify the data")

	// ---------------------------------------------------------------------------------------------------------------

	// Try updating with higher height
	coins = sdk.NewDecCoins(sdk.NewDecCoin("atom", sdk.NewInt(500)))
	amount = types.NewValidatorCommissionAmount(validator.GetOperator(), validator.GetSelfDelegateAddress(), coins, 11)
	err = suite.database.SaveValidatorCommissionAmount(amount)
	suite.Require().NoError(err)

	// Verify the data
	expected = dbtypes.NewValidatorCommissionAmountRow(validator.GetConsAddr(), types.NewDbDecCoins(coins), 11)
	rows = []dbtypes.ValidatorCommissionAmountRow{}
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM validator_commission_amount`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().Equal(expected, rows[0], "updating with higher height should modify the data")
}

func (suite *DbTestSuite) TestBigDipperDb_SaveDelegatorsRewardsAmounts() {
	delegator1 := suite.getAccount("cosmos1z4hfrxvlgl4s8u4n5ngjcw8kdqrcv43599amxs")
	delegator2 := suite.getAccount("cosmos184ma3twcfjqef6k95ne8w2hk80x2kah7vcwy4a")
	validator1 := suite.getValidator(
		"cosmosvalcons1qqqqrezrl53hujmpdch6d805ac75n220ku09rl",
		"cosmosvaloper1rcp29q3hpd246n6qak7jluqep4v006cdsc2kkl",
		"cosmosvalconspub1zcjduepq7mft6gfls57a0a42d7uhx656cckhfvtrlmw744jv4q0mvlv0dypskehfk8",
	)
	validator2 := suite.getValidator(
		"cosmosvalcons1qq92t2l4jz5pt67tmts8ptl4p0jhr6utx5xa8y",
		"cosmosvaloper1000ya26q2cmh399q4c5aaacd9lmmdqp90kw2jn",
		"cosmosvalconspub1zcjduepqe93asg05nlnj30ej2pe3r8rkeryyuflhtfw3clqjphxn4j3u27msrr63nk",
	)

	// Save the data
	rewards := []types.DelegatorRewardAmount{
		types.NewDelegatorRewardAmount(
			delegator1.String(),
			validator1.GetOperator(),
			delegator1.String(),
			sdk.NewDecCoins(sdk.NewDecCoin("cosmos", sdk.NewInt(100))),
			10,
		),
		types.NewDelegatorRewardAmount(
			delegator1.String(),
			validator2.GetOperator(),
			delegator1.String(),
			sdk.NewDecCoins(sdk.NewDecCoin("cosmos", sdk.NewInt(200))),
			11,
		),
		types.NewDelegatorRewardAmount(
			delegator2.String(),
			validator2.GetOperator(),
			delegator2.String(),
			sdk.NewDecCoins(sdk.NewDecCoin("cosmos", sdk.NewInt(200))),
			12,
		),
	}
	err := suite.database.SaveDelegatorsRewardsAmounts(rewards)
	suite.Require().NoError(err)

	// Verify the data
	expected := []dbtypes.DelegationRewardRow{
		dbtypes.NewDelegationRewardRow(
			delegator1.String(),
			validator1.GetConsAddr(),
			delegator1.String(),
			types.NewDbDecCoins(sdk.NewDecCoins(sdk.NewDecCoin("cosmos", sdk.NewInt(100)))),
			10,
		),
		dbtypes.NewDelegationRewardRow(
			delegator1.String(),
			validator2.GetConsAddr(),
			delegator1.String(),
			types.NewDbDecCoins(sdk.NewDecCoins(sdk.NewDecCoin("cosmos", sdk.NewInt(200)))),
			11,
		),
		dbtypes.NewDelegationRewardRow(
			delegator2.String(),
			validator2.GetConsAddr(),
			delegator2.String(),
			types.NewDbDecCoins(sdk.NewDecCoins(sdk.NewDecCoin("cosmos", sdk.NewInt(200)))),
			12,
		),
	}

	var rows []dbtypes.DelegationRewardRow
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM delegation_reward ORDER BY height`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, len(expected))

	for index, row := range rows {
		suite.Require().True(row.Equals(expected[index]))
	}

	// -------------------------------------------------------------------------------------------------------------------

	// Update the data
	rewards = []types.DelegatorRewardAmount{
		types.NewDelegatorRewardAmount(
			delegator1.String(),
			validator1.GetOperator(),
			delegator1.String(),
			sdk.NewDecCoins(sdk.NewDecCoin("cosmos", sdk.NewInt(120))),
			10,
		),
		types.NewDelegatorRewardAmount(
			delegator1.String(),
			validator2.GetOperator(),
			delegator1.String(),
			sdk.NewDecCoins(sdk.NewDecCoin("cosmos", sdk.NewInt(180))),
			9,
		),
		types.NewDelegatorRewardAmount(
			delegator2.String(),
			validator2.GetOperator(),
			delegator2.String(),
			sdk.NewDecCoins(sdk.NewDecCoin("cosmos", sdk.NewInt(50))),
			13,
		),
	}
	err = suite.database.SaveDelegatorsRewardsAmounts(rewards)
	suite.Require().NoError(err)

	// Verify the data
	expected = []dbtypes.DelegationRewardRow{
		dbtypes.NewDelegationRewardRow(
			delegator1.String(),
			validator1.GetConsAddr(),
			delegator1.String(),
			types.NewDbDecCoins(sdk.NewDecCoins(sdk.NewDecCoin("cosmos", sdk.NewInt(120)))),
			10,
		),
		dbtypes.NewDelegationRewardRow(
			delegator1.String(),
			validator2.GetConsAddr(),
			delegator1.String(),
			types.NewDbDecCoins(sdk.NewDecCoins(sdk.NewDecCoin("cosmos", sdk.NewInt(200)))),
			11,
		),
		dbtypes.NewDelegationRewardRow(
			delegator2.String(),
			validator2.GetConsAddr(),
			delegator2.String(),
			types.NewDbDecCoins(sdk.NewDecCoins(sdk.NewDecCoin("cosmos", sdk.NewInt(50)))),
			13,
		),
	}

	rows = []dbtypes.DelegationRewardRow{}
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM delegation_reward ORDER BY height`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, len(expected))

	for index, row := range rows {
		suite.Require().True(row.Equals(expected[index]))
	}
}
