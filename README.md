# Guessing-Game
A Simple number guessing game build with Go. The program randomly selects a number between 1 and 100, and the player tries to guess it. After each guess, the game provides feedback ("Too High" or "To Low") until the correct number is guessed. A fun way to practice conditionals, loops, user input, and random number generation in Go.


# 🎮 Guessing Game in Go

Welcome to the **Guessing Game**, a fun little terminal-based game written in Go!  
Your mission: **guess the secret number** (between 1 and 100) within limited attempts.  
Can you beat the game and claim glory? 😎

---

## 🚀 Features

- Random number generation every time you play
- Colorful and styled terminal output (using ANSI escape codes)
- Input validation with friendly error messages
- Feedback on each guess: 📉 Too Low / 📈 Too High
- Retry system with a set number of attempts
- Reveals the correct number if you fail
- Fun emojis for a playful experience

---

## 🛠️ How to Run

### 1. Clone the repository (or copy the code)
```bash
git clone https://github.com/devops-gautamjha/Guessing-Game
cd Guessing-Game
```

### 2. Run the game
```bash
go run main.go
```
> Make sure you have go installed

## 📦 Requirements
- Go 1.24 or higher (but should work on older versions too)

- A terminal that supports ANSI escape sequences (for color output)

## 🎯 Gameplay
- You get 4 attempts to guess a number between 1 and 100

- The game gives you clues:

    - 📉 If your guess is too low

    - 📈 If your guess is too high

- If you guess correctly, you win! 🎉

- If not, you’ll be shown the correct number at the end


## 🙏 Acknowledgments
- Built with ❤️ and Go

- Terminal styling with ANSI escape codes


## 🧙 Author
Gautam Jha

> Feel free to connect or fork the project!