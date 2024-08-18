# solution_ch_part1

Saya berasumsi bahwa Task Breakdown pertama adalah membuat REST API yang menerima request ndjson untuk melakukan perhitungan OHLC, dengan mengunakan kafka dan juga redis.
service rest tersebut dimulai pada directory "src" pada solution_ch_part1
> untuk menjalankannya, buka terminal pada solution_ch_part1 lalu ketikan go mod tidy untuk menginstall package yang dibutuhkan
> kemudian ketikan go run main.go maka service akan berjalan pada port 9000
struktur pada directory src adalah struktur yg biasa saya gunakan sebagai boilerplate saya.

Untuk Task Breakdown yang kedua, saya membuat directory "grpc" pada solution_ch_part1.
> untuk menjalannya, buka terminal pada directory "grpc" lalu ketikan go mod tidy untuk menginstall package yang dibutuhkan
> kemudian ketikan go run main.go maka service akan berjalan pada port 9005

keduanya memiliki file .env dan dapat diganti isinya sesuai kebutuhan

# disclaimer
saya masih baru untuk GRPC karena belum pernah menggunakanya di pekerjaan. jadi masih sebatas belajar autodidak saja. jadi mungkin dari segi struktur, skeleton, dll, belum merupakan best practice. dan saya tertarik untuk mempraktikannya lebih lanjut dipekerjaan.

> untuk dokumentasi gambar, saya membuat 2 gambar. yang pertama adalah service diagram yang menggambarkan alur dari service ini. Dan saya juga membuat flow chart diagram dari proses perhitungan OHLC. Kedua file tersebut ada pada root directory solution_ch_part1

> untuk unit test sudah saya buat pada usecase layer dan sudah coverage 100%. untuk melihat code coverage bila menggunakan vscode yaitu dengan membuka file src/app/usecase/ohlc/ohlc_test.go lalu klik kanan dan pilih Go : toogle test coverage in current package. pada tab output di terminal akan menampilkan code coverage sebagai berikut :
> Running tool: /usr/local/go/bin/go test -timeout 30s -coverprofile=/var/folders/h_/tjhvlj3n32sc9lvvfbm8x9ym0000gn/T/vscode-goKZDCYC/go-code-cover solution_ch_part1/src/app/usecase/ohlc

> ok  	solution_ch_part1/src/app/usecase/ohlc	1.093s	coverage: 100.0% of statements

> untuk kafka dan redis, saya juga menyediakan pada file docker-compose.yaml

Terimakasih.

Best Regards,
Jody Almaida Putra.
