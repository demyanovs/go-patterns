# Generative
1. [Generator](./01-generative/01-generator/main.go)
2. [Fan in](./01-generative/02-fan-in/main_ver2.go)
3. [Fan out](./01-generative/04-fan-out/main_ver2.go)
4. [Pipeline](./01-generative/05-pipeline/main.go)

# Sync
1. [Mutex (rate limiter)](02-sync/01-mutex/main.go)
2. [Semaphore (rate limiter)](02-sync/02-semaphore/main.go)

# Parallel Computing
1. [Worker pool](./03-parallel-computing/01-worker-pool/main3.go)
2. [Queuing](./03-parallel-computing/02-queuing/main.go)
3. [Parallel for loop](./03-parallel-computing/03-parallel-for-loop/main.go)
4. [Error group](./03-parallel-computing/05-errorgroup/main.go)

# Delayed Computing
1. [Futures (Promise)](./04-delayed-computing/01-futures/main.go)
2. Lazy evaluation

# General Patterns
1. Circuit breaker
2. [Retry/timeout](stability/retry/main.go) / [Retry/timeout ver 2](stability/retry_ver2/main.go)
3. Cache
4. [Configurable object](00-general_design/configurable_object/main.go)
5. [Functional options](00-general_design/functional_options/main.go)
6. Error group. Several requests, if one fails, all fail.
7. RW map