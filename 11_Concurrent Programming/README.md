## Summary of Concurrent Programming

### Poin-poin yang dipelajari dari materi ini yaitu.
- *Concurrent* program adalah sebuah program yang dapat berjalan secara independen dan memungkinkan berjalan secara bersamaan. *Concurrent* pada Golang adalah Go's *concurrency* (**goroutines**). 
- *Channel* adalah sebuah *object* yang memungkinkan goroutines dapat berkomunikasi satu sama lain. *Select* adalah sebuah tools yang mempermudah dalam mengontrol komunikasi data antarchannel.
- *Race condition* adalah sebuah kondisi dimana 2 *threads* mengakses sebuah memori secara bersamaan. *Race condition* terjadi karena adanya akses yang tidak sinkron. *Race condition* dapat diatasi dengan 3 cara, yaitu WaitGroups, channel blocking, dan mutex.