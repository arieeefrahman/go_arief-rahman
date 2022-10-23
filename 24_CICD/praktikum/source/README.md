### Praktikum CI/CD
#### Implementasi CI/CD menggunakan Github Actions
Berisi jawaban weekly task golang 3 dengan konten source code pada folder source-code dan screenshot tabel pada folder screenshots.

### How to Run
1. Clone this repository.
2. Copy the `.env` file.
   ```shell
   cp .env.example .env
   ```
3. Fill the values inside the `.env` file for the database configurations.
4. Create a new database called `echo_notes`.
   ```sql
   CREATE DATABASE "echo_notes";
   ```
5. Run the application with this command.
   ```shell
   go run main.go
   ```

### About Github Actions
1. First, copy the content in `go.yml` in `.github` folder, then delete `.github` folder.
2. Initiate Github Actions.
3. Setup your `DOCKER_USERNAME` and `DOCKER_PASSWORD` in your repository by accessing Settings > Actions > New repository secret.
4. Paste into `[file-name].yml` file.
5. Wait for Github Actions to run.