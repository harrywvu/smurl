- POST /api/v1/urls — body is { "original_url": "..." }. Success: 201 with code, short_url, original_url, created_at. Failure: 400 with invalid_json or invalid_url.
- GET /:code — success: 302 with Location header. Unknown code: 404 with not_found.
- HEAD /:code — same status and headers as GET, no body.
- GET /healthz — 200. Every error uses the same shape: { "error": { "code": "...", "message": "..." } }.
1. Success body, exactly: code (string), short_url (absolute, from BASE_URL), original_url (string), created_at (RFC3339 UTC). No extra fields.
2. Validation rules: trim outer whitespace, scheme must be http/https, host required, max 2048 chars. Body that won't parse or is missing original_url → invalid_json; parses but fails those rules → invalid_url.
3. Reserved codes: healthz and metrics are both exactly 7 characters, so your generator can produce codes that collide with real routes. Add them to a blocklist and regenerate on hit. Also {"status":"ok"} is the /healthz body.
