# gRPC !

# Data structures
     * Типы данных в го
    [*] Устройство slice
        [*] Как работает append?
        [*] Default value for a slice/map? (nil)
        [*] Slice/arrays comparable? (Array values are comparable if values of the array element type are comparable. Two array values are equal if their corresponding elements are equal).
    [*] Устройство map 
        [*] Разрешение коллизий: метод открытой адресации, метод цепочек
        [*] Какой порядок обхода ключей map? 
        [*] В каком порядке выводит map функция fmt.Println? (всегда отсортированный список по ключам)
        [*] Почему стандартный порядок обхода ключей именно такой? (Зависит от многих факторов - какая хэш ф-ция использовалась, были ли эвакуации (итератор ходит по бакетам - старым и новым))
        [*] Можно ли взять адрес от элемента map? Почему? (Нет, Может произойти эвакуация данных в новые бакеты и в ссылка изменится)
        [*] Зачем нужна sync map, если обычную map можно обернуть в мьютекс?
        [*] Доступ к элементам осуществляется за О(1) в идеальном случае.
    [*] Устройство string
    [*] есть ли потокобезопасные структуры данных в го? Map? (Go не имеет предопределенных lock-free структур данных в стандартной библиотеке). sync.Pool потокобезопасный (но дорогой), но не гарантирует хранение данных, GC может удалить их.

# Error
    * Error is vs Error as

# Interfaces
    [*] Сравнение интерфейса с nil (false, more answers/interfaces.txt)

# Горутины
    [*] Почему го рутины легковесны. Сравнение го рутины и потока ОС.
    [*] Что такое контекст свитч, стек, куча. Где аллоцируется стек, где куча.
    [*] Почему контекст свитч горутин быстрее.


# Каналы
    [*] Что такое каналы (https://www.youtube.com/watch?v=ZTJcaP4G4JM)
    [*] Разница буферизированные / не буферизированные каналы
    [*] Как устроены, как внутри устроен кольцевой буфер. Какие есть очереди. (Simple Queue or Linear Queue, Circular Queue, Priority Queue, Dequeue (Double-Ended Queue))
    * Почему нельзя сравнивать каналы?
    [*] Невалидные операции. Что будет, если дважды закрыть канал. Запись в закрытый канал. (Запись и чтение из нулевого канала блокирует горутину навсегда)

# Примитивы синхронизации
    [*] sync.WaitGroup
    * sync.Mutex. Что такое мьютекс, как он устроен, работает.
    [*] errgroup (https://www.codingexplorations.com/blog/mastering-concurrency-in-go-with-errgroup-simplifying-goroutine-management)
    * Что происходит с го рутиной, когда она вызывает метод Лок.
    * Что такое дэллоки, лайв Локи.

# Модели многозадачности: кооперативная и вытесняющая.
    * Какую модель многозадачности использует планировщик го. (preemptive)
    * Что происходит с го рутиной, когда она вызывает сискол.

# Параллельные вычисления
    [*] Data race vs race condition. 
        A data race occurs when multiple goroutines simultaneously access the same memory location and at least one of them is writing.
        A race condition is a behavior that depends on uncontrolled events (such as goroutine execution, how fast a message is published to a channel, or how long a call to a database lasts).

# Memory model
    [*] Stack vs heap

# Scheduler https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html
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

# Garbage collector
    [*] По какой модели устроен. Какие фазы. На каких фазах происходит полная остановка программы. (STW на стадии подготовки перед маркировкой, и на стадии завершения маркировки. Во время самой маркировки исполнение кода не останавливается). [GOGC, GOMEMLIMIT]. https://blog.golang.org/ismmkeynote


# Паттерны
    [*] Pipe
    [*] Fan in
    [*] Fan out
    [*] Worker pool
    [ ] Queuing

    * Semaphore
    * Rate limiter
    * Cache using timeAfter
    * Circuit breaker
    [*] Retry

## Разные вопросы

# Алгоритмы
    [ ] Quick sort
    [ ] Binary search

## Транзакции и уровни изоляции

## HTTP vs HTTPS / HTTP2
GraphQL

TCP UDP



# Speed


# Database

# System design
    * Event-driven design
    * Saga pattern
    * CQRS
    * Event sourcing
    * Circuit breaker
    * Service discovery
    * Kafka гаранитии доставки сообщений


# Load balancer
    * https://selectel.ru/blog/load-balancer-review/
    * https://habr.com/ru/companies/vk/articles/347026/



    * ACID
    * CAP
    * SOLID
    * DRY
    * KISS


* Pub/Sub




* Leader election
* Consistent hashing
* Sharding
* REST vs GraphQL


# Kubernetes

Аксиомы каналов:
https://dave.cheney.net/2014/03/19/channel-axioms


# Задачи
Развернуть односвязный список golang
Good understanding and experience in L2 to L7