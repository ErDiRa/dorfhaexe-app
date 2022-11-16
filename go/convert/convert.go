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

	// Get all the rows in the Sheet1.
	rows, err := xls.GetRows(sheetList[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	narrenfahrplan := []model.Date{}
	kampagne := []model.Date{}
	aufAbbau := []model.Date{}

	for _, sheet := range sheetList {
		for idx, row := range rows {

			if idx == 0 {
				//skip header
				continue
			}

			switch sheet {
			case "Narrenfahrplan":
				date := model.Date{
					Date:         clean(row[0]),
					Event:        clean(row[1]),
					Time:         clean(row[2]),
					MeetingPoint: clean(row[3]),
					Outfit:       clean(row[4]),
				}

				dateTime := date.Time
				if dateTime == "-" {
					dateTime = "12:00"
				}

				fmt.Println(date.Event, date.Date, dateTime, sheet)

				startTime, err := time.Parse("2006-01-02 15:04", fmt.Sprintf("%s %s", date.Date, dateTime))
				if err != nil {
					fmt.Println("could not parse start time", err, date.Event, sheet)
					return
				}
				endTime := startTime.Add(1 * time.Hour)

				cal := ics.NewCalendar()
				cal.SetMethod(ics.MethodRequest)
				event := cal.AddEvent(date.Event + date.Date)
				event.SetCreatedTime(time.Now())
				event.SetDtStampTime(time.Now())
				event.SetStartAt(startTime)
				event.SetEndAt(endTime)
				event.SetSummary(date.Event)
				event.SetLocation(date.MeetingPoint)
				event.SetDescription("Description")
				event.SetOrganizer("Dorfhäxe Rümmingen 1975 e.V.", ics.WithCN("This Machine"))
				calendarEvent := cal.Serialize()

				filename := fmt.Sprintf("%s/%s-%s.ics", outputPath, strings.Trim(date.Event, " "), startTime.Format("2006-01-02_1506"))
				err = os.WriteFile(filename, []byte(calendarEvent), 0644)
				if err != nil {
					fmt.Println("coudl not save ics file for ", date.Event, err)
					return
				}
				fmt.Println("saved ics file to ", filename)

				narrenfahrplan = append(narrenfahrplan, date)

			case "Kampagne":
				date := model.Date{
					Date:  clean(row[0]),
					Event: clean(row[1]),
					Time:  clean(row[2]),
				}

				kampagne = append(kampagne, date)
			case "Auf-Abbau":
				date := model.Date{
					Date:         clean(row[0]),
					Event:        clean(row[1]),
					Time:         clean(row[2]),
					MeetingPoint: clean(row[3]),
				}

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
