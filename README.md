
# 📸 Event-O-Matic 3000™  
_A program that does everything you were too lazy to do after an event._

---

## 🧠 What Is This?

So you had an event. People showed up. Photos were taken. Now you’re drowning in image files like `IMG_3298123.JPG` and have zero energy to write that “Thanks to everyone who came…” post for LinkedIn.  
**Enter: this glorious Go program.**

It does all the boring stuff for you using Gemini AI. Because you deserve to chill.

---

## 🔮 Features (aka Why This Is Better Than You)

- **Auto Event Naming** – Takes your chaotic description and turns it into a nice, lowercase, filename-friendly name. No more `final_final_v2_reallyfinal.pdf`.
- **Markdown + PDF Summary** – So you can pretend you wrote a meaningful recap. Impress your boss or your cat.
- **Social Media Wizardry** – Generates platform-specific posts so you sound like a pro on:
  - LinkedIn (professional mode: ON)
  - Twitter/X/Whatever Elon’s doing
  - Instagram (caption-worthy lines incoming)
  - Forums (hello, boomers)
  - Blogs (because Medium still exists apparently)
- **AI-Powered Image Naming** – Never again deal with `IMG_20230413_143933.jpg`. Now you get stuff like `pizza_party.jpg` or `confused_dog_on_stage.jpg` (kidding, kinda).
- **Organized Output** – Dumps everything neatly into a folder in `out/`, unlike your Downloads folder.

---

## ⚙️ How to Use (in 4 steps because 3 is too mainstream)

1. Put your Gemini API key in a `.env` file like so:
   ```
   GEMINI_API_KEY=your_magical_key_goes_here
   ```

2. Write an **epic** event description in `input.txt`. The more dramatic, the better.

3. Dump your `.jpg` and `.png` files in the same folder. No, not `.webp`, we don’t like those.

4. Run the thing:
   ```bash
   go run main.go
   ```

---

## 📁 What You’ll Get (aka AI Did This, Not You)

Inside `out/<cool_event_name>/`:
```
├── README.md         # Event summary (totally not written by AI)
├── summary.pdf       # The same thing but fancier
├── linkedin.txt      # Corporate buzzwords ✅
├── instagram.txt     # #vibes #aesthetic
├── twitter.txt       # Now with 30% more sass
├── blog.txt          # Looks like you tried
├── forum.txt         # Hello old-school internet
└── renamed_images/   # Sensible filenames, finally
```

---

## 🧾 Requirements (We Don't Make The Rules)

- Go 1.18+
- Your soul (just kidding, just a Gemini API key)
- These libraries:
  - `github.com/google/generative-ai-go/genai`
  - `github.com/joho/godotenv`
  - `github.com/jung-kurt/gofpdf`
  - `google.golang.org/api/option`

---

## 🧙‍♂️ Pro Tips

- The better your `input.txt`, the better the output. But hey, even if it sucks, Gemini will try its best.
- Works great for college fests, hackathons, family reunions, or your dog’s birthday party.

---

## 🙏 Credits

Written in Go because Python devs need to suffer a little.

Built by Manappattil Chaandi Enterprises, who probably had too much caffeine and not enough sleep.

---

## 🪦 License

MIT. Use it. Break it. Fork it. Blame someone else.

---

> “Work smart, not hard.”  
> – This program, probably
