package requests

type Text struct {
	ShopType     		string  `json:"ShopType"`
	Language          	string  `json:"Language"`
	ShopTypeName		string  `json:"ShopTypeName"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
