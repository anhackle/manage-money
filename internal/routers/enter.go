package routers

import (
	"github.com/anle/codebase/internal/routers/account"
	"github.com/anle/codebase/internal/routers/group"
	"github.com/anle/codebase/internal/routers/groupdistributed"
	"github.com/anle/codebase/internal/routers/transaction"
	"github.com/anle/codebase/internal/routers/user"
)

type RouterGroup struct {
	User        user.UserRouterGroup
	Account     account.AccountRouterGroup
	Transaction transaction.TransactionRouterGroup
	Group       group.GroupRouterGroup
	GroupDis    groupdistributed.GroupDisRouterGroup
}

var RouterGroupApp = new(RouterGroup)
