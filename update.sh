# need "gem install asciidoctor"
goagen bootstrap -d github.com/uzimith/goa-learn/design

echo swagger2markup.markupLanguage=MARKDOWN > swagger/config.properties
docker pull swagger2markup/swagger2markup
docker run --rm -v $(pwd)/swagger/:/opt swagger2markup/swagger2markup convert -i /opt/swagger.yaml -f /opt/swagger
docker run --rm -v $(pwd)/swagger/:/opt swagger2markup/swagger2markup convert -i /opt/swagger.yaml -f /opt/swagger -c /opt/config.properties
asciidoctor -a toc=left ./swagger/swagger.adoc -D docs
