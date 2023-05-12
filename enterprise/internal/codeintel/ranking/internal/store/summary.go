package store

import (
	"context"
	"time"

	"github.com/keegancsmith/sqlf"

	"github.com/sourcegraph/sourcegraph/enterprise/internal/codeintel/ranking/shared"
	"github.com/sourcegraph/sourcegraph/internal/database/basestore"
	"github.com/sourcegraph/sourcegraph/internal/database/dbutil"
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
			graphKey                     string
			mappersStartedAt             time.Time
			mapperCompletedAt            *time.Time
			seedMapperCompletedAt        *time.Time
			reducerStartedAt             *time.Time
			reducerCompletedAt           *time.Time
			numPathRecordsTotal          int
			numReferenceRecordsTotal     int
			numCountRecordsTotal         int
			numPathRecordsProcessed      int
			numReferenceRecordsProcessed int
			numCountRecordsProcessed     int
		)
		if err := rows.Scan(
			&graphKey,
			&mappersStartedAt,
			&mapperCompletedAt,
			&seedMapperCompletedAt,
			&reducerStartedAt,
			&reducerCompletedAt,
			&dbutil.NullInt{N: &numPathRecordsTotal},
			&dbutil.NullInt{N: &numReferenceRecordsTotal},
			&dbutil.NullInt{N: &numCountRecordsTotal},
			&dbutil.NullInt{N: &numPathRecordsProcessed},
			&dbutil.NullInt{N: &numReferenceRecordsProcessed},
			&dbutil.NullInt{N: &numCountRecordsProcessed},
		); err != nil {
			return nil, err
		}

		pathMapperProgress := shared.Progress{
			StartedAt:   mappersStartedAt,
			CompletedAt: seedMapperCompletedAt,
			Processed:   numPathRecordsProcessed,
			Total:       numPathRecordsTotal,
		}

		referenceMapperProgress := shared.Progress{
			StartedAt:   mappersStartedAt,
			CompletedAt: mapperCompletedAt,
			Processed:   numReferenceRecordsProcessed,
			Total:       numReferenceRecordsTotal,
		}

		var reducerProgress *shared.Progress
		if reducerStartedAt != nil {
			reducerProgress = &shared.Progress{
				StartedAt:   *reducerStartedAt,
				CompletedAt: reducerCompletedAt,
				Processed:   numCountRecordsProcessed,
				Total:       numCountRecordsTotal,
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
	p.reducer_completed_at,
	p.num_path_records_total,
	p.num_reference_records_total,
	p.num_count_records_total,
	p.num_path_records_processed,
	p.num_reference_records_processed,
	p.num_count_records_processed
FROM codeintel_ranking_progress p
ORDER BY p.mappers_started_at DESC
`
