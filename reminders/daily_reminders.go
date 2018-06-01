package reminders

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	eorzeaConstant float64 = 20.571428571428573

	dailyReset      = "15:00 -0000 GMT" // 15:00 GMT
	weeklyReset     = "09:00 -0000 GMT" // 9:00 GMT
	notificationFmt = "15:04 -0700 MST"
)

type reminderMeta struct {
	channelID string
	guildID   string
}

var (
	l *log.Logger

	reminderChannels = make([]*reminderMeta, 0)
	// A singleton reference to our session to send out reminders
	session *discordgo.Session
)

func getEorzeaTime() time.Time {
	epoch := time.Now().Unix()
	eorzeaTime := time.Unix(int64(float64(epoch)*eorzeaConstant), 0)
	return eorzeaTime.UTC()
}

// RegisterReminders registers a channel ID to get daily reset reminders
// These should be registered on deserialization of guild
func RegisterReminders(channelID, guildID string) error {
	if _, err := receivingReminders(reminderChannels, channelID, guildID); err != nil {
		return err
	}
	meta := &reminderMeta{
		channelID: channelID,
		guildID:   guildID,
	}
	reminderChannels = append(reminderChannels, meta)
	l.Printf("Registered reminders to %s", guildID)
	return nil
}

// DeregisterReminders deregisters reminders here. Make sure the calling function
// is one that can edit the FC struct to actually remove in case of bot reboot/persistence
func DeregisterReminders(guildID string) error {
	for idx, meta := range reminderChannels {
		if meta.guildID == guildID {
			reminderChannels = append(reminderChannels[:idx], reminderChannels[idx+1:]...)
			l.Printf("Deregistered reminders to %s", guildID)
			return nil
		}
	}
	return fmt.Errorf("guild is not receiving notifications")
}

func receivingReminders(slice []*reminderMeta, channelID, guildID string) (bool, error) {
	for _, meta := range slice {
		if strings.EqualFold(meta.channelID, channelID) {
			return true, fmt.Errorf("channel already receiving notifications")
		}
		if strings.EqualFold(meta.guildID, guildID) {
			return true, fmt.Errorf("guild already receiving notifications")
		}
	}
	return false, nil
}

// RegisterSession registers the singleton session reference
func RegisterSession(s *discordgo.Session) {
	if session == nil {
		session = s
	}
}

func init() {
	l = log.New(os.Stderr, "reminders: ", log.LstdFlags|log.Lshortfile)
	// Startup the daily/weekly timer functions
	go func() {
		dailySleep := durationUntilTimer(dailyReset)
		l.Printf("Sleeping %s until daily notifications go out\n", dailySleep)
		time.Sleep(dailySleep)
		notifTicker := time.NewTicker(24 * time.Hour)
		msg := []string{
			"**Dailies have reset!**",
			" - Beastmen quests (12)",
			" - Duty roulette",
			" - Daily repeatable quests",
			" - Squadron training (3)",
		}
		// First time at start doesn't tick...
		for _, meta := range reminderChannels {
			err := doNotification(meta.channelID, strings.Join(msg, "\n"))
			if err != nil {
				l.Println(err)
			}
			time.Sleep(500 * time.Millisecond)
		}
		for range notifTicker.C {
			for _, meta := range reminderChannels {
				err := doNotification(meta.channelID, strings.Join(msg, "\n"))
				if err != nil {
					l.Println(err)
				}
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	go func() {
		weeklySleep := durationUntilWeekly(weeklyReset, time.Tuesday)
		l.Printf("Sleeping %s until weekly notifications go out (%f days)\n", weeklySleep, weeklySleep.Hours()/24)
		time.Sleep(weeklySleep)
		notifTicker := time.NewTicker(24 * time.Hour * 7)
		msg := []string{
			"**Weeklies have reset!**",
			" - Tomestone currency",
			" - Weekly repeatable quests",
			" - Challenge Log",
			" - New Wondrous Tails",
		}
		// First time at start doesn't tick...
		for _, meta := range reminderChannels {
			err := doNotification(meta.channelID, strings.Join(msg, "\n"))
			if err != nil {
				l.Println(err)
			}
			time.Sleep(500 * time.Millisecond)
		}
		for range notifTicker.C {
			for _, meta := range reminderChannels {
				err := doNotification(meta.channelID, strings.Join(msg, "\n"))
				if err != nil {
					l.Println(err)
				}
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
}

func doNotification(channelID, msg string) error {
	if session == nil {
		return fmt.Errorf("singleton session not registered")
	}
	if _, err := session.ChannelMessageSend(channelID, msg); err != nil {
		return err
	}
	return nil
}

func durationUntilTimer(rawTimer string) time.Duration {
	timer, _ := time.Parse(notificationFmt, rawTimer)
	today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), timer.Hour(), timer.Minute(), 0, 0, timer.Location())
	diff := today.Sub(time.Now())
	if diff < 0 {
		return today.Add(24 * time.Hour).Sub(time.Now())
	}
	return diff
}

func durationUntilWeekly(rawTimer string, resetWeekday time.Weekday) time.Duration {
	timer, _ := time.Parse(notificationFmt, rawTimer)
	today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), timer.Hour(), timer.Minute(), 0, 0, timer.Location())
	if today.Weekday() != resetWeekday {
		diff := resetWeekday - today.Weekday()
		if diff < 0 {
			diff = 7 + diff
		}
		today = today.Add(time.Hour * 24 * time.Duration(diff))
	}
	diff := today.Sub(time.Now())
	if diff < 0 {
		return today.Add(24 * time.Hour * 7).Sub(time.Now())
	}
	return diff
}
