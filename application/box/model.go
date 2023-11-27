package box

// MyBoxesRes ...
type MyBoxesRes struct {
	Boxes []*Box `json:"boxes,omitempty"`
}

// Box ...
type Box struct {
	Id          *string `json:"id,omitempty"`
	Name        *string `json:"name" binding:"required,min=3,max=30"`
	Desc        *string `json:"desc,omitempty"`
	CreatedByID *string `json:"created_by_id,omitempty"`
}
