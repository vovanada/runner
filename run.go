package runner

import (
	"os"
	"os/exec"
)

type Runner struct {
	process string
	args    []string
	envs    []string
	cmd     *exec.Cmd
	stdout  []byte
	stderr  []byte
	close   chan struct{}
	die     chan struct{}
}

func (r *Runner) StdOut() []byte {
	return r.stdout
}

func (r *Runner) StdErr() []byte {
	return r.stderr
}

func (r *Runner) Stop() {
	r.close <- struct{}{}
}

func (r *Runner) PID() int {
	return r.cmd.Process.Pid
}

func (r *Runner) healthCheck() {
	for {
		select {
		case <-r.die:
			r.run()
		case <-r.close:
			r.cmd.Process.Kill()
			return
		}
	}
}

func (r *Runner) Signal(sig os.Signal) error {
	return r.cmd.Process.Signal(sig)
}

func (r *Runner) run() error {
	r.cmd = exec.Command(r.process, r.args...)
	r.cmd.Env = r.envs

	stdoutIn, err := r.cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderrIn, err := r.cmd.StderrPipe()
	if err != nil {
		return err
	}
	err = r.cmd.Start()

	if err != nil {
		return err
	}

	go r.healthCheck()

	go func() {
		captureOut(stdoutIn, &r.stdout)
	}()

	go func() {
		captureOut(stderrIn, &r.stderr)
	}()

	go func() {
		r.cmd.Wait()
		r.die <- struct{}{}
	}()

	return nil
}
