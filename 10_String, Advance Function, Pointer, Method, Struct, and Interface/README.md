## Summary of String, Advance Function, Pointer, Method, Struct, and Interface

### Tiga poin yang dipelajari dari materi ini adalah sebagai berikut.
- Ada beberapa function yang terdapat pada string yaitu `len()`, `strings.Compare()`, `strings.Contains()`, `strings.Substring()`, `strings.Replace()`, dan `strings.Insert()`. Untuk menggunakan *functions* tersebut, kita harus *import library strings* terlebih dahulu dengan cara `import "strings"`.
- Bentuk-bentuk *function* seperti *variadic function*, *anonymous function* atau *literal function*, dan *closure*. *Variadic function* dapat dideklarasikan dengan caa menambahkan `...` setelah nama variabel, contohnya `func sum(numbers ...int) int {}`. *Anonymous function* atau *literal function* adalah sebuah fungsi yang tidak memiliki nama. *Closure* adalah *special type* dari *anonymous function* yang mereferensikan variabel yang dideklarasikan diluar dari fungsinya sendiri.
- *Pointer* adalah variabel yang menyimpan alamat memori dari variabel lain. *Value* pada *pointer* bergantung pada *value* dari variabel lain yang diambil. Cara mendeklarasikannya adalah `var <variable_name> *<data_type> = &<reference_variable>.`
- Struct adalah sebuah tipe data yang dapat menampung berbagai macam tipe data variabel dan method. Cara kerja struct mirip dengan *object oriented programming* (OOP) pada bahasa pemrograman lain. Cara mendeklarasikannya yaitu.
```
type struct_variable_name struct {
    field <data_type>
    field <data_type>
    ...
    field <data_type>
}
```
- Method adalah sebuah function yang biasanya terdapat dalam struct. Cara mengdeklarasikannya yaitu.
```
func (receiver StructType) MethodName(parameterList) (returnTypes) {
    // block statement
}
```