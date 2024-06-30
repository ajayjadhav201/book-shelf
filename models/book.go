package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID             primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	BookID         string             `json:"id,omitempty"`
	Title          string             `json:"title,omitempty"`
	Author         []string           `json:"author,omitempty"`
	Language       string             `json:"language,omitempty"`
	Publisher      string             `json:"publisher,omitempty"`
	PublishingDate string             `json:"publishing date,omitempty"`
	Genre          string             `json:"genre,omitempty"`
	ISBN           string             `json:"isbn,omitempty"`
	Edition        string             `json:"edition,omitempty"`
	Pages          int                `json:"pages,omitempty"`
	Summery        string             `json:"summery,omitempty"`
	AboutAuthor    []string           `json:"about author,omitempty"`
}

func NewBook() {
	//
}

func SampleBooks() []*Book {

	return []*Book{
		{
			BookID:         "B007",
			Title:          "The Go Programming Language",
			Author:         []string{"Alan A. A. Donovan", "Brian W. Kernighan"},
			Language:       "English",
			Publisher:      "Addison-Wesley Professional",
			PublishingDate: "2015-10-26",
			Genre:          "Programming",
			ISBN:           "9780134190440",
			Edition:        "1st",
			Pages:          400,
			Summery:        "The Go Programming Language is the authoritative resource for any programmer who wants to learn Go. It covers the language, its libraries, and the many ways in which they can be used to build powerful software.",
			AboutAuthor:    []string{"Alan A. A. Donovan is a principal engineer at Google, working on Go.", "Brian W. Kernighan is a professor of computer science at Princeton University and co-author of several classic programming books."},
		},
		{
			BookID:         "B006",
			Title:          "Rich Dad Poor Dad",
			Author:         []string{"Robert T. Kiyosaki"},
			Language:       "English",
			Publisher:      "Warner Books Ed",
			PublishingDate: "2000-04-01",
			Genre:          "Personal Finance",
			ISBN:           "0446677450",
			Edition:        "1st",
			Pages:          207,
			Summery:        "Rich Dad Poor Dad is a book by Robert T. Kiyosaki that advocates the importance of financial literacy, financial independence, and building wealth through investing in assets, real estate investing, starting and owning businesses, as well as increasing one's financial intelligence.",
			AboutAuthor:    []string{"Robert T. Kiyosaki is an American businessman and author. He is the founder of Rich Global LLC and the Rich Dad Company, a private financial education company that provides personal finance and business education to people through books and videos."},
		},
		{
			BookID:         "B002",
			Title:          "Introduction to Algorithms",
			Author:         []string{"Thomas H. Cormen", "Charles E. Leiserson", "Ronald L. Rivest", "Clifford Stein"},
			Language:       "English",
			Publisher:      "MIT Press",
			PublishingDate: "2009-07-31",
			Genre:          "Computer Science",
			ISBN:           "0262033844",
			Edition:        "3rd",
			Pages:          1312,
			Summery:        "A comprehensive textbook covering a wide range of algorithms in depth.",
			AboutAuthor:    []string{"Thomas H. Cormen is a Professor of Computer Science.", "Charles E. Leiserson is a Professor of Computer Science and Engineering."},
		},
		{
			BookID:         "B003",
			Title:          "Clean Code: A Handbook of Agile Software Craftsmanship",
			Author:         []string{"Robert C. Martin"},
			Language:       "English",
			Publisher:      "Prentice Hall",
			PublishingDate: "2008-08-01",
			Genre:          "Software Engineering",
			ISBN:           "0132350882",
			Edition:        "1st",
			Pages:          464,
			Summery:        "A guide to producing clean, readable, and maintainable code.",
			AboutAuthor:    []string{"Robert C. Martin, also known as Uncle Bob, is a software engineer and author."},
		},
		{
			BookID:         "B004",
			Title:          "The Pragmatic Programmer: Your Journey to Mastery",
			Author:         []string{"Andrew Hunt", "David Thomas"},
			Language:       "English",
			Publisher:      "Addison-Wesley Professional",
			PublishingDate: "2019-09-13",
			Genre:          "Software Development",
			ISBN:           "0135957052",
			Edition:        "2nd",
			Pages:          352,
			Summery:        "A guide to the principles and practices of pragmatic programming.",
			AboutAuthor:    []string{"Andrew Hunt is a programmer and author.", "David Thomas is a software developer and author."},
		},
		{
			BookID:         "B005",
			Title:          "Design Patterns: Elements of Reusable Object-Oriented Software",
			Author:         []string{"Erich Gamma", "Richard Helm", "Ralph Johnson", "John Vlissides"},
			Language:       "English",
			Publisher:      "Addison-Wesley Professional",
			PublishingDate: "1994-10-31",
			Genre:          "Software Engineering",
			ISBN:           "0201633612",
			Edition:        "1st",
			Pages:          416,
			Summery:        "A catalog of simple and succinct solutions to commonly occurring design problems.",
			AboutAuthor:    []string{"Erich Gamma is a software engineer and author.", "Richard Helm is a consultant and author."},
		},
	}
}
