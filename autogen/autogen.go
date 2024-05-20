package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

const (
	// Headers constants: Section Modules Year Authors Title URL Optional MLA
	HeaderSection  = "Section"
	HeaderModules  = "Modules"
	HeaderYear     = "Year"
	HeaderAuthors  = "Authors"
	HeaderTitle    = "Title"
	HeaderURL      = "URL"
	HeaderOptional = "Optional"
	HeaderMLA      = "MLA"
)

var description = `This is a comprehensive list of papers on database theory for understanding and building database systems. It covers various aspects of database systems, including the essential theoretical background, classic system design, and multiple modules within the database.

The list is organized into different categories and subcategories for easy navigation. Each paper is accompanied by a title, author, and publication year, along with a link to the full text if available.

This collection serves as a learning and training resource primarily for the Tencent Cloud Database Team and is also open to external researchers, students, and learners interested in database systems.

**In case you are reading this and making the effort to comprehend these papers, we would really like to have a conversation with you regarding opportunities at Tencent Cloud Database Team ([@Henry L.](https://x.com/henrylonng)).**

## Contribution

This list is generated from a Sheet document automatically. If you have any suggestions or would like to contribute to this list, please feel free to file an issue. And we will update our sheet to make the chagnes available for public.

Any contribution that can help improve this list and make it more comprehensive and useful to the community are welcome. Here are some ways you can contribute:

- **Add a new paper**: If you have a paper that you think should be included in this list, please file an issue to provide the paper's title, author, publication year, and a link to the full text (if available).
- **Update an existing paper**: If you find any errors or outdated information in the list, please file an issue to provide the correct information.
- **Remove a paper**: If you think a paper is no longer relevant or useful, please file an issue to suggest its removal.
- **General suggestions**: If you have any general suggestions or feedback on how to improve this list, please file an issue to share your thoughts.`

var badges = []string{
	"[![README autogen](https://github.com/lonng/db-papers/actions/workflows/autogen.yml/badge.svg)](https://github.com/lonng/db-papers/actions/workflows/autogen.yml)",
}

type (
	options struct {
		addrFormat  string
		sheetID     string
		sheetName   string
		title       string
		output      string
		description string
		directory   string
	}

	section struct {
		name    string
		modules []*module
	}

	module struct {
		name    string
		records []*record
	}

	record struct {
		year     string
		authors  string
		title    string
		url      string
		optional string
		mla      string
	}
)

func main() {
	opt := &options{}
	cmd := &cobra.Command{
		Use:   "db-papers",
		Short: "Generate README.md file based on database papers table",
		RunE: func(cmd *cobra.Command, args []string) error {
			return convert(opt)
		},
	}

	// Flags
	cmd.Flags().StringVarP(&opt.addrFormat, "addr-format", "a", "", "Online table address format")
	cmd.Flags().StringVarP(&opt.sheetID, "sheet-id", "i", "", "Online table sheet ID")
	cmd.Flags().StringVarP(&opt.sheetName, "sheet-name", "n", "", "Online table sheet Name")
	cmd.Flags().StringVarP(&opt.title, "title", "T", "Database Papers", "Title of the README.md")
	cmd.Flags().StringVarP(&opt.description, "description", "d", description, "Description of the README.md")
	cmd.Flags().StringVarP(&opt.directory, "directory", "D", "papers", "Directory to save the papers")
	cmd.Flags().StringVarP(&opt.output, "output", "o", "README.md", "Output file name (default to README.md)")

	if err := cmd.Execute(); err != nil {
		cobra.CheckErr(err)
	}
}

func convert(opt *options) error {
	if opt.addrFormat == "" || opt.sheetID == "" || opt.sheetName == "" {
		return errors.New("missing required flags")
	}

	url := fmt.Sprintf(opt.addrFormat, opt.sheetID, url.QueryEscape(opt.sheetName))
	content, err := http.Get(url)
	if err != nil {
		return err
	}

	data, err := csv.NewReader(content.Body).ReadAll()
	if err != nil {
		return err
	}
	if len(data) < 1 {
		return errors.New("no data found")
	}

	indexes := map[string]int{}
	for index, header := range data[0] {
		indexes[header] = index
	}

	fmt.Println("Headers", data[0])
	fmt.Println("Header Indexes", indexes)

	var sections []*section

	// Process data
	for _, row := range data[1:] {
		// Illegal row if section name is empty and no sections exist
		if row[indexes[HeaderSection]] == "" && len(sections) == 0 {
			return errors.New("section name is empty")
		}

		// Found new section
		if sectionName := row[indexes[HeaderSection]]; sectionName != "" {
			sections = append(sections, &section{name: sectionName})
		}

		// Illegal row if module name is empty and no modules exist in the last section
		lastSection := sections[len(sections)-1]
		if row[indexes[HeaderModules]] == "" && len(lastSection.modules) == 0 {
			return errors.New("module name is empty")
		}

		// Found new module
		if moduleName := row[indexes[HeaderModules]]; moduleName != "" {
			lastSection.modules = append(lastSection.modules, &module{name: moduleName})
		}

		// Add record to the last module
		lastModule := lastSection.modules[len(lastSection.modules)-1]
		lastModule.records = append(lastModule.records, &record{
			year:     strings.TrimSpace(row[indexes[HeaderYear]]),
			authors:  strings.TrimSpace(row[indexes[HeaderAuthors]]),
			title:    strings.TrimSpace(row[indexes[HeaderTitle]]),
			url:      strings.TrimSpace(row[indexes[HeaderURL]]),
			optional: row[indexes[HeaderOptional]],
			mla:      strings.TrimSpace(row[indexes[HeaderMLA]]),
		})
	}

	var output string
	generate := func(s string, newline int) {
		output += s + strings.Repeat("\n", newline)
	}

	// Generate Badges if exists
	if len(badges) > 0 {
		for _, badge := range badges {
			generate(badge, 1)
		}
		generate("", 1)
		generate("---", 2)
	}

	// Generate Title and Description
	generate(fmt.Sprintf("# %s", opt.title), 2)
	generate(opt.description, 2)

	// Generate Table Of Contents
	generate("## Table of Contents", 2)
	for _, s := range sections {
		generate(fmt.Sprintf("- [%s](#%s)", s.name, strings.ToLower(strings.ReplaceAll(s.name, " ", "-"))), 1)
		for _, m := range s.modules {
			generate(fmt.Sprintf("  - [%s](#%s)", m.name, strings.ToLower(strings.ReplaceAll(m.name, " ", "-"))), 1)
		}
	}

	// Add a new line
	generate("", 1)

	padding := func(space int) string {
		return strings.Repeat(" ", space)
	}

	knownFiles := map[string]struct{}{}

	// Generate Sections
	for _, s := range sections {
		generate(fmt.Sprintf("## %s", s.name), 2)
		for _, m := range s.modules {
			generate(fmt.Sprintf("### %s", strings.ReplaceAll(m.name, ":", " -")), 2)
			// Sort all records by year
			sort.Slice(m.records, func(i, j int) bool {
				return m.records[i].year < m.records[j].year
			})
			for i, r := range m.records {
				// Append dot to authors if not present
				if len(r.authors) > 0 && r.authors[len(r.authors)-1] != '.' {
					r.authors = r.authors + "."
				}

				var local string

				// Download the paper
				if opt.directory != "" && r.url != "" {
					// Create directory if not exists
					dir := filepath.Join(opt.directory,
						strings.ToLower(strings.ReplaceAll(m.name, " ", "-")))
					file := strings.ToLower(strings.ReplaceAll(r.title, " ", "-")) + ".pdf"
					path := filepath.Join(dir, strings.ReplaceAll(file, "/", "-"))
					knownFiles[path] = struct{}{}
					if _, err := os.Stat(dir); os.IsNotExist(err) {
						if err := os.MkdirAll(dir, 0755); err != nil {
							return err
						}
					}

					// Download the paper is not exists
					if _, err := os.Stat(path); os.IsNotExist(err) {
						paper, err := http.Get(r.url)
						if err != nil {
							return err
						}

						fmt.Println("Download Paper", r.title)
						fmt.Println(padding(2), "url", r.url)
						fmt.Println(padding(2), "path", path)

						data, err := io.ReadAll(paper.Body)
						if err != nil {
							return err
						}

						// Write to file
						if err := os.WriteFile(path, data, 0644); err != nil {
							return err
						}
					}

					local = path

					// Check if the file is a PDF
					pdf, err := os.OpenFile(path, os.O_RDONLY, 0644)
					if err == nil {
						magicNumber := make([]byte, 4)
						_, err := io.ReadFull(pdf, magicNumber)
						if err != nil || string(magicNumber) != "%PDF" {
							_ = os.Remove(path)
							local = ""
						}
					}

				}

				if local != "" {
					generate(fmt.Sprintf("- [%s](%s) (%s) - %s", r.title, local, r.year, r.authors), 1)
				} else {
					generate(fmt.Sprintf("- [%s](%s) (%s) - %s", r.title, r.url, r.year, r.authors), 1)
				}

				if i == len(m.records)-1 {
					generate("", 1)
				}
			}
		}
	}

	// Clean redundant files
	if opt.directory != "" {
		err := filepath.WalkDir(opt.directory, func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				return nil
			}
			if _, ok := knownFiles[path]; !ok {
				fmt.Println("Delete unknown file", path)
				return os.Remove(path)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	if opt.output != "" {
		// Write to file
		return os.WriteFile(opt.output, []byte(output), 0644)
	}

	// Write to stdout
	fmt.Println(output)

	return nil
}
