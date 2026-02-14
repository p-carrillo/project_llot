# Nginx Integration

- Install drop-in config snippets only.
- Validate candidate config with `nginx -t`.
- Reload only after successful validation.
- Keep previous known-good config for rollback.
