# gpt-gitcommits

GPT Git Commit is a command-line tool that leverages OpenAI's GPT-3 language model to simplify the process of creating commit messages for Git commits. It addresses the common challenge of spending excessive time crafting commit messages that effectively describe code changes. With AI Commit, developers can generate a variety of high-quality commit message options within seconds, reducing the time spent on writing commit messages and allowing them to focus more on coding.

## Prerequisites :clipboard:

Before installing AI Commit, you will need to generate your OpenAI API key.

You will also need to have GoLang installed. To check if you have Go installed, open your terminal window and run the following command:
```sh
go --version
```

If you see a version number in the output, it means that Go is installed. If not, you can download and install go using the below documentation:
```sh
https://go.dev/doc/install
```
## Installation :inbox_tray:
To install AI Commit, please follow the steps below:
1. Clone the GPT GIT Commit repository to your local machine by running the following command in your terminal:
```sh
git clone https://github.com/chrispeterjeyaraj/gpt-gitcommits
```
2. Navigate into the cloned repository by running the following command in your terminal:
```sh
cd gpt-gitcommits
```
3. Run the install.sh script by running the following command in your terminal:
```sh
./install.sh
```
This will install the required Go packages, copy the gptcommit executable to your /usr/local/bin directory, and add an alias to your .bashrc file.

## Usage :zap:
To use GPT GIT Commit, do the following:
1. Export your OpenAI API key as an environment variable by running:
```sh
export OPENAI_API_KEY="YOUR_API_KEY_HERE"
```
2. Navigate to the directory of the Git repository you wish to commit changes to, and run the following command in your terminal:
```sh
aigitcommit

```
This will prompt AI Commit to generate a commit message based on your changes. The generated message will be displayed, along with the option to either accept it (`y`), reject it and generate a new one (`n`), or edit it manually (`e`). If you choose to edit it manually, you will be prompted to enter your desired commit message.

**Note:** AI Commit generates one commit message at a time, with the option to generate more if the initial message is rejected.

## Uninstallation :outbox_tray:
To uninstall AI Commit, please follow the steps below:
1. Navigate to the directory where you cloned the AI Commit repository by running the following command in your terminal:
```sh
cd <path-to-aicommits-repo>
```
2. Run the uninstall.sh script by running the following command in your terminal:
```sh
./uninstall.sh
```
This will remove the gitcommit executable from your /usr/local/bin directory, remove the alias from your .bashrc file, and uninstall the required Python packages.

## Contribution :raised_hands:
I welcome contributions from the community! If you find a bug 🐛 or have an idea for a new feature 💡, please open an issue or submit a pull request.

## LICENSE :scroll:
AI Commit is licensed under the MIT License. See the [LICENSE](./LICENSE) file for more information.

----------
- If you found GPT GIT Commit useful, please consider giving this repo a star ⭐️!

