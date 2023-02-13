package embeddings

import (
	"context"
	"time"

	"github.com/sourcegraph/sourcegraph/cmd/worker/job"
	workerdb "github.com/sourcegraph/sourcegraph/cmd/worker/shared/init/db"
	edb "github.com/sourcegraph/sourcegraph/enterprise/internal/database"
	"github.com/sourcegraph/sourcegraph/enterprise/internal/embeddings"
	embeddingsbg "github.com/sourcegraph/sourcegraph/enterprise/internal/embeddings/background"
	"github.com/sourcegraph/sourcegraph/internal/actor"
	"github.com/sourcegraph/sourcegraph/internal/env"
	"github.com/sourcegraph/sourcegraph/internal/gitserver"
	"github.com/sourcegraph/sourcegraph/internal/goroutine"
	"github.com/sourcegraph/sourcegraph/internal/observation"
	"github.com/sourcegraph/sourcegraph/internal/uploadstore"
	"github.com/sourcegraph/sourcegraph/internal/workerutil"
	"github.com/sourcegraph/sourcegraph/internal/workerutil/dbworker"
	dbworkerstore "github.com/sourcegraph/sourcegraph/internal/workerutil/dbworker/store"
)

type repoEmbeddingJob struct{}

func NewRepoEmbeddingJob() job.Job {
	return &repoEmbeddingJob{}
}

func (s *repoEmbeddingJob) Description() string {
	return ""
}

func (s *repoEmbeddingJob) Config() []env.Config {
	return []env.Config{embeddings.EmbeddingsUploadStoreConfigInst}
}

func (s *repoEmbeddingJob) Routines(_ context.Context, observationCtx *observation.Context) ([]goroutine.BackgroundRoutine, error) {
	// TODO: Check if embeddings are enabled
	db, err := workerdb.InitDB(observationCtx)
	if err != nil {
		return nil, err
	}

	workCtx := actor.WithInternalActor(context.Background())
	uploadStore, err := embeddings.NewEmbeddingsUploadStore(workCtx, observationCtx, embeddings.EmbeddingsUploadStoreConfigInst)
	if err != nil {
		return nil, err
	}

	return []goroutine.BackgroundRoutine{
		newRepoEmbeddingJobWorker(
			workCtx,
			observationCtx,
			embeddingsbg.NewRepoEmbeddingJobWorkerStore(observationCtx, db.Handle()),
			edb.NewEnterpriseDB(db),
			uploadStore,
			gitserver.NewClient(),
		),
	}, nil
}

func newRepoEmbeddingJobWorker(
	ctx context.Context,
	observationCtx *observation.Context,
	workerStore dbworkerstore.Store[*embeddingsbg.RepoEmbeddingJob],
	db edb.EnterpriseDB,
	uploadStore uploadstore.Store,
	gitserverClient gitserver.Client,
) *workerutil.Worker[*embeddingsbg.RepoEmbeddingJob] {
	handler := &handler{db, uploadStore, gitserverClient}
	return dbworker.NewWorker[*embeddingsbg.RepoEmbeddingJob](ctx, workerStore, handler, workerutil.WorkerOptions{
		Name:              "repo_embedding_job_worker",
		Interval:          time.Second, // Poll for a job once per second
		NumHandlers:       1,           // Process only one job at a time (per instance)
		HeartbeatInterval: 10 * time.Second,
		Metrics:           workerutil.NewMetrics(observationCtx, "repo_embedding_job_worker"),
	})
}
