package main

import (
	"fmt"
	"os"

	"tudoo.app/cli/db"
	"tudoo.app/cli/types"
	"tudoo.app/cli/utils"
)

func main() {
	var args = os.Args[1:]
	for i := range args {
		utils.Log(args[i])
	}

	var name string
	var age int
	var height float32

	fmt.Print("Enter name: ")
	name = utils.ReadString()
	fmt.Print("Enter age: ")
	fmt.Scanln(&age)
	fmt.Print("Enter height: ")
	fmt.Scanln(&height)

	satvik := types.CreatePerson(name, age, height)
	var insertedName, insertedAge, err = db.InsertPerson(satvik)
	if err != nil {
		panic(err)
	}
	utils.Heading("Inserted person:")
	utils.Log("Name: " + insertedName)
	utils.Log("Age: " + insertedAge)
}
