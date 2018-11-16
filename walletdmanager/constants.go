// Copyright (c) 2018, The TurtleCoin Developers
// Copyright (c) 2018, The Calex Developers
// Copyright (c) 2018, The CyprusCoin Developers
//
// Please see the included LICENSE file for more information.
//

package walletdmanager

const (
	// DefaultTransferFee is the default fee. It is expressed in XCY
	DefaultTransferFee float64 = 0.001

	logWalletdCurrentSessionFilename     = "turtle-service-session.log"
	logWalletdAllSessionsFilename        = "turtle-service.log"
	logCyprusCoindCurrentSessionFilename = "CyprusCoind-session.log"
	logCyprusCoindAllSessionsFilename    = "CyprusCoind.log"
	walletdLogLevel                      = "3" // should be at least 3 as I use some logs messages to confirm creation of wallet
	walletdCommandName                   = "turtle-service"
	cypruscoindCommandName               = "CyprusCoind"
)
