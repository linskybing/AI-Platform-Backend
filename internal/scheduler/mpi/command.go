package mpi

import (
	"fmt"
)

// MPICommand generates the mpirun command for job execution
type MPICommand struct {
	HostfilePath string
	NumProcesses int
	WorkingDir   string
	Executable   string
	Args         []string
}

// Build constructs the mpirun command string
func (c *MPICommand) Build() string {
	cmd := fmt.Sprintf("mpirun -np %d", c.NumProcesses)

	if c.HostfilePath != "" {
		cmd += fmt.Sprintf(" --hostfile %s", c.HostfilePath)
	}

	if c.WorkingDir != "" {
		cmd += fmt.Sprintf(" --wdir %s", c.WorkingDir)
	}

	cmd += " --allow-run-as-root"
	cmd += " --mca btl_tcp_if_exclude lo,docker0"

	cmd += fmt.Sprintf(" %s", c.Executable)

	for _, arg := range c.Args {
		cmd += fmt.Sprintf(" %s", arg)
	}

	return cmd
}

// NewMPICommand creates a new MPI command builder
func NewMPICommand(executable string, numProcesses int) *MPICommand {
	return &MPICommand{
		Executable:   executable,
		NumProcesses: numProcesses,
		Args:         []string{},
	}
}

// WithHostfile sets the hostfile path
func (c *MPICommand) WithHostfile(path string) *MPICommand {
	c.HostfilePath = path
	return c
}

// WithWorkingDir sets the working directory
func (c *MPICommand) WithWorkingDir(dir string) *MPICommand {
	c.WorkingDir = dir
	return c
}

// WithArgs sets the arguments
func (c *MPICommand) WithArgs(args []string) *MPICommand {
	c.Args = args
	return c
}
