package btc_test

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/shiftdevices/godbb/backend/coins/btc"
	blockchainMock "github.com/shiftdevices/godbb/backend/coins/btc/blockchain/mocks"
	"github.com/shiftdevices/godbb/backend/keystore/mocks"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type accountSuite struct {
	suite.Suite

	net            *chaincfg.Params
	keyStoreMock   mocks.Keystore
	blockchainMock blockchainMock.Interface
	onEvent        func(btc.Event)
	account        *btc.Account

	log *logrus.Entry
}

// func (s *accountSuite) SetupTest() {
// 	s.log = logging.Log.WithGroup("btc_test")
// 	s.net = &chaincfg.TestNet3Params
// 	s.onEvent = func(btc.Event) {}
// 	var err error

// 	const xpubSerialized = "tpubDEXZPZzoVxHQdZg6ndWKoDXwsPtfTKpYsF6SDCm2dHxydcNvoKM58RmA7FDj3hXqy8BrxfwoTNaV5SzWgCzurTaQmDNywHVvv5tPSj6Evgr"
// 	xpub, err := hdkeychain.NewKeyFromString(xpubSerialized)
// 	if err != nil || xpub.IsPrivate() {
// 		panic(err)
// 	}

// 	db, err := transactionsdb.NewDB(test.TstTempFile("godbb-db-"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	s.keyStoreMock.On("XPub").Return(xpub)
// 	s.account, err = btc.NewAccount(
// 		s.net,
// 		db,
// 		signing.NewEmptyAbsoluteKeypath(),
// 		&s.keyStoreMock,
// 		&s.blockchainMock,
// 		&headersMock.Interface{},
// 		addresses.AddressTypeP2PKH,
// 		s.onEvent,
// 		s.log,
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func TestAccountSuite(t *testing.T) {
// 	suite.Run(t, &accountSuite{})
// }
