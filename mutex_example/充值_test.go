package mutex_example

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Account struct {
	balance float64
	mu      sync.Mutex
}

type UserManager struct {
	accounts map[int]*Account
	mu       sync.Mutex
	locks    []*sync.Mutex
}

func NewUserManager() *UserManager {
	return &UserManager{
		accounts: make(map[int]*Account),
		locks:    make([]*sync.Mutex, 100), // 假设总共有100个用户，每个用户对应一个锁
	}
}

func (u *UserManager) GetUserAccount(userID int) *Account {
	u.mu.Lock()
	defer u.mu.Unlock()

	account, ok := u.accounts[userID]
	if !ok {
		account = &Account{balance: 0}
		u.accounts[userID] = account
	}

	return account
}

func (u *UserManager) LockUser(userID int) {
	lock := u.locks[userID%len(u.locks)]
	lock.Lock()
}

func (u *UserManager) UnlockUser(userID int) {
	lock := u.locks[userID%len(u.locks)]
	lock.Unlock()
}

func (u *UserManager) Deposit(userID int, amount float64) {
	account := u.GetUserAccount(userID)

	userIDString := fmt.Sprintf("%d", userID)
	fmt.Printf("用户%s开始充值，充值金额: %.2f\n", userIDString, amount)
	time.Sleep(time.Second * 3) // 模拟充值过程

	account.mu.Lock()
	defer account.mu.Unlock()

	account.balance += amount

	fmt.Printf("用户%s充值完成\n", userIDString)
}

func TestDeposit(t *testing.T) {

	userManager := NewUserManager()
	wg := sync.WaitGroup{}

	// 启动10个goroutine并发进行充值操作
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			userID := i % 4 // 模拟4个用户
			userManager.LockUser(userID)
			defer userManager.UnlockUser(userID)

			userManager.Deposit(userID, 100.0)
		}(i)
	}

	// 等待所有充值任务完成
	wg.Wait()

	// 输出每个用户的账户余额
	for userID, account := range userManager.accounts {
		fmt.Printf("用户%d的账户余额: %.2f\n", userID, account.balance)
	}
}
