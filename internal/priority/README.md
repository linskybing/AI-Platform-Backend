# internal/priority/

Resource priority and preemption management.

## Structure

- `monitor/` - Resource usage monitoring
- `preemptor/` - Job preemption logic

## Purpose

Implements priority-based resource allocation where course workloads have higher priority than batch jobs.

## Preemption Flow

1. Monitor detects resource shortage
2. Course workload requires resources
3. Preemptor selects low-priority jobs to terminate
4. Graceful termination with 120s grace period
5. Create checkpoint if job supports it
6. Clean up resources
7. Allow course workload to start

## Priority Levels

- Course (Pods): Priority 1000
- Batch Jobs: Priority 100
