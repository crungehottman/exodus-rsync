package rsync

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/release-engineering/exodus-rsync/internal/args"
	"github.com/release-engineering/exodus-rsync/internal/log"
)

func addTestBinPath(t *testing.T) {
	oldPath := os.Getenv("PATH")
	newPath := "../../test/bin:" + oldPath

	t.Cleanup(func() {
		os.Setenv("PATH", oldPath)
	})
	err := os.Setenv("PATH", newPath)
	if err != nil {
		t.Fatalf("could not set PATH, err = %v", err)
	}
}

func TestRawExec(t *testing.T) {
	addTestBinPath(t)

	argv := []string{"some-src", "some-dest"}
	expectedArgv := []string{"../../test/bin/rsync", "some-src", "some-dest"}

	var gotArgv0 string
	var gotArgv []string

	oldExec := ext.exec
	ext.exec = func(argv0 string, argv, _ []string) error {
		gotArgv0 = argv0
		gotArgv = argv
		return fmt.Errorf("simulated error")
	}
	defer func() { ext.exec = oldExec }()

	ctx := context.Background()
	ctx = log.NewContext(ctx, log.Package.NewLogger(args.Config{}))

	err := Package.RawExec(ctx, argv)
	if err.Error() != "simulated error" {
		t.Error("error not propagated from exec, got =", err)
	}

	if gotArgv0 != "../../test/bin/rsync" {
		t.Error("invoked unexpected rsync command", gotArgv0)
	}

	if !reflect.DeepEqual(gotArgv, expectedArgv) {
		t.Error("rsync invoked with wrong arguments", gotArgv)
	}
}

func TestExec(t *testing.T) {
	addTestBinPath(t)

	tests := []struct {
		name         string
		args         args.Config
		expectedArgv []string
	}{
		{"basic",
			args.Config{
				Src:  "some-src",
				Dest: "some-dest",
			},
			[]string{"../../test/bin/rsync", "some-src", "some-dest"},
		},

		{"all args",
			args.Config{
				Src:     "src",
				Dest:    "dest",
				Verbose: 3,
				IgnoredConfig: args.IgnoredConfig{
					Archive:        true,
					Recursive:      true,
					Links:          true,
					CopyLinks:      true,
					KeepDirlinks:   true,
					HardLinks:      true,
					Perms:          true,
					Executability:  true,
					Acls:           true,
					Xattrs:         true,
					Owner:          true,
					Group:          true,
					Devices:        true,
					Specials:       true,
					Times:          true,
					Atimes:         true,
					Crtimes:        true,
					OmitDirTimes:   true,
					Rsh:            "some-rsh",
					Delete:         true,
					PruneEmptyDirs: true,
					Timeout:        1234,
					Compress:       true,
					Stats:          true,
					ItemizeChanges: true,
				},
				Relative:       true,
				IgnoreExisting: true,
				Filter:         []string{"some-filter"},
				Exclude:        []string{".*"},
				Include:        []string{"**/dir"},
				FilesFrom:      "sources.txt",
			},
			[]string{
				"../../test/bin/rsync", "-vvv",
				"--archive", "--recursive", "--relative", "--links", "--copy-links",
				"--keep-dirlinks", "--hard-links", "--perms", "--executability", "--acls",
				"--xattrs", "--owner", "--group", "--devices", "--specials", "--times",
				"--atimes", "--crtimes", "--omit-dir-times", "--rsh", "some-rsh",
				"--ignore-existing", "--delete", "--prune-empty-dirs", "--timeout", "1234",
				"--compress", "--filter", "some-filter", "--exclude", ".*", "--include", "**/dir",
				"--files-from", "sources.txt", "--stats", "--itemize-changes",
				"src", "dest",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotArgv0 string
			var gotArgv []string

			oldExec := ext.exec
			ext.exec = func(argv0 string, argv, _ []string) error {
				gotArgv0 = argv0
				gotArgv = argv
				return fmt.Errorf("simulated error")
			}
			defer func() { ext.exec = oldExec }()

			ctx := context.Background()
			ctx = log.NewContext(ctx, log.Package.NewLogger(tt.args))

			err := Package.Exec(ctx, tt.args)
			if err.Error() != "simulated error" {
				t.Error("error not propagated from exec, got =", err)
			}

			if gotArgv0 != "../../test/bin/rsync" {
				t.Error("invoked unexpected rsync command", gotArgv0)
			}

			if !reflect.DeepEqual(gotArgv, tt.expectedArgv) {
				t.Error("rsync invoked with wrong arguments", gotArgv)
			}
		})
	}
}
