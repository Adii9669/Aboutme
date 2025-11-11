package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Println("Got request:", r.URL.Path)
		fmt.Fprint(w, `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Aditya Raj Singh — Software Engineer</title>
<style>
body {
    font-family: sans-serif;
    max-width: 800px;
    margin: 50px auto;
    line-height: 1.6;
    background: #fafafa;
    color: #222;
}
h1, h2, h3 { color: #333; }
a { color: #007acc; text-decoration: none; }
a:hover { text-decoration: underline; }
section { margin-bottom: 2em; }
hr { border: 0; height: 1px; background: #ddd; margin: 2em 0; }
</style>
</head>
<body>
<h1>Aditya Raj Singh</h1>
<p><strong>Software Engineer</strong> | Backend & Fullstack Developer</p>

<section>
<p>
  <a href="/resume" download>Download Resume</a>  
</p>
</section>

<section>
<h2>About Me</h2>
<p>Hi! I'm Aditya, a software engineer passionate about building scalable systems and experimenting with modern technologies like Go, Next.js, and databases such as NeonDB. I enjoy combining ideas from different domains to create something interesting and different.</p>
</section>

<section>
<h2>Current Projects</h2>
<ul>
  <li><strong>ChatApp</strong> – A real-time communication app built with Go and Next.js, supporting group and personal messaging.</li>
  <li><strong>StudyRoom</strong> – A Django-based collaborative environment for students to work together.</li>
</ul>
</section>

<section>
<h2>Experience</h2>
<ul>
  <li>Software Engineer @ Tescra (2024–Present)</li>
  <li>Focus: Scalable backend APIs, user authentication, and communication systems.</li>
</ul>
</section>

<section>
<h2>Links</h2>
<p>
  <a href="https://github.com/Adii9669">GitHub</a> · 
  <a href="https://www.linkedin.com/in/adii9669/">LinkedIn</a> · 
  <a href="mailto:adityarajsingh799@gmail.com">Email</a>
</p>
</section>

<hr>
<p><em>This page is fully self-contained. No external requests. One HTTP call. Served by a Go binary.</em></p>
</body>
</html>
        `)
	})

	http.HandleFunc("/resume", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=resume.pdf")
		http.ServeFile(w, r, "resume.pdf")

	})

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
