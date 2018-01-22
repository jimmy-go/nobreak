package nobreak

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// NoBreak type.
type NoBreak struct {
	Config *Config
}

// New returns a no-break from config filename.
func New(filename string) (*NoBreak, error) {
	c, err := LoadConfig(filename)
	if err != nil {
		return nil, err
	}
	nb := &NoBreak{
		c,
	}
	return nb, nil
}

// Run runs nobreak server.
func (nb *NoBreak) Run() error {
	log.Printf("Listen port: %d", nb.Config.Port)
	log.Printf("Admin port: %d", nb.Config.AdminPort)
	var wg sync.WaitGroup
	wg.Add(2)
	errc := make(chan error, 2)
	// Start server.
	go func() {
		defer wg.Done()
		errc <- nb.runTunnel()
	}()
	// Start admin.
	go func() {
		defer wg.Done()
		errc <- nb.runAdmin()
	}()
	wg.Wait()
	err := <-errc
	return err
}

func (nb *NoBreak) runTunnel() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Handler(nb.Config))
	addr := fmt.Sprintf(":%d", nb.Config.Port)
	var err error
	if nb.Config.TLSEnabled {
		err = http.ListenAndServeTLS(addr, nb.Config.TLSCert, nb.Config.TLSKey, mux)
	} else {
		err = http.ListenAndServe(addr, mux)
	}
	if err != nil {
		return err
	}
	return nil
}

func (nb *NoBreak) runAdmin() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello admin. (WIP)")
	}))
	addr := fmt.Sprintf(":%d", nb.Config.AdminPort)
	err := http.ListenAndServe(addr, mux)
	return err
}
