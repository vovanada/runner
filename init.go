package runner

func RunProcess(process string, args []string, envs []string) (*Runner, error) {
	r := &Runner{
		process: process,
		args:    args,
		envs:    envs,
	}

	r.close = make(chan struct{})
	r.die = make(chan struct{})

	err := r.run()

	return r, err
}
