package decorator

import "fmt"

//Decorator
type DataSource interface {
	WriteData(data string)
}

type FileDataSource struct{}

func (f FileDataSource) WriteData(data string) {
	fmt.Println("Saving to file:", data)
}

type EncryptedDataSource struct {
	fSD DataSource
}

func (e EncryptedDataSource) WriteData(data string) {
	runes := []rune(data)

	fmt.Println("encrypting...")
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	e.fSD.WriteData(string(runes))
}

type LoggerDataSource struct {
	fSD DataSource
}

func (l LoggerDataSource) WriteData(data string) {
	fmt.Println("Log: saving data...")
	l.fSD.WriteData(data)
}

func main() {
	fds := FileDataSource{}
	enc := EncryptedDataSource{fSD: fds}
	log := LoggerDataSource{fSD: enc}

	log.WriteData("Adan Ayaz <==> zayA nadA")
}
