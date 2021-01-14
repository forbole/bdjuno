package auth

import (
	"github.com/desmos-labs/juno/client"
	"github.com/forbole/bdjuno/database"
	"github.com/forbole/bdjuno/x/utils"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

// RegisterOps returns the AdditionalOperation that periodically runs fetches from
// the LCD to make sure that constantly changing data are synced properly.
func RegisterOps(scheduler *gocron.Scheduler, cp *client.Proxy, db *database.BigDipperDb) error {
	log.Debug().Str("module", "auth").Msg("setting up periodic tasks")

	// Setup a cron job to run every midnight that updates the accounts
	if _, err := scheduler.Every(1).Day().At("00:00").StartImmediately().Do(func() {
		utils.WatchMethod(func() error { return updateAccounts(cp, db) })
	}); err != nil {
		return err
	}

	return nil

}

// updateAccounts gets all the accounts stored inside the database, and refreshes their
// balances by fetching the LCD endpoint.
func updateAccounts(cp *client.Proxy, db *database.BigDipperDb) error {
	height, err := db.GetLastBlockHeight()
	if err != nil {
		return err
	}

	addresses, err := db.GetAccounts()
	if err != nil {
		return err
	}

	return RefreshAccounts(addresses, height, cp, db)
}
