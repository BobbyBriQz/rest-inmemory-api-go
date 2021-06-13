package models

type Book struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	NoOfPages int64  `json:"noOfPages"`
	Author    Author `json:"author"`
}

func (b Book) GetID() int64 {
	return b.ID
}
func (b *Book) SetID(i int64) {
	b.ID = i
}

func (b Book) GetTitle() string {
	return b.Title
}

func (b *Book) SetTitle(t string) {
	b.Title = t
}
func (b Book) GetNoOfPages() int64 {
	return b.NoOfPages
}
func (b *Book) SetNoOfPages(n int64) {
	b.NoOfPages = n
}
func (b Book) GetAuthor() Author {
	return b.Author
}
func (b *Book) SetAuthor(a Author) {
	b.Author = a
}
