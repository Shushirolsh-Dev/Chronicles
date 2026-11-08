package main

import (
	"fmt"
	"os"

	gitrepo "github.com/Shushirolsh-Dev/Chronicles/internal/git"
	"github.com/Shushirolsh-Dev/Chronicles/internal/journal"
	"github.com/Shushirolsh-Dev/Chronicles/internal/render"
	"github.com/spf13/cobra"
)

type options struct {
	days int
	path string
	author string
}

func main() {
	if err := newRootCommand().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "chronicle:", err)
		os.Exit(1)
	}
}

func newRootCommand() *cobra.Command {
	opts := &options{}
	root := &cobra.Command{
		Use: "chronicle",
		Short: "Turn your Git history into a development journal",
		Long: "Chronicles creates a readable development journal from commits in a local Git repository.",
		SilenceUsage: true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, _ []string) error { return runJournal(cmd, opts) },
	}
	addFlags(root, opts)
	root.AddCommand(newChangelogCommand(opts))
	return root
}

func newChangelogCommand(opts *options) *cobra.Command {
	cmd := &cobra.Command{
		Use: "changelog",
		Short: "Generate a Markdown changelog from Git history",
		RunE: func(cmd *cobra.Command, _ []string) error {
			entries, name, err := loadEntries(opts)
			if err != nil { return err }
			_, err = fmt.Fprint(cmd.OutOrStdout(), render.Markdown(name, entries, opts.days))
			return err
		},
	}
	addFlags(cmd, opts)
	return cmd
}

func addFlags(cmd *cobra.Command, opts *options) {
	cmd.Flags().IntVar(&opts.days, "days", 14, "Include commits from the last N days")
	cmd.Flags().StringVar(&opts.path, "path", ".", "Path to a local Git repository")
	cmd.Flags().StringVar(&opts.author, "author", "", "Filter by author name or email")
}

func loadEntries(opts *options) ([]journal.Entry, string, error) {
	repo, err := gitrepo.NewRepository(opts.path)
	if err != nil { return nil, "", err }
	entries, err := repo.Log(opts.days, opts.author)
	if err != nil { return nil, "", err }
	return entries, repo.Name(), nil
}

func runJournal(cmd *cobra.Command, opts *options) error {
	entries, name, err := loadEntries(opts)
	if err != nil { return err }
	_, err = fmt.Fprint(cmd.OutOrStdout(), render.Journal(name, entries, opts.days))
	return err
}
