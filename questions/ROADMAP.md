# gRPC !

# Data structures
    [*] Устройство slice
        [*] Как работает append?
        [*] Default value for a slice?
        [*] Slice/arrays comparable? (Array values are comparable if values of the array element type are comparable. Two array values are equal if their corresponding elements are equal.)
    [*] Устройство map 
        [*] Разрешение коллизий: метод открытой адресации, метод цепочек
        [*] Какой порядок обхода ключей map? 
        [*] В каком порядке выводит map функция fmt.Println? (всегда отсортированный список по ключам)
        [*] Почему стандартный порядок обхода ключей именно такой? (Зависит от многих факторов - какая хэш ф-ция использовалась, были ли эвакуации (итератор ходит по бакетам - старым и новым))
        [*] Можно ли взять адрес от элемента map? Почему? (Может произойти эвакуация данных и в ссылка изменится)
        [*] Зачем нужна sync map, если обычную map можно обернуть в мьютекс?
    [*] Устройство string
    [*] есть ли потокобезопасные структуры данных в го? Map? (Go не имеет предопределенных lock-free структур данных в стандартной библиотеке)

# Error
    * Error is vs Error as

# Interfaces
    [*] Сравнение интерфейса с nil (false, more answers/interfaces.txt)

# Горутины
    [*] Почему го рутины легковесны. Сравнение го рутины и потока ОС.
    [*] Что такое контекст свитч, стек, куча. Где аллоцируется стек, где куча.

# Каналы
    [*] Что такое каналы (https://www.youtube.com/watch?v=ZTJcaP4G4JM)
    [*] Разница буферизированные / не буферизированные каналы
    [*] Как устроены, как внутри устроен кольцевой буфер. Какие есть очереди. (Simple Queue or Linear Queue, Circular Queue, Priority Queue, Dequeue (Double-Ended Queue))
    * Почему нельзя сравнивать каналы?
    [*] Невалидные операции. Что будет, если дважды закрыть канал. Запись в закрытый канал. (Запись и чтение из нулевого канала блокирует горутину навсегда)

# Примитивы синхронизации
    [*] sync.WaitGroup
    * sync.Mutex
    * Что такое мьютекс, как он устроен, работает.
    [*] errgroup (https://www.codingexplorations.com/blog/mastering-concurrency-in-go-with-errgroup-simplifying-goroutine-management)
    * Что происходит с го рутиной, когда она вызывает метод Лок.
    * Что такое дэллоки, лайв Локи.

# Модели многозадачности: кооперативная и вытесняющая.
Какую модель многозадачности использует планировщик го.
Что происходит с го рутиной, когда она вызывает сискол.

# Параллельные вычисления
    * Data race vs race condition


## Разные вопросы
Как работает Scheduler https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html
Как работает Garbage collector 
По какой модели устроен. Какие фазы. На каких фазах происходит полная остановка программы.

Сравнение интерфейса с нил
Что такое контекст.

Data race vs race condition
* Gorouting
* Stack vs heap


## Транзакции и уровни изоляции

## HTTP vs HTTPS / HTTP2

TCP UDP
Deadlock

# Паттерны
Pipe
Fan in 
fan out
Worker pool

# Speed

# Database

# System design
* Event-driven design
* Saga pattern
* CQRS
* Event sourcing
* ACID
* CAP
* Rate limiter
* Cache
* Pub/Sub
* Load balancer
    * https://selectel.ru/blog/load-balancer-review/
    * https://habr.com/ru/companies/vk/articles/347026/
* Circuit breaker
* Service discovery
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