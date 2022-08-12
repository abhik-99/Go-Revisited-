package models

type Book struct {
	Isbn        string  `bson:"isbn" json:"isbn"`
	Name        string  `bson:"name" json:"name"`
	Publication string  `bson:"pub" json:"pub"`
	Author      *Author `bson:"author" json:"author"`
}

type Author struct {
	Firstname    string `bson:"firstname" json:"firstname"`
	Lastname     string `bson:"lastname" json:"lastname"`
	BooksWritten uint   `bson:"numBooks" json:"numBooks"`
}
