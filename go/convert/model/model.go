package model

type Date struct {
	Date         string `json:"date,omitempty"`
	Event        string `json:"event,omitempty"`
	Time         string `json:"time,omitempty"`
	Outfit       string `json:"outfit,omitempty"`
	MeetingPoint string `json:"meeting_point,omitempty"`
	ICSFile      string `json:"ics_file,omitempty"`
}
