package routers

import (
	"github.com/anle/codebase/internal/routers/account"
	"github.com/anle/codebase/internal/routers/transaction"
	"github.com/anle/codebase/internal/routers/user"
)

type RouterGroup struct {
	User        user.UserRouterGroup
	Account     account.AccountRouterGroup
	Transaction transaction.TransactionRouterGroup
}

var RouterGroupApp = new(RouterGroup)
