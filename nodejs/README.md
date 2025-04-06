# Nodejs
Is a collection of dependencies that allow you to run JavaScript on the server side.

## V8 project
JavaScript engine created by Google. 
The goal of the project is to execute JavaScript code outside of a browser.

## Libuv project
C++ open-source project gives Node access to the OS's file system, concurrency, networking, and other operating system functions.
V8 is used to interpret and execute Javascript code, while libuv is used for accessing the filesystem and some aspects of concurrency.

-----
There is a process binding which binds js code and c++ functions and V8 coverts values between JS and C++.
Libuv gives Node access to underlying OS.

# Event loop
[node_04.png](scr/node_04.png)
[node_05.png](scr/node_05.png)
Event loop is a single thread which checks different events:
 * pendingTimers (setTimeout, setInterval, setImmediate)
 * pendingOSTasks
 * pending long running operations
 * callbacks

While event loop is single-threaded, all I/O and calls to native APIs are either asynchronous, or run on a separate thread.

# Thread pool
[node_06.png](scr/node_06.png)
[node_07.png](scr/node_07.png)
Libuv has a thread pool with size 4 by default

Some functions calls in a standard library are delegated to the underlying OS entirely outside of event loop.

# How to block event loop
let flag = false;
setTimeout(() => {
    flag = true;
}, 1000);

while (!flag) {
    // do something
}

or some long read file operation will block for a while.

# Nodejs problems by design
 * Dependency management

# OSI
https://www.youtube.com/watch?v=5F1MA9JCfjM

7. Прикладной (application)
6. Представления (presentation)    (Данные (data) | Преобразование данных)
5. Сеансовый (session)
4. Транспортный (transport) (Сегменты (segment) | Контроль над передачей данных)
3. Сетевой (network) (Пакеты (packet) | Определение маршрута и логическая адресация)
2. Канальный (data link) (Биты (bit)/ Кадры (frame) | Физическая адресация)
1. Физический (physical) (Биты (bit) | Работа со средой передачи, сигналами и двоичными данными)

## REST
tell what to do (GET, POST) in headers what to do and it's transport and not message.