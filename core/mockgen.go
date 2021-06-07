package internal

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozoncp/ocp-progress-api/core/repo Repo
//go:generate mockgen -destination=./mocks/saver_mock.go -package=mocks github.com/ozoncp/ocp-progress-api/core/saver Saver
//go:generate mockgen -destination=./mocks/alarmer_mock.go -package=mocks github.com/ozoncp/ocp-progress-api/core/alarmer Alarmer
