package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("App started...")
	//creating app, window and adjusting settings
	a := app.New()
	w := a.NewWindow("Deployment app")
	w.Resize(fyne.NewSize(800, 600))
	//setup whole UI
	SetupUI(w)

	w.ShowAndRun()
}

func executeSSHCommand(user, host, commands string) (string, error) {
	var output bytes.Buffer
	cmd := exec.Command("ssh", fmt.Sprintf("%s@%s", user, host), commands)

	cmd.Stdout = &output
	cmd.Stderr = &output

	err := cmd.Run()
	return output.String(), err
}

func deployPortfolio() {
	err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
        return
    }
	// VPS connection details
	vpsUser := os.Getenv("VPS_USER")
    vpsHost := os.Getenv("VPS_HOST")
    projectPath := os.Getenv("PROJECT_PATH")

	commands := fmt.Sprintf(`
		export LC_ALL=C.UTF-8
		export LANG=C.UTF-8
		export NVM_DIR="$HOME/.nvm"
		[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh" # Load nvm
		cd %s && \
		git pull && \
		nvm install 20.17.0 && \ # Ensure Node.js 20 is installed
		nvm use 20.17.0 && \ # Use Node.js 20
		npm install && \
		npm run build && \
		pm2 restart portfolio
	`, projectPath)

	// Execute the SSH command
	output, err := executeSSHCommand(vpsUser, vpsHost, commands)
	if err != nil {
		fmt.Println("Output:", output)
		fmt.Printf("Deployment failed: %s\n", err.Error())
	} else {
		fmt.Println("Output:", output)
		fmt.Println("Deployment succeeded!")
	}
}

func SetupUI(w fyne.Window) {
	// Set window content
	w.SetContent(
		container.NewVBox(
			container.NewBorder(
				layout.NewSpacer(),
				nil,
				widget.NewLabel("Portfolio!"),
				container.NewBorder(
					nil,nil,
					widget.NewButton("Deploy!", func() {
						deployPortfolio()
					}),
					nil,
				),
			),
			container.NewBorder(
				layout.NewSpacer(),
				nil,
				widget.NewLabel("Weather app!"),
				container.NewBorder(
					nil,nil,
					widget.NewButton("Deploy!", func() {
						w.Close()
					}),
					nil,
				),
			),
			container.NewBorder(
				layout.NewSpacer(),
				nil,
				widget.NewLabel("Notes Frontend!"),
				container.NewBorder(
					nil,nil,
					widget.NewButton("Deploy!", func() {
						w.Close()
					}),
					nil,
				),
			),
			container.NewBorder(
				layout.NewSpacer(),
				nil,
				widget.NewLabel("Notes Api!"),
				container.NewBorder(
					nil,nil,
					widget.NewButton("Deploy!", func() {
						w.Close()
					}),
					nil,
				),
			),
			container.NewBorder(
				layout.NewSpacer(),
				nil,
				widget.NewLabel("Playground Frontend!"),
				container.NewBorder(
					nil,nil,
					widget.NewButton("Deploy!", func() {
						w.Close()
					}),
					nil,
				),
			),
			container.NewBorder(
				layout.NewSpacer(),
				nil,
				widget.NewLabel("Playground Api!"),
				container.NewBorder(
					nil,nil,
					widget.NewButton("Deploy!", func() {
						w.Close()
					}),
					nil,
				),
			),
		),
	)

	// Set window icon
	iconData, err := os.ReadFile("icon.png")
	if err == nil {
		w.SetIcon(fyne.NewStaticResource("icon.png", iconData))
	}
}