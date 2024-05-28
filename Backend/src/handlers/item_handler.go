package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wichijan/InventoryPro/src/controllers"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/utils"
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
// @Router /items/:id [get]
func GetItemByIdHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
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

// @Summary Create Item
// @Description Create Item
// @Tags Items
// @Accept  json
// @Produce  json
// @Param room body model.Rooms true "ItemWithStatus model"
// @Success 201 {object} models.ItemWithStatus
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /items [post]
func CreateItemHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item models.ItemWithStatus
		err := c.ShouldBindJSON(&item)
		if err != nil || item.Name == "" {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		itemId, inv_err := itemCtrl.CreateItem(&item)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusCreated, itemId)
	}
}

// @Summary Update Item
// @Description Update Item
// @Tags Items
// @Accept  json
// @Produce  json
// @Param item body models.ItemWithStatus true "ItemWithStatus model"
// @Success 201 {object} models.ItemWithStatus
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /items [put]
func UpdateItemHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item models.ItemWithStatus
		err := c.ShouldBindJSON(&item)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		inv_err := itemCtrl.UpdateItem(&item)
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
// @Router /items/addkeyword [post]
func AddKeywordToItemHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var itemAndKeyword models.ItemWithKeywordName
		err := c.ShouldBindJSON(&itemAndKeyword)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
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
// @Router /items/removekeyword [post]
func RemoveKeywordFromItemHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var itemAndKeyword models.ItemWithKeywordName
		err := c.ShouldBindJSON(&itemAndKeyword)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
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
// @Router /items/addsubject [post]
func AddSubjectToItemHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var itemAndKeyword models.ItemWithSubjectName
		err := c.ShouldBindJSON(&itemAndKeyword)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
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
// @Router /items/removesubject [post]
func RemoveSubjectFromItemHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var itemAndKeyword models.ItemWithSubjectName
		err := c.ShouldBindJSON(&itemAndKeyword)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		inv_err := itemCtrl.RemoveSubjectFromItem(itemAndKeyword)
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
// @Param ItemReserveODT body models.ItemReserveODT true "ItemReserveODT model"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Router /items/reserve [post]
func ReserveItemHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {

		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
			return
		}

		var itemReserveODT models.ItemReserveODT
		err := c.ShouldBindJSON(&itemReserveODT)
		if err != nil || utils.ContainsEmptyString(itemReserveODT.ItemID) {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		itemReserve := models.ItemReserve{
			ItemID:   itemReserveODT.ItemID,
			UserID:   userId.String(),
			Quantity: itemReserveODT.Quantity,
		}

		inv_err := itemCtrl.ReserveItem(itemReserve)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Cancel Reserve Item
// @Description Cancel Reserve Item when logged-in
// @Tags Items
// @Accept  json
// @Produce  json
// @Param id path string true "item id"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Router /items/reserve-cancel/:id [post]
func CancelReserveItemHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {

		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
			return
		}

		itemId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		inv_err := itemCtrl.CancelReserveItem(userId, &itemId)
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
		if err != nil || utils.ContainsEmptyString(itemReserveODT.ItemID) {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		itemReserve := models.ItemBorrow{
			ItemID:   itemReserveODT.ItemID,
			UserID:   userId.String(),
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
// @Router /items/return/:id [post]
func ReturnReserveItemHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {

		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
			return
		}

		itemId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
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
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read the form data"})
			return
		}
		file := form.File["file"][0]
		itemId, err := uuid.Parse(form.Value["id"][0])
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		log.Print("Uploading image for item: ", itemId.String())

		imageId, inv_err := itemCtrl.UploadImage(&itemId)
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
// @Success 200 {object} models.ItemPicturePath
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /items-picture/:id [get]
func GetImagePathForItemHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		// single file
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}
		log.Print("ID: ", id.String())

		imageId, inv_err := itemCtrl.GetImageIdFromItem(&id)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		imageName := "./../uploads/" + imageId.String() + ".jpeg"
		log.Print("Reading image: ", imageName)

		c.JSON(http.StatusOK, models.ItemPicturePath{Path: imageName})
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
// @Router /items-picture/:id [delete]
func RemoveImageForItemHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		// single file
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
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
