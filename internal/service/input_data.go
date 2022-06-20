package service

import "fmt"

func InputData() (repositoryName string, repositoryOwner string) {
	fmt.Print("Who own the repository: ")
	fmt.Scanf("%s", &repositoryOwner)
	fmt.Print("What is the name of the repository: ")
	fmt.Scanf("%s", &repositoryName)

	return
}
