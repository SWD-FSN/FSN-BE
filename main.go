package main

import "social_network/cmd"

func main() {
	cmd.Execute()
}

// if err := godotenv.Load(); err != nil {
// 	log.Fatal(fmt.Sprintf(noti.EnvLoadErr, "") + err.Error())
// }

// service, _ := service.GeneratePostService()
// res := service.GetPosts(context.Background())
// log.Println(res)
