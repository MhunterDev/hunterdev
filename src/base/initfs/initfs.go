package initFS

import (
	"os"
	"time"

	easy "github.com/MhunterDev/hunterdev/src/base/encoder"
	logsalot "github.com/MhunterDev/hunterdev/src/base/logs"
)

func BuildFS() error {

	// Build Logging directories
	os.MkdirAll("/usr/mhdev/logs", 077)
	os.Create("/usr/mhdev/logs/Install.log")
	os.Create("/usr/mhdev/logs/Base-processes.log")
	os.Create("/usr/mhdev/logs/pgrunner.log")
	logsalot.LogInit("Log directories")
	time.Sleep(1 * time.Second)

	//Build the Keychains
	err := os.MkdirAll("/usr/mhdev/keychain/tls/secret", 077)
	if err != nil {
		return err
	}

	os.Create("/usr/mhdev/keychain/tls/secret/CA.key")
	os.Create("/usr/mhdev/keychain/tls/CA.crt")
	os.Create("/usr/mhdev/keychain/secret.pem")
	logsalot.LogInit("Keychain")
	time.Sleep(1 * time.Second)

	//Populate the secrets
	easy.MakeSecret()
	easy.GenerateHTTPS()
	logsalot.LogInit("Secrets")
	time.Sleep(2 * time.Second)

	return nil
}
