// Package yiigo makes Golang development easier !
//
// Basic usage
//
// MySQL
//
//    // default db
//    yiigo.RegisterDB(yiigo.AsDefault, yiigo.MySQL, "root:root@tcp(localhost:3306)/test")
//
//    yiigo.DB().Get(&User{}, "SELECT * FROM `user` WHERE `id` = ?", 1)
//    yiigo.Orm().First(&User{}, 1)
//
//    // other db
//    yiigo.RegisterDB("foo", yiigo.MySQL, "root:root@tcp(localhost:3306)/foo")
//
//    yiigo.DB("foo").Get(&User{}, "SELECT * FROM `user` WHERE `id` = ?", 1)
//    yiigo.Orm("foo").First(&User{}, 1)
//
// MongoDB
//
//    // default mongodb
//    yiigo.RegisterMongoDB(yiigo.AsDefault, "mongodb://localhost:27017")
//
//    ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
//    yiigo.Mongo().Database("test").Collection("numbers").InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
//
//    // other mongodb
//    yiigo.RegisterMongoDB("foo", "mongodb://localhost:27017")
//
//    ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
//    yiigo.Mongo("foo").Database("test").Collection("numbers").InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
//
// Redis
//
//    // default redis
//    yiigo.RegisterRedis(yiigo.AsDefault, "localhost:6379")
//
//    conn, err := yiigo.Redis().Get()
//
//    if err != nil {
//        log.Fatal(err)
//    }
//
//    defer yiigo.Redis().Put(conn)
//
//    conn.Do("SET", "test_key", "hello world")
//
//    // other redis
//    yiigo.RegisterRedis("foo", "localhost:6379")
//
//    conn, err := yiigo.Redis("foo").Get()
//
//    if err != nil {
//        log.Fatal(err)
//    }
//
//    defer yiigo.Redis("foo").Put(conn)
//
//    conn.Do("SET", "test_key", "hello world")
//
// Config
//
//    // env.toml
//    //
//    // [app]
//    // env = "dev"
//    // debug = true
//    // port = 50001
//
//    yiigo.SetEnvFile("env.toml")
//
//    yiigo.Env.GetBool("app.debug", true)
//    yiigo.Env.GetInt("app.port", 12345)
//    yiigo.Env.GetString("app.env", "dev")
//
//
// Zipkin
//
//    reporter := yiigo.NewZipkinHTTPReporter("http://localhost:9411/api/v2/spans")
//
//    err := yiigo.RegisterZipkinTracer(yiigo.AsDefault, reporter)
//
//    if err != nil {
//        log.Fatal(err)
//    }
//
//    client, err := yiigo.ZTracer().HTTPClient()
//
//    if err != nil {
//        log.Fatal(err)
//    }
//
//    b, err := client.Get(context.Background(), "url...",
//        yiigo.WithRequestHeader("Content-Type", "application/json"),
//        yiigo.WithRequestTimeout(5*time.Second),
//    )
//
//    if err != nil {
//        log.Fatal(err)
//    }
//
//    fmt.Println(string(b))
//
// Logger
//
//    // default logger
//    yiigo.RegisterLogger(yiigo.AsDefault, "app.log")
//    yiigo.Logger().Info("hello world")
//
//    // other logger
//    yiigo.RegisterLogger("foo", "foo.log")
//    yiigo.Logger("foo").Info("hello world")
//
// For more details, see the documentation for the types and methods.
//
package yiigo
