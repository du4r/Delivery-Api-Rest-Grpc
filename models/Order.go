package models

type Order struct{
	Costumer
	Title string `json:"title"`
	Price float32 `json:"price"`
	Status Status
}

type Status int

const ( 
	Created Status = iota
	Preparing 
	Pending
	OnWay
	Delivered
	Cancelled
)

func(status Status) String() string{
	switch status{
	case Created:
		return "created"
	case Preparing:
		return "preparing"
	case Pending:
		return "pending"
	case OnWay:
		return "onWay"
	case Delivered:
		return "Delivered"
	case Cancelled:
		return "Cancelled"
	}
	return "unknown"
}