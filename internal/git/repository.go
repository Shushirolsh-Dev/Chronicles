package git

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/Shushirolsh-Dev/Chronicles/internal/journal"
)

type Repository struct { path, name string }

func NewRepository(path string) (*Repository, error) {
	if strings.TrimSpace(path) == "" { path = "." }
	absolute, err := filepath.Abs(path)
	if err != nil { return nil, fmt.Errorf("resolve repository path: %w", err) }
	info, err := os.Stat(absolute)
	if err != nil { return nil, fmt.Errorf("access repository path %q: %w", path, err) }
	if !info.IsDir() { return nil, fmt.Errorf("repository path %q is not a directory", path) }

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "git", "-C", absolute, "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil { return nil, fmt.Errorf("%q is not a Git repository (or Git is unavailable)", path) }
	root := strings.TrimSpace(string(output))
	return &Repository{path: root, name: filepath.Base(root)}, nil
}

func (r *Repository) Name() string { return r.name }

func (r *Repository) Log(days int, author string) ([]journal.Entry, error) {
	if days < 1 { return nil, fmt.Errorf("--days must be greater than zero") }
	args := []string{
		"-C", r.path, "log", "--since=" + fmt.Sprintf("%d days ago", days),
		"--date=iso-strict",
		"--pretty=format:%H%x1f%an%x1f%ae%x1f%aI%x1f%s%x1f%b%x1e",
	}
	if strings.TrimSpace(author) != "" { args = append(args, "--author="+author) }

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "git", args...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	output, err := cmd.Output()
	if err != nil {
		detail := strings.TrimSpace(stderr.String())
		if detail == "" { return nil, fmt.Errorf("read Git history: %w", err) }
		return nil, fmt.Errorf("read Git history: %s", detail)
	}
	entries, err := journal.ParseLog(string(output))
	if err != nil { return nil, fmt.Errorf("parse Git history: %w", err) }
	return entries, nil
}
