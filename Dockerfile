FROM httpd:2.4
COPY ["simple web app.html", "/usr/local/apache2/htdocs/index.html"]
