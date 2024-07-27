package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

type Resume struct {
	Name             string
	Title            string
	Location         string
	Links            []Link
	Skills           []ScoredContent
	WorkExperiences  []WorkExperience
	EducationItems   []EducationItem
	OOSContributions []template.HTML
	Awards           []Award
	Languages        []ScoredContent
	Interests        []string
}

type Link struct {
	IconLocation string
	Content      string
	HREF         string
}

type ScoredContent struct {
	Content string
	Score   int
}

type WorkExperience struct {
	Company     string
	Roles       []Role
	Description []string
}

type Role struct {
	Title    string
	Duration string
}

type EducationItem struct {
	Course    string
	Institute string
	Location  string
	Score     string
	Duration  string
}

type Award struct {
	Title     string
	Institute string
	Time      string
	Details   string
}

func main() {
	if len(os.Args) == 2 && os.Args[1] == "server" {
		http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "style.css") })
		http.HandleFunc("/", RenderResumeHandler)
		http.ListenAndServe(":8080", nil)
	} else {
		if f, err := os.Create("index.html"); err != nil {
			log.Panic(err)
		} else {
			renderResume(f)
		}
	}
}

func RenderResumeHandler(w http.ResponseWriter, r *http.Request) {
	if err := renderResume(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderResume(w io.Writer) error {
	resumeTemplPath := "index.html.tmpl"
	resume := Resume{
		Name:     "Nilesh Kevlani",
		Title:    "Software Engineer",
		Location: "Bangalore, India",
		Links: []Link{
			{IconLocation: "", Content: "njkevlani.dev", HREF: "https://njkevlani.dev"},
			{IconLocation: "", Content: "njkevlani@gmail.com", HREF: "mailto://njkevlani@gmail.com"},
		},
		Skills: []ScoredContent{
			{Content: "algorithms", Score: 8},
			{Content: "protobuf", Score: 9},
		},
		WorkExperiences: []WorkExperience{
			{
				Company: "ShareChat",
				Roles: []Role{
					{Title: "Software Devlopment Engineer 3", Duration: "2024 March - Present"},
					{Title: "Software Devlopment Engineer 2", Duration: "September 2021 - 2024 March"},
				},
				Description: []string{
					"Played primary role in developing end-to-end systems for data warehousing and streaming use cases.",
					"Played primary role in planning, developing toolding and moving data from JSON to Protobuf format.",
				},
			},
			{
				Company: "Razorpay",
				Roles: []Role{
					{Title: "Software Devlopment Engineer", Duration: "June 2020 - September"},
				},
				Description: []string{
					"lorem",
					"ipsum",
				},
			},
		},
		EducationItems: []EducationItem{
			{
				Course:    "BE Computer Engineering",
				Institute: "BVM Engineering College",
				Location:  "Anand",
				Score:     "8.27 CGPA",
				Duration:  "2014 - 2018",
			},
			{
				Course:    "HSE",
				Institute: "BM Commerse High School",
				Location:  "Bhavnagar",
				Score:     "84.66 %",
				Duration:  "2012 - 2014",
			},
			{
				Course:    "SSC",
				Institute: "Vishudhanand Vidya Mandir",
				Location:  "Bhavnagar",
				Score:     "85.66 %",
				Duration:  "2010 - 2012",
			},
		},
		OOSContributions: []template.HTML{
			template.HTML("<a href=\"https://github.com/njkevlani/hst\">hst</a> - A simple <a href=\"https://gohugo.io\">hugo</a> theme"),
		},
		Awards: []Award{
			{Title: "Smart India Hackathon", Institute: "Governmemt Of India", Time: "March 2018", Details: "Participated in the final round of SIH 2018."},
			{Title: "Best Final Year Project", Institute: "BVM Engineering College", Time: "November 2017", Details: "Selected as the best project among 30 projects at the college level."},
		},
		Languages: []ScoredContent{
			{Content: "English", Score: 8},
			{Content: "Hindi", Score: 8},
			{Content: "Gujarati", Score: 8},
			{Content: "Sindhi", Score: 6},
		},
		Interests: []string{
			"Travelling",
			"Badminton",
			"Swimming",
			"Listening to Music",
		},
	}

	tmpl, err := template.ParseFiles(resumeTemplPath)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(w, resume); err != nil {
		return err
	}

	return nil
}
