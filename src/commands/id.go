package commands

import (
	"fmt"
	"random/src/common"
	"slices"
	"strings"

	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/urfave/cli/v2"
)

var IdCommand = &cli.Command{
	Name:        "id",
	Description: "Prints a random ID",

	Flags: []cli.Flag{
		common.CountFlag,

		&cli.StringFlag{
			Name:    "type",
			Aliases: []string{"t"},
			Usage:   fmt.Sprintf("Type of the id, must be one of [%s]", strings.Join(common.ID_TYPES, ", ")),
			Value:   common.ID_TYPE_UUID4,
			Action: func(ctx *cli.Context, s string) error {
				if !slices.Contains(common.ID_TYPES, s) {
					return fmt.Errorf("invalid id type %q", s)
				}

				return nil
			},
		},
	},

	Action: func(ctx *cli.Context) error {
		count := ctx.Int("count")
		idType := ctx.String("type")

		ids := make([]string, 0, count)

		for range count {
			var id string
			var err error

			if idType == common.ID_TYPE_UUID4 {
				id = uuid.New().String()
			}

			if idType == common.ID_TYPE_UUID7 {
				var _id uuid.UUID
				_id, err = uuid.NewV7()

				id = _id.String()

			}

			if idType == common.ID_TYPE_NANO {
				id, err = gonanoid.New()
			}

			if err != nil {
				return fmt.Errorf("error generating nanoid: %w", err)
			}

			ids = append(ids, id)
		}

		// Print the generated random ids
		common.Std.Println(strings.Join(ids, "\n"))

		return nil
	},
}
