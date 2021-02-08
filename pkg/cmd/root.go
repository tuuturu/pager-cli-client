package cmd

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/spf13/cobra"
	"github.com/tuuturu/event-client/pkg/core"
	"github.com/tuuturu/event-client/pkg/oauth2"
	"github.com/tuuturu/event-client/pkg/pager"
	"github.com/tuuturu/pager-event-service/pkg/core/models"
	"os"
)

type Arguments struct {
	models.Event
}

func (a Arguments) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Title, validation.Required),
		validation.Field(&a.Timestamp, is.UTFNumeric),
		validation.Field(&a.ReadMoreURL, is.URL),
		validation.Field(&a.ImageURL, is.URL),
	)
}

var (
	arguments = Arguments{}
	cfg       core.Config
)

var rootCmd = &cobra.Command{
	Use:   "pager",
	Short: "pager sends notifications to a Pager service",
	Long:  `A simple CLI tool to send notifications to a Pager service`,
	Args:  cobra.ExactArgs(2),
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		cfg, err = core.LoadConfig()
		if err != nil {
			return err
		}

		arguments.Title = args[0]
		arguments.Description = args[1]

		if err = arguments.Validate(); err != nil {
			return err
		}

		return cfg.Validate()
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		token, err := oauth2.AcquireToken(cfg.DiscoveryURL, cfg.ClientID, cfg.ClientSecret)
		if err != nil {
			return err
		}

		err = pager.CreateEvent(cfg.EventsServiceURL, token, models.Event{
			Timestamp:   arguments.Timestamp,
			Read:        arguments.Read,
			Title:       arguments.Title,
			Description: arguments.Description,
			ImageURL:    arguments.ImageURL,
			ReadMoreURL: arguments.ReadMoreURL,
		})
		if err != nil {
			return err
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	flags := rootCmd.Flags()

	flags.StringVarP(&arguments.Timestamp, "timestamp", "t", "", "sets the Unix Nano timestamp for the event")
	flags.StringVarP(&arguments.ImageURL, "image-url", "i", "", "sets the icon URL for the event")
	flags.StringVarP(&arguments.ReadMoreURL, "read-more-url", "u", "", "sets the read more URL for the event")
	flags.BoolVarP(&arguments.Read, "read", "r", false, "sets the read/unread flag for the event")
}
