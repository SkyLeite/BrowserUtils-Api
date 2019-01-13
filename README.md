# BrowserUtils-API

This is the API for BrowserUtils.

# Routes
GET `/about/` - Returns version information
POST `/play/mpv` - Plays URL in MPV - `{ "url": "https://www.youtube.com/watch?v=1tAFXThjzsY", "isOnTop": true, "geometry": "50%:50%" }`
POST `/youtubedl` - Downloads URL with Youtube-DL - `{ "url": "https://soundcloud.com/user-650170319-492500742/sophie-take-me-to-dubai-demo", "workingDir": "/home/rodrigo/Music" }`