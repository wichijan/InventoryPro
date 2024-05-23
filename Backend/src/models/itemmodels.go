package models

import (
	"time"

	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
)

type ItemWithStatus struct {
	ID                 string `alias:"items.id" sql:"primary_key"`
	Name               string `alias:"items.name"`
	Description        string `alias:"items.description"`
	ClassOne           bool   `alias:"items.class_one"`
	ClassTwo           bool   `alias:"items.class_two"`
	ClassThree         bool   `alias:"items.class_three"`
	ClassFour          bool   `alias:"items.class_four"`
	Damaged            bool   `alias:"items.damaged"`
	DamagedDesc        string `alias:"items.damaged_description"`
	Picture            string `alias:"items.picture"`
	Status             string `alias:"item_status.status_name"`
	BorrowedByUserID   string `alias:"users.id"`
	BorrowedByUserName string `alias:"users.username"`
}

type ItemWithEverything struct {
	ID                 string `alias:"items.id" sql:"primary_key"`
	Name               string `alias:"items.name"`
	Description        string `alias:"items.description"`
	ClassOne           bool   `alias:"items.class_one"`
	ClassTwo           bool   `alias:"items.class_two"`
	ClassThree         bool   `alias:"items.class_three"`
	ClassFour          bool   `alias:"items.class_four"`
	Damaged            bool   `alias:"items.damaged"`
	DamagedDesc        string `alias:"items.damaged_description"`
	QuantityInShelve   int32  `alias:"items_in_shelve.quantity"`
	Picture            string `alias:"items.picture"`
	Status             string `alias:"item_status.status_name"`
	BorrowedByUserID   string `alias:"users.id"`
	BorrowedByUserName string `alias:"users.username"`

	Keywords []struct {
		model.KeywordsForItems
	}

	Subject []struct {
		model.ItemSubjects
	}
}

type ItemWithKeyword struct {
	ItemID    string `json:"itemId"`
	KeywordID string `json:"keywordId"`
}

type ItemWithKeywordName struct {
	ItemID      string `json:"itemId"`
	KeywordName string `json:"keywordName"`
}

type ItemWithSubject struct {
	ItemID    string `json:"itemId"`
	SubjectID string `json:"keywordId"`
}

type ItemWithSubjectName struct {
	ItemID      string `json:"itemId"`
	SubjectName string `json:"keywordName"`
}

type ItemWithUser struct {
	Item_id          string `alias:"items.id"`
	Item_Name        string `alias:"items.name"`
	Item_Description string `alias:"items.description"`
	User_id          string `alias:"users.id"`
	User_Username    string `alias:"users.username"`
}

type ItemReserveODT struct {
	ItemID   string `json:"itemId"`
	Quantity int32  `json:"quantity"`
}

type ItemReserve struct {
	ItemID      string
	UserID      string
	Quantity    int32
	StatusID    string
	ReserveDate time.Time
}
