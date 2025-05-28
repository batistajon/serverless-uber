package types

type Receipt struct {
	EventID      string `json:"event_id" dynamodb:"EventID"`
	EventTime    int    `json:"event_time" dynamodb:"EventTime"`
	EventType    string `json:"event_type" dynamodb:"EventType"`
	ResourceHref string `json:"resource_href" dynamodb:"ResourceHref"`
	Meta         struct {
		UserID   string `json:"user_id" dynamodb:"UserID"`
		OrgUUID  string `json:"org_uuid" dynamodb:"OrgUUID"`
		Resource string `json:"resource_id" dynamodb:"Resource"`
		Status   string `json:"status" dynamodb:"Status"`
	} `json:"meta" dynamodb:"Meta"`
}

type ReceiptDTable struct {
	Name        string `json:"table_name" dynamodb:"Name"`
	ReceiptId   string `json:"receipt_id" dynamodb:"ReceiptId"`
	ReceiptData Receipt
}

func (r ReceiptDTable) GetName() string {
	return r.Name
}
