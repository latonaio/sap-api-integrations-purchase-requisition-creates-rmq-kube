package requests

type Header struct {
            PurchaseRequisition     string `json:"PurchaseRequisition"`
			PurchaseRequisitionType *string `json:"PurchaseRequisitionType"`
			SourceDetermination     *bool   `json:"SourceDetermination"`
}
