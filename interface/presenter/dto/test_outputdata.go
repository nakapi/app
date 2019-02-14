package dto

type TestOuputData struct {
	Id    int
	Name  string
	Error error
}

func NewTestOuputData(id int, name string) *TestOuputData {
	testData := new(TestOuputData)
	testData.Id = id
	testData.Name = name
	return testData
}
