package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wichijan/InventoryPro/src/controllers"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/utils"
	"github.com/wichijan/InventoryPro/src/websocket"
)

// @Summary Get items
// @Description Get items
// @Tags Items
// @Accept  json
// @Produce  json
// @Success 200 {array} models.ItemWithEverything
// @Failure 500 {object} models.INVErrorMessage
// @Router /items [get]
func GetItemsHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, inv_err := itemCtrl.GetItems()
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

// @Summary Get item by id
// @Description Get item by id
// @Tags Items
// @Accept  json
// @Produce  json
// @Param id path string true "item id"
// @Success 200 {object} models.ItemWithEverything
// @Failure 500 {object} models.INVErrorMessage
// @Router /items/{id} [get]
func GetItemByIdHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid id"))
			return
		}

		item, inv_err := itemCtrl.GetItemById(&id)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, item)
	}
}

// @Summary Get book by ItemId
// @Description Get book by ItemId
// @Tags Items
// @Accept  json
// @Produce  json
// @Param id path string true "item id"
// @Success 200 {object} model.Books
// @Failure 500 {object} models.INVErrorMessage
// @Router /items/book/{id} [get]
func GetBookByIdHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid id"))
			return
		}

		book, inv_err := itemCtrl.GetBookById(&id)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, book)
	}
}

// @Summary Get SingleObject by ItemId
// @Description Get SingleObject by ItemId
// @Tags Items
// @Accept  json
// @Produce  json
// @Param id path string true "item id"
// @Success 200 {object} model.SingleObject
// @Failure 500 {object} models.INVErrorMessage
// @Router /items/single-object/{id} [get]
func GetSingleObjectByIdHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid id"))
			return
		}

		singleObject, inv_err := itemCtrl.GetSingleObjectById(&id)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, singleObject)
	}
}

// @Summary Get SetsOfObjects by ItemId
// @Description Get SetsOfObjects by ItemId
// @Tags Items
// @Accept  json
// @Produce  json
// @Param id path string true "item id"
// @Success 200 {object} model.SetsOfObjects
// @Failure 500 {object} models.INVErrorMessage
// @Router /items/set-of-objects/{id} [get]
func GetSetOfObjectsByIdHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid id"))
			return
		}

		setOfObjects, inv_err := itemCtrl.GetSetOfObjectsById(&id)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, setOfObjects)
	}
}

// @Summary Create Item with book
// @Description Create Item with book
// @Tags Items
// @Accept  json
// @Produce  json
// @Param item body models.ItemCreateWithBook true "ItemCreateWithBook model"
// @Success 201 {object} uuid.UUID
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /items/book [post]
func CreateItemWithBookHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item models.ItemCreateWithBook
		err := c.ShouldBindJSON(&item)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid item object"))
			return
		}
		if item.BaseQuantityInShelf < 0 {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid quantity"))
			return
		}
		if item.Name == "" {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid name"))
			return
		}

		itemId, inv_err := itemCtrl.CreateItemWithBook(&item)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusCreated, itemId)
	}
}

// @Summary Create Item with single object
// @Description Create Item with single object
// @Tags Items
// @Accept  json
// @Produce  json
// @Param item body models.ItemCreateWithSingleObject true "ItemCreateWithSingleObject model"
// @Success 201 {object} uuid.UUID
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /items/single-object [post]
func CreateItemWithSingleObjectHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item models.ItemCreateWithSingleObject
		err := c.ShouldBindJSON(&item)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid item object"))
			return
		}
		if item.BaseQuantityInShelf < 0 {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid quantity"))
			return
		}
		if item.Name == "" {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid name"))
			return
		}

		itemId, inv_err := itemCtrl.CreateItemWithSingleObject(&item)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusCreated, itemId)
	}
}

// @Summary Create Item with sets of objects
// @Description Create Item with sets of objects
// @Tags Items
// @Accept  json
// @Produce  json
// @Param item body models.ItemCreateWithSetOfObject true "ItemCreateWithSetOfObject model"
// @Success 201 {object} uuid.UUID
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /items/set-of-objects [post]
func CreateItemWithSetOfObjectsHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item models.ItemCreateWithSetOfObject
		err := c.ShouldBindJSON(&item)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid item object"))
			return
		}
		if item.BaseQuantityInShelf < 0 {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid quantity"))
			return
		}
		if item.Name == "" {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid name"))
			return
		}

		itemId, inv_err := itemCtrl.CreateItemWithSetOfObject(&item)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusCreated, itemId)
	}
}

// @Summary Update Item with Book
// @Description Update Item with Book
// @Tags Items
// @Accept  json
// @Produce  json
// @Param item body models.ItemUpdateWithBook true "ItemUpdateWithBook model"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /items/book [put]
func UpdateItemWithBookHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item models.ItemUpdateWithBook
		err := c.ShouldBindJSON(&item)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid item object"))
			return
		}

		if strings.ToLower(item.ItemTypes.String()) != "book" {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Not a book"))
			return
		}

		inv_err := itemCtrl.UpdateItemWithBook(&item)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Update Item with Single Object
// @Description Update Item with Single Object
// @Tags Items
// @Accept  json
// @Produce  json
// @Param item body models.ItemUpdateWithSingleObject true "ItemUpdateWithSingleObject model"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /items/single-object [put]
func UpdateItemWithSingleObjectHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item models.ItemUpdateWithSingleObject
		err := c.ShouldBindJSON(&item)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid item object"))
			return
		}

		if strings.ToLower(item.ItemTypes.String()) != "single_object" {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Not a single_object"))
			return
		}

		inv_err := itemCtrl.UpdateItemWithSingleObject(&item)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Update Item
// @Description Update Item
// @Tags Items
// @Accept  json
// @Produce  json
// @Param item body models.ItemUpdateWithSetsOfObjects true "ItemUpdateWithSetsOfObjects model"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /items/set-of-objects [put]
func UpdateItemWithSetOfObjectsHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item models.ItemUpdateWithSetsOfObjects
		err := c.ShouldBindJSON(&item)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid item object"))
			return
		}

		if strings.ToLower(item.ItemTypes.String()) != "set_of_objects" {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Not a set_of_objects"))
			return
		}

		inv_err := itemCtrl.UpdateItemWithSetOfObject(&item)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Delete Item
// @Description Delete Item
// @Tags Items
// @Accept  json
// @Produce  json
// @Param id path string true "Item Id"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /items/{id} [delete]
func DeleteItemHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		itemId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid room id"))
			return
		}

		inv_err := itemCtrl.DeleteItem(&itemId)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Add keyword to item
// @Description Add keyword to item
// @Tags Items
// @Accept  json
// @Produce  json
// @Param item_keyword body models.ItemWithKeywordName true "ItemWithKeywordName model"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /items/add-keyword [post]
func AddKeywordToItemHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var itemAndKeyword models.ItemWithKeywordName
		err := c.ShouldBindJSON(&itemAndKeyword)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid keyword object"))
			return
		}

		inv_err := itemCtrl.AddKeywordToItem(itemAndKeyword)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Remove keyword to item
// @Description Remove keyword to item
// @Tags Items
// @Accept  json
// @Produce  json
// @Param item_keyword body models.ItemWithKeywordName true "ItemWithKeywordName model"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /items/remove-keyword [delete]
func RemoveKeywordFromItemHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var itemAndKeyword models.ItemWithKeywordName
		err := c.ShouldBindJSON(&itemAndKeyword)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid keyword object"))
			return
		}

		inv_err := itemCtrl.RemoveKeywordFromItem(itemAndKeyword)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Add Subject to item
// @Description Add Subject to item
// @Tags Items
// @Accept  json
// @Produce  json
// @Param item_keyword body models.ItemWithSubjectName true "ItemWithSubjectName model"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /items/add-subject [post]
func AddSubjectToItemHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var itemAndKeyword models.ItemWithSubjectName
		err := c.ShouldBindJSON(&itemAndKeyword)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid subject object"))
			return
		}

		inv_err := itemCtrl.AddSubjectToItem(itemAndKeyword)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Remove Subject to item
// @Description Remove Subject to item
// @Tags Items
// @Accept  json
// @Produce  json
// @Param item_subject body models.ItemWithSubjectName true "ItemWithSubjectName model"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /items/remove-subject [delete]
func RemoveSubjectFromItemHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var itemAndSubject models.ItemWithSubjectName
		err := c.ShouldBindJSON(&itemAndSubject)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid subject object"))
			return
		}

		inv_err := itemCtrl.RemoveSubjectFromItem(itemAndSubject)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Reserve Item
// @Description Reserve Item when logged-in
// @Tags Items
// @Accept  json
// @Produce  json
// @Param ReservationCreateODT body models.ReservationCreateODT true "ReservationCreateODT model"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Router /items/reserve [post]
func ReserveItemHandler(reservationCtrl controllers.ReservationControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {

		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
			return
		}

		var itemReserveODT models.ReservationCreateODT
		err := c.ShouldBindJSON(&itemReserveODT)
		if err != nil || utils.ContainsEmptyString(itemReserveODT.ItemID) || itemReserveODT.Quantity <= 0 {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid reservation object"))
			return
		}

		newTimeFrom, inv_error := time.Parse("2006-01-02", itemReserveODT.TimeFrom)
		if inv_error != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid time from"))
			return
		}

		newTimeTo, inv_error := time.Parse("2006-01-02", itemReserveODT.TimeTo)
		if inv_error != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid time to"))
			return
		}

		itemReserve := models.ReservationCreate{
			ItemID:   itemReserveODT.ItemID,
			UserID:   userId.String(),
			Quantity: itemReserveODT.Quantity,
			TimeFrom: newTimeFrom,
			TimeTo:   newTimeTo,
		}

		reservationId, inv_err := reservationCtrl.CreateReservation(&itemReserve)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, reservationId)
	}
}

// @Summary Cancel Reserve Item
// @Description Cancel Reserve Item when logged-in
// @Tags Items
// @Accept  json
// @Produce  json
// @Param id path string true "reservation id"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Router /items/reserve-cancel/{id} [post]
func CancelReserveItemHandler(reservationCtrl controllers.ReservationControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {

		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
			return
		}

		reservationId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid reservation id"))
			return
		}

		inv_err := reservationCtrl.DeleteReservation(userId, &reservationId)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Borrow Item
// @Description Borrow Item when logged-in
// @Tags Items
// @Accept  json
// @Produce  json
// @Param ItemReserveODT body models.ItemReserveODT true "ItemReserveODT model"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Router /items/borrow [post]
func BorrowItemHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {

		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
			return
		}

		var itemReserveODT models.ItemReserveODT
		err := c.ShouldBindJSON(&itemReserveODT)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid reservation object"))
			return
		}

		itemReserve := models.ItemBorrowCreate{
			ItemID:   itemReserveODT.ItemID,
			UserID:   *userId,
			Quantity: itemReserveODT.Quantity,
		}

		inv_err := itemCtrl.BorrowItem(itemReserve)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Return Reserve Item
// @Description Return Reserve Item when logged-in
// @Tags Items
// @Accept  json
// @Produce  json
// @Param id path string true "item id"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Router /items/return/{id} [delete]
func ReturnReserveItemHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {

		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
			return
		}

		itemId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid item id"))
			return
		}

		inv_err := itemCtrl.ReturnItem(userId, &itemId)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Upload Img for Item
// @Description Upload Img for Item. Form with enctype="multipart/form-data" <input type="file" name="file" /> & <input type="hidden" name="id" /> for item id
// @Tags Items
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /items-picture [post]
func UploadImageForItemHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		// single file
		form, err := c.MultipartForm()
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Unable to read the form data"))
			return
		}
		file := form.File["file"][0]
		itemId, err := uuid.Parse(form.Value["id"][0])
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid item id"))
			return
		}

		log.Print("Uploading image for item: ", itemId.String())

		imageId, inv_err := itemCtrl.UploadItemImage(&itemId)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		imageName := "./../uploads/" + imageId.String() + ".jpeg"
		c.SaveUploadedFile(file, imageName)

		c.JSON(http.StatusOK, imageName)
	}
}

// @Summary Get ImagePath For Item
// @Description Get ImagePath For Item
// @Tags Items
// @Accept  json
// @Produce  json
// @Param id path string true "item id"
// @Success 200 {object} models.PicturePath
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /items-picture/{id} [get]
func GetImagePathForItemHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// single file
		id, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(ctx, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid id"))
			return
		}
		log.Print("ID: ", id.String())

		imageId, inv_err := itemCtrl.GetImageIdFromItem(&id)
		if inv_err != nil {
			utils.HandleErrorAndAbort(ctx, inv_err)
			return
		}
		imageName := "./../uploads/" + imageId.String() + ".jpeg"
		log.Print("Reading image: ", imageName)

		// New ---------------------------------------------------------------

		// Open the file
		fileData, err := os.Open(imageName)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
			return
		}
		defer fileData.Close()

		// Read the first 512 bytes of the file to determine its content type
		fileHeader := make([]byte, 512)
		_, err = fileData.Read(fileHeader)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image"})
			return
		}
		fileContentType := http.DetectContentType(fileHeader)

		// Get the file info
		fileInfo, err := fileData.Stat()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get image info"})
			return
		}

		ctx.Header("Content-Description", "File Transfer")
		ctx.Header("Content-Transfer-Encoding", "binary")
		ctx.Header("Content-Disposition", fmt.Sprintf("inline; filename=%s", imageId.String()))
		ctx.Header("Content-Type", fileContentType)
		ctx.Header("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))
		ctx.File(imageName)
	}
}

// @Summary Delete Img for Item
// @Description Delete Picture from item and replace with ""
// @Tags Items
// @Accept  json
// @Produce  json
// @Param id path string true "item id"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /items-picture/{id} [delete]
func RemoveImageForItemHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		// single file
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid id"))
			return
		}
		log.Print("ID: ", id.String())

		inv_err := itemCtrl.RemoveImageIdFromItem(&id)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Create Transfer Request
// @Description Create Transfer Request to move item from User A to User B
// @Tags Items-Transfer
// @Accept  json
// @Produce  json
// @Param item body models.ItemMoveRequest true "ItemMoveRequest model"
// @Success 200 {object} models.TransferRequestResponse
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /items/transfer-request [post]
func MoveItemRequestHandler(itemCtrl controllers.ItemControllerI, hub *websocket.Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
			return
		}

		var item models.ItemMoveRequest
		err := c.ShouldBindJSON(&item)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid item object"))
			return
		}

		itemMove := models.ItemMove{
			ItemID:    &item.ItemID,
			UserID:    userId,
			NewUserID: &item.NewUserID,
		}

		transferRequestId, inv_err := itemCtrl.MoveItemRequest(itemMove)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		// Inform user that item move request created
		hub.HandleMessage(websocket.Message{
			Type:         utils.MESSAGE_TYPE_TO_USER,
			SentToUserId: item.NewUserID.String(),
			Sender:       userId.String(),
			Content:      "Item move request created",
		})

		c.JSON(http.StatusOK, models.TransferRequestResponse{
			TransferRequestID: transferRequestId.String(),
		})
	}
}

// @Summary Accept Transfer Request
// @Description Accept Transfer Request
// @Tags Items-Transfer
// @Accept  json
// @Produce  json
// @Param id path string true "transfer request id"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /items/transfer-accept/{id} [post]
func MoveItemAcceptedHandler(itemCtrl controllers.ItemControllerI, hub *websocket.Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
			return
		}

		transferRequestId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid id"))
			return
		}

		transferAccept := models.TransferAccept{
			TransferRequestID: &transferRequestId,
			UserId:            userId,
		}

		transferRequest, inv_err := itemCtrl.MoveItemAccepted(transferAccept)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		// Inform user that item has been moved
		hub.HandleMessage(websocket.Message{
			Type:         utils.MESSAGE_TYPE_TO_USER,
			SentToUserId: transferRequest.UserID.String(),
			Sender:       userId.String(),
			Content:      "Item move request accepted",
		})

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Get Transfer Requests by UserId
// @Description Get Transfer Requests by UserId
// @Tags Items-Transfer
// @Accept  json
// @Produce  json
// @Success 200 {array} models.TransferRequestSelect
// @Failure 500 {object} models.INVErrorMessage
// @Router /items/transfer-requests [get]
func GetTransferRequestByIdHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
			return
		}

		transferRequest, inv_err := itemCtrl.GetTransferRequestByUserId(*userId)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, transferRequest)
	}
}
