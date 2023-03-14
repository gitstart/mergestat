// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

// git blame of all lines in all files of a repo
type GitBlame struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// email of the author who modified the line
	AuthorEmail sql.NullString
	// name of the author who modified the line
	AuthorName sql.NullString
	AuthorWhen sql.NullTime
	// hash of the commit the modification was made in
	CommitHash sql.NullString
	// line number of the modification
	LineNo int32
	Line   sql.NullString
	// path of the file the modification was made in
	Path string
	// timestamp when record was synced into the MergeStat database
	MergestatSyncedAt time.Time
}

type GitBranch struct {
	// foreign key for public.repos.id
	RepoID   uuid.UUID
	FullName string
	// hash of the commit for refs that are not of type tag
	Hash sql.NullString
	// name of the ref
	Name sql.NullString
	// remote of the ref
	Remote sql.NullString
	// target of the ref
	Target sql.NullString
	// type of the ref
	Type sql.NullString
	// hash of the commit for refs that are of type tag
	TagCommitHash sql.NullString
	// timestamp when record was synced into the MergeStat database
	MergestatSyncedAt time.Time
}

// git commit history of a repo
type GitCommit struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// hash of the commit
	Hash string
	// message of the commit
	Message string
	// name of the author of the the modification
	AuthorName string
	// email of the author of the modification
	AuthorEmail string
	// timestamp of when the modifcation was authored
	AuthorWhen time.Time
	// name of the author who committed the modification
	CommitterName string
	// email of the author who committed the modification
	CommitterEmail string
	// timestamp of when the commit was made
	CommitterWhen time.Time
	// the number of parents of the commit
	Parents int32
	// timestamp when record was synced into the MergeStat database
	MergestatSyncedAt time.Time
}

// git commit stats of a repo
type GitCommitStat struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// hash of the commit
	CommitHash string
	// path of the file the modification was made in
	FilePath string
	// the number of additions in this path of the commit
	Additions int32
	// the number of deletions in this path of the commit
	Deletions int32
	// timestamp when record was synced into the MergeStat database
	MergestatSyncedAt time.Time
	// old file mode derived from git mode. possible values (unknown, none, regular_file, symbolic_link, git_link)
	OldFileMode string
	// new file mode derived from git mode. possible values (unknown, none, regular_file, symbolic_link, git_link)
	NewFileMode string
}

// git files (content and paths) of a repo
type GitFile struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// path of the file
	Path string
	// boolean to determine if the file is an executable
	Executable bool
	// contents of the file
	Contents sql.NullString
	// timestamp when record was synced into the MergeStat database
	MergestatSyncedAt time.Time
}

// git refs of a repo
type GitRef struct {
	// foreign key for public.repos.id
	RepoID   uuid.UUID
	FullName string
	// hash of the commit for refs that are not of type tag
	Hash sql.NullString
	// name of the ref
	Name sql.NullString
	// remote of the ref
	Remote sql.NullString
	// target of the ref
	Target sql.NullString
	// type of the ref
	Type sql.NullString
	// hash of the commit for refs that are of type tag
	TagCommitHash sql.NullString
	// timestamp when record was synced into the MergeStat database
	MergestatSyncedAt time.Time
}

// table of git repo remotes
type GitRemote struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// name of the remote
	Name string
	// url of the remote
	Url string
	// timestamp when record was synced into the MergeStat database
	MergestatSyncedAt time.Time
}

type GitTag struct {
	// foreign key for public.repos.id
	RepoID   uuid.UUID
	FullName string
	// hash of the ref that are of type tag
	Hash sql.NullString
	// name of the ref
	Name sql.NullString
	// remote of the ref
	Remote sql.NullString
	// target of the ref
	Target sql.NullString
	// type of the ref
	Type sql.NullString
	// hash of the commit for refs that are of type tag
	TagCommitHash sql.NullString
	// timestamp when record was synced into the MergeStat database
	MergestatSyncedAt time.Time
}

type GithubActionsWorkflow struct {
	RepoID            uuid.UUID
	ID                int64
	WorkflowNodeID    sql.NullString
	Name              sql.NullString
	Path              sql.NullString
	State             sql.NullString
	CreatedAt         sql.NullTime
	UpdatedAt         sql.NullTime
	Url               sql.NullString
	HtmlUrl           sql.NullString
	BadgeUrl          sql.NullString
	MergestatSyncedAt time.Time
}

type GithubActionsWorkflowRun struct {
	RepoID            uuid.UUID
	ID                int64
	WorkflowRunNodeID sql.NullString
	Name              sql.NullString
	HeadBranch        sql.NullString
	RunNumber         sql.NullInt32
	RunAttempt        sql.NullInt32
	Event             sql.NullString
	Status            sql.NullString
	Conclusion        sql.NullString
	WorkflowID        sql.NullInt64
	CheckSuiteID      sql.NullInt64
	CheckSuiteNodeID  sql.NullString
	Url               sql.NullString
	HtmlUrl           sql.NullString
	PullRequests      pgtype.JSONB
	CreatedAt         sql.NullTime
	UpdatedAt         sql.NullTime
	RunStartedAt      sql.NullTime
	JobsUrl           sql.NullString
	LogsUrl           sql.NullString
	CheckSuiteUrl     sql.NullString
	ArtifactsUrl      sql.NullString
	CancelUrl         sql.NullString
	RerunUrl          sql.NullString
	HeadCommit        pgtype.JSONB
	WorkflowUrl       sql.NullString
	RepositoryUrl     sql.NullString
	HeadRepositoryUrl sql.NullString
	MergestatSyncedAt time.Time
}

type GithubActionsWorkflowRunJob struct {
	RepoID            uuid.UUID
	ID                int64
	RunID             sql.NullInt64
	Log               sql.NullString
	RunUrl            sql.NullString
	JobNodeID         sql.NullString
	HeadSha           sql.NullString
	Url               sql.NullString
	HtmlUrl           sql.NullString
	Status            sql.NullString
	Conclusion        sql.NullString
	StartedAt         sql.NullTime
	CompletedAt       sql.NullTime
	WorkflowName      sql.NullString
	Steps             pgtype.JSONB
	CheckRunUrl       sql.NullString
	Labels            pgtype.JSONB
	RunnerID          sql.NullInt64
	RunnerName        sql.NullString
	RunnerGroupID     sql.NullInt64
	RunnerGroupName   sql.NullString
	MergestatSyncedAt time.Time
}

// issues of a GitHub repo
type GithubIssue struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// login of the author of the issue
	AuthorLogin sql.NullString
	// body of the issue
	Body sql.NullString
	// boolean to determine if the issue is closed
	Closed sql.NullBool
	// timestamp of when the issue was closed
	ClosedAt     sql.NullTime
	CommentCount sql.NullInt32
	CreatedAt    sql.NullTime
	// boolean to determine if the issue was created via email
	CreatedViaEmail sql.NullBool
	// GitHub database_id of the issue
	DatabaseID int32
	// login of the editor of the issue
	EditorLogin sql.NullString
	// boolean to determine if the issue was edited and includes an edit with the creation data
	IncludesCreatedEdit sql.NullBool
	// number of labels associated to the issue
	LabelCount sql.NullInt32
	// timestamp when the issue was edited
	LastEditedAt sql.NullTime
	// boolean to determine if the issue is locked
	Locked sql.NullBool
	// number of milestones associated to the issue
	MilestoneCount sql.NullInt32
	// GitHub number for the issue
	Number int32
	// number of participants associated to the issue
	ParticipantCount sql.NullInt32
	// timestamp when the issue was published
	PublishedAt sql.NullTime
	// number of reactions associated to the issue
	ReactionCount sql.NullInt32
	// state of the issue
	State sql.NullString
	// title of the issue
	Title sql.NullString
	// timestamp when the issue was updated
	UpdatedAt sql.NullTime
	// GitHub URL of the issue
	Url sql.NullString
	// labels associated to the issue
	Labels pgtype.JSONB
	// timestamp when record was synced into the MergeStat database
	MergestatSyncedAt time.Time
}

// GitHub Workflow Run Jobs
type GithubPullRequest struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// the number of additions in the pull request
	Additions sql.NullInt32
	// login of the author of the pull request
	AuthorLogin sql.NullString
	// author association to the repo
	AuthorAssociation sql.NullString
	// URL to the avatar of the author of the pull request
	AuthorAvatarUrl sql.NullString
	// name of the author of the pull request
	AuthorName sql.NullString
	// the base ref object id associated with the pull request
	BaseRefOid sql.NullString
	// the name of base ref associated with the pull request
	BaseRefName sql.NullString
	// the name of the base repo for the pull request
	BaseRepositoryName sql.NullString
	// body of the pull request
	Body sql.NullString
	// the number of files changed/modified in the pull request
	ChangedFiles sql.NullInt32
	// boolean to determine if the pull request is closed
	Closed sql.NullBool
	// timestamp of when the pull request was closed
	ClosedAt sql.NullTime
	// the number of comments in the pull request
	CommentCount sql.NullInt32
	CommitCount  sql.NullInt32
	// timestamp of when the pull request was created
	CreatedAt sql.NullTime
	// boolean to determine if the pull request was created via email
	CreatedViaEmail sql.NullBool
	// GitHub database_id of the pull request
	DatabaseID int32
	// the number of deletions in the pull request
	Deletions sql.NullInt32
	// login of the editor of the pull request
	EditorLogin sql.NullString
	// the name of head ref associated with the pull request
	HeadRefName sql.NullString
	// the head ref object id associated with the pull request
	HeadRefOid sql.NullString
	// the name of the head repo for the pull request
	HeadRepositoryName sql.NullString
	// boolean to determine if the pull request is a draft
	IsDraft sql.NullBool
	// number of labels associated to the pull request
	LabelCount sql.NullInt32
	// timestamp of when the pull request was last edited
	LastEditedAt sql.NullTime
	// boolean to determine if the pull request is locked
	Locked sql.NullBool
	// boolean to determine if a maintainer can modify the pull request
	MaintainerCanModify sql.NullBool
	// mergeable state of the pull request
	Mergeable sql.NullString
	// boolean to determine if the pull request is merged
	Merged sql.NullBool
	// timestamp of when the pull request was merged
	MergedAt sql.NullTime
	// actor who merged the pull request
	MergedBy sql.NullString
	// GitHub number of the pull request
	Number sql.NullInt32
	// number of participants associated to the pull request
	ParticipantCount sql.NullInt32
	// timestamp of when the pull request was published
	PublishedAt sql.NullTime
	// review decision of the pull request
	ReviewDecision sql.NullString
	// state of the pull request
	State sql.NullString
	// title of the pull request
	Title sql.NullString
	// timestamp of when the pull request was updated
	UpdatedAt sql.NullTime
	// GitHub URL of the pull request
	Url sql.NullString
	// labels associated to the pull request
	Labels pgtype.JSONB
	// timestamp when record was synced into the MergeStat database
	MergestatSyncedAt time.Time
}

// commits for all pull requests of a GitHub repo
type GithubPullRequestCommit struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// GitHub pull request number of the commit
	PrNumber int32
	// hash of the commit
	Hash sql.NullString
	// message of the commit
	Message sql.NullString
	// name of the author of the the modification
	AuthorName sql.NullString
	// email of the author of the modification
	AuthorEmail sql.NullString
	// timestamp of when the modifcation was authored
	AuthorWhen sql.NullTime
	// name of the author who committed the modification
	CommitterName sql.NullString
	// email of the author who committed the modification
	CommitterEmail sql.NullString
	// timestamp of when the commit was made
	CommitterWhen sql.NullTime
	// the number of additions in the commit
	Additions sql.NullInt32
	// the number of deletions in the commit
	Deletions sql.NullInt32
	// the number of files changed/modified in the commit
	ChangedFiles sql.NullInt32
	// GitHub URL of the commit
	Url sql.NullString
	// timestamp when record was synced into the MergeStat database
	MergestatSyncedAt time.Time
}

// reviews for all pull requests of a GitHub repo
type GithubPullRequestReview struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// GitHub pull request number of the review
	PrNumber int32
	// GitHub id of the review
	ID string
	// login of the author of the review
	AuthorLogin sql.NullString
	// URL to the profile of the author of the review
	AuthorUrl sql.NullString
	// author association to the repo
	AuthorAssociation sql.NullString
	// boolean to determine if author can push to the repo
	AuthorCanPushToRepository sql.NullBool
	// body of the review
	Body sql.NullString
	// number of comments associated to the review
	CommentCount sql.NullInt32
	// timestamp of when the review was created
	CreatedAt sql.NullTime
	// boolean to determine if the review was created via email
	CreatedViaEmail sql.NullBool
	// login of the editor of the review
	EditorLogin sql.NullString
	// timestamp of when the review was last edited
	LastEditedAt sql.NullTime
	// timestamp of when the review was published
	PublishedAt sql.NullTime
	// state of the review
	State sql.NullString
	// timestamp of when the review was submitted
	SubmittedAt sql.NullTime
	// timestamp of when the review was updated
	UpdatedAt sql.NullTime
	// timestamp when record was synced into the MergeStat database
	MergestatSyncedAt time.Time
}

// info/metadata of a GitHub repo
type GithubRepoInfo struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// the user or organization that owns the repo
	Owner string
	// the name of the repo
	Name string
	// timestamp of when the repo was created
	CreatedAt sql.NullTime
	// the name of the default branch for the repo
	DefaultBranchName sql.NullString
	// the description for the repo
	Description sql.NullString
	// the number of kilobytes on disk for the repo
	Size sql.NullInt32
	// number of forks associated to the repo
	ForkCount sql.NullInt32
	// the GitHub homepage URL for the repo
	HomepageUrl sql.NullString
	// boolean to determine if the repo is archived
	IsArchived sql.NullBool
	// boolean to determine if the repo is disabled
	IsDisabled sql.NullBool
	// boolean to determine if the repo is private
	IsPrivate sql.NullBool
	// number of issues associated to the repo
	TotalIssuesCount sql.NullInt32
	// the author of the latest release in the repo
	LatestReleaseAuthor sql.NullString
	// timestamp of the creation of the latest release in the repo
	LatestReleaseCreatedAt sql.NullTime
	// the name of the latest release in the repo
	LatestReleaseName sql.NullString
	// timestamp of the publishing of the latest release in the repo
	LatestReleasePublishedAt sql.NullTime
	// the license key for the repo
	LicenseKey sql.NullString
	// the license name for the repo
	LicenseName sql.NullString
	// the primary language for the repo
	PrimaryLanguage sql.NullString
	// timestamp of the latest push to the repo
	PushedAt sql.NullTime
	// number of releases associated to the repo
	ReleasesCount sql.NullInt32
	// number of stargazers associated to the repo
	StargazersCount sql.NullInt32
	// timestamp of the latest update to the repo
	UpdatedAt sql.NullTime
	// number of watchers associated to the repo
	WatchersCount sql.NullInt32
	// timestamp when record was synced into the MergeStat database
	MergestatSyncedAt time.Time
	// advanced security availability
	AdvancedSecurity sql.NullString
	// secret scanning availability
	SecretScanning sql.NullString
	// secret scanning push protection availability
	SecretScanningPushProtection sql.NullString
	MirrorUrl                    sql.NullString
}

// stargazers of a GitHub repo
type GithubStargazer struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// login of the user who starred the repo
	Login string
	// email of the user who starred the repo
	Email sql.NullString
	// name of the user who starred the repo
	Name sql.NullString
	// public profile bio of the user who starred the repo
	Bio sql.NullString
	// company of the user who starred the repo
	Company sql.NullString
	// URL to the avatar of the user who starred the repo
	AvatarUrl sql.NullString
	// timestamp of when the user was created who starred the repo
	CreatedAt sql.NullTime
	// timestamp of the latest update to the user who starred the repo
	UpdatedAt sql.NullTime
	// twitter of the user who starred the repo
	Twitter sql.NullString
	// website of the user who starred the repo
	Website sql.NullString
	// location of the user who starred the repo
	Location sql.NullString
	// timestamp the user starred the repo
	StarredAt sql.NullTime
	// timestamp when record was synced into the MergeStat database
	MergestatSyncedAt time.Time
}

type GitleaksRepoDetection struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// description of the detection
	Description interface{}
	// detection start line
	StartLine interface{}
	// detection end line
	EndLine interface{}
	// detection start column
	StartColumn interface{}
	// detection end column
	EndColumn interface{}
	// detection match
	Match interface{}
	// detection secret
	Secret interface{}
	// detection file
	File interface{}
	// detected symlink file
	SymlinkFile interface{}
	// detection commit
	Commit interface{}
	// detection entropy
	Entropy interface{}
	// detection author
	Author interface{}
	// detection email
	Email interface{}
	// detection date
	Date interface{}
	// detection message
	Message interface{}
	// detection tags
	Tags interface{}
	// detection rule id
	RuleID interface{}
	// detection fingerprint
	Fingerprint interface{}
}

// scan output of a Gitleaks repo scan
type GitleaksRepoScan struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// JSON output of a Gitleaks scan
	Results pgtype.JSONB
}

type GosecRepoDetection struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// detection severity
	Severity interface{}
	// detection confidence
	Confidence interface{}
	// detection CWE (Common Weakness Enumeration) ID
	CweID interface{}
	// detection rule ID
	RuleID interface{}
	// detection details
	Details interface{}
	// detection file
	File interface{}
	// detection line in file
	Line interface{}
	// detection column in line
	Column interface{}
	// flag to determine if #nosec annotation was used
	Nosec interface{}
}

// Table of gosec repo scans
type GosecRepoScan struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// JSON issues from gosec repo scan
	Issues pgtype.JSONB
	// timestamp when record was synced into the MergeStat database
	MergestatSyncedAt time.Time
}

// Table for Grype repo scan results
type GrypeRepoScan struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// JSON results of Grype repo scanner
	Results pgtype.JSONB
	// timestamp when record was synced into the MergeStat database
	MergestatSyncedAt time.Time
}

type GrypeRepoVulnerability struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// id of the current vulnerability
	ID interface{}
	// level of severity
	Severity interface{}
	// description  of vulnerability
	Description interface{}
	// current version of package vulnerable
	Version interface{}
	// type of vulnerability
	Type interface{}
	// programming language associated to vulnerability
	Language interface{}
	// path to file of current vulnerability
	Path interface{}
}

type MergestatContainerImage struct {
	ID         uuid.UUID
	Name       string
	Type       string
	Url        string
	Version    string
	Parameters pgtype.JSONB
}

type MergestatContainerImageType struct {
	Name        string
	Displayname string
	Description sql.NullString
}

type MergestatContainerSync struct {
	ID         uuid.UUID
	RepoID     uuid.UUID
	ImageID    uuid.UUID
	Parameters pgtype.JSONB
}

type MergestatContainerSyncExecution struct {
	SyncID    uuid.UUID
	JobID     uuid.UUID
	CreatedAt sql.NullTime
}

type MergestatContainerSyncSchedule struct {
	ID        uuid.UUID
	SyncID    uuid.UUID
	CreatedAt sql.NullTime
}

type MergestatLatestRepoSync struct {
	ID         int64
	CreatedAt  time.Time
	RepoSyncID uuid.UUID
	Status     string
	StartedAt  sql.NullTime
	DoneAt     sql.NullTime
}

type MergestatProvider struct {
	ID          uuid.UUID
	Name        string
	Vendor      string
	Settings    pgtype.JSONB
	CreatedAt   time.Time
	Description sql.NullString
}

type MergestatQueryHistory struct {
	ID    uuid.UUID
	RunAt sql.NullTime
	RunBy string
	Query string
}

// Table for "dynamic" repo imports - regularly loading from a GitHub org for example
type MergestatRepoImport struct {
	ID                  uuid.UUID
	CreatedAt           time.Time
	UpdatedAt           time.Time
	Settings            pgtype.JSONB
	LastImport          sql.NullTime
	ImportInterval      sql.NullString
	LastImportStartedAt sql.NullTime
	ImportStatus        sql.NullString
	ImportError         sql.NullString
	Provider            uuid.UUID
}

// Types of repo imports
type MergestatRepoImportType struct {
	Type        string
	Description string
}

type MergestatRepoSync struct {
	RepoID                       uuid.UUID
	SyncType                     string
	Settings                     pgtype.JSONB
	ID                           uuid.UUID
	ScheduleEnabled              bool
	Priority                     int32
	LastCompletedRepoSyncQueueID sql.NullInt64
}

type MergestatRepoSyncLog struct {
	ID              int64
	CreatedAt       time.Time
	LogType         string
	Message         string
	RepoSyncQueueID int64
}

type MergestatRepoSyncLogType struct {
	Type        string
	Description sql.NullString
}

type MergestatRepoSyncQueue struct {
	ID            int64
	CreatedAt     time.Time
	RepoSyncID    uuid.UUID
	Status        string
	StartedAt     sql.NullTime
	DoneAt        sql.NullTime
	LastKeepAlive sql.NullTime
	Priority      int32
	TypeGroup     string
}

type MergestatRepoSyncQueueStatusType struct {
	Type        string
	Description sql.NullString
}

type MergestatRepoSyncType struct {
	Type        string
	Description sql.NullString
	ShortName   string
	Priority    int32
	TypeGroup   string
}

type MergestatRepoSyncTypeGroup struct {
	Group           sql.NullString
	ConcurrentSyncs sql.NullInt32
}

// @name labels
type MergestatRepoSyncTypeLabel struct {
	Label       string
	Description sql.NullString
	Color       string
}

// @name labelAssociations
type MergestatRepoSyncTypeLabelAssociation struct {
	Label        string
	RepoSyncType string
}

// Table to save queries
type MergestatSavedQuery struct {
	ID uuid.UUID
	// query creator
	CreatedBy sql.NullString
	// timestamp when query was created
	CreatedAt sql.NullTime
	// query name
	Name string
	// query description
	Description sql.NullString
	// query sql
	Sql string
	// query metadata
	Metadata pgtype.JSONB
}

type MergestatSchemaIntrospection struct {
	Schema            interface{}
	TableName         interface{}
	TableType         interface{}
	ColumnName        interface{}
	OrdinalPosition   interface{}
	IsNullable        interface{}
	DataType          interface{}
	UdtName           interface{}
	ColumnDescription string
}

type MergestatServiceAuthCredential struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Type        string
	Credentials []byte
	Provider    uuid.UUID
	IsDefault   sql.NullBool
	Username    []byte
}

type MergestatServiceAuthCredentialType struct {
	Type        string
	Description string
}

type MergestatUserMgmtPgUser struct {
	Rolname        interface{}
	Rolsuper       sql.NullBool
	Rolinherit     sql.NullBool
	Rolcreaterole  sql.NullBool
	Rolcreatedb    sql.NullBool
	Rolcanlogin    sql.NullBool
	Rolconnlimit   sql.NullInt32
	Rolvaliduntil  sql.NullTime
	Rolreplication sql.NullBool
	Rolbypassrls   sql.NullBool
	Memberof       interface{}
}

type MergestatVendor struct {
	Name        string
	DisplayName string
	Description sql.NullString
	Type        string
}

type MergestatVendorType struct {
	Name        string
	DisplayName string
	Description sql.NullString
}

type OssfScorecardRepoCheckResult struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// name of the check in the scan run
	Name interface{}
	// score of the check in the scan run
	Score interface{}
	// reason for the score of the check in the scan run
	Reason interface{}
	// details of the check in the scan run
	Details interface{}
	// URL to the documentation of the check in the scan run
	DocumentationUrl interface{}
	// description of the check in the scan run
	DocumentationShort interface{}
	// JSON of the check results in the scan run
	Value pgtype.JSONB
}

// Output of OSSF scorecard scans on a git repository
type OssfScorecardRepoScan struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// JSON results of the scan
	Results pgtype.JSONB
	// timestamp when record was synced into the MergeStat database
	MergestatSyncedAt time.Time
}

type OssfScorecardRepoScore struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// URL to repo scan was run on
	Repo interface{}
	// commit of repo that the scan was run on
	Commit interface{}
	// date of the scan
	Date interface{}
	// resulting total score of the scan
	Score interface{}
	// scorecard version used to run the scan
	ScorecardVersion interface{}
}

// git repositories to track
type Repo struct {
	// MergeStat identifier for the repo
	ID uuid.UUID
	// URL for the repo
	Repo string
	// ref for the repo
	Ref sql.NullString
	// timestamp of when the MergeStat repo entry was created
	CreatedAt time.Time
	// JSON settings for the repo
	Settings pgtype.JSONB
	// array of tags for the repo for topics in GitHub as well as tags added in MergeStat
	Tags pgtype.JSONB
	// foreign key for mergestat.repo_imports.id
	RepoImportID uuid.NullUUID
	Provider     uuid.UUID
}

// MergeStat internal table to track schema migrations
type SchemaMigration struct {
	Version int64
	Dirty   bool
}

// MergeStat internal table to track schema migrations history
type SchemaMigrationsHistory struct {
	ID        int32
	Version   int64
	AppliedAt time.Time
}

type SyftRepoArtifact struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// artifact JSON results from Syft repo scan
	Artifact pgtype.JSONB
	// artifact id
	ID interface{}
	// artifact name
	Name interface{}
	// artifact version
	Version interface{}
	// artifact type
	Type interface{}
	// artifact found_by
	FoundBy interface{}
	// artifact locations
	Locations interface{}
	// artifact licenses
	Licenses interface{}
	// artifact language
	Language interface{}
	// artifact cpes
	Cpes interface{}
	// artifact purl
	Purl interface{}
}

// Syft repo scans
type SyftRepoScan struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// JSON results from Syft repo scan
	Results pgtype.JSONB
	// timestamp when record was synced into the MergeStat database
	MergestatSyncedAt time.Time
}

// Trivy repo scans
type TrivyRepoScan struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// JSON results from Trivy repo scan
	Results pgtype.JSONB
	// timestamp when record was synced into the MergeStat database
	MergestatSyncedAt time.Time
}

type TrivyRepoVulnerability struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// vulnerability JSON results of Trivy scan
	Vulnerability pgtype.JSONB
	// vulnerability target
	Target interface{}
	// vulnerability class
	Class interface{}
	// vulnerability type
	Type interface{}
	// vulnerability id
	VulnerabilityID interface{}
	// vulnerability package name
	VulnerabilityPkgName interface{}
	// vulnerability installed version
	VulnerabilityInstalledVersion interface{}
	// vulnerability severity
	VulnerabilitySeverity interface{}
	// vulnerability title
	VulnerabilityTitle interface{}
	// vulnerability description
	VulnerabilityDescription interface{}
	// timestamp when record was synced into the MergeStat database
	MergestatSyncedAt time.Time
}

type YelpDetectSecretsRepoDetection struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// detection type
	Type interface{}
	// detection filename
	Filename interface{}
	// detection is verified
	IsVerified interface{}
	// detection line number
	LineNumber interface{}
	// detection secret
	HashedSecret interface{}
	// detection version
	Version interface{}
	// detection generated at
	GeneratedAt interface{}
	// detection filters used
	FiltersUsed interface{}
	// detection plugins used
	PluginsUsed interface{}
}

// scan output of a Yelp detect-secrets repo scan
type YelpDetectSecretsRepoScan struct {
	// foreign key for public.repos.id
	RepoID uuid.UUID
	// JSON output of a Yelp detect-secrets scan
	Results pgtype.JSONB
}
