#optimach-be

*Backend: Optimach - Core API & Database*

Intro
Optimach is a comprehensive nutrition calculating application born out of a personal mission. Initially designed as a mobile solution to help my mother track calories for weight management—crucial for managing her osteoarthritis—the project evolved into a Progressive Web App (PWA). This transition was driven by her need for a desktop-accessible tool while working, ensuring her health goals remained within reach regardless of the device. What started as a focused mobile utility has expanded into a full-stack ecosystem, refined by my experience as a Full-Stack Intern at Kalbe Consumer Health and a rigorous "Monk Mode" remastering period.

*Backend Tech Stack*
Language: Golang 1.23.5 (Chosen for high-performance concurrency and type safety)


*Data Architecture*
Optimach leverages a hybrid data model that merges two distinct worlds:

The Optimach Core Database: This is a proprietary database I personally curated based on my own history of calorie counting. I spent years finding the exact calculations and nutritional profiles that actually yielded results, and I've baked that experience into the logic of this app.

FatSecret Premier-Free API Food Database: To ensure global coverage, I integrated the FatSecret premiere food database, allowing users access to millions of verified items.

*Database Optimization*
Rather than forcing the API to run massive, expensive join statements for the "Food Diary" feature, I implemented PostgreSQL Views, which act as a "virtual table" that pre-calculates daily totals and macronutrient splits within the database itself. This approach keeps the Go code clean—the backend simply queries a single "View" instead of managing complex logic, significantly reducing execution time and API complexity.

*Distributed System Challenges*
Moving from a monolithic architecture to a distributed one (Client on Vercel, API on Leapcell, DB on Neon) required a much higher attention to detail. Coordinating these independent units meant handling partial failures and network timeouts that a single-server setup doesn't face.

Development is still is ongoing. There are bugs, and there will be more as we scale. Currently, I am focused on manual testing and rigorous edge-case logging to refine the system stability even more in the future.
