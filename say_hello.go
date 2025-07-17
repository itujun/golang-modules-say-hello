package golang_modules_say_hello

// nama function harus diawali dengan huruf besar agar bisa diakses di luar package
func SayHelloi(user string) string {
	return "Hello," + user
}