### Weekly Task 3
#### Aplikasi blogging sederhana dalam bentuk RESTful API
Berisi jawaban weekly task golang 3 dengan konten source code pada folder source-code dan screenshot tabel pada folder screenshots.

### How to Run
1. Clone this repository.
2. Copy the `.env` file.
   ```shell
   cp .env.example .env
   ```
3. Fill the values inside the `.env` file for the database configurations.
4. Create a new database called `simple-blog`.
   ```sql
   CREATE DATABASE "simple-blog";
   ```
5. Run the application with this command.
   ```shell
   go run main.go
   ```