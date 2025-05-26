package types

type Receipt struct {
	EventID      string `json:"event_id"`
	EventTime    int    `json:"event_time"`
	EventType    string `json:"event_type"`
	ResourceHref string `json:"resource_href"`
	Meta         struct {
		UserID   string `json:"user_id"`
		OrgUUID  string `json:"org_uuid"`
		Resource string `json:"resource_id"`
		Status   string `json:"status"`
	} `json:"meta"`
}
