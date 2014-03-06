package routes

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/codegangsta/martini"
	"github.com/coopernurse/gorp"
	"github.com/zachlatta/southbayfession/misc"
	"github.com/zachlatta/southbayfession/models"
)

type School struct {
	Id     int            `json:"id"`
	Name   string         `json:"name"`
	Tweets []models.Tweet `json:"tweets,omitempty"`
}

func GetSchools(enc Encoder, db gorp.SqlExecutor) (int, string) {
	schools := make([]School, len(misc.Schools))

	schoolNames := make([]string, len(misc.Schools))
	i := 0
	for k, _ := range misc.Schools {
		schoolNames[i] = k
		i++
	}
	sort.Strings(schoolNames)

	for i, name := range schoolNames {
		schools[i].Id = i
		schools[i].Name = name
	}

	return http.StatusOK, Must(enc.EncodeOne(schools))
}

func GetSchool(enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
	var school School
	id, err := strconv.Atoi(parms["id"])
	if err != nil {
		return http.StatusConflict, ""
	}

	schoolNames := make([]string, len(misc.Schools))
	i := 0
	for k, _ := range misc.Schools {
		schoolNames[i] = k
		i++
	}
	sort.Strings(schoolNames)

	school.Id = id
	school.Name = schoolNames[id]

	_, err = db.Select(&school.Tweets, fmt.Sprintf("select * from tweets where `school` = '%s' order by id desc", school.Name))
	if err != nil {
		return http.StatusConflict, ""
	}

	return http.StatusOK, Must(enc.EncodeOne(school))
}
