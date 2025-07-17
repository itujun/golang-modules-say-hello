## Repository Berkaitan

[golang-app-say-hello](https://github.com/itujun/golang-app-say-hello)

## Langkah-langkah

1. buat repository git (golang-modules-say-hello)

2. repository boleh set ke public

3. buka folder projek yang telah terbuat di pc/laptop (golang-modules-say-hello) lalu ketik dan jalankan perintah:

```bash
go mod init github.com/itujun/golang-modules-say-hello
```

4. buat file pada projek dengan nama say_hello.mod yang berisi kode berikut:

```bash
package golang_modules_say_hello

// nama function harus diawali dengan huruf besar agar bisa diakses di luar package
func SayHello() string {
	return "Hello"
}
```

5. init projek ke git

```bash
git init
```

6. git add

```bash
git add .
```

7. git commit

```bash
git commit -m "membuat module say-hello"
```

8. git remote

```bash
git remote add origin https://github.com/itujun/golang-modules-say-hello.git
```

9. git push

```bash
git push -u origin main
```

10. rilis versi module

```bash
git tag v1.0.0
```

11. push lagi

```bash
git push origin v1.0.0
```

# Menambahkan Dependency

12. buat repository git (golang-app-say-hello)

13. repository boleh set ke public

Langkah ke-14 lanjut repository terkait [golang-app-say-hello](https://github.com/itujun/golang-app-say-hello)
