package main

import "fmt"

type bookmarkMap = map[string]string

func main() {
	bookmarks := bookmarkMap{}
	fmt.Println("__Приложение для закладок__")
Menu:
	for {
		variant := menuScan()
		switch variant {
		case 1:
			printBookmarks(bookmarks)
		case 2:
			bookmarks = addBooksmark(bookmarks)
		case 3:
			bookmarks = deleteBookmark(bookmarks)
		case 4:
			break Menu
		}
	}
}

func menuScan() int {
	var variant int
	fmt.Println("Выберите пункт меню: ")
	fmt.Println("1. Посмотреть закладку")
	fmt.Println("2. Добавить закладку ")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
	fmt.Scan(&variant)
	return variant
}

func checkInput(bookmarks bookmarkMap) bool {
	var variant int
	fmt.Scan(&variant)
	switch variant {
	case 1:
		printBookmarks(bookmarks)
	case 2:
		bookmarks = addBooksmark(bookmarks)
	case 3:
		bookmarks = deleteBookmark(bookmarks)
	case 4:
		return false
	}
	return true
}

func printBookmarks(bookmarks bookmarkMap) {
	if len(bookmarks) == 0 {
		fmt.Println("Пока нет закладок")
	}
	for key, value := range bookmarks {
		fmt.Println(key, ": ", value)
	}
}

func addBooksmark(bookmarks bookmarkMap) bookmarkMap {
	var newBookmarkKey string
	var newBookmarkValue string
	fmt.Print("Введите название: ")
	fmt.Scan(&newBookmarkKey)
	fmt.Print("Введите ссылку: ")
	fmt.Scan(&newBookmarkValue)
	bookmarks[newBookmarkKey] = newBookmarkValue
	return bookmarks
}

func deleteBookmark(bookmarks bookmarkMap) bookmarkMap {
	var BookmarkKeyToDelete string
	fmt.Print("Введите название: ")
	fmt.Scan(&BookmarkKeyToDelete)
	delete(bookmarks, BookmarkKeyToDelete)
	return bookmarks
}
