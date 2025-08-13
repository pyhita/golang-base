package wait

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail.
func waitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)

	for retry := 0; time.Now().Before(deadline); retry++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s);retrying…", err)
		time.Sleep(time.Second << uint(retry))
	}

	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
