package main

import "fmt"

type Editor struct {
	text string
}

func (e *Editor) Write(text string) {
	e.text += text
}

func (e *Editor) Read() string {
	return e.text
}

type Memento struct {
	text string
}

func (m *Memento) Text() string {
	return m.text
}

type Caretaker struct {
	memento *Memento
}

func (c *Caretaker) Save(editor *Editor) {
	c.memento = &Memento{text: editor.Read()}
}

func (c *Caretaker) Restore(editor *Editor) {
	editor.Write(c.memento.Text())
}

func main() {
	editor := &Editor{}
	caretaker := &Caretaker{}

	editor.Write("Hello, ")
	caretaker.Save(editor)

	editor.Write("world!")
	fmt.Println(editor.Read())

	caretaker.Restore(editor)
	fmt.Println(editor.Read())

}
