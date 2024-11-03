In Go, the goroutine context switching is faster than an OS-level context switch due to a few key factors related to its design and runtime. Here’s why:

# Lightweight Goroutines:
Goroutines are extremely lightweight compared to OS threads. 
They require only a small stack (around 2 KB, which can grow and shrink), whereas an OS thread usually has a 
larger fixed stack size. This lightweight nature reduces memory overhead and the time required to create, switch, 
and manage them.

# User-Space Scheduling:
Go has its own scheduler in the runtime, which operates in user space and is optimized for goroutine management. 
Unlike OS threads, which are managed by the kernel, goroutines are scheduled by the Go runtime, 
avoiding costly system calls that would otherwise be required to involve the OS in context switching.

# Avoiding Mode Switching:
OS thread context switches involve switching between user mode and kernel mode (a mode switch) when the OS scheduler 
is involved. In Go, the context switch between goroutines happens entirely in user space, avoiding this transition 
and saving significant time.

# Efficient M
    Scheduling:
Go uses an M
scheduling model, where many goroutines (M) are multiplexed onto fewer OS threads (N). 
This model is managed within the Go runtime, allowing multiple goroutines to run on a single thread or 
be reassigned to different threads if needed. This flexibility enables Go to avoid creating too many threads, 
which would lead to more expensive OS-level context switches.

# Low-Cost Stack Management:
Goroutines have dynamically growing stacks, meaning they start small and only increase as needed. 
This stack management is efficient and keeps the memory footprint low, allowing Go to handle millions of goroutines 
with low overhead.

These design choices make Go’s goroutine context switching faster and more efficient compared to 
OS-level thread context switching. However, this speed is most beneficial for applications with high concurrency
needs and many lightweight tasks that can be managed by goroutines rather than heavyweight OS threads.


# В какой момент проходит переключение контекста?
    * The use of keyword go
    * Grabage collection
    * System calls
    * Syncronization and Orchestration