package sqlds

import (
	"context"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

// QueryLifeCycle implements various lifecycle hooks to the sqldatasource QueryData method
type QueryLifeCycle interface {
	PreProcessQuery(ctx context.Context, query backend.DataQuery) (context.Context, backend.DataQuery)
	PostProcessQuery(ctx context.Context, q *Query) (context.Context, *Query)
	PostProcessResults(ctx context.Context, frames data.Frames, err error) (context.Context, data.Frames, error)
}

func (ds *sqldatasource) PreProcessQuery(ctx context.Context, query backend.DataQuery) (context.Context, backend.DataQuery) {
	if ds.QueryLifeCycle == nil {
		return ctx, query
	}
	return ds.QueryLifeCycle.PreProcessQuery(ctx, query)
}
func (ds *sqldatasource) PostProcessQuery(ctx context.Context, q *Query) (context.Context, *Query) {
	if ds.QueryLifeCycle == nil {
		return ctx, q
	}
	return ds.QueryLifeCycle.PostProcessQuery(ctx, q)
}
func (ds *sqldatasource) PostProcessResults(ctx context.Context, frames data.Frames, err error) (context.Context, data.Frames, error) {
	if ds.QueryLifeCycle == nil {
		return ctx, frames, err
	}
	return ds.QueryLifeCycle.PostProcessResults(ctx, frames, err)
}
