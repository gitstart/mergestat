package syncer

import (
	"context"
	"errors"
	"fmt"
	"os/exec"

	"github.com/jackc/pgx/v4"
	"github.com/mergestat/fuse/internal/db"
)

// handleTrivyRepoScan executes `trivy repo {git-repo} -f json` for a repo
// and inserts the output JSON into the DB
func (w *worker) handleTrivyRepoScan(ctx context.Context, j *db.DequeueSyncJobRow) error {
	l := w.loggerForJob(j)

	var ghToken string
	var err error
	if ghToken, err = w.fetchGitHubTokenFromDB(ctx); err != nil {
		return err
	}

	if err := w.sendBatchLogMessages(ctx, []*syncLog{
		{
			Type:            SyncLogTypeInfo,
			RepoSyncQueueID: j.ID,
			Message:         "starting to execute trivy scan",
		},
	}); err != nil {
		return fmt.Errorf("log messages: %w", err)
	}

	cmd := exec.CommandContext(ctx, "trivy", "repository", j.Repo, "-q", "-f", "json")
	cmd.Env = append(cmd.Env, fmt.Sprintf("GITHUB_TOKEN=%s", ghToken))

	var output []byte
	if output, err = cmd.Output(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			w.logger.Err(exitErr).Msgf("error running trivy scan")
		}
		return fmt.Errorf("running trivy scan: %w", err)
	}

	var tx pgx.Tx
	if tx, err = w.pool.BeginTx(ctx, pgx.TxOptions{}); err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			if !errors.Is(err, pgx.ErrTxClosed) {
				w.logger.Err(err).Msgf("could not rollback transaction")
			}
		}
	}()

	if _, err := tx.Exec(ctx, "DELETE FROM trivy_repo_scans WHERE repo_id = $1;", j.RepoID.String()); err != nil {
		return fmt.Errorf("exec delete: %w", err)
	}

	if _, err := tx.Exec(ctx, "INSERT INTO trivy_repo_scans (repo_id, results) VALUES ($1, $2)", j.RepoID, output); err != nil {
		return fmt.Errorf("inserting trivy results: %w", err)
	}

	l.Info().Msg("inserted trivy scan results")

	if err := w.db.WithTx(tx).SetSyncJobStatus(ctx, db.SetSyncJobStatusParams{Status: "DONE", ID: j.ID}); err != nil {
		return fmt.Errorf("update status done: %w", err)
	}
	if err := w.sendBatchLogMessages(ctx, []*syncLog{
		{
			Type:            SyncLogTypeInfo,
			RepoSyncQueueID: j.ID,
			Message:         "finished",
		},
	}); err != nil {
		return fmt.Errorf("log messages: %w", err)
	}

	return tx.Commit(ctx)
}