package db

import (
	"bufio"
	"github.com/satori/go.uuid"
	"log"
	"math/rand"
	"os"
)

const currentDir = "../../test/resource/"

func GetNewRandomAccountParams(n int) []CreateAccountParams {
	accounts := make([]CreateAccountParams, 0)

	for i := 0; i < n; i++ {
		accounts = append(accounts, CreateAccountParams{
			ID:       uuid.NewV4().String(),
			Owner:    randomNOwner(currentDir),
			Balance:  randomMoney(),
			Currency: randomCurrency(),
		})
	}
	return accounts
}

func randomNOwner(dir string) string {
	firstNameList := readFileLines(dir, "first-name.txt")
	lastNameList := readFileLines(dir, "last-name.txt")

	return firstNameList[rand.Intn(len(firstNameList))] +
		" " + lastNameList[rand.Intn(len(lastNameList))] +
		" " + lastNameList[rand.Intn(len(lastNameList))]
}

func readFileLines(dir, fileName string) (lines []string) {
	file, err := os.Open(dir + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	sc := bufio.NewScanner(file)
	lines = make([]string, 0)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return
}

func randomMoney() int64 {
	return int64(rand.Intn(10000-100) + 100)
}

func randomCurrency() string {
	currencies := []string{"BRL", "CAD", "ARS", "GQT", "CNH"}
	size := len(currencies)

	return currencies[rand.Intn(size)]
}
