package rest

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"
)

func HandleInfo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Always set content type and status code
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	// And write your response to w
	var ts = time.Now().UTC()
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)


	ma := float32(mem.Alloc) / 1024.0 / 1024.0                    //  convert to MB
	ta := float32(mem.TotalAlloc) / 1024.0 / 1024.0               //  convert to MB
	sm := float32(mem.Sys) / 1024.0 / 1024.0                      //  convert to MB
	fm := float32(mem.Sys-mem.Alloc) / 1024.0 / 1024.0            //  convert to MB

	message := []string{
		fmt.Sprintf("Received Info on ... %s \n", ts.Format(time.RFC850)),
		fmt.Sprint("Memory Stats\n"),
		fmt.Sprintf("  Alloc:       %f  (MB allocated and not yet freed)\n", ma),
		fmt.Sprintf("  Total Alloc: %f  (MB allocated (even if freed))\n", ta),
		fmt.Sprintf("  Sys:         %f  (MB obtained from system)\n", sm),
		fmt.Sprintf("  Free Memory: %f  (MB Sys - Alloc)\n", fm),
		fmt.Sprintf("  Lookups:     %d  (number of pointer lookups)\n", mem.Lookups),
		fmt.Sprintf("  Mallocs:     %d  (number of mallocs)\n", mem.Mallocs),
		fmt.Sprintf("  Frees:       %d  (number of frees)\n", mem.Frees),

	}

	fmt.Fprintf(w, strings.Join(message, ""))
}
