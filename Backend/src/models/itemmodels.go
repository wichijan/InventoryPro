package models

import "github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"

type ItemWithStatus struct {
	ID          string `alias:"items.id"`
	Name        string `alias:"items.name"`
	Description string `alias:"items.description"`
	ClassOne    string `alias:"items.class_one"`
	ClassTwo    string `alias:"items.class_two"`
	ClassThree  string `alias:"items.class_three"`
	ClassFour   string `alias:"items.class_four"`
	Damaged     bool   `alias:"items.damaged"`
	DamagedDesc string `alias:"items.damaged_description"`
	Quantity    int    `alias:"items.quantity"`
	Status      string `alias:"item_status.status_name"`
}

type ItemWithEverything struct {
	ID          string `alias:"items.id"`
	Name        string `alias:"items.name"`
	Description string `alias:"items.description"`
	ClassOne    string `alias:"items.class_one"`
	ClassTwo    string `alias:"items.class_two"`
	ClassThree  string `alias:"items.class_three"`
	ClassFour   string `alias:"items.class_four"`
	Damaged     bool   `alias:"items.damaged"`
	DamagedDesc string `alias:"items.damaged_description"`
	Quantity    int    `alias:"items.quantity"`
	Status      string `alias:"item_status.status_name"`

	Keywords []struct {
		model.KeywordsForItems
	}

	Subject []struct {
		model.ItemSubjects
	}

	Pictures []struct {
		model.ItemPictures
	}
}
