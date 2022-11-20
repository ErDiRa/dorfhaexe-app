package convert

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"erdira.com/date-converter/convert/model"
	ics "github.com/arran4/golang-ical"
	"github.com/thatisuday/commando"
	"github.com/xuri/excelize/v2"
)

func Run(_ map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
	inputPath, _ := flags["input"].GetString()
	outputPath, _ := flags["output"].GetString()

	xls, err := excelize.OpenFile(inputPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := xls.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// get name of all sheets (we only use the first one)
	sheetList := xls.GetSheetList()

	narrenfahrplan := []model.Date{}
	kampagne := []model.Date{}
	aufAbbau := []model.Date{}

	for sheetIdx, sheet := range sheetList {
		// Get all the rows in the Sheet1.
		rows, err := xls.GetRows(sheetList[sheetIdx])
		if err != nil {
			fmt.Println(err)
			return
		}

		for idx, row := range rows {

			if idx == 0 {
				//skip header
				continue
			}

			switch sheet {
			case "Narrenfahrplan":
				date := model.Date{
					Date:         format(clean(row[0]), "02.01.2006"),
					Event:        clean(row[1]),
					Time:         clean(row[2]),
					MeetingPoint: clean(row[3]),
					Outfit:       clean(row[4]),
				}

				dateTime := date.Time
				if dateTime == "-" {
					dateTime = "12:00"
				}

				startTime, err := time.Parse("02.01.2006 15:04", fmt.Sprintf("%s %s", date.Date, dateTime))
				if err != nil {
					fmt.Println("could not parse start time", err, date.Event, sheet)
					return
				}
				endTime := startTime.Add(1 * time.Hour)

				cal := ics.NewCalendar()
				cal.SetMethod(ics.MethodRequest)
				event := cal.AddEvent(date.Event + "_" + date.Date)
				event.SetCreatedTime(time.Now())
				event.SetDtStampTime(time.Now())
				event.SetStartAt(startTime)
				event.SetEndAt(endTime)
				event.SetSummary(date.Event)
				event.SetLocation(date.MeetingPoint)
				event.SetOrganizer("Dorfhäxe Rümmingen 1975 e.V.")
				calendarEvent := cal.Serialize()

				date.ICSFile = calendarEvent
				date.ICSFileName = date.Event + "_" + date.Date + ".ics"

				narrenfahrplan = append(narrenfahrplan, date)

			case "Kampagne":
				date := model.Date{
					Date:  format(clean(row[0]), "02.01.2006"),
					Event: clean(row[1]),
					Time:  clean(row[2]),
				}
				dateTime := date.Time
				if dateTime == "-" {
					dateTime = "12:00"
				}

				startTime, err := time.Parse("02.01.2006 15:04", fmt.Sprintf("%s %s", date.Date, dateTime))
				if err != nil {
					fmt.Println("could not parse start time", err, date.Event, sheet)
					return
				}
				endTime := startTime.Add(1 * time.Hour)

				cal := ics.NewCalendar()
				cal.SetMethod(ics.MethodRequest)
				event := cal.AddEvent(date.Event + "_" + date.Date)
				event.SetCreatedTime(time.Now())
				event.SetDtStampTime(time.Now())
				event.SetStartAt(startTime)
				event.SetEndAt(endTime)
				event.SetSummary(date.Event)
				event.SetLocation(date.MeetingPoint)
				event.SetOrganizer("Dorfhäxe Rümmingen 1975 e.V.")
				calendarEvent := cal.Serialize()

				date.ICSFile = calendarEvent
				date.ICSFileName = date.Event + "_" + date.Date + ".ics"

				kampagne = append(kampagne, date)

			case "Auf-Abbau":
				date := model.Date{
					Date:         format(clean(row[0]), "02.01.2006"),
					Event:        clean(row[1]),
					Time:         clean(row[2]),
					MeetingPoint: clean(row[3]),
				}
				dateTime := date.Time
				if dateTime == "-" {
					dateTime = "12:00"
				}

				startTime, err := time.Parse("02.01.2006 15:04", fmt.Sprintf("%s %s", date.Date, dateTime))
				if err != nil {
					fmt.Println("could not parse start time", err, date.Event, sheet)
					return
				}
				endTime := startTime.Add(1 * time.Hour)

				cal := ics.NewCalendar()
				cal.SetMethod(ics.MethodRequest)
				event := cal.AddEvent(date.Event + "_" + date.Date)
				event.SetCreatedTime(time.Now())
				event.SetDtStampTime(time.Now())
				event.SetStartAt(startTime)
				event.SetEndAt(endTime)
				event.SetSummary(date.Event)
				event.SetLocation(date.MeetingPoint)
				event.SetOrganizer("Dorfhäxe Rümmingen 1975 e.V.")
				calendarEvent := cal.Serialize()

				date.ICSFile = calendarEvent
				date.ICSFileName = date.Event + "_" + date.Date + ".ics"
				aufAbbau = append(aufAbbau, date)
			}

		}
	}

	toJsons := [][]model.Date{narrenfahrplan, kampagne, aufAbbau}

	for idx, toJson := range toJsons {
		bytes, err := json.MarshalIndent(toJson, "", "    ")
		if err != nil {
			fmt.Println("failed to marshal data at idx", idx, err)
			return
		}

		err = os.WriteFile(fmt.Sprintf("%s/output-%d.json", outputPath, idx), bytes, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}

	}

	fmt.Println("saved files to", outputPath)

}

func clean(input string) string {
	output := input
	output = strings.Trim(output, " ")
	output = strings.ReplaceAll(output, "\n", "")
	output = strings.ReplaceAll(output, "Uhr", "")
	output = strings.Trim(output, " ")

	return output
}

func format(input, layout string) string {

	if input == "" || input == "-" {
		return input
	}

	timeVal, err := time.Parse("2006-01-02", input)
	if err != nil {
		fmt.Println("error formatting time: ", err, "input: ", input)
		return input
	}

	return timeVal.Format(layout)

}
