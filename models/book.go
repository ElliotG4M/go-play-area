package models

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) ToString() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n",
		b.Title,
		b.Author,
		b.YearPublished,
	)
}

var Books = []Book{
	{
		ID:            1,
		Title:         "His Dark Materials",
		Author:        "Phillip Pullman",
		YearPublished: 1999,
	},
	{
		ID:            2,
		Title:         "Attack on Titan Complete Collection",
		Author:        "Hajime Isayama",
		YearPublished: 2021,
	},
	{
		ID:            3,
		Title:         "Epher",
		Author:        "Elliot Parker",
		YearPublished: 2041,
	},
	{
		ID:            4,
		Title:         "Harry Potter Complete Collection",
		Author:        "JK Rowling",
		YearPublished: 2029,
	},
	{
		ID:            5,
		Title:         "The Bible",
		Author:        "God (Apparently)",
		YearPublished: 0,
	},
	{
		ID:            6,
		Title:         "OCR A Level Computing",
		Author:        "Chris Leadbetter",
		YearPublished: 2010,
	},
	{
		ID:            7,
		Title:         "Harmonica for Physicists Who Are Actually Chemists",
		Author:        "Jon Burr",
		YearPublished: 2000,
	},
	{
		ID:            8,
		Title:         "Something That Happened GDD",
		Author:        "John Pokey, The Pilgrim, Nasty Bowe Face",
		YearPublished: 2011,
	},
	{
		ID:            9,
		Title:         "Inside Out",
		Author:        "Nick Mason",
		YearPublished: 2004,
	},
	{
		ID:            10,
		Title:         "All of Wikipedia",
		Author:        "The World",
		YearPublished: 9999,
	},
}
