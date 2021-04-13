swagger:
	docker run --rm -p 8088:8080 -e BASE_URL=/swagger -e URLS='[\
		{url:"/docs/swagger.yaml", name:"API v1"},\
	]' -v $(PWD)/docs:/etc/nginx/html/docs swaggerapi/swagger-ui
