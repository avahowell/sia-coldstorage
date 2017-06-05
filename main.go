package main

import (
	"fmt"
	"log"
	"os"
	"github.com/NebulousLabs/Sia/crypto"
	"github.com/NebulousLabs/Sia/modules"
	"github.com/NebulousLabs/Sia/types"
	"github.com/NebulousLabs/fastrand"
)

const nAddresses = 20;

// getAddress returns an address generated from a seed at the index specified
// by `index`.
func getAddress(seed modules.Seed, index uint64) types.UnlockHash {
	_, pk := crypto.GenerateKeyPairDeterministic(crypto.HashAll(seed, index))
	return types.UnlockConditions{
		PublicKeys:         []types.SiaPublicKey{types.Ed25519PublicKey(pk)},
		SignaturesRequired: 1,
	}.UnlockHash()
}

func main() {
	// generate a seed and a few addresses from that seed
	var seed modules.Seed
	fastrand.Read(seed[:])
	var addresses []types.UnlockHash
	seedStr, err := modules.SeedToString(seed, "english")

	// log if error
	if err != nil {
		log.Fatal(err)
	}

	// create new addresses and append to array
	for i := uint64(0); i < nAddresses; i++ {
		addresses = append(addresses, getAddress(seed, i))
	}

	// @todo review go docs and figure out a better
	// method of doing this so it is cleaner.
	// this is just a dirty hack to get things working
	fmt.Println("{");
	fmt.Println("    \"seed\": \"" + seedStr + "\", ");
	fmt.Println("    \"addresses\": [");

	//print addresses
	for _, address := range addresses {
		fmt.Println("        \"", address, "\",");
	}

	// close JSON object
	fmt.Println("    ]\n}");

	// exit application
	os.Exit(0)
}
