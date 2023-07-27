package api

import (
	authregistry "synergize/api/auth"
	userregistry "synergize/api/user"

	"synergize/api/auth/repository/repositorygorm"
	repositorybankaccount "synergize/api/bankaccount/repository/repositorygorm"
	transactionservice "synergize/api/transaction"

	"synergize/api/auth/repository/repositoryredis"
	"synergize/api/auth/usecase/usecaseimpl"
	usecaseimplbankaccount "synergize/api/bankaccount/usecase/usecaseimpl"
	repositorytransaction "synergize/api/transaction/repository/repositorygorm"
	usecaseimpltransaction "synergize/api/transaction/usecase/usecaseimpl"
	repositoryuser "synergize/api/user/repository/repositorygorm"
	usecaseimpluser "synergize/api/user/usecase/usecaseimpl"

	"synergize/api/bankaccount"

	"github.com/go-chi/jwtauth/v5"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func deliveryAuth(db *gorm.DB, cacheRds *redis.Client, jwtAuth *jwtauth.JWTAuth) *authregistry.AuthHttpRouterRegistry {
	authRepoRds := repositoryredis.NewRepositoryRedis(cacheRds)
	authRepoGorm := repositorygorm.NewRepositoryGorm(db)
	authImpl := usecaseimpl.NewAuthService(authRepoGorm, jwtAuth, authRepoRds)
	authRoute := authregistry.NewAuthHttpRouterRegistry(authImpl)
	return authRoute
}

func deliveryBankAccount(db *gorm.DB) *bankaccount.BankAccountHttpRouterRegistry {
	bankAccountGorm := repositorybankaccount.NewBankAccountRepository(db)
	bankAccountImpl := usecaseimplbankaccount.NewBankAccountService(bankAccountGorm)
	bankAccountRoute := bankaccount.NewBankAccountHttpRouterRegistry(bankAccountImpl)
	return bankAccountRoute
}

func deliveryTransaction(db *gorm.DB) *transactionservice.TransactionHttpRouterRegistry {
	bankAccountGorm := repositorybankaccount.NewBankAccountRepository(db)
	transactionGorm := repositorytransaction.NewRepositoryGormTransaction(db)
	transactionImpl := usecaseimpltransaction.NewTransactionService(transactionGorm, bankAccountGorm)
	transactionAccountRoute := transactionservice.NewTransactionHttpRouterRegistry(transactionImpl)
	return transactionAccountRoute
}

func deliveryUser(db *gorm.DB) *userregistry.UserHttpRouterRegistry {
	userRepo := repositoryuser.NewRepositoryGormUser(db)
	userimpl := usecaseimpluser.NewServiceUser(userRepo)
	userRoute := userregistry.NewUserHttpRouterRegistry(userimpl)
	return userRoute
}
