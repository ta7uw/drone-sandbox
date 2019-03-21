  local Pipeline(name, image) = {
  kind: "pipeline",
  name: name,
  workspace: [
  {
    base: "/go",
    path: "src/github.com/takaaki82/drone-sandbox",
    }
  ],
  steps: [
  {
    name: "test",
    image: image,
    commands: [
      "go build -v github.com/takaaki82/drone-sandbox/go",
      "go test -v ./...",
    ]
  }
  ]
  };

    [
      Pipeline("go-1-11", "golang:1.11"),
      Pipeline("go-1-12", "golang:1.12"),
    ]
