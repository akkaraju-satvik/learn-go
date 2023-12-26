package types

type Person struct {
	Name   string
	Age    int
	Height float32
}

func (p Person) Iterate() map[string]interface{} {
	var m = make(map[string]interface{})
	m["Name"] = p.Name
	m["Age"] = p.Age
	m["Height"] = p.Height
	return m
}

func (p Person) IsUnderage() bool {
	return p.Age < 18
}

type Student struct {
	Person
	RollNo int
}

func CreatePerson(name string, age int, height float32) Person {
	return Person{
		Name:   name,
		Age:    age,
		Height: height,
	}
}

func (s Student) Iterate() map[string]interface{} {
	var m = make(map[string]interface{})
	for key, value := range s.Person.Iterate() {
		m[key] = value
	}
	m["RollNo"] = s.RollNo
	return m
}
