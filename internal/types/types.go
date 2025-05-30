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
	Name        string  `json:"table_name"`
	ReceiptId   string  `json:"receipt_id" dynamodb:"ReceiptId"`
	ReceiptData Receipt `json:"Receipt_data" dynamodb:"ReceiptData"`
}

func (r ReceiptDTable) GetName() string {
	return r.Name
}

func (r ReceiptDTable) AddItem() {}
func (r ReceiptDTable) GetItem() {}

type Credentials struct {
	GoogleApiKey     string `json:"google_api_key" dynamodb:"GoogleApiKey"`
	UberClientSecret string `json:"uber_client_secret" dynamodb:"UberClientSecret"`
}

type User struct {
	Name        string      `json:"table_name"`
	UserId      string      `json:"user_id" dynamodb:"UserId"`
	UberUserId  string      `json:"uber_user_id" dynamodb:"UberUserId"`
	credentials Credentials `json:"credentials" dynamodb:"credentials"`
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetCredentials() Credentials {
	return u.credentials
}

func (u *User) AddItem() {}
func (u *User) GetItem() {}
