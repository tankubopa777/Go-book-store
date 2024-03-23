package userbooks

type (
	Userbooks struct {
		Id string `json:"id" bson:"_id,omitempty"`
		UserId string `json:"user_id" bson:"user_id"`
		BookId string `json:"book_id" bson:"book_id"`
	}
)