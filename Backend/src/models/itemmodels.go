package models

import "github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"

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
	Quantity           int32  `alias:"items.quantity"`
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
	Quantity           int32  `alias:"items.quantity"`
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
	item_id          string `alias:"items.id"`
	item_Name        string `alias:"items.name"`
	item_Description string `alias:"items.description"`
	user_id          string `alias:"users.id"`
	user_Username    string `alias:"users.username"`
}
