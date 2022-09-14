package helper

var Application = "golang" // able to call from outside because of lowcase
var letter = "huruf kecil" // cant call from outside because of lowcase

func SayHello(name string) string {
	return "Hello" + name
}

func SayHi(name string) string {
	return "Hi" + name
}
