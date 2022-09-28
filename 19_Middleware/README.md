## Summary of Middleware

### Poin-poin yang dapat dipelajari dari materi ini adalah sebagai berikut.
- *Middleware* adalah sebuah *entity* yang kita pasangkan di dalam proses server. Proses yang dimaksud adalah ketika *client* melakukan *request* dan server memberikan *response*, maka akan dipasangkan sebuah layer/*middleware*. Pada layer tadi dipasangkan fungsi-fungsi khusus. Contoh tools *middleware* yaitu Negroni, Echo, Interpose, Alice, dsb. Pada echo *middleware* terdapat 2 jenis *setup* yaitu `Echo#Pre()` dan `Echo#Use()`.
- Log pada *Middleware*. Log berguna untuk mencatat log ketika ada HTTP *request*, dan log juga dapat digunakan untuk analisis.
- *Authentication* pada *Middleware*. *Authentication* digunakan untuk *user identification* dan mengamankan data. Jenis *authentication* pada *middleware* terbagi dua yaitu basic *authentication* dan JSON web token.