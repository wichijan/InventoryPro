package models

type TransferRequestCreate struct {
	UserID string `alias:"transfer_requests.user_id"`
	ItemID string `alias:"transfer_requests.item_id"`
	TargetUserID string `alias:"transfer_requests.target_user_id"`
}

type TransferRequestUpdate struct {
	TransferRequestID string `alias:"transfer_requests.transfer_request_id"`
	UserID string `alias:"transfer_requests.user_id"`
	ItemID string `alias:"transfer_requests.item_id"`
	TargetUserID string `alias:"transfer_requests.target_user_id"`
	IsAccepted bool `alias:"transfer_requests.is_accepted"`
}
