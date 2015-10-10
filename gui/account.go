package gui

import (
	"errors"
	"strconv"
	"time"

	"github.com/twstrike/coyim/config"
	"github.com/twstrike/coyim/session"
	"github.com/twstrike/go-gtk/glib"
)

// someone who knows how to persist account configuration
type configManager interface {
	Save() error
}

type Account struct {
	ID                 string
	ConnectedSignal    *glib.Signal
	DisconnectedSignal *glib.Signal

	configManager
	*config.Config
	*session.Session
}

func (acc *Account) Connected() bool {
	return acc.ConnStatus == session.CONNECTED
}

var (
	errFingerprintAlreadyAuthorized = errors.New("the fingerprint is already authorized")
)

func (acc *Account) AuthorizeFingerprint(uid string, fingerprint []byte) error {
	existing := acc.UserIdForFingerprint(fingerprint)
	if len(existing) != 0 {
		return errFingerprintAlreadyAuthorized
	}

	acc.KnownFingerprints = append(acc.KnownFingerprints, config.KnownFingerprint{
		Fingerprint: fingerprint, UserId: uid,
	})

	return nil
}

func BuildAccountsFrom(multiAccConfig *config.MultiAccountConfig, manager configManager) []Account {
	accounts := make([]Account, 0, len(multiAccConfig.Accounts))

	for i := range multiAccConfig.Accounts {
		conf := &multiAccConfig.Accounts[i]
		account := newAccount(conf)
		account.configManager = manager
		accounts = append(accounts, account)
	}

	return accounts
}

func newAccount(conf *config.Config) Account {
	//id := conf.Account + "-" + strconv.FormatUint(uint64(time.Now().UnixNano()), 10)
	id := strconv.FormatUint(uint64(time.Now().UnixNano()), 10)

	return Account{
		ID:      id,
		Config:  conf,
		Session: session.NewSession(conf),

		ConnectedSignal:    glib.NewSignal(signalName(id, "connected")),
		DisconnectedSignal: glib.NewSignal(signalName(id, "disconnected")),
	}
}

func signalName(id, signal string) string {
	return "coyim-account-" + signal + "-" + id
}

func (u *gtkUI) showAddAccountWindow() {
	conf := config.NewConfig()
	account := Account{
		Config: conf,
	}

	accountDialog(account, func() error {
		err := u.configFileManager.Add(*conf)
		if err != nil {
			return err
		}

		return u.SaveConfig()
	})
}