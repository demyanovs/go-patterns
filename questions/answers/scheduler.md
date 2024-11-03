
    GMP model.
    P - Virtual Cores
    M - Machine. OS Thread, is assigned to a P. Managed by OS, OS is responsible for placing the Tread on a core.
    G - Goroutine. Every Go programm is also givent an initial Gourotine (G). 
        Gourintes are application level threads. They are managed by the Go runtime. G are context-switched on and off an M.
        GRQ (Global Run Queue) / LRQ
        Each P has its own LRQ. When a G is created, it is placed on the local run queue of the P that created it.
        GRQ is for Goroutines that are not assigned to a P yet.
        When a P runs out of Gs in its local run queue, it will attempt to steal Gs from other P's run queues.
        
    Net poller is responsible for asynchronous system calls. The net poller has OS Thread and it's handling an efficient event loop.