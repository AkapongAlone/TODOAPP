package requests

type Request struct {
	Name     string `validate:"acceptlist=aek|earth"`
	SomeType string `validate:"duplicate"`
}

type MainTest struct {
	Intro string `validate:"nonzero"`
	Data  []SubTest
}

type SubTest struct {
	ContractNumber string `validate:"unique=contract_info|contract_number,max=3"`
	AcceptList     string `validate:"acceptlist=aek|earth,nonzero"`
	TimeNow        string `validate:"date"`
	SubTest        []SubTest
}
