package service

import (
	"context"
	"fmt"
	"os"

	"github.com/sourcegraph/sourcegraph/enterprise/internal/batches/store"
	btypes "github.com/sourcegraph/sourcegraph/enterprise/internal/batches/types"
	"github.com/sourcegraph/sourcegraph/internal/api"
	"github.com/sourcegraph/sourcegraph/internal/authz"
	"github.com/sourcegraph/sourcegraph/internal/gitserver"
	"github.com/sourcegraph/sourcegraph/internal/trace"
	"github.com/sourcegraph/sourcegraph/internal/types"
	"github.com/sourcegraph/sourcegraph/lib/errors"
)

// GetRepoMetadata returns the repo metadata for the given repo, recalculating
// it if necessary.
//
// Note that this may block while a gitserver request is made, and therefore
// probably shouldn't be called from the hot path of GraphQL resolvers and the
// like.
func GetRepoMetadata(ctx context.Context, tx *store.Store, client gitserver.Client, repo *types.Repo) (*btypes.RepoMetadata, error) {
	meta, err := tx.GetRepoMetadata(ctx, repo.ID)
	if err != nil && err != store.ErrNoResults {
		return nil, errors.Wrap(err, "getting repo metadata")
	}

	// Check if we need to refresh the metadata.
	if err == store.ErrNoResults || meta.UpdatedAt.Before(repo.UpdatedAt) {
		meta, err = calculateRepoMetadata(ctx, client, repo)
		if err != nil {
			return nil, errors.Wrap(err, "refreshing repo metadata")
		}

		if err := tx.UpsertRepoMetadata(ctx, meta); err != nil {
			return nil, errors.Wrap(err, "upserting repo metadata")
		}
	}

	return meta, nil
}

const batchIgnoreFilePath = ".batchignore"

func calculateRepoMetadata(ctx context.Context, client gitserver.Client, repo *types.Repo) (meta *btypes.RepoMetadata, err error) {
	traceTitle := fmt.Sprintf("RepoID: %q", repo.ID)
	tr, ctx := trace.New(ctx, "hasBatchIgnoreFile", traceTitle)
	defer func() {
		tr.SetError(err)
		tr.Finish()
	}()

	// Figure out the head commit, since we need it to stat the file.
	commit, ok, err := client.Head(ctx, repo.Name, authz.DefaultSubRepoPermsChecker)
	if err != nil {
		return nil, errors.Wrapf(err, "resolving head commit in repo %q", string(repo.Name))
	}
	if !ok {
		return nil, errors.Newf("no head commit for repo %q", string(repo.Name))
	}

	meta = &btypes.RepoMetadata{RepoID: repo.ID, Ignored: false}
	meta.Ignored, err = hasBatchIgnoreFile(ctx, client, repo, api.CommitID(commit))
	if err != nil {
		return nil, errors.Wrapf(err, "looking for %s file in repo %q", batchIgnoreFilePath, string(repo.Name))
	}

	return meta, nil
}

func hasBatchIgnoreFile(ctx context.Context, client gitserver.Client, repo *types.Repo, commit api.CommitID) (bool, error) {
	stat, err := client.Stat(ctx, authz.DefaultSubRepoPermsChecker, repo.Name, commit, batchIgnoreFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	if !stat.Mode().IsRegular() {
		return false, errors.Errorf("not a blob: %q", batchIgnoreFilePath)
	}
	return true, nil
}
