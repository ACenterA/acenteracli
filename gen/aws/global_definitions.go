package aws

import (
	"os"

	// "github.com/wallix/awless/cloud/rdf"
)

type glob struct {
	AwlessLabel   string
	Label         string
	Value         string
}

func IfThenElse(condition bool, a string, b string) string {
        if condition {
                return a
        }
        return b
}

var (
        ENC_PWD              = IfThenElse(os.Getenv("ACENTERA_ENCRYPT_KEY") != "", os.Getenv("ACENTERA_ENCRYPT_KEY"), "ACenter4")
        API_ENDPOINT         = IfThenElse(os.Getenv("ACENTERA_API_ENDPOINT") != "", os.Getenv("ACENTERA_API_ENDPOINT"), "https://cms-api-v1.dev.acentera/")
        Verbose              = IfThenElse(os.Getenv("ACENTERA_CLI_VERBOSE") != "", os.Getenv("ACENTERA_CLI_VERBOSE"), "")
)

func init() {
        if (ENC_PWD == "") {
          ENC_PWD = "ACenter4"
        }
}

var GlobalDefinitions = []glob{
	{AwlessLabel: "ENC_PWD", Value: ENC_PWD},
	{AwlessLabel: "API_ENDPOINT", Value: API_ENDPOINT},
	{AwlessLabel: "Verbose", Value: Verbose},
}
