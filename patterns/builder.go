package patterns

import "fmt"

type Essay struct {
}

type Writer interface {
	Intro()
	Main()
	Concl()
	WriteEssay() *Essay
}

type MyWriter struct {
	Essay *Essay
}

func (writer *MyWriter) Intro() {
	fmt.Println("Introduction is written!")
}
func (writer *MyWriter) Main() {
	fmt.Println("Main part is written!")
}
func (writer *MyWriter) Concl() {
	fmt.Println("Conclusion is written!")
}

func (writer *MyWriter) WriteEssay() *Essay {
	fmt.Println("Essay is written!")
	return writer.Essay
}

func NewWriter() *MyWriter {
	return &MyWriter{
		&Essay{},
	}
}

type Redactor struct {
	Writer Writer
}

func NewRedactor(writer Writer) *Redactor {
	return &Redactor{writer}
}

func (r *Redactor) Publish() *Essay {
	r.Writer.Intro()
	r.Writer.Main()
	r.Writer.Concl()
	return r.Writer.WriteEssay()
}
