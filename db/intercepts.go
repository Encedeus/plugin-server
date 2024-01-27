package db

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/Encedeus/pluginServer/ent"
	"github.com/Encedeus/pluginServer/ent/publication"
)

func PublicationOrderIntercept() ent.InterceptFunc {
	return ent.InterceptFunc(func(next ent.Querier) ent.Querier {
		return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
			if q, ok := query.(*ent.PublicationQuery); ok {
				q.Order(publication.ByCreatedAt(sql.OrderDesc()))
			}
			return next.Query(ctx, query)
		})
	})
}
