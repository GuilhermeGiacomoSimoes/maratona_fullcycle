#cria o container a partir da imagem do golang
FROM golang:1.14

#diretório que será o diretório pai quando a imagem for executada
WORKDIR /go/src/

#copia todos os arquivos locais da máquina e joga dentro do container
COPY . . 

#comando de build do golang
RUN GOOS=linux go build main.go 

#quando essa imagem for gerada, vai executar o seguinte arquivo
ENTRYPOINT ["./main"]

