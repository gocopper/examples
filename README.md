<p align="center">
  <a href="https://gocopper.dev" target="_blank" rel="noopener noreferrer">
    <img width="180" src="https://gocopper.dev/static/logo.svg" alt="Copper logo">
  </a>
</p>

<br />

## Copper Examples

<p>
A collection of example projects using Copper. 
</p>

### Projects

| Project                                                                             | Description                                      |
|-------------------------------------------------------------------------------------|--------------------------------------------------|
| <a href="https://github.com/gocopper/examples/tree/main/hackernews">Hacker News</a> | A minimal HN clone built using Copper + Tailwind |

### Run Any Project Locally

1. Install Copper & Wire
```shell
❯ go install github.com/gocopper/cli/cmd/copper@v1
❯ go install github.com/google/wire/cmd/wire@latest
```

2. Clone Repository
```
❯ git clone https://github.com/gocopper/examples.git copper-ex
```

3. Run App Server
```
❯ cd copper-ex/<project>
❯ copper run
```

4. Run NPM (if applicable)
```
❯ cd copper-ex/<project>/web
❯ npm run dev
```

5. Open http://localhost:5901 in browser