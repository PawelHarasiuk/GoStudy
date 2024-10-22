# My Path Studing Golang

## My Approach 
Go is not my first language so i decided to take diffrent approach to learning go and avoid mistakes I made with other languages.
- I do not follow tutorials unless I get stack and can't find solution in documentation. After I unstack I immediately drop tutorial and get back to coding
- I do not use LLMs for code generation. I've used ai only to brainstorm ideas or make sure I understand concepts

**Reason for avoiding both generated code and tutorial is similar - by using LLM or tutorials I tend to turn off critical thinking and basically I outsource learning nad thinking. 
Which might be good for short term but really damaging for long term career.**

## Techonologies explored 
- Go
  - gin
  - gorilla
- Sql
  - Sqlite
  - Postgres
- AWS
  - AWS Lambda
  - API Gateway
  - EventBrige
  - S3
  - ECR
- Docker
- JavaScript
- HTML/CSS

## Projects 
Most projects here are just demos, because I used them just to explore technologies or to better understand how to work with them.
I try now to use this knowledge, for something bigger and actually finish it, for example here: 
- <a href="https://github.com/PawelHarasiuk/CurrencyNewsletter">Currency Newsletter</a> 

### <a href="https://github.com/PawelHarasiuk/GoStudy/tree/main/GoCrawler">GoCrawler</a>
- After learning basic syntax in TourOfGo I wanted to write code so I jumped into recreating things from my collage. So this is basic web crawler I've done in C# previously. It enters website and looks for all emails on page
### <a href="https://github.com/PawelHarasiuk/GoStudy/tree/main/UniSys">UniSys</a>
- Again exercise from my C# course in collage. Program takes csv as an input, clean it and return in json format. It helped me play around with files and become more confident with syntax
### <a href="https://github.com/PawelHarasiuk/GoStudy/tree/main/moviesServer">moviesServer</a>
- App with server managing movies. I just wanted to learn gorilla framework
### <a href="https://github.com/PawelHarasiuk/GoStudy/tree/main/RestDemo">RestDemo</a>
- Rest api app. I wanted to create rest api server that will work with database. I used strategy pattern to create dependency injection of repository. I can inject csv_repository or postgres_repository and with changing one line of code change behavior 
### <a href="https://github.com/PawelHarasiuk/GoStudy/tree/main/TemplatesDemo">TemplatesDemo</a>
- It is simple demo that i used to learn about go templates and html 
### <a href="https://github.com/PawelHarasiuk/GoStudy/tree/main/Form">Form</a>
- App that gives form to fill and send it to backend. I continued learning working with html/css/javascript and go templates
### <a href="https://github.com/PawelHarasiuk/GoStudy/tree/main/TodoApp">TodoApp</a>
- Just another to do app, but I just wanted to test and use all skills I learned and do not use any tutorial. I've used there net/http, postgres, templates
### <a href="https://github.com/PawelHarasiuk/GoStudy/tree/main/LambdaTest">LambdaTest</a>
- I started learning AWS Lambda, additionally I started learning docker and now it is pretty natural for me similar to git
### <a href="https://github.com/PawelHarasiuk/GoStudy/tree/main/ConnectToS3">ConnectToS3</a>
- I've learned reading and writing to file stored on aws S3
### <a href="https://github.com/PawelHarasiuk/GoStudy/tree/main/StaticWebsite">StaticWebsite</a>
- I wanted to test hosting static website on aws S3 for my other project. Additionally i managed to send request from HTML and javascript to aws lambda go backend, and communicate between those two. PS turned out I did not care about CORS because aws allows seemless communication between lambda and S3
