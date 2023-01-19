package pagemaker

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Properties struct {
	data map[string]string
}

func (p Properties) Get(key string) string {
	return p.data[key]
}

type Database struct {
}

func (d Database) GetProperties(dbName string) (Properties, error) {
	fileName := dbName + ".txt"
	f, err := os.Open(fileName)
	if err != nil {
		return Properties{}, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	properties := map[string]string{}
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "=")
		if len(data) != 2 {
			return Properties{}, fmt.Errorf("ファイル形式が異なっています。")
		}
		properties[data[0]] = data[1]
	}
	//open file here

	return Properties{data: properties}, nil
}

type HtmlWriter struct {
	writer *bytes.Buffer
}

func NewHtmlWriter(writer *bytes.Buffer) HtmlWriter {
	return HtmlWriter{writer: writer}
}
func (hw *HtmlWriter) Write(str string) {
	hw.writer.Write([]byte(str))
}
func (hw *HtmlWriter) Title(title string) {
	hw.Write("<html>")
	hw.Write("<head>")
	hw.Write(fmt.Sprintf("<titile>%s</title>", title))
	hw.Write("</head>")
	hw.Write("<body>¥n")
	hw.Write(fmt.Sprintf("<h1>%s</h1>", title))
}
func (hw *HtmlWriter) Paragraph(msg string) {
	hw.Write(fmt.Sprintf("<p>%s</p>", msg))
}

type PageMaker struct {
}

func (pw PageMaker) MakeWelComPage(mailAddr string) error {
	mailDrop, err := Database{}.GetProperties("maildata")
	if err != nil {
		return err
	}
	userName := mailDrop.Get(mailAddr)
	var fileWriter bytes.Buffer
	writer := NewHtmlWriter(&fileWriter)
	writer.Title(fmt.Sprintf("Welcome to %s's page", userName))
	writer.Paragraph(fmt.Sprintf("%sのページへようこそ", userName))
	file, err := os.Create("./result.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(fileWriter.String())
	return err
}
