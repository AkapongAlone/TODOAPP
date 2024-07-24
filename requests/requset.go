package requests

type Request struct {
	Name     string `validate:"acceptlist=aek|earth"`
	SomeType string `validate:"duplicate"`
}


type Test struct {
	ContractNumber string `validate:"unique=contract_info|contract_number,max=3"`
	AcceptList     string `validate:"acceptlist=aek|earth"`
	TimeNow        string `validate:"date"`
}
