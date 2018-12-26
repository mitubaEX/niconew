# niconew

niconew get newest video list from [nicovideo](https://www.nicovideo.jp/).

## Usage

```sh
# install package
dep ensure

# run
go run main.go
```

## heroku to deploy

Please change main.go.

```
- log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
+ log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), api.MakeHandler()))
```

OK. push to heroku.

```
# heroku configuration
heroku login
heroku container:login
heroku create

# push
docker build -t server .
docker tag server registry.heroku.com/<your app>/web
docker push registry.heroku.com/<your app>/web
heroku container:release web
```

## Dependencies

- [PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery)
