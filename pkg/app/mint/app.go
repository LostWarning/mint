package app

import (
	"os"
)

func Run() {

	cli := newCLI()
	// acc := account.NewSavingsAccount("001001570736")

	// transact.Deposit(&acc)
	// fmt.Printf("Balance in account : %d\n", acc.Balance())
	// fmt.Println("Running mint")

	cli.Run(os.Args)

}
