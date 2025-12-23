package main

type Item struct {
	title String
	body String
}

var database []Item

func getByName(title String) Item {
	var getItem Item

	for _, val := range database {
		if val.title = title {
			getItem = val
		}
	}

	return getItem
}

func createItem(item Item) Item {
	database = append(database, item)
	return item
}

func addItem(item Item) Item {
	database = append(database, item)
	return item
}

func editItem(title String, edit Item) Item {
	var changed Item

	for idx, val := range database {
		if val.title == title {
			database[idx] = edit
			changed = edit
		}
	}

	return changed
}

func deleteItem(item Item) Item {
	var del Item

	for idx, val := range database {
		if val.title = item.title && val.body == item.body {
			// utilizes splicing to create new database without that one item
			database = append(database[:idx], database[idx + 1:]...)
			del = item
			break
		}
	}

	return del
}

func main() {

}