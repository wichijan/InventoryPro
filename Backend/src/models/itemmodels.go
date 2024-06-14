package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
)

type ItemWThin struct {
	model.Items

	QuantityInShelf    int32  `alias:"items_in_shelf.quantity"`
	BorrowedByUserID   string `alias:"users.id"`
	BorrowedByUserName string `alias:"users.username"`
}

type ItemWithEverything struct {
	model.Items

	QuantityInShelf    int32  `alias:"items_in_shelf.quantity"`
	BorrowedByUserID   string `alias:"users.id"`
	BorrowedByUserName string `alias:"users.username"`

	Keywords []struct {
		model.KeywordsForItems
	}

	Subject []struct {
		model.ItemSubjects
	}

	Reservations []struct {
		model.Reservations
	}
}

type ItemCreate struct {
	Name                string               `json:"name" binding:"required"`
	Description         string               `json:"description"`
	ItemTypeName        model.ItemsItemTypes `json:"itemTypeName" binding:"required"`
	RegularShelfId      uuid.UUID            `json:"regularShelfId" binding:"required"`
	ClassOne            bool                 `json:"classOne"`
	ClassTwo            bool                 `json:"classTwo"`
	ClassThree          bool                 `json:"classThree"`
	ClassFour           bool                 `json:"classFour"`
	Damaged             bool                 `json:"damaged"`
	DamagedDesc         string               `json:"damagedDesc"`
	HintText            string               `json:"hintText"`
	BaseQuantityInShelf int32                `json:"BaseQuantityInShelf" binding:"required"`
}

type ItemUpdate struct {
	model.Items
	QuantityInShelf int32 `json:"QuantityInShelf"`
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
	SubjectName string `json:"subjectName"`
}

type ItemWithUser struct {
	Item_id          string `alias:"items.id"`
	Item_Name        string `alias:"items.name"`
	Item_Description string `alias:"items.description"`
	User_id          string `alias:"users.id"`
	User_Username    string `alias:"users.username"`
}

type ItemReserveODT struct {
	ItemID   uuid.UUID `binding:"required"`
	Quantity int32     `binding:"required"`
}

type ItemBorrow struct {
	ItemID          string
	UserID          string
	Quantity        int32
	TransactionDate time.Time
}

type ItemBorrowCreate struct {
	ItemID   uuid.UUID
	UserID   uuid.UUID
	Quantity int32
}

type ItemMove struct {
	ItemID    *uuid.UUID
	UserID    *uuid.UUID
	NewUserID *uuid.UUID
}

type ItemMoveRequest struct {
	ItemID    uuid.UUID
	NewUserID uuid.UUID
}

type ItemPicture struct {
	PictureId string `alias:"items.picture"`
}

type PicturePath struct {
	Path string
}

type ItemTypes struct {
	TypeName string `alias:"item_types.type_name"`
}

type TransferRequestResponse struct {
	TransferRequestID string
}

type TransferAccept struct {
	TransferRequestID *uuid.UUID
	UserId            *uuid.UUID
}

type TransferRequestSelect struct {
	TransferRequestID *uuid.UUID `alias:"transfer_requests.transfer_request_id" sql:"primary_key"`
	ItemID            *uuid.UUID `alias:"transfer_requests.item_id"`
	UserID            *uuid.UUID `alias:"transfer_requests.user_id"`
	TargetUserID      *uuid.UUID `alias:"transfer_requests.target_user_id"`
	RequestDate       *time.Time `alias:"transfer_requests.request_date"`
	IsAccepted        *bool      `alias:"transfer_requests.is_accepted"`
}
