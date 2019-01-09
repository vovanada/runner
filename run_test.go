package runner

import (
	"syscall"
	"testing"
	"time"
)

func TestRunProcess(t *testing.T) {
	r, _:=RunProcess("ping", []string{"127.0.0.1", "-n", "100"}, []string{})

	pid := r.PID()
	time.Sleep(time.Second)
	syscall.Kill(pid, syscall.SIGKILL)

	if pid == r.PID() {
		t.Errorf("pid must be changed, %v - %v", pid, r.PID())
	}
}