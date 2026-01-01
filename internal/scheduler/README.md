# internal/scheduler/

Job scheduling system with support for multiple job types.

## Structure

- `queue/` - Priority queue management
- `executor/` - Job execution logic for different job types
- `mpi/` - MPI-specific job handling

## Features

- Priority-based scheduling (Course > Job)
- Extensible job type system (Normal, MPI, GPU)
- Resource availability checking
- Job lifecycle management (Queued → Scheduling → Running → Completed/Failed)
- MPI job coordination with OpenMPI

## Job Types

- **Normal**: Standard containerized jobs
- **MPI**: Distributed computing jobs using OpenMPI
- **GPU**: Jobs requiring dedicated GPU resources

## Priority Levels

- Course workloads: 1000 (pods, higher priority)
- Batch jobs: 100 (can be preempted)
