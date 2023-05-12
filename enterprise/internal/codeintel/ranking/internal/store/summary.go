package store

import (
	"context"
	"time"

	"github.com/keegancsmith/sqlf"

	"github.com/sourcegraph/sourcegraph/enterprise/internal/codeintel/ranking/shared"
	"github.com/sourcegraph/sourcegraph/internal/database/basestore"
)

func (s *store) Summaries(ctx context.Context) (_ []shared.Summary, err error) {
	// TODO - add observability

	rows, err := s.db.Query(ctx, sqlf.Sprintf(summariesQuery))
	if err != nil {
		return nil, err
	}
	defer func() { err = basestore.CloseRows(rows, err) }()

	// TODO - extract scanner
	var summaries []shared.Summary
	for rows.Next() {
		var (
			graphKey              string
			mappersStartedAt      time.Time
			mapperCompletedAt     *time.Time
			seedMapperCompletedAt *time.Time
			reducerStartedAt      *time.Time
			reducerCompletedAt    *time.Time
		)
		if err := rows.Scan(
			&graphKey,
			&mappersStartedAt,
			&mapperCompletedAt,
			&seedMapperCompletedAt,
			&reducerStartedAt,
			&reducerCompletedAt,
		); err != nil {
			return nil, err
		}

		pathMapperProgress := shared.Progress{
			StartedAt:   mappersStartedAt,
			CompletedAt: seedMapperCompletedAt,
			Processed:   0, // TODO
			Total:       0, // TODO
		}

		referenceMapperProgress := shared.Progress{
			StartedAt:   mappersStartedAt,
			CompletedAt: mapperCompletedAt,
			Processed:   0, // TODO
			Total:       0, // TODO
		}

		var reducerProgress *shared.Progress
		if reducerStartedAt != nil {
			reducerProgress = &shared.Progress{
				StartedAt:   *reducerStartedAt,
				CompletedAt: reducerCompletedAt,
				Processed:   0, // TODO
				Total:       0, // TODO
			}
		}

		summaries = append(summaries, shared.Summary{
			GraphKey:                graphKey,
			PathMapperProgress:      pathMapperProgress,
			ReferenceMapperProgress: referenceMapperProgress,
			ReducerProgress:         reducerProgress,
		})
	}

	return summaries, nil
}

//
// TODO - progress janitor

const summariesQuery = `
SELECT
	p.graph_key,
	p.mappers_started_at,
	p.mapper_completed_at,
	p.seed_mapper_completed_at,
	p.reducer_started_at,
	p.reducer_completed_at
FROM codeintel_ranking_progress p
ORDER BY p.mappers_started_at desc
`
