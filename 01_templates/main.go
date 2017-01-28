package main

import "text/template"
import "net/http"

var (
	tmpFunc = template.FuncMap{
		"active":   activeOnly,
		"finished": finishedOnly,
		"duration": countDuration,
	}
	tpl = template.Must(template.New("").Funcs(tmpFunc).ParseGlob("gohtml/*"))
)

type UserData struct {
	Name     string
	UserName string
}

type Task struct {
	Name   string
	Status string
	Events []Event
}

type Event struct {
	Duration int64
	Type     string
}

func main() {

	server := http.NewServeMux()
	server.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {

		data := struct {
			User  UserData
			Tasks []Task
		}{
			UserData{
				"Radomir Sohlich",
				"sola_cz",
			},
			[]Task{
				Task{
					Name:   "Working",
					Status: "running",
				},
				Task{
					Name: "Reading",
				},
			},
		}

		tpl.ExecuteTemplate(rw, "Index", data)
	})
	http.ListenAndServe(":8080", server)

}

func activeOnly(tasks []Task) Task {
	for _, val := range tasks {
		if val.Status == "running" {
			return val
		}
	}

	return Task{}
}

func finishedOnly(tasks []Task) []Task {
	filtered := []Task{}
	for _, val := range tasks {
		if val.Status != "running" {
			filtered = append(filtered, val)
		}
	}
	return filtered
}

func countDuration(t Task) int64 {
	total := int64(0)
	for _, val := range t.Events {
		total += val.Duration
	}
	return total
}
