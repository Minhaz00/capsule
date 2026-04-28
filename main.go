package main

import (
	"fmt"
	"os"
	"os/exec"
)

const inspectScript = `
echo "========== CHILD PROCESS INSPECTION =========="
echo ""

echo "--- Process Identity ---"
echo "  PID:  $$"
echo "  PPID: $(grep PPid /proc/$$/status | awk '{print $2}')"
echo ""

echo "--- Hostname ---"
echo "  $(hostname)"
echo ""

echo "--- Filesystem Root ---"
ls /
echo ""

echo "--- Mount Table (first 10 entries) ---"
cat /proc/self/mounts | head -10
echo ""

echo "--- Network Interfaces ---"
cat /proc/net/dev | head -5
echo ""

echo "========== END INSPECTION =========="
`

func main() {
	fmt.Println("=== Containers Are Just Processes ===")
	fmt.Printf("Parent process PID: %d\n\n", os.Getpid())

	// Create a child process running /bin/sh with our inspection script
	cmd := exec.Command("/bin/sh", "-c", inspectScript)

	// Connect child's I/O to parent's I/O
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the child process and wait for it to finish
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running child process: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Child process exited. All output above came from the child.")
}
