package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/skratchdot/open-golang/open"

	"github.com/NebulousLabs/Sia/crypto"
	"github.com/NebulousLabs/Sia/modules"
	"github.com/NebulousLabs/Sia/types"
	"github.com/NebulousLabs/fastrand"
)

const outputTmpl = `
<html>
	<head>
		<title> Sia Cold Storage </title>
	</head>
	<style>
		body {
			font-family: "Gotham A", "Gotham B", Helvetica, Arial, sans-serif;
			margin-left: auto;
			margin-right: auto;
			max-width: 900px;
			text-align: center;
		}
		.info {
			margin-top: 75px;
		}
	</style>
	<body>
		<h3>Sia cold wallet successfully generated.</h3>
		<p> Please save the information below in a safe place. You can use the Seed to recover any money sent to any of the addresses, without an online or synced wallet. Make sure to keep the seed safe, and secret.</p>
		<section class="info">
			<section class="seed">
				<h4>Seed: </h4>
				<p>{{.Seed}}</p>
			</section>
			<section class="addresses">
				<h4>Addresses: </h4>
				<ul>
				{{ range .Addresses }}
					<li>{{.}}</li>
				{{ end }}
			</section>
		</section>
	</body>
</html>
`

const nAddresses = 20

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
	var seed modules.Seed
	var seedStr string

	// get a seed
	var seedErr error
	if len(os.Args) > 1 {
		// non-zero arguments: read seed words
		var words []string
		if len(os.Args[1:]) == 1 {
			words = strings.Fields(os.Args[1])
		} else {
			words = os.Args[1:]
		}
		if len(words) != 29 {
			log.Fatal("29 seed words required")
		}
		seedStr = strings.Join(words[:], " ")
		seed, seedErr = modules.StringToSeed(seedStr, "english")
	} else {
		// zero arguments: generate a seed
		fastrand.Read(seed[:])
		seedStr, seedErr = modules.SeedToString(seed, "english")
	}
	if seedErr != nil {
		log.Fatal(seedErr)
	}

	// generate a few addresses from that seed
	var addresses []types.UnlockHash
	for i := uint64(0); i < nAddresses; i++ {
		addresses = append(addresses, getAddress(seed, i))
	}

	templateData := struct {
		Seed      string
		Addresses []types.UnlockHash
	}{
		Seed:      seedStr,
		Addresses: addresses,
	}
	t, err := template.New("output").Parse(outputTmpl)
	if err != nil {
		log.Fatal(err)
	}
	l, err := net.Listen("tcp", "localhost:8087")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, templateData)
		l.Close()
		close(done)
	})
	go http.Serve(l, handler)

	err = open.Run("http://localhost:8087")
	if err != nil {
		// fallback to console, clean up the server and exit
		l.Close()
		fmt.Println("Seed:", seedStr)
		fmt.Println("Addresses:")
		for _, address := range addresses {
			fmt.Println(address)
		}
		os.Exit(0)
	}
	<-done
}
