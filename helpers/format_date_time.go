package helpers

import (
	"fmt"
	"time"
)

/**
https://api.slack.com/reference/surfaces/formatting#date-formatting

{date_num} is displayed as 2014-02-18. It will include leading zeros before the month and date and is probably best for more technical integrations that require a developer-friendly date format.
{date} is displayed as February 18th, 2014. The year will be omitted if the date is less than six months in the past or future.
{date_short} is displayed as Feb 18, 2014. The year will be omitted if the date is less than six months in the past or future.
{date_long} is displayed as Tuesday, February 18th, 2014. The year will be omitted if the date is less than six months in the past or future.
{date_pretty} displays the same as {date} but uses "yesterday", "today", or "tomorrow" where appropriate.
{date_short_pretty} displays the same as {date_short} but uses "yesterday", "today", or "tomorrow" where appropriate.
{date_long_pretty} displays the same as {date_long} but uses "yesterday", "today", or "tomorrow" where appropriate.
{time} is displayed as 6:39 AM or 6:39 PM in 12-hour format. If the client is set to show 24-hour format, it is displayed as 06:39 or 18:39.
{time_secs} is displayed as 6:39:45 AM 6:39:42 PM in 12-hour format. In 24-hour format it is displayed as 06:39:45 or 18:39:42.
*/

func FormatDateTime(t *time.Time, format string) string {
	if t == nil {
		return ""
	}

	if format == "" {
		format = "{date_short_pretty}, {time}"
	}
	return fmt.Sprintf("<!date^%d^%s|%s>", t.Unix(), format, t.Format(time.RFC3339))
}
