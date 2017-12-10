java -jar ./swagger/swagger2markup-cli-1.3.1.jar convert -i ./swagger/swagger.json -f swagger/swagger -c ./swagger/config.properties
java -jar ./swagger/swagger2markup-cli-1.3.1.jar convert -i ./swagger/swagger.json -f swagger/swagger
asciidoctor -a toc=left ../swagger/swagger.adoc -D docs
