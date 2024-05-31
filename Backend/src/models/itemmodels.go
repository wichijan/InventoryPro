package models

import (
	"time"

	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
)

type ItemWThin struct {
	ID                 string `alias:"items.id" sql:"primary_key"`
	Name               string `alias:"items.name"`
	Description        string `alias:"items.description"`
	ItemTypeName 	   string `alias:"item_types.type_name"`
	RegularShelfId  string `alias:"items_in_shelf.shelf_id"`
	QuantityInShelf   int32  `alias:"items_in_shelf.quantity"`
	ClassOne           bool   `alias:"items.class_one"`
	ClassTwo           bool   `alias:"items.class_two"`
	ClassThree         bool   `alias:"items.class_three"`
	ClassFour          bool   `alias:"items.class_four"`
	Damaged            bool   `alias:"items.damaged"`
	DamagedDesc        string `alias:"items.damaged_description"`
	Picture            string `alias:"items.picture"`
	HintText 		 string `alias:"items.hint_text"`
	
	BorrowedByUserID   string `alias:"users.id"`
	BorrowedByUserName string `alias:"users.username"`
}

type ItemWithEverything struct {
	ID                 string `alias:"items.id" sql:"primary_key"`
	Name               string `alias:"items.name"`
	Description        string `alias:"items.description"`
	ItemTypeName 	   string `alias:"item_types.type_name"`
	RegularShelfId  string `alias:"items_in_shelf.shelf_id"`
	QuantityInShelf   int32  `alias:"items_in_shelf.quantity"`
	ClassOne           bool   `alias:"items.class_one"`
	ClassTwo           bool   `alias:"items.class_two"`
	ClassThree         bool   `alias:"items.class_three"`
	ClassFour          bool   `alias:"items.class_four"`
	Damaged            bool   `alias:"items.damaged"`
	DamagedDesc        string `alias:"items.damaged_description"`
	Picture            string `alias:"items.picture"`
	HintText 		 string `alias:"items.hint_text"`

	BorrowedByUserID   string `alias:"users.id"`
	BorrowedByUserName string `alias:"users.username"`

	Keywords []struct {
		model.KeywordsForItems
	}

	Subject []struct {
		model.ItemSubjects
	}
}

type ItemCreate struct {
	Name               string `alias:"items.name"`
	Description        string `alias:"items.description"`
	ItemTypeName 	   string `alias:"item_types.type_name"`
	RegularShelfId  string `alias:"items_in_shelf.shelf_id"`
	QuantityInShelf   int32  `alias:"items_in_shelf.quantity"`
	ClassOne           bool   `alias:"items.class_one"`
	ClassTwo           bool   `alias:"items.class_two"`
	ClassThree         bool   `alias:"items.class_three"`
	ClassFour          bool   `alias:"items.class_four"`
	Damaged            bool   `alias:"items.damaged"`
	DamagedDesc        string `alias:"items.damaged_description"`
	Picture            string `alias:"items.picture"`
	HintText 		 string `alias:"items.hint_text"`
}
type ItemUpdate struct {
	ID                 string `alias:"items.id"`
	Name               string `alias:"items.name"`
	Description        string `alias:"items.description"`
	RegularShelfId  string `alias:"items_in_shelf.shelf_id"`
	QuantityInShelf   int32  `alias:"items_in_shelf.quantity"`
	ClassOne           bool   `alias:"items.class_one"`
	ClassTwo           bool   `alias:"items.class_two"`
	ClassThree         bool   `alias:"items.class_three"`
	ClassFour          bool   `alias:"items.class_four"`
	Damaged            bool   `alias:"items.damaged"`
	DamagedDesc        string `alias:"items.damaged_description"`
	Picture            string `alias:"items.picture"`
	HintText 		 string `alias:"items.hint_text"`
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

type ItemBorrow struct {
	ItemID     string
	UserID     string
	Quantity   int32
	TransactionDate time.Time
}

type ItemPicture struct {
	PictureId string `alias:"items.picture"`
}

type ItemPicturePath struct {
	Path string
}

type ItemTypes struct {
	TypeName string `alias:"item_types.type_name"`
}