package templates

import "fmt"

func RenderDockerfileTemplate(version string) string {
	const dockerTemplate = `FROM golang:%v

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

EXPOSE 8080

CMD ["/docker-gs-ping"]`

	return fmt.Sprintf(dockerTemplate, version)
}
