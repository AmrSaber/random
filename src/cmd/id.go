package cmd

import (
	"fmt"
	"slices"
	"strings"

	"github.com/AmrSaber/random/v3/src/common"

	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/spf13/cobra"
)

// NewIDCommand wires the id subcommand.
func NewIDCommand(opts *common.RootOptions) *cobra.Command {
	var idType string

	cmd := &cobra.Command{
		Use:   "id",
		Short: "Prints a random ID",
		RunE: func(_ *cobra.Command, _ []string) error {
			if !slices.Contains(common.IDTypes, idType) {
				return fmt.Errorf("invalid id type %q", idType)
			}

			count := opts.Count
			ids := make([]string, 0, count)
			for range count {
				switch idType {
				case common.IDTypeUUID4:
					ids = append(ids, uuid.New().String())
				case common.IDTypeUUID7:
					u, err := uuid.NewV7()
					if err != nil {
						return fmt.Errorf("error generating id: %w", err)
					}
					ids = append(ids, u.String())
				case common.IDTypeNano:
					id, err := gonanoid.New()
					if err != nil {
						return fmt.Errorf("error generating id: %w", err)
					}
					ids = append(ids, id)
				}
			}

			common.Std.Println(strings.Join(ids, "\n"))
			return nil
		},
	}

	cmd.Flags().StringVarP(&idType, "type", "t", common.IDTypeUUID4, fmt.Sprintf("Type of the id, must be one of [%s]", strings.Join(common.IDTypes, ", ")))

	return cmd
}
