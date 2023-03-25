package db

import (
	"bufio"
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestQueries_CreateAccount_and_FindLastInsertedId_and_FindAccountById(t *testing.T) {
	account := CreateAccountParams{
		Owner:    randomNOwner(1)[0],
		Balance:  randomMoney(),
		Currency: randomCurrency(),
	}

	err := testQueries.CreateAccount(context.Background(), account)

	require.NoError(t, err)

	lastInsertedId, err := testQueries.FindLastInsertedId(context.Background())
	accountInserted, err := testQueries.FindAccountById(context.Background(), lastInsertedId)

	require.NoError(t, err)

	require.NotZero(t, accountInserted.ID)
	require.NotZero(t, accountInserted.CreatedAt)

	require.NotEmpty(t, accountInserted)
	require.Equal(t, account.Owner, accountInserted.Owner)
	require.Equal(t, account.Balance, accountInserted.Balance)
	require.Equal(t, account.Currency, accountInserted.Currency)
}

func TestQueries_UpdateAccount(t *testing.T) {
	account := CreateAccountParams{
		Owner:    randomNOwner(1)[0],
		Balance:  randomMoney(),
		Currency: randomCurrency(),
	}

	err := testQueries.CreateAccount(context.Background(), account)

	require.NoError(t, err)

	lastInsertedId, err := testQueries.FindLastInsertedId(context.Background())
	accountInserted, err := testQueries.FindAccountById(context.Background(), lastInsertedId)
	require.NoError(t, err)

	updateAccountParams := UpdateAccountParams{ID: lastInsertedId, Balance: randomMoney()}

	err = testQueries.UpdateAccount(context.Background(), updateAccountParams)
	require.NoError(t, err)

	accountUpdated, err := testQueries.FindAccountById(context.Background(), lastInsertedId)
	require.NoError(t, err)

	require.NotZero(t, accountInserted.ID)
	require.NotZero(t, accountInserted.CreatedAt)

	require.NotEmpty(t, accountInserted)
	require.Equal(t, account.Owner, accountInserted.Owner)
	require.NotEqual(t, account.Balance, accountUpdated.Balance)
	require.Equal(t, account.Balance, accountInserted.Balance)
	require.Equal(t, account.Currency, accountInserted.Currency)
}

func TestQueries_FindAllAccounts(t *testing.T) {
	owners := randomNOwner(3)
	account1 := CreateAccountParams{
		Owner:    owners[0],
		Balance:  randomMoney(),
		Currency: randomCurrency(),
	}
	account2 := CreateAccountParams{
		Owner:    owners[1],
		Balance:  randomMoney(),
		Currency: randomCurrency(),
	}
	account3 := CreateAccountParams{
		Owner:    owners[2],
		Balance:  randomMoney(),
		Currency: randomCurrency(),
	}

	err := testQueries.CreateAccount(context.Background(), account1)
	require.NoError(t, err)

	err = testQueries.CreateAccount(context.Background(), account2)
	require.NoError(t, err)

	err = testQueries.CreateAccount(context.Background(), account3)
	require.NoError(t, err)

	accounts, err := testQueries.FindAllAccounts(context.Background())
	require.NoError(t, err)

	require.NotZero(t, len(accounts))
	fmt.Print(accounts)
}

func TestQueries_FindAllAccounts_ErrorOnRetrieveData(t *testing.T) {
	_, err := testQueries.FindAllAccounts(context.Background())
	require.NoError(t, err)
}

func TestQueries_DeleteAccount(t *testing.T) {
	account := CreateAccountParams{
		Owner:    randomNOwner(1)[0],
		Balance:  randomMoney(),
		Currency: randomCurrency(),
	}

	err := testQueries.CreateAccount(context.Background(), account)

	require.NoError(t, err)

	lastInsertedId, err := testQueries.FindLastInsertedId(context.Background())
	accountInserted, err := testQueries.FindAccountById(context.Background(), lastInsertedId)
	require.NoError(t, err)

	err = testQueries.DeleteAccount(context.Background(), lastInsertedId)
	require.NoError(t, err)

	accountDeleted, err := testQueries.FindAccountById(context.Background(), lastInsertedId)
	require.Error(t, err)

	require.NotZero(t, accountInserted.ID)
	require.NotZero(t, accountInserted.CreatedAt)
	require.NotEmpty(t, accountInserted)

	require.Zero(t, accountDeleted.ID)
	require.Zero(t, accountDeleted.CreatedAt)
	require.Empty(t, accountDeleted)
}

func randomNOwner(n int) []string {
	names := make([]string, 0)
	firstNameList := readFileLines("first-name.txt")
	lastNameList := readFileLines("last-name.txt")

	for i := 0; i < n; i++ {
		names = append(
			names,
			firstNameList[rand.Intn(len(firstNameList))]+
				" "+lastNameList[rand.Intn(len(lastNameList))]+
				" "+lastNameList[rand.Intn(len(lastNameList))])
	}

	return names
}

func readFileLines(fileName string) (lines []string) {
	file, err := os.Open("../../test/resource/" + fileName)
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
	return int64(rand.Intn(10000-0) + 0)
}

func randomCurrency() string {
	currencies := []string{"BRL", "CAD", "ARS", "GQT", "CNH"}
	size := len(currencies)

	return currencies[rand.Intn(size)]
}
