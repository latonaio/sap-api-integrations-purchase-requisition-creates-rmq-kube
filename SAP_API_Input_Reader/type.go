package sap_api_input_reader

type EC_MC struct {
	ConnectionKey       string `json:"connection_key"`
	Result              bool   `json:"result"`
	RedisKey            string `json:"redis_key"`
	Filepath            string `json:"filepath"`
	PurchaseRequisition struct {
		PurchaseRequisition      string `json:"document_no"`
		Plant                    string `json:"deliver_to"`
		RequestedQuantity        string `json:"quantity"`
		OrderedQuantity          string `json:"picked_quantity"`
		PurchaseRequisitionPrice string `json:"price"`
		Batch                    string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema                      string `json:"api_schema"`
	MaterialCode                   string `json:"material_code"`
	Supplier_SupplyingPlant        string `json:"plant/supplier"`
	Stock                          string `json:"stock"`
	PurchaseRequisitionType        string `json:"document_type"`
	PRNumber                       string `json:"document_no"`
	DeliveryDate                   string `json:"planned_date"`
	PurchaseRequisitionReleaseDate string `json:"validated_date"`
	Deleted                        bool   `json:"deleted"`
}


type SDC struct {
	ConnectionKey       string `json:"connection_key"`
	Result              bool   `json:"result"`
	RedisKey            string `json:"redis_key"`
	Filepath            string `json:"filepath"`
	PurchaseRequisition struct {
		PurchaseRequisition     string `json:"PurchaseRequisition"`
		PurchaseRequisitionType *string `json:"PurchaseRequisitionType"`
		SourceDetermination     *bool   `json:"SourceDetermination"`
		PurchaseRequisitionItem []struct {
			PurchaseRequisitionItem        string `json:"PurchaseRequisitionItem"`
			PurchasingDocument             *string `json:"PurchasingDocument"`
			PurchasingDocumentItem         *string `json:"PurchasingDocumentItem"`
			PurReqnReleaseStatus           *string `json:"PurReqnReleaseStatus"`
			PurchasingDocumentItemCategory *string `json:"PurchasingDocumentItemCategory"`
			PurchaseRequisitionItemText    *string `json:"PurchaseRequisitionItemText"`
			AccountAssignmentCategory      *string `json:"AccountAssignmentCategory"`
			Material                       *string `json:"Material"`
			MaterialGroup                  *string `json:"MaterialGroup"`
			PurchasingDocumentCategory     *string `json:"PurchasingDocumentCategory"`
			RequestedQuantity              *string `json:"RequestedQuantity"`
			BaseUnit                       *string `json:"BaseUnit"`
			PurchaseRequisitionPrice       *string `json:"PurchaseRequisitionPrice"`
			PurReqnPriceQuantity           *string `json:"PurReqnPriceQuantity"`
			MaterialGoodsReceiptDuration   *string `json:"MaterialGoodsReceiptDuration"`
			ReleaseCode                    *string `json:"ReleaseCode"`
			PurchaseRequisitionReleaseDate *string `json:"PurchaseRequisitionReleaseDate"`
			PurchasingOrganization         *string `json:"PurchasingOrganization"`
			PurchasingGroup                *string `json:"PurchasingGroup"`
			Plant                          *string `json:"Plant"`
			CompanyCode                    *string `json:"CompanyCode"`
			SourceOfSupplyIsAssigned       *bool   `json:"SourceOfSupplyIsAssigned"`
			SupplyingPlant                 *string `json:"SupplyingPlant"`
			OrderedQuantity                *string `json:"OrderedQuantity"`
			DeliveryDate                   *string `json:"DeliveryDate"`
			CreationDate                   *string `json:"CreationDate"`
			ProcessingStatus               *string `json:"ProcessingStatus"`
			ExternalApprovalStatus         *string `json:"ExternalApprovalStatus"`
			PurchasingInfoRecord           *string `json:"PurchasingInfoRecord"`
			Supplier                       *string `json:"Supplier"`
			FixedSupplier                  *string `json:"FixedSupplier"`
			RequisitionerName              *string `json:"RequisitionerName"`
			PurReqnItemCurrency            *string `json:"PurReqnItemCurrency"`
			MaterialPlannedDeliveryDurn    *string `json:"MaterialPlannedDeliveryDurn"`
			StorageLocation                *string `json:"StorageLocation"`
			PurReqnSourceOfSupplyType      *string `json:"PurReqnSourceOfSupplyType"`
			ConsumptionPosting             *string `json:"ConsumptionPosting"`
			PurReqnOrigin                  *string `json:"PurReqnOrigin"`
			IsPurReqnBlocked               *string `json:"IsPurReqnBlocked"`
			PurchaseRequisitionStatus      *string `json:"PurchaseRequisitionStatus"`
			Batch                          *string `json:"Batch"`
			GoodsReceiptIsExpected         *bool   `json:"GoodsReceiptIsExpected"`
			GoodsReceiptIsNonValuated      *bool   `json:"GoodsReceiptIsNonValuated"`
			RequirementTracking            *string `json:"RequirementTracking"`
			MRPController                  *string `json:"MRPController"`
			Reservation                    *string `json:"Reservation"`
			LastChangeDateTime             *string `json:"LastChangeDateTime"`
			IsDeleted                      *string `json:"IsDeleted"`
		} `json:"PurchaseRequisitionItem"`
	} `json:"PurchaseRequisition"`
	APISchema string   `json:"api_schema"`
	Accepter  []string `json:"accepter"`
	PRNumber  string   `json:"purchase_requisition"`
	Deleted   bool     `json:"deleted"`
}
