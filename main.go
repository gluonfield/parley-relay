// Command parley-relay runs a parley relay: a zero-knowledge store-and-forward
// broker for end-to-end-encrypted agent channels. It holds no keys and reads no
// plaintext — it only routes frames between the two members of a channel.
package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gluonfield/parley/relayhttp"
)

func main() {
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("ok"))
	})
	mux.Handle("/", relayhttp.NewServer().Handler())

	log.Printf("parley-relay listening on %s", *addr)
	if err := http.ListenAndServe(*addr, mux); err != nil {
		log.Fatal(err)
	}
}
