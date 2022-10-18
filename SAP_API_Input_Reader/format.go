package sap_api_input_reader

import (
	"sap-api-integrations-purchase-requisition-creates-rmq-kube/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeaderItem() *requests.HeaderItem {
	data := sdc.PurchaseRequisition
	results := make([]requests.Item, 0, len(data.PurchaseRequisitionItem))

	header := sdc.ConvertToHeader()

	for i := range data.PurchaseRequisitionItem {
		results = append(results, *sdc.ConvertToItem(i))
	}

	return &requests.HeaderItem{
		Header: *header,
		ToItem: requests.ToItem{
			Results: results,
		},
	}
}

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.PurchaseRequisition
	return &requests.Header{
		PurchaseRequisition:     data.PurchaseRequisition,
		PurchaseRequisitionType: data.PurchaseRequisitionType,
		SourceDetermination:     data.SourceDetermination,
	}
}

func (sdc *SDC) ConvertToItem(num int) *requests.Item {
	dataPurchaseRequisition := sdc.PurchaseRequisition
	data := sdc.PurchaseRequisition.PurchaseRequisitionItem[num]
	return &requests.Item{
		PurchaseRequisition:            dataPurchaseRequisition.PurchaseRequisition,
		PurchaseRequisitionItem:        data.PurchaseRequisitionItem,
		PurchasingDocument:             data.PurchasingDocument,
		PurchasingDocumentItem:         data.PurchasingDocumentItem,
		PurReqnReleaseStatus:           data.PurReqnReleaseStatus,
		PurchasingDocumentItemCategory: data.PurchasingDocumentItemCategory,
		PurchaseRequisitionItemText:    data.PurchaseRequisitionItemText,
		AccountAssignmentCategory:      data.AccountAssignmentCategory,
		Material:                       data.Material,
		MaterialGroup:                  data.MaterialGroup,
		PurchasingDocumentCategory:     data.PurchasingDocumentCategory,
		RequestedQuantity:              data.RequestedQuantity,
		BaseUnit:                       data.BaseUnit,
		// PurchaseRequisitionPrice:       data.PurchaseRequisitionPrice,
		// PurReqnPriceQuantity:           data.PurReqnPriceQuantity,
		// MaterialGoodsReceiptDuration:   data.MaterialGoodsReceiptDuration,
		ReleaseCode: data.ReleaseCode,
		// PurchaseRequisitionReleaseDate: data.PurchaseRequisitionReleaseDate,
		PurchasingOrganization:   data.PurchasingOrganization,
		PurchasingGroup:          data.PurchasingGroup,
		Plant:                    data.Plant,
		CompanyCode:              data.CompanyCode,
		SourceOfSupplyIsAssigned: data.SourceOfSupplyIsAssigned,
		SupplyingPlant:           data.SupplyingPlant,
		// OrderedQuantity:             data.OrderedQuantity,
		DeliveryDate: data.DeliveryDate,
		// CreationDate:                data.CreationDate,
		ProcessingStatus:       data.ProcessingStatus,
		ExternalApprovalStatus: data.ExternalApprovalStatus,
		PurchasingInfoRecord:   data.PurchasingInfoRecord,
		Supplier:               data.Supplier,
		FixedSupplier:          data.FixedSupplier,
		RequisitionerName:      data.RequisitionerName,
		PurReqnItemCurrency:    data.PurReqnItemCurrency,
		// MaterialPlannedDeliveryDurn: data.MaterialPlannedDeliveryDurn,
		StorageLocation:           data.StorageLocation,
		PurReqnSourceOfSupplyType: data.PurReqnSourceOfSupplyType,
		ConsumptionPosting:        data.ConsumptionPosting,
		PurReqnOrigin:             data.PurReqnOrigin,
		IsPurReqnBlocked:          data.IsPurReqnBlocked,
		PurchaseRequisitionStatus: data.PurchaseRequisitionStatus,
		Batch:                     data.Batch,
		GoodsReceiptIsExpected:    data.GoodsReceiptIsExpected,
		GoodsReceiptIsNonValuated: data.GoodsReceiptIsNonValuated,
		RequirementTracking:       data.RequirementTracking,
		MRPController:             data.MRPController,
		// Reservation:               data.Reservation,
		// LastChangeDateTime: data.LastChangeDateTime,
		IsDeleted: data.IsDeleted,
	}
}
