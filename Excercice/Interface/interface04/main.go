package main

import "fmt"

type File interface {
	Open() string
	Close() string
}
type TextFile struct{
	fileName string
}

func (tf TextFile) Open() string {
	return "TextFile " + tf.fileName + " opened"
}
func (tf TextFile) Close() string {
	return "TextFile " + tf.fileName + " closed"
}

type ImageFile struct{
	imageName string
}

func (i ImageFile) Open() string {
	return "ImageFile " + i.imageName + " opened"
}
func (i ImageFile) Close() string {
	return "ImageFile " + i.imageName + " closed"
}
func ManageFile(f File) {
	fmt.Println(f.Open())
	fmt.Println(f.Close())
}

func main() {
	text := TextFile{"report.txt"}
	image := ImageFile{"photo.jpg"}

	ManageFile(text)
	ManageFile(image)
}