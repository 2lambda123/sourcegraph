package graphqlbackend

import (
	"context"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-middleware/providers/openmetrics/v2"

	"github.com/sourcegraph/sourcegraph/internal/database"
	"github.com/sourcegraph/sourcegraph/internal/search"
)

type searchAlertResolver struct {
	metrics *grpc_prometheus.ClientMetrics
	alert   *search.Alert
}

func NewSearchAlertResolver(alert *search.Alert, metrics *grpc_prometheus.ClientMetrics) *searchAlertResolver {
	if alert == nil {
		return nil
	}
	return &searchAlertResolver{alert: alert, metrics: metrics}
}

func (a searchAlertResolver) Title() string { return a.alert.Title }

func (a searchAlertResolver) Description() *string {
	if a.alert.Description == "" {
		return nil
	}
	return &a.alert.Description
}

func (a searchAlertResolver) Kind() *string {
	if a.alert.Kind == "" {
		return nil
	}
	return &a.alert.Kind
}

func (a searchAlertResolver) PrometheusType() string {
	return a.alert.PrometheusType
}

func (a searchAlertResolver) ProposedQueries() *[]*searchQueryDescriptionResolver {
	if len(a.alert.ProposedQueries) == 0 {
		return nil
	}
	var proposedQueries []*searchQueryDescriptionResolver
	for _, q := range a.alert.ProposedQueries {
		proposedQueries = append(proposedQueries, &searchQueryDescriptionResolver{q})
	}
	return &proposedQueries
}

func (a searchAlertResolver) wrapSearchImplementer(db database.DB) *alertSearchImplementer {
	return &alertSearchImplementer{
		db:      db,
		metrics: a.metrics,
		alert:   a,
	}
}

// alertSearchImplementer is a light wrapper type around an alert that implements
// SearchImplementer. This helps avoid needing to have a db on the searchAlert type
type alertSearchImplementer struct {
	db      database.DB
	metrics *grpc_prometheus.ClientMetrics
	alert   searchAlertResolver
}

func (a alertSearchImplementer) Results(context.Context) (*SearchResultsResolver, error) {
	return &SearchResultsResolver{db: a.db, metrics: a.metrics, SearchAlert: a.alert.alert}, nil
}

func (alertSearchImplementer) Stats(context.Context) (*searchResultsStats, error) { return nil, nil }
