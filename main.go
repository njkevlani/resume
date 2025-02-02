package main

import (
	"html/template"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Resume struct {
	Name     string `yaml:"name"`
	Title    string `yaml:"title"`
	Location string `yaml:"location"`
	Links    []struct {
		IconLocation string `yaml:"icon_location"`
		Content      string `yaml:"content"`
		HREF         string `yaml:"href"`
	} `yaml:"links"`
	Skills []struct {
		Content string `yaml:"content"`
		Score   int    `yaml:"score"`
	} `yaml:"skills"`
	WorkExperiences []struct {
		Company string `yaml:"company"`
		Roles   []struct {
			Title    string `yaml:"title"`
			Duration string `yaml:"duration"`
		} `yaml:"roles"`
		Description []string `yaml:"description"`
	} `yaml:"work_experiences"`
	EducationItems []struct {
		Course    string `yaml:"course"`
		Institute string `yaml:"institute"`
		Location  string `yaml:"location"`
		Score     string `yaml:"score"`
		Duration  string `yaml:"duration"`
	} `yaml:"education_items"`
	OOSContributions []template.HTML `yaml:"oos_contributions"`
	Languages        []struct {
		Content string `yaml:"content"`
		Score   int    `yaml:"score"`
	} `yaml:"languages"`
	Interests []string `yaml:"interests"`
}

func main() {
	if f, err := os.Create("index.html"); err != nil {
		log.Panic(err)
	} else {
		defer f.Close()
		renderResume(f)
	}
}

func renderResume(w io.Writer) error {
	resumeTemplPath := "index.html.tmpl"

	resumeYAMLContent, err := os.ReadFile("resume.yaml")
	if err != nil {
		return err
	}

	resume := Resume{}
	if err := yaml.Unmarshal(resumeYAMLContent, &resume); err != nil {
		return err
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
