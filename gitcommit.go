package main

import (
    "fmt"
    "log"
    "os"
    "os/exec"

    "github.com/go-resty/resty/v2"
)

func main() {
    openaiAPIKey := os.Getenv("OPENAI_API_KEY")

    if openaiAPIKey == "" {
        log.Fatal("ERROR: Please set your OpenAI API key as an environment variable named 'OPENAI_API_KEY'")
    }

    // Choose the GPT-3 model you want to use
    modelEngine := "text-davinci-003"

    // Define the input data
    diff, err := exec.Command("git", "diff", "--cached").Output()
    if err != nil {
        log.Fatal("Failed to execute git diff command:", err)
    }

    for {
        // Generate the commit message
        prompt := fmt.Sprintf("Generate a commit message for the following changes:\n%s", diff)

        generatedText, err := generateCommitMessage(prompt, modelEngine, openaiAPIKey)
        if err != nil {
            log.Fatal("Failed to generate commit message:", err)
        }

        // Modify the commit message and present it to the user
        message := generatedText[len("Commit message"):]

        fmt.Printf("Generated commit message:\n%s\n", message)
        userInput := getUserInput("Accept commit message? (y/n/e): ")

        // Commit the changes or regenerate the message
        if userInput == "y" {
            commitChanges(message)
            fmt.Println("Changes committed!")
            break
        } else if userInput == "e" {
            editedText := getUserInput("Enter edited commit message: ")
            commitChanges(editedText)
            fmt.Println("Changes committed!")
            break
        } else {
            fmt.Println("Regenerating commit message...")
        }
    }
}

type OpenAIResponse struct {
    Choices []struct {
        Text string `json:"text"`
    } `json:"choices"`
}

func generateCommitMessage(prompt, modelEngine, openaiAPIKey string) (string, error) {
    client := resty.New()

    resp, err := client.R().
        SetHeaders(map[string]string{
            "Content-Type": "application/json",
            "Authorization": fmt.Sprintf("Bearer %s", openaiAPIKey),
        }).
        SetBody(map[string]interface{}{
            "model":     modelEngine,
            "prompt":    prompt,
            "max_tokens": 50,       // Adjust the max tokens and other parameters as needed
            "temperature": 0.5,     // Adjust the temperature value as needed
        }).
        SetResult(&OpenAIResponse{}).
        Post("https://api.openai.com/v1/completions")

    if err != nil {
        return "", err
    }

    if resp.StatusCode() != 200 {
        return "", fmt.Errorf("API request failed with status code: %d", resp.StatusCode())
    }

    openAIResp, ok := resp.Result().(*OpenAIResponse)
    if !ok {
        return "", fmt.Errorf("Unable to access the desired fields from resp.Result()")
    }

    if len(openAIResp.Choices) == 0 {
        return "", fmt.Errorf("No choices returned in the response")
    }

    return openAIResp.Choices[0].Text, nil
}

func commitChanges(commitMessage string) {
    cmd := exec.Command("git", "commit", "-m", commitMessage)
    err := cmd.Run()
    if err != nil {
        log.Fatal("Failed to execute git commit command:", err)
    }
}

func getUserInput(prompt string) string {
    var userInput string
    fmt.Print(prompt)
    _, err := fmt.Scanln(&userInput)
    if err != nil {
        log.Fatal("Failed to read user input:", err)
    }
    return userInput
}
