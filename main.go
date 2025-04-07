package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"github.com/jung-kurt/gofpdf"
	"google.golang.org/api/option"
)

const (
	InputTextFile  = "input.txt"
	OutputBasePath = "out"
	APIKeyEnvVar   = "GEMINI_API_KEY"
)

func callGemini(ctx context.Context, model *genai.GenerativeModel, prompt string) (string, error) {
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}
	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no response from Gemini")
	}
	return fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0]), nil
}

func readInputText() (string, error) {
	data, err := os.ReadFile(InputTextFile)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func writeOutputFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

func generatePDF(outputPath, content string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.MultiCell(0, 10, content, "", "L", false)
	return pdf.OutputFileAndClose(outputPath)
}

func processImages(ctx context.Context, model *genai.GenerativeModel, imgPaths []string, eventDir string) []string {
	var renamed []string
	for i, path := range imgPaths {
		imgData, err := os.ReadFile(path)
		if err != nil {
			log.Println("Failed to read image:", path)
			continue
		}

		prompt := "Give a short name for this image"
		resp, err := model.GenerateContent(ctx,
			genai.Text(prompt),
			genai.ImageData("jpeg", imgData),
		)
		var newName string
		if err != nil || len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
			newName = fmt.Sprintf("image_%02d%s", i+1, filepath.Ext(path))
		} else {
			slug := strings.ToLower(strings.ReplaceAll(fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0]), " ", "_"))
			slug = strings.Map(func(r rune) rune {
				if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_' {
					return r
				}
				return -1
			}, slug)
			newName = fmt.Sprintf("%s%s", slug, filepath.Ext(path))
		}

		dest := filepath.Join(eventDir, newName)
		os.Rename(path, dest)
		renamed = append(renamed, newName)
	}
	return renamed
}

func main() {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv(APIKeyEnvVar)
	if apiKey == "" {
		log.Fatal("Missing GEMINI_API_KEY environment variable")
	}
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Failed to create Gemini client: %v", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	inputText, err := readInputText()
	if err != nil {
		log.Fatal("Error reading input text:", err)
	}

	eventNamePrompt := "Generate a short, filename-friendly event name from this description. Only return the name, no formatting, no punctuation: " + inputText
	eventName, err := callGemini(ctx, model, eventNamePrompt)
	if err != nil {
		log.Fatal("Error generating event name:", err)
	}
	eventName = strings.TrimSpace(eventName)
	eventName = strings.ToLower(strings.ReplaceAll(eventName, " ", "_"))
	eventDir := filepath.Join(OutputBasePath, eventName)
	os.MkdirAll(eventDir, 0755)

	readmePrompt := "Generate a markdown summary for this event:\n" + inputText
	readmeContent, _ := callGemini(ctx, model, readmePrompt)
	writeOutputFile(filepath.Join(eventDir, "README.md"), readmeContent)
	generatePDF(filepath.Join(eventDir, "summary.pdf"), readmeContent)

	platforms := []string{"linkedin", "forum", "blog", "instagram", "twitter"}
	for _, platform := range platforms {
		prompt := fmt.Sprintf("Write a %s post about this event:\n%s", platform, inputText)
		content, _ := callGemini(ctx, model, prompt)
		writeOutputFile(filepath.Join(eventDir, platform+".txt"), content)
	}

	files, _ := ioutil.ReadDir(".")
	var images []string
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".jpg") || strings.HasSuffix(f.Name(), ".png") {
			images = append(images, f.Name())
		}
	}
	processImages(ctx, model, images, eventDir)

	fmt.Println("✅ Event generated in:", eventDir)
}
