## The Standard Go Packages

- fmt [yes]

    用于格式化I/O。提供了标准I/O的支持。

    <>f  格式化

    F<>  文件

    S<>  字符串操作

- os
- io

    io包提供基本的I/O操作功能。

    Reader, Writer接口。

- io/ioutil

    ioutil实现了一些常用I/O工具函数。文件、路径操作。

- builtin
- bufio

    bufio包实现缓冲I/O。它封装了io.Reader, io.Writer对象，用于创建另一个也实现了该接口的对象，但提供了缓冲以及其他文本I/O功能。

- log

    
- strings
- bytes
- container
- errors
- math
- sort
- time
- unsafe
- more ...
