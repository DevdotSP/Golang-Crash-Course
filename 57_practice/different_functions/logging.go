package diffuc

import (
	"log"
	"net/smtp"
)

func init(){
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}

func ExampleLog() {
// Println writes to the standard logger.
log.Println("main started")

// Fatalln is Println() followed by a call to os.Exit(1)
	log.Fatalln("fatal message")

// Panicln is Println() followed by a call to panic()
	log.Panicln("panic message")

	otherExample()
}



func otherExample() {
	// Connect to the remote SMTP server.
client, err := smtp.Dial("hisnameisraph@gmail.com")
if err != nil {	
	log.Fatalln(err)
}
client.Data()
}